package Controllers

import "flame"

type MainController struct {
	Request *flame.Request
}

func (c MainController) Index(res *flame.Response) *flame.Response {
	res.SendView("welcome")
	return res
}
