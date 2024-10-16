package models

import (
	"ecommerce/constants"
	"ecommerce/logger"
	"fmt"

	"github.com/google/uuid"
)

type Price float32
type Name string
type UUID uuid.UUID

type Product struct {
	Id            UUID
	Name          Name
	Price         Price
	Availablility constants.AVAILABILITY
	StockCount    uint8
}

func (product *Product) ChangeAvailability(availablity constants.AVAILABILITY) {
	logger.Log(fmt.Sprintf("ChangeAvailability: availablity = %v", availablity))

	product.Availablility = availablity

	if availablity == constants.OUT_OF_STOCK || availablity == constants.DISCONTINUED {
		product.StockCount = 0
	}
}

func (product *Product) AlterStockCount(operation constants.AlterOperations, quantity uint8) {
	logger.Log(fmt.Sprintf("AlterStockCount: operation = %v | quanity = %v", operation, quantity))

	if operation == constants.ADDITION {
		product.StockCount += quantity

		if quantity > uint8(9) {
			logger.Log("AlterStockCount: ChangeAvailability is called with constants.AVAILABLE")

			product.ChangeAvailability(constants.AVAILABLE)
		}
	} else {
		if product.StockCount-quantity == uint8(0) {
			logger.Log("AlterStockCount: ChangeAvailability is called with constants.OUT_OF_STOCK")

			product.StockCount = 0
			product.ChangeAvailability(constants.OUT_OF_STOCK)
		} else if int(product.StockCount-quantity) < 0 {
			logger.Log("AlterStockCount: Operation cannot be performed, stock count cannot be negative")
		} else {
			product.StockCount -= quantity
		}
	}
}

func (product Product) GetProductId() UUID {
	return product.Id
}
