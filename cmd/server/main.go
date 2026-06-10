package main

import (
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/config"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db"
)

func main() {
	cfg := config.Load()

	pool := db.InitDB(cfg.DBurl)
	defer pool.Close()

}
