package data

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the code for this poduct
	//
	// required: true
	// max length: 255
	Code string `json:"code" validate:"required"`

	// the description for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the sub_category for the product
	//
	// required: true
	// min: 0.01
	Sub_Category string `json:"sub_category" validate:"required"`

	// the brand for the product
	//
	// required: true
	// max length: 255
	Brand string `json:"brand" validate:"required"`

	// the retail_price for the product
	//
	// required: true
	// min: 1000
	// max: 99999999
	Retail_Price float64 `json:"retail_price" validate:"required"`

	// the status for the product
	//
	// required: true
	// pattern: y/
	Status string `json:"status" validate:"required"`

	Max_Count int `json:"max_count"`
}

type Products []*Product

var ProductList = []*Product{}
