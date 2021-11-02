package main

import "livecrudl/web/app"

func main() {
	app := app.NewApp()
	serverClose := make(chan bool)
	go app.Start("8000", serverClose)
	println("listening on 8000")
	<-serverClose
}
