package interfaces

import "ecommerce/models"

type IProduct interface {
	GetProductId() models.UUID
}
