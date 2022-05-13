package middlewares

import (
	"blog/pkg/auth"
	"blog/pkg/flash"

	"net/http"
)

// Auth 登录用户才可访问
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next(w, r)
	}
}
