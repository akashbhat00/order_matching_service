// repositories/seller_repository.go
package repositories

import (
	"order_matching_service/dto"
	"order_matching_service/models"

	"gorm.io/gorm"
)

type SellerRepository struct {
	DB *gorm.DB
}

func NewSellerRepository(db *gorm.DB) *SellerRepository {
	return &SellerRepository{DB: db}
}

// CreateSeller creates a new seller in the database.
func (repo *SellerRepository) CreateSeller(seller *models.Sellers) error {
	return repo.DB.Create(seller).Error
}

func (repo *SellerRepository) GetSellerByID(sellerID uint) (*dto.SellerWithProducts, error) {
	var sellerWithProducts dto.SellerWithProducts
	db := repo.DB.Model(&models.Sellers{}).
		Preload("Products").
		Where("id = ?", sellerID)

		// Retrieve the seller along with associated products
	if err := db.First(&sellerWithProducts).Error; err != nil {
		return nil, err
	}

	// Define a custom SQL query that joins sellers and products tables
	// query := `
	//     SELECT
	//         s.id AS seller_id,
	//         s.name AS seller_name,
	//         s.email AS seller_email,
	//         p.id AS product_id,
	//         p.name AS product_name,
	//         p.price AS product_price,
	//         p.quantity AS product_quantity
	//     FROM sellers s
	//     LEFT JOIN products p ON s.id = p.seller_id
	//     WHERE s.id = ?`

	// // Execute the query and scan the result into the DTO structure
	// if err := repo.DB.Raw(query, sellerID).Scan(&sellerWithProducts).Error; err != nil {
	//     return nil, err
	// }
	// fields := []string{
	// 	"s.id AS seller_id",
	// 	"s.name AS seller_name",
	// 	"s.email AS seller_email",
	// 	"p.id AS product_id",
	// 	"p.name AS product_name",
	// 	"p.price AS product_price",
	// 	"p.quantity AS product_quantity",
	// }

	// // Build the query using GORM's methods
	// db := repo.DB.Table("sellers s").
	// 	Select(fields).
	// 	Joins("LEFT JOIN products p ON s.id = p.seller_id").
	// 	Where("s.id = ?", sellerID)

	// // Scan the result into the DTO structure
	// if err := db.Scan(&sellerWithProducts).Error; err != nil {
	// 	return nil, err

	return &sellerWithProducts, nil
}

func (r *SellerRepository) AddProductToCatalog(sellerID uint, productDTO *dto.ProductRequestDto) (*models.Products, error) {
	// Perform the database operation to add the product to the seller's catalog
	product := &models.Products{
		Name:     productDTO.Name,
		Price:    productDTO.Price,
		Quantity: productDTO.Quantity,
		SellerID: sellerID,
	}
	if err := r.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
