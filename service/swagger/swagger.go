package swagger

import (
	"path"

	"github.com/labstack/echo"
)

// MountSwaggerAPI mounts swagger API
func MountSwaggerAPI(g *echo.Group) {
	h := func(indexFile string) echo.HandlerFunc {
		staticRoot := "public"
		return func(ctx echo.Context) error {
			f := path.Join(staticRoot, ctx.Request().RequestURI)
			ext := path.Ext(f)
			if len(ext) == 0 {
				f = path.Join(f, indexFile)
			}
			ctx.Response().Header().Set("cache-control", "no-cache")

			return ctx.File(f)
		}
	}

	g.GET("/swagger-resource*", h("index.json"))
	g.GET("/swagger-ui*", h("index.html"))
}
