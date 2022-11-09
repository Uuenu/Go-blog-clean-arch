package app

import (
	"context"
	"go-blog-ca/config"
	v1 "go-blog-ca/internal/controller/http/v1"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/internal/usecases/repo"
	"go-blog-ca/pkg/client/mongodb"
	logger "go-blog-ca/pkg/logging"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	// Logger

	l := *logger.GetLogger()

	// MongoDB

	mongodbClient, err := mongodb.NewClient(context.TODO(), cfg.Mongodb.Host, cfg.Mongodb.Port, cfg.Mongodb.Username, cfg.Mongodb.Password,
		cfg.Mongodb.Database, cfg.Mongodb.AuthDb)
	if err != nil {
		panic("implement me with logger")
	}
	mdb := mongodbClient.Database(cfg.Mongodb.Database)

	//Use case
	sessionRepo := repo.NewSessionRepo(mdb)
	authorRepo := repo.NewAuthorRepo(mdb)
	articleRepo := repo.NewArticleRepo(mdb)

	sessionUseCase := usecases.NewSessionUseCase(sessionRepo)
	authorUseCase := usecases.NewAuthorUseCase(authorRepo, sessionUseCase)
	authUseCase := usecases.NewAuthUseCase(authorUseCase, sessionUseCase)
	articleUseCase := usecases.NewArticleUseCase(articleRepo, authorUseCase, sessionUseCase)

	//HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, authorUseCase, articleUseCase, authUseCase, sessionUseCase, cfg)

	handler.Run()

}
