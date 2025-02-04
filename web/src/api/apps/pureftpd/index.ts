import { request } from '@/utils'

export default {
  // 列表
  list: (page: number, limit: number): any =>
    request.get('/apps/pureftpd/users', { params: { page, limit } }),
  // 添加
  add: (username: string, password: string, path: string): any =>
    request.post('/apps/pureftpd/users', { username, password, path }),
  // 删除
  delete: (username: string): any => request.delete(`/apps/pureftpd/users/${username}`),
  // 修改密码
  changePassword: (username: string, password: string): any =>
    request.post(`/apps/pureftpd/users/${username}/password`, { password }),
  // 获取端口
  port: (): any => request.get('/apps/pureftpd/port'),
  // 修改端口
  setPort: (port: number): any => request.post('/apps/pureftpd/port', { port })
}
