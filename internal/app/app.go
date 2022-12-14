package app

import (
	"context"
	"fmt"
	"go-blog-ca/config"
	v1 "go-blog-ca/internal/controller/http/v1"
	"go-blog-ca/internal/usecases"
	"go-blog-ca/internal/usecases/repo"
	"go-blog-ca/pkg/client/mongodb"
	logger "go-blog-ca/pkg/logging"
	"go-blog-ca/pkg/postgres"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {

	// Logger

	l := *logger.GetLogger()

	// MongoDB

	mongodbClient, err := mongodb.NewClient(context.TODO(), cfg.Mongodb.Host, cfg.Mongodb.Port, cfg.Mongodb.Username, cfg.Mongodb.Password,
		cfg.Mongodb.Database, cfg.Mongodb.AuthDb)
	if err != nil {
		l.Errorf("mongodbClient error:%v", err)
	}
	mdb := mongodbClient.Database(cfg.Mongodb.Database)

	//Postgresql
	p := cfg.Postgres
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", p.Username, p.Password, p.Host, p.Port, p.Database)
	postgres, err := postgres.New(dsn)
	if err != nil {
		l.Errorf("app - Run - postgres.New. error: %v", err)
	}

	// Repo based on Mongodb
	sessionRepo := repo.NewSessionRepo(mdb, l)
	//authorRepo := repo.NewAuthorRepo(mdb)
	//articleRepo := repo.NewArticleRepo(mdb)

	// Repo based on Postgres
	authorRepo := repo.New(postgres)
	articleRepo := repo.NewArticlePostgresRepo(postgres)

	//Use case
	sessionUseCase := usecases.NewSessionUseCase(sessionRepo, cfg, l)
	authorUseCase := usecases.NewAuthorUseCase(authorRepo, sessionUseCase)
	authUseCase := usecases.NewAuthUseCase(authorUseCase, sessionUseCase, l)
	articleUseCase := usecases.NewArticleUseCase(articleRepo, authorUseCase, sessionUseCase)

	//HTTP Server
	handler := gin.Default()
	v1.NewRouter(handler, l, authorUseCase, articleUseCase, authUseCase, sessionUseCase, cfg)

	handler.Run(":3000")

}
