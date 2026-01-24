package main

import "cashier-api/internal/app"

func main() {
	app := app.NewApp()
	app.Run()
}
