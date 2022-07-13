package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	tweet "portfolio/pkg/tweet/application"
	"strconv"
)

type TweetHandler struct {
	tweetApplication *tweet.TweetApplication
}

func NewTweetHandler(tweetApplication *tweet.TweetApplication) *TweetHandler {
	return &TweetHandler{tweetApplication}
}

func (h *TweetHandler) GetTweet(ctx *gin.Context) {
	maxQuery := ctx.Query("max_results")
	max, err := strconv.ParseInt(maxQuery, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%s: %v", "max_results query params has to be a number", err.Error()),
		})
		return
	}

	tweets, err := h.tweetApplication.FindTweets(max)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSONP(http.StatusOK, tweets)
}
