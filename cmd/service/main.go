package main

import (
	"fmt"
	"main/internal/config"
	"main/internal/storage/postgres"
	"os"
)

func main() {
	cfg := config.MustLoad()
	
	storage, err := postgres.New(cfg.StoragePath)
	if err != nil{
		fmt.Println("Database error")
		os.Exit(1)
	} else {
		fmt.Println("Database connected")
	}
	_=storage

	//router

	//run server
}