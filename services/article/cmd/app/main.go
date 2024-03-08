package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/article/configs"
	"architecture_go/services/article/internal/delivery/http"
	postgres2 "architecture_go/services/article/internal/repository/storage/postgres/article"
	contact2 "architecture_go/services/article/internal/repository/storage/postgres/comment"
	"architecture_go/services/article/internal/useCase/article"
	"architecture_go/services/article/internal/useCase/group"
	"log"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	storage, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to database!")

	articleRepo := postgres2.NewArticleRepository(storage.GetDB())
	commentRepo := contact2.NewCommentRepository(storage.GetDB())
	articleUseCase := article.NewArticleUseCase(articleRepo)
	commentUseCase := group.NewCommentUseCase(commentRepo)
	delivery := http.NewArticleHTTP(articleUseCase, commentUseCase)
	delivery.Run(cfg)
}
