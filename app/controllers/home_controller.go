package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	renderEngine := render.New(render.Options{
		Layout: "layout",
	})

	_ = renderEngine.HTML(w, http.StatusOK, "home", map[string]interface{}{
		"title": "Home Title",
		"body":  "Home Description",
	})
}
