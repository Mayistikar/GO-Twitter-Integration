package rest

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	r "portfolio/cmd/rest_api/router"
	profileApplication "portfolio/pkg/profile/application"
	profile "portfolio/pkg/profile/infrastructure/rest"
	storage "portfolio/pkg/profile/infrastructure/storage/postgres"
	tweetApplication "portfolio/pkg/tweet/application"
	"portfolio/pkg/tweet/infrastructure/integrations/twitter"
	tweet "portfolio/pkg/tweet/infrastructure/rest"
	"portfolio/shared/config"
	"portfolio/shared/rest"
)

type App struct {
	Config *config.Config
	Router r.Router
}

func New(config *config.Config) (*App, error) {

	dsn := "host=localhost user=root password=root dbname=portfolio port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return &App{}, err
	}

	// Profile
	profileRepository := storage.NewProfileRepository(db)
	profileApplication := profileApplication.NewProfileApplication(profileRepository)
	profileHandler := profile.NewProfileHandler(profileApplication)
	profileRoutes := profile.NewProfileRoutes(profileHandler)

	// Tweet
	twitterRestClient := rest.NewClient(config)
	twitterRepository := twitter.NewRepository(twitterRestClient)
	tweetApplication := tweetApplication.NewTweetApplication(twitterRepository)
	tweetHandler := tweet.NewTweetHandler(tweetApplication)
	tweetRoutes := tweet.NewTweetRoutes(tweetHandler)

	routes := r.RoutesGroups{
		Profile: profileRoutes,
		Tweet:   tweetRoutes,
	}
	router := r.NewRouter(routes)

	return &App{Config: config, Router: router}, nil
}

func (a *App) Run() error {
	return a.Router.Run(fmt.Sprintf(":%s", a.Config.Application.Port))
}
