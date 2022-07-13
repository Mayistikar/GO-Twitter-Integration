package rest

import "github.com/gin-gonic/gin"

type ProfileRoutes struct {
	handler *ProfileHandler
}

func NewProfileRoutes(handler *ProfileHandler) *ProfileRoutes {
	return &ProfileRoutes{handler}
}

func (p *ProfileRoutes) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("profile/:uuid", p.handler.GetProfile)
	group.POST("profile", p.handler.CreateProfile)
	group.PATCH("profile/:uuid", p.handler.UpdateProfile)
}
