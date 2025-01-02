package application

import (
	"net/http"
	"os"
	"fmt"
	"github.com/vizurth/task_manager_Go/pkg/task"
)



type Config struct{
	Addr string
}

func ConfigFromEnv() *Config{
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == ""{
		config.Addr = "8080"
	}
	return config
}

type Application struct{
	config *Config
}

func New() *Application{
	return &Application{
		config: ConfigFromEnv(),
	}
}

func TaskHandle(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodGet:
			err := task.GetTask(w,r)
			if err != nil{
				http.Error(w, "Something wrong method", 500)
			}
	}
}

func (a *Application) RunServer() error{
	http.HandleFunc("/api/task", TaskHandle)
	fmt.Println("Server is running on port 8080...")
	return http.ListenAndServe(":" + a.config.Addr, nil)
}