package twitter

import "errors"

var (
	ErrorTweetDTOInvalidID   = errors.New("invalid tweet id")
	ErrorTweetDTOInvalidText = errors.New("invalid text id")
)
