package controllers

import (
	"net/http"
	"strconv"

	"github.com/gieart87/gotoko/app/models"

	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	renderEngine := render.New(render.Options{
		Layout: "layout",
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLinks(server.AppConfig, PaginationParams{
		Path:        "products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	_ = renderEngine.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products":   products,
		"pagination": pagination,
	})
}
