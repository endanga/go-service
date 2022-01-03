package repository

import (
	"example/test/data"
	"example/test/database"
)

// func GetProducts() data.Products {

// 	listOfProds := getProductList()

// 	return listOfProds
// }

// getProductList returns all products from the database feed to productList
func GetProductList() data.Products {
	// if checkConnection(db) {
	// 	openConnection()
	// }

	prods := []*data.Product{}
	result, err := database.DBCon.Query("SELECT ID, Name, Description, Price, SKU FROM products")
	if err != nil {
		// defer result.Close()
		panic(err)
	}
	// defer db.Close()
	for result.Next() {
		p := &data.Product{}
		err = result.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.SKU)
		if err != nil {
			panic(err.Error())
		}
		prods = append(prods, p)

	}
	return prods
}
