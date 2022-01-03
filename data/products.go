package data

import (
	"database/sql"
	"encoding/json"
	"example/test/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: ^[0-9]{4}$
	SKU string `json:"sku" validate:"sku"`
}

type Products []*Product

var ProductList = []*Product{}

// GetProducts returns all products from the database

// AddProduct adds a new product to the database
func AddProduct(p *Product) {
	// get the next id in sequence
	p.ID = getLastID()
	fmt.Println("last ID is ", p.ID)
	err := postProduct(p)
	if err != nil {
		fmt.Println("RESULT ", err)
	}
}

func UpdateProduct(id int, p *Product) error {

	result := findProductById(id)
	fmt.Println("find_id: ", result)
	if (Product{}) == result {
		fmt.Println("Desc ", p.Description)
	} else {
		if (p.Description == "") || (len(p.Description) == 0) {
			p.Description = result.Description
		}
		fmt.Println("this is price ", p.Price)
		p.ID = id
		err := updateProductData(p)
		if err != nil {
			fmt.Println("something want wrong", err)
			return err
		}
	}
	return nil
}

func DeleteProduct(id int) error {
	err := deleteProductById(id)
	if err != nil {
		fmt.Println("something want wrong", err)
		return err
	}
	return nil
}

var ErrProductNotFound = fmt.Errorf("id is not found")

// func findProduct(id int) (*Product, int, error) {
// 	// getProductList()
// 	for i, p := range productList {
// 		if p.ID == id {
// 			return p, i, nil
// 		}
// 	}
// 	return nil, -1, ErrProductNotFound
// }

var db *sql.DB

func openConnection() {
	dbCon, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		panic(err)
	}

	db = dbCon
	// return db
}

func checkConnection(dbCheck *sql.DB) bool {
	fmt.Println("DB CHECK ", dbCheck)
	if dbCheck != nil {
		err := dbCheck.Ping()
		if err != fmt.Errorf("sql: databse is closed") {
			fmt.Println("ping.....", dbCheck.Ping())
			return false
		}
	}

	return true
}

func getLastID() int {
	// if checkConnection(db) {
	// 	openConnection()
	// }

	curId := 0
	result, err := database.DBCon.Query("SELECT MAX(ID) as id FROM products")
	if err != nil {
		fmt.Println("Cannot insert data: ", err)
	}

	for result.Next() {
		err = result.Scan(&curId)
		if err != nil {
			fmt.Println("Cannot insert data: ", err)
		}
	}
	// defer db.Close()

	return curId + 1
}

func postProduct(p *Product) error {
	// if checkConnection(db) {
	// 	openConnection()
	// }

	sql := `INSERT INTO products (ID, Name, Description, Price, SKU) VALUES (?, ?, ?, ?, ?)`
	err := database.DBCon.QueryRow(sql, p.ID, p.Name, p.Description, p.Price, p.SKU)
	if err.Err() != nil {
		return fmt.Errorf("cannot insert data: %v", err.Err())
	}
	// defer db.Close()
	return nil
}

func findProductById(id int) Product {

	// if checkConnection(db) {
	// 	openConnection()
	// }

	fmt.Println("jajajajjajj")

	var p Product
	fmt.Println("ID:", id)
	result, err := database.DBCon.Query("SELECT ID, Name, Description, Price, SKU FROM products where ID = ? ", id)
	if err != nil {
		// defer db.Close()
		fmt.Println("error when execute query:", err)
		return Product{}
	}
	for result.Next() {
		err = result.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.SKU)
		if err != nil {
			// defer db.Close()
			fmt.Println("error when mapping value:", err)
			return Product{}
		}
	}

	fmt.Println("Desc:", p.Description)
	// defer db.Close()
	return p
}

func updateProductData(p *Product) error {

	fmt.Println("is not nilll")

	sql := `UPDATE products SET Name=?, Description=?, Price=?, SKU=? WHERE ID=?`
	result, err := database.DBCon.Exec(sql, p.Name, p.Description, p.Price, p.SKU, p.ID)
	if err != nil {
		return fmt.Errorf("error when execute query: %v", err)
	}

	count, _ := result.RowsAffected()
	if count <= 0 {
		// defer db.Close()
		return fmt.Errorf("id is not found")
	}

	// defer db.Close()
	return nil
}

func deleteProductById(id int) error {

	sql := `DELETE FROM products where ID = ?`

	result, err := database.DBCon.Exec(sql, id)
	if err != nil {
		// defer db.Close()
		return fmt.Errorf("error when execute query: %v", err)
	}

	count, _ := result.RowsAffected()
	if count <= 0 {
		// defer db.Close()
		return fmt.Errorf("id is not found")
	}
	// defer db.Close()
	return nil
}

type NullString struct {
	sql.NullString
}

func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte(""), nil
	}
	return json.Marshal(ns.String)
}

// var productList = []*Product{
// 	&Product{
// 		ID:          1,
// 		Name:        "Latte",
// 		Description: "Light Coffee with milk",
// 		Price:       20000,
// 		SKU:         "1001",
// 		CreatedOn:   time.Now().Format(time.RFC822),
// 		UpdatedOn:   time.Now().Format(time.RFC822),
// 	},
// 	&Product{
// 		ID:          2,
// 		Name:        "Espresso",
// 		Description: "Strong Coffee without milk",
// 		Price:       15000,
// 		SKU:         "1002",
// 		CreatedOn:   time.Now().Format(time.RFC822),
// 		UpdatedOn:   time.Now().Format(time.RFC822),
// 	},
// }
