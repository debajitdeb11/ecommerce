package models

import (
	"ecommerce/constants"
	"ecommerce/logger"
	"fmt"
)

type Laptop struct {
	Mobile
	OS constants.OS
}

func (laptop *Laptop) IsUpgradeAble() bool {
	logger.Log(fmt.Sprintf("IsUpgradeable is called & OS type is %s", laptop.OS))
	if laptop.OS == constants.WINDOWS || laptop.OS == constants.UNIX {
		return true
	}

	return false
}

func NewLaptop(
	id UUID,
	name Name,
	price Price,
	availability constants.AVAILABILITY,
	stockCount uint8,
	screenSize ScreenSize,
	batteryCapacity BatteryCapacity,
	modelName Model,
	exclusive bool,
	os constants.OS,
) Laptop {
	l := Laptop{}

	l.Id = id
	l.Name = name
	l.Price = price
	l.Availablility = availability
	l.StockCount = stockCount
	l.ScreenSize = screenSize
	l.BatteryCapacity = batteryCapacity
	l.ModelName = modelName
	l.Exclusive = exclusive
	l.OS = os

	return l
}

func (laptop Laptop) GetProductId() UUID {
	return laptop.Id
}
