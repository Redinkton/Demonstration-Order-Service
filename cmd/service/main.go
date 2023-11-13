package main

import (
	"fmt"
	"main/internal/config"
	"main/internal/storage/postgres"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
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


	nc, err := nats.Connect("nats://localhost:4222")
    if err != nil {
        fmt.Println(err.Error())
    }

    defer nc.Close()
	sc, err := stan.Connect("test-cluster", "client-1", stan.NatsConn(nc))


	sub, _ := sc.Subscribe("my-channel", func(msg *stan.Msg) {
		//msg
	})
    defer sub.Unsubscribe()


	//router

	//run server
}