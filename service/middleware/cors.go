package middleware

import (
	"github.com/labstack/echo"
	echomw "github.com/labstack/echo/middleware"
)

// RequestCORS configures the cross domain
func RequestCORS() echo.MiddlewareFunc {
	// "http://localhost:3000"
	return echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "No-Cache", "X-Requested-With", "If-Modified-Since",
			"Pragma", "Last-Modified", "Cache-Control", "Expires", "Content-Type", "X-E4M-With", "Authorization"},
		AllowCredentials: true,
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	})
}
