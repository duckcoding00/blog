package api

import (
	"github.com/duckcoding00/blog/internal/handler"
)

func InitServer() {
	handler := handler.NewHandler()

	config := Application{
		config: AppConfig{
			handler: handler,
			addr:    ":8080",
		},
	}

	app := NewApp(config.config)
	app.RegisterRouter()
	app.Run()
}
