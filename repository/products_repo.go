package repository

import (
	"github.com/endanga/product-service/data"
	"github.com/endanga/product-service/database"
)

// func GetProducts() data.Products {

// 	listOfProds := getProductList()

// 	return listOfProds
// }

// getProductList returns all products from the database feed to productList
func GetProductList(limit int, offset int) data.Products {
	// if checkConnection(db) {
	// 	openConnection()
	// }

	prods := []*data.Product{}
	stmt := `SELECT id, product_code, product_name, sub_category, brand, retail_price, status FROM master.products limit $1 offset $2`
	result, err := database.DBCon.Query(stmt, limit, offset)
	if err != nil {
		// defer result.Close()
		panic(err)
	}
	// defer db.Close()
	for result.Next() {
		p := &data.Product{}
		err = result.Scan(&p.ID, &p.Name, &p.Code, &p.Sub_Category, &p.Brand, &p.Retail_Price, &p.Status)
		if err != nil {
			panic(err.Error())
		}
		prods = append(prods, p)

	}
	return prods
}

func GetMaxRowCount() int {
	result, err := database.DBCon.Query("SELECT reltuples::bigint AS estimat FROM pg_class WHERE oid = 'master.products'::regclass;")
	if err != nil {
		// defer result.Close()
		panic(err)
	}

	var count int
	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			panic(err.Error())
		}
	}
	return count
}
