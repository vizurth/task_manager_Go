package main

import "github.com/vizurth/task_manager_Go/internal/application"

func main(){
	app := application.New()
	app.RunServer()
}