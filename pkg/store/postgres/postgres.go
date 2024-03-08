package postgres

import (
	"architecture_go/services/article/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Storage struct {
	db *sql.DB
}

var lock = &sync.Mutex{}

var singleInstance *Storage

func Connect(cfg *configs.Config) (*Storage, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			log.Println("Creating single instance now.")
			connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.DatabasePort, cfg.User, cfg.Password, cfg.DBName)
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				return nil, err
			}

			err = db.Ping()
			if err != nil {
				return nil, err
			}

			return &Storage{db: db}, nil
		} else {
			log.Println("Single instance already created.")
		}
	} else {
		log.Println("Single instance already created.")
	}

	return singleInstance, nil
}

func (s *Storage) GetDB() *sql.DB {
	return s.db
}
