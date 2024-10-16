package models

import (
	"ecommerce/constants"
	"ecommerce/logger"
	"fmt"
)

type ScreenSize float32
type BatteryCapacity uint32
type Model string

type Mobile struct {
	Product
	ScreenSize      ScreenSize
	BatteryCapacity BatteryCapacity
	ModelName       Model
	Exclusive       bool
}

func (mobile *Mobile) ChangeExclusivity(isExeclusive bool) {
	logger.Log(fmt.Sprintf("Changing exclusivity type to %t", isExeclusive))

	mobile.Exclusive = isExeclusive
}

func NewMobile(
	id UUID,
	name Name,
	price Price,
	availability constants.AVAILABILITY,
	stockCount uint8,
	screenSize ScreenSize,
	batteryCapacity BatteryCapacity,
	modelName Model,
	exclusive bool,
) Mobile {
	m := Mobile{}

	m.Name = name
	m.Price = price
	m.Availablility = availability
	m.StockCount = stockCount
	m.ScreenSize = screenSize
	m.BatteryCapacity = batteryCapacity
	m.ModelName = modelName
	m.Exclusive = exclusive

	return m
}

func (mobile Mobile) GetProductId() UUID {
	return mobile.Id
}
