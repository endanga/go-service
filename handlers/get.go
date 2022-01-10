package handlers

import (
	"net/http"
	"strconv"

	"github.com/endanga/product-service/data"
	"github.com/endanga/product-service/repository"
)

// swagger:route GET /products product_repo listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	rw.Header().Set("Content-Type", "application/json")
	var offset, limit int
	{
		y := r.URL.Query().Get("limit")
		if y != "" {
			x, err := strconv.Atoi(y)
			if err != nil {
				http.Error(rw, "Something want worng!", http.StatusBadRequest)
				return
			}
			limit = x
		}
		y = r.URL.Query().Get("offset")
		if y != "" {
			x, err := strconv.Atoi(y)
			if err != nil {
				http.Error(rw, "Something want worng!", http.StatusBadRequest)
				return
			}
			offset = x
		}
	}

	println("thisis imit", limit)
	println("thisis offset", offset)
	prods := repository.GetProductList(limit, offset)

	err := data.ToJSON(prods, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}

func (p *Products) GetMaxRowCount(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	rw.Header().Set("Content-Type", "application/json")
	count := repository.GetMaxRowCount()
	mapDt := data.Product{
		Max_Count: count,
	}
	err := data.ToJSON(mapDt, rw)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] something want wrong!", err)
	}
}
