package rest

import "github.com/gin-gonic/gin"

type TweetRoutes struct {
	handler *TweetHandler
}

func NewTweetRoutes(handler *TweetHandler) *TweetRoutes {
	return &TweetRoutes{handler}
}

func (t *TweetRoutes) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("tweet", t.handler.GetTweet)
}
