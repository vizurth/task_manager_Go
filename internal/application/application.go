package application

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vizurth/task_manager_Go/pkg/task"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
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
			err := task.getList(w,r)
			if err != nil{
				http.Error(w, "Something wrong method", 500)
			}
	}
}

func (a *Application) RunServer() error{
	http.HandleFunc("/api/task", TaskHandle)

	return http.ListenAndServe(":" + a.config.Addr, nil)
}