package main

import "github.com/Mendes11/go-blueprints/simple/internal/app"

func main() {
	conf := app.LoadConfig()
	err := app.StartAPI(conf)
	if err != nil {
		panic(err)
	}
}
