package router

import (
	"github.com/gin-gonic/gin"
	profile "portfolio/pkg/profile/infrastructure/rest"
	tweet "portfolio/pkg/tweet/infrastructure/rest"
)

type Router interface {
	Run(addr ...string) error
}

func NewRouter(routes RoutesGroups) Router {
	router := gin.Default()

	// Registering routes
	base := router.Group("v1")
	routes.Profile.RegisterRoutes(base)
	routes.Tweet.RegisterRoutes(base)

	return router
}

type RoutesGroups struct {
	Profile *profile.ProfileRoutes
	Tweet   *tweet.TweetRoutes
}
