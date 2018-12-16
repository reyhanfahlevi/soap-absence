package main

import (
	"context"
	"log"

	"github.com/reyhanfahlevi/soap-absence/config"
	"github.com/reyhanfahlevi/soap-absence/container"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	conf := config.Get()
	if len(conf.App.DeviceIPList) == 0 {
		log.Fatal("Empty Device")
	}

	svc, err := container.InitializeService(conf.App.DeviceIPList[0])
	if err != nil {
		log.Fatal(err)
	}

	result, err := svc.GetAllUserInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range result.Users {
		//err = svc.SaveUserInfo(context.Background(), u)
		//if err != nil {
		//	log.Fatal(err)
		//}
		log.Println(u.Name)
	}
}
