package main

import (
	"github.com/hiram66/user-service/cmd/web"
	"github.com/hiram66/user-service/configs"
)

func main() {
	web.Start(configs.AppPort)
}
