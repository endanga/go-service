package data

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCheckForValidation(t *testing.T) {
	// p := &Product{
	// 	Name:  "Test",
	// 	Price: 23000,
	// 	SKU:   "1234",
	// }
	// err := p.Validate()
	// if err != nil {
	// 	t.Fatal(err)
	// }

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err)
	}
	defer db.Close()

}
