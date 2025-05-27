package main

import (
	"github.com/google/wire"
	"thuanle/cse-mark/internal/configs"
	http2 "thuanle/cse-mark/internal/delivery/http"
	"thuanle/cse-mark/internal/delivery/tele/views"
	"thuanle/cse-mark/internal/domain/course"
	"thuanle/cse-mark/internal/infra/http"
	"thuanle/cse-mark/internal/infra/mongo"
	"thuanle/cse-mark/internal/usecases/iam"
	"thuanle/cse-mark/internal/usecases/markimport"
)

type App struct {
	Config      *configs.Config
	MongoClient *mongo.Client
	HttpService http2.Service
}

func InitializeApp() (*App, error) {
	wire.Build(
		//configurations
		configs.LoadConfig,

		//infrastructures
		mongo.NewClient,
		mongo.NewCourseRepo,
		mongo.NewMarkRepo,
		mongo.NewUserRepo,
		http.NewSimpleDownloader,

		//domain repositories and rules
		course.NewRules,

		//usecases
		markimport.NewService,
		iam.NewAuthzService,

		//delivery-view
		views.NewTeacherRenderer,

		//delivery
		http2.NewHttpService,
		//app
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
