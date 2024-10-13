package data

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/shirou/gopsutil/disk"

	"github.com/TheTNB/panel/internal/app"
	"github.com/TheTNB/panel/internal/biz"
	"github.com/TheTNB/panel/pkg/db"
	"github.com/TheTNB/panel/pkg/io"
	"github.com/TheTNB/panel/pkg/shell"
	"github.com/TheTNB/panel/pkg/str"
	"github.com/TheTNB/panel/pkg/types"
)

type backupRepo struct {
	setting biz.SettingRepo
	website biz.WebsiteRepo
}

func NewBackupRepo() biz.BackupRepo {
	return &backupRepo{
		setting: NewSettingRepo(),
		website: NewWebsiteRepo(),
	}
}

// List 备份列表
func (r *backupRepo) List(typ string) ([]*types.BackupFile, error) {
	backupPath, err := r.GetPath(typ)
	if err != nil {
		return nil, err
	}

	files, err := io.ReadDir(backupPath)
	if err != nil {
		return nil, err
	}

	var backupList []*types.BackupFile
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}
		backupList = append(backupList, &types.BackupFile{
			Name: file.Name(),
			Size: str.FormatBytes(float64(info.Size())),
		})
	}

	return backupList, nil
}

// Create 创建备份
func (r *backupRepo) Create(typ, target string, path ...string) error {
	defPath, err := r.GetPath(typ)
	if err != nil {
		return err
	}
	if len(path) > 0 && path[0] != "" {
		defPath = path[0]
	}

	switch typ {
	case "website":
		return r.createWebsite(defPath, target)
	case "mysql":
		return r.createMySQL(defPath, target)
	case "postgres":
		return r.createPostgres(defPath, target)
	case "panel":
		return r.createPanel(defPath)

	}

	return errors.New("未知备份类型")
}

// Delete 删除备份
func (r *backupRepo) Delete(typ, name string) error {
	path, err := r.GetPath(typ)
	if err != nil {
		return err
	}

	file := filepath.Join(path, name)
	return io.Remove(file)
}

// CutoffLog 切割日志
// path 保存目录绝对路径
// target 待切割日志文件绝对路径
func (r *backupRepo) CutoffLog(path, target string) error {
	if !io.Exists(target) {
		return errors.New("日志文件不存在")
	}

	to := filepath.Join(path, fmt.Sprintf("%s_%s.zip", time.Now().Format("20060102150405"), filepath.Base(target)))
	if err := io.Compress([]string{target}, to, io.Zip); err != nil {
		return err
	}

	return io.Remove(target)
}

// CleanExpired 清理过期备份
// path 备份目录绝对路径
// prefix 目标文件前缀
// save 保存份数
func (r *backupRepo) CleanExpired(path, prefix string, save int) error {
	files, err := io.ReadDir(path)
	if err != nil {
		return err
	}

	var filtered []os.FileInfo
	for _, file := range files {
		if strings.HasPrefix(file.Name(), prefix) && strings.HasSuffix(file.Name(), ".zip") {
			info, err := os.Stat(filepath.Join(path, file.Name()))
			if err != nil {
				continue
			}
			filtered = append(filtered, info)
		}
	}

	// 排序所有备份文件，从新到旧
	slices.SortFunc(filtered, func(a, b os.FileInfo) int {
		if a.ModTime().After(b.ModTime()) {
			return -1
		}
		if a.ModTime().Before(b.ModTime()) {
			return 1
		}
		return 0
	})
	if len(filtered) <= save {
		return nil
	}

	// 切片保留 save 份，删除剩余
	toDelete := filtered[save:]
	for _, file := range toDelete {
		filePath := filepath.Join(path, file.Name())
		if err = os.Remove(filePath); err != nil {
			return err
		}
	}

	return nil
}

// GetPath 获取备份路径
func (r *backupRepo) GetPath(typ string) (string, error) {
	backupPath, err := r.setting.Get(biz.SettingKeyBackupPath)
	if err != nil {
		return "", err
	}

	backupPath = filepath.Join(backupPath, typ)
	if !io.Exists(backupPath) {
		if err = io.Mkdir(backupPath, 0644); err != nil {
			return "", err
		}
	}

	return backupPath, nil
}

// createWebsite 创建网站备份
func (r *backupRepo) createWebsite(to string, name string) error {
	website, err := r.website.GetByName(name)
	if err != nil {
		return err
	}

	if err = r.preCheckPath(to, website.Path); err != nil {
		return err
	}

	backup := filepath.Join(to, fmt.Sprintf("%s_%s.zip", website.Name, time.Now().Format("20060102150405")))
	if _, err = shell.Execf(`cd '%s' && zip -r '%s' .`, website.Path, backup); err != nil {
		return err
	}

	return nil
}

// createMySQL 创建 MySQL 备份
func (r *backupRepo) createMySQL(to string, name string) error {
	rootPassword, err := r.setting.Get(biz.SettingKeyMySQLRootPassword)
	if err != nil {
		return err
	}
	mysql, err := db.NewMySQL("root", rootPassword, "/tmp/mysql.sock")
	if err != nil {
		return err
	}
	if exist, _ := mysql.DatabaseExists(name); !exist {
		return fmt.Errorf("数据库不存在：%s", name)
	}
	size, err := mysql.DatabaseSize(name)
	if err != nil {
		return err
	}
	if err = r.preCheckDB(to, size); err != nil {
		return err
	}

	if err = os.Setenv("MYSQL_PWD", rootPassword); err != nil {
		return err
	}
	backup := filepath.Join(to, fmt.Sprintf("%s_%s.sql", name, time.Now().Format("20060102150405")))
	if _, err = shell.Execf(`mysqldump -u root '%s' > '%s'`, name, backup); err != nil {
		return err
	}
	if err = os.Unsetenv("MYSQL_PWD"); err != nil {
		return err
	}

	if err = io.Compress([]string{backup}, backup+".zip", io.Zip); err != nil {
		return err
	}
	if err = io.Remove(backup); err != nil {
		return err
	}

	return nil
}

// createPostgres 创建 PostgreSQL 备份
func (r *backupRepo) createPostgres(to string, name string) error {
	postgres, err := db.NewPostgres("postgres", "", "127.0.0.1", fmt.Sprintf("%s/server/postgresql/data/pg_hba.conf", app.Root), 5432)
	if err != nil {
		return err
	}
	if exist, _ := postgres.DatabaseExist(name); !exist {
		return fmt.Errorf("数据库不存在：%s", name)
	}
	size, err := postgres.DatabaseSize(name)
	if err != nil {
		return err
	}
	if err = r.preCheckDB(to, size); err != nil {
		return err
	}

	backup := filepath.Join(to, fmt.Sprintf("%s_%s.sql", name, time.Now().Format("20060102150405")))
	if _, err = shell.Execf(`su - postgres -c "pg_dump '%s'" > '%s'`, name, backup); err != nil {
		return err
	}

	if err = io.Compress([]string{backup}, backup+".zip", io.Zip); err != nil {
		return err
	}
	if err = io.Remove(backup); err != nil {
		return err
	}

	return nil
}

// createPanel 创建面板备份
func (r *backupRepo) createPanel(to string) error {
	backup := filepath.Join(to, fmt.Sprintf("panel_%s.zip", time.Now().Format("20060102150405")))

	if err := r.preCheckPath(to, filepath.Join(app.Root, "panel")); err != nil {
		return err
	}

	return io.Compress([]string{
		filepath.Join(app.Root, "panel"),
		"/usr/local/sbin/panel-cli",
		"/usr/local/etc/panel/config.yml",
	}, backup, io.Zip)
}

// restoreWebsite 恢复网站备份
func (r *backupRepo) restoreWebsite(backup, name string) error {
	if !io.Exists(backup) {
		return errors.New("备份文件不存在")
	}

	website, err := r.website.GetByName(name)
	if err != nil {
		return err
	}
	format, err := io.FormatArchiveByPath(backup)
	if err != nil {
		return err
	}

	if err = io.Remove(website.Path); err != nil {
		return err
	}
	if err = io.UnCompress(backup, website.Path, format); err != nil {
		return err
	}
	if err = io.Chmod(website.Path, 0755); err != nil {
		return err
	}
	if err = io.Chown(website.Path, "www", "www"); err != nil {
		return err
	}

	return nil
}

// restoreMySQL 恢复 MySQL 备份
func (r *backupRepo) restoreMySQL(backup, name string) error {
	if !io.Exists(backup) {
		return errors.New("备份文件不存在")
	}

	rootPassword, err := r.setting.Get(biz.SettingKeyMySQLRootPassword)
	if err != nil {
		return err
	}
	mysql, err := db.NewMySQL("root", rootPassword, "/tmp/mysql.sock")
	if err != nil {
		return err
	}
	if exist, _ := mysql.DatabaseExists(name); !exist {
		return fmt.Errorf("数据库不存在：%s", name)
	}
	if err = os.Setenv("MYSQL_PWD", rootPassword); err != nil {
		return err
	}

	if !strings.HasSuffix(backup, ".sql") {
		backup, err = r.autoUnCompressSQL(backup)
		if err != nil {
			return err
		}
		defer io.Remove(filepath.Dir(backup))
	}

	if _, err = shell.Execf(`mysql -u root '%s' < '%s'`, name, backup); err != nil {
		return err
	}

	return os.Unsetenv("MYSQL_PWD")
}

// restorePostgres 恢复 PostgreSQL 备份
func (r *backupRepo) restorePostgres(backup, name string) error {
	if !io.Exists(backup) {
		return errors.New("备份文件不存在")
	}

	postgres, err := db.NewPostgres("postgres", "", "127.0.0.1", fmt.Sprintf("%s/server/postgresql/data/pg_hba.conf", app.Root), 5432)
	if err != nil {
		return err
	}
	if exist, _ := postgres.DatabaseExist(name); !exist {
		return fmt.Errorf("数据库不存在：%s", name)
	}

	if !strings.HasSuffix(backup, ".sql") {
		backup, err = r.autoUnCompressSQL(backup)
		if err != nil {
			return err
		}
		defer io.Remove(filepath.Dir(backup))
	}

	if _, err = shell.Execf(`su - postgres -c "psql '%s'" < '%s'`, name, backup); err != nil {
		return err
	}

	return nil
}

// preCheckPath 预检空间和 inode 是否足够
// to 备份保存目录
// path 待备份目录
func (r *backupRepo) preCheckPath(to, path string) error {
	size, err := io.SizeX(path)
	if err != nil {
		return err
	}
	files, err := io.CountX(path)
	if err != nil {
		return err
	}

	usage, err := disk.Usage(to)
	if err != nil {
		return err
	}

	color.Greenln(fmt.Sprintf("|-目标大小：%s", str.FormatBytes(float64(size))))
	color.Greenln(fmt.Sprintf("|-目标文件数：%d", files))
	color.Greenln(fmt.Sprintf("|-备份目录可用空间：%s", str.FormatBytes(float64(usage.Free))))
	color.Greenln(fmt.Sprintf("|-备份目录可用Inode：%d", usage.InodesFree))

	if uint64(size) > usage.Free {
		return errors.New("备份目录空间不足")
	}
	if uint64(files) > usage.InodesFree {
		return errors.New("备份目录Inode不足")
	}

	return nil
}

// preCheckDB 预检空间和 inode 是否足够
// to 备份保存目录
// size 数据库大小
func (r *backupRepo) preCheckDB(to string, size int64) error {
	usage, err := disk.Usage(to)
	if err != nil {
		return err
	}

	color.Greenln(fmt.Sprintf("|-目标大小：%s", str.FormatBytes(float64(size))))
	color.Greenln(fmt.Sprintf("|-备份目录可用空间：%s", str.FormatBytes(float64(usage.Free))))
	color.Greenln(fmt.Sprintf("|-备份目录可用Inode：%d", usage.InodesFree))

	if uint64(size) > usage.Free {
		return errors.New("备份目录空间不足")
	}

	return nil
}

// autoUnCompressSQL 自动处理压缩文件
func (r *backupRepo) autoUnCompressSQL(backup string) (string, error) {
	temp, err := io.TempDir(backup)
	if err != nil {
		return "", err
	}

	format, err := io.FormatArchiveByPath(backup)
	if err != nil {
		return "", err
	}
	if err = io.UnCompress(backup, temp, format); err != nil {
		return "", err
	}

	backup = "" // 置空，防止干扰后续判断
	if files, err := os.ReadDir(temp); err == nil {
		if len(files) != 1 {
			return "", fmt.Errorf("压缩文件中包含的文件数量不为1，实际为%d", len(files))
		}
		if strings.HasSuffix(files[0].Name(), ".sql") {
			backup = filepath.Join(temp, files[0].Name())
		}
	}

	if backup == "" {
		return "", errors.New("无法找到.sql备份文件")
	}

	return backup, nil
}