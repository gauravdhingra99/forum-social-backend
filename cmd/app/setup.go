package app

import (
	config "socialForumBackend/internal/config"
	db "socialForumBackend/internal/database"
)

func Init() {
	config.Load()
}

func ShutDown() {
	err := db.Close()
	if err != nil {
		return
	}

}
