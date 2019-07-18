package routing

import (
	"fmt"
	"git.sdkeji.top/share/sdlib/log"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"net/http"
	"sdkeji/person/pkg/app"
	open "sdkeji/person/pkg/routing/app"
)

func Register(logger log.Logger) http.Handler {
	router := routing.New()
	router.NotFound(notFound)
	router.Use(
		routingLogger(logger),
		errorHandler(logger),
		content.TypeNegotiator(content.JSON),
	)

	api := router.Group("/" + app.System)

	{
		versionHandler := NewVersionHandler()
		api.Get("/version", versionHandler.Version)
		open.Register(api.Group("/app/v1"))

	}

	for _, route := range router.Routes() {
		logger.Debug(fmt.Sprintf("register route: \"%-6s -> %s\".", route.Method(), route.Path()))
	}

	return router
}
