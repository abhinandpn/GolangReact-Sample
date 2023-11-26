package interfaces

type ProductUseCase interface {

	// Product
	CreateProduct()
	UpdateProduct()
	DeletProduct()

	// Listing
	ViewFullProducts()
	ViewProductById()
	ViewProductByName()
	ViewProductsByCategory()
	ViewProductsBySUbCategory()
	ViewProductsByBrands()

}
