package main

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/article/configs"
	configs2 "architecture_go/services/user/configs"
	"architecture_go/services/user/internal/delivery/http"
	"architecture_go/services/user/internal/repository/storage/postgres/user"
	user2 "architecture_go/services/user/internal/useCase/user"
	"github.com/alexedwards/scs/v2"
	"log"
	"time"
)

func main() {
	cfg, err := configs2.NewUserConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	storage, err := postgres.Connect((*configs.Config)(cfg))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer storage.GetDB().Close()

	log.Println("Successfully connected to database!")

	userRepo := user.NewUserRepository(storage.GetDB())
	userUseCase := user2.NewUserUseCase(userRepo)

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true

	sessionManager.Cookie.Secure = true

	delivery := http.NewUserHTTP(userUseCase, storage.GetDB(), sessionManager)
	delivery.Run((*configs.Config)(cfg))
}
