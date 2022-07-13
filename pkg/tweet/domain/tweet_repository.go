package domain

type TweetRepository interface {
	FindMany(max int64) ([]Tweet, error)
}
