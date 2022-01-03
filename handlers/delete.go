package handlers

import (
	"example/test/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Hanlde DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Something want wrong!", http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(rw, "Something want wrong!", http.StatusInternalServerError)
		return
	}
}
