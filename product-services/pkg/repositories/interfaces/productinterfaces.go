package interfaces

type ProductRepositorie interface {

	// Product
	AddProduct()
	UpdateProduct()
	DeleteProduct()

	// Filtering
	GetFullProducts()
	GetProductById()
	GetProductByName()
	GetProductByCategory()
	GetProductByBrand()
}
