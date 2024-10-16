package main

import (
	"ecommerce/constants"
	"ecommerce/interfaces"
	"ecommerce/logger"
	"ecommerce/models"
	"fmt"

	"github.com/google/uuid"
)

func createProduct() []interfaces.IProduct {
	iPhone13 := models.NewMobile(
		models.UUID(uuid.New()),
		"iPhone-13",
		models.Price(78_000),
		constants.AVAILABLE,
		10,
		6.5,
		4000,
		"iPhone 13",
		false,
	)

	macbookAir := models.NewLaptop(
		models.UUID(uuid.New()),
		"MacBook Air",
		models.Price(1_20_000),
		constants.UPCOMING,
		0,
		13,
		72_000,
		"Apple MacBook Air 13",
		true,
		constants.MAC_OS,
	)

	var products []interfaces.IProduct
	products = append(products, iPhone13)
	products = append(products, macbookAir)

	return products
}

func displayList(products []interfaces.IProduct) {
	for _, product := range products {
		logger.Log(fmt.Sprintf("%v", product))
	}
}

func Start() {
	pList := createProduct()
	displayList(pList)
}
