package routes

import (
	"Controllers"
	"flame"
)

func RegisterWebRoutes(router *flame.Router) {
	router.Get("/", Controllers.MainController.Index)
}
