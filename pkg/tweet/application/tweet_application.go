package application

import "portfolio/pkg/tweet/domain"

type TweetApplication struct {
	tweetRepository domain.TweetRepository
}

func NewTweetApplication(tweetRepository domain.TweetRepository) *TweetApplication {
	return &TweetApplication{tweetRepository}
}

func (t *TweetApplication) FindTweets(max int64) ([]domain.Tweet, error) {
	return t.tweetRepository.FindMany(max)
}
