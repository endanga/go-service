package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name" validate:"required"`
	Description string     `json:"description"`
	Price       float32    `json:"price" validate:"gt=0"`
	SKU         string     `json:"sku" validate:"required,sku"`
	CreatedOn   NullString `json:"_"`
	UpdatedOn   NullString `json:"_"`
	DeletedOn   NullString `json:"_"`
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^[0-9]{4}$`)
	macthes := re.FindAllString(fl.Field().String(), -1)

	if len(macthes) != 1 {
		return false
	}

	return true
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	getProductList()
	return productList
}

func AddProduct(p *Product) {
	getProductList()
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	// getProductList()
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	getProductList()
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	// getProductList()
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func openConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")

	if err != nil {
		panic(err)
	}
	return db
}

func getProductList() {
	db := openConnection()
	prods := []*Product{}
	result, err := db.Query("SELECT ID, Name, Description, Price, SKU, CreatedOn, UpdatedOn, DeletedOn FROM products")
	if err != nil {
		defer result.Close()
		panic(err)
	}
	for result.Next() {
		// var p *Product
		p := &Product{}
		// if p != nil {
		// 	fmt.Println("TEST ", *p)
		// }
		err = result.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.SKU, &p.CreatedOn, &p.UpdatedOn, &p.DeletedOn)
		if err != nil {
			panic(err.Error())
		}
		prods = append(prods, p)

	}
	productList = prods
	defer db.Close()
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

var productList = []*Product{}

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
