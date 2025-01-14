package middleware

import (
	"net/http"

	"github.com/go-rat/chix"

	"github.com/tnb-labs/panel/internal/app"
)

// Status 检查程序状态
func Status(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch app.Status {
		case app.StatusUpgrade:
			render := chix.NewRender(w)
			defer render.Release()
			render.Status(http.StatusServiceUnavailable)
			render.JSON(chix.M{
				"message": "面板更新中，请稍后刷新",
			})
			return
		case app.StatusMaintain:
			render := chix.NewRender(w)
			defer render.Release()
			render.Status(http.StatusServiceUnavailable)
			render.JSON(chix.M{
				"message": "面板正在运行维护任务，请稍后刷新",
			})
			return
		case app.StatusClosed:
			render := chix.NewRender(w)
			defer render.Release()
			render.Status(http.StatusForbidden)
			render.JSON(chix.M{
				"message": "面板已关闭",
			})
			return
		case app.StatusFailed:
			render := chix.NewRender(w)
			defer render.Release()
			render.Status(http.StatusInternalServerError)
			render.JSON(chix.M{
				"message": "面板运行出错，请检查排除或联系支持",
			})
			return
		default:
			next.ServeHTTP(w, r)
		}
	})
}
