package interfaces

import (
	"fmt"
)

var (
	DefaultMenu = map[int]string{
		1: "Fresh Fruit Salad",
		2: "Curry sirloin",
		3: "Fried eggs with tomatoes",
		4: "dumplings",
	}

	DefaultPrice = map[int]float64{
		1: 12.03,
		2: 22.00,
		3: 2.00,
		4: 3.20,
	}
)

// define OrderRoBot interface
type RoBot interface {
	Welcome(tableNumber int)
	Name(key int) string
	GetPrice(kye int) float64
	SetDish(key, number int) bool
	DelDish(key, number int) bool
	TotalDish() float64
	Pay() bool
}

// define adapters use map key to value
var adapters = make(map[string]RoBot)

// register adapters
func Register(name string, adapter RoBot) {
	if adapter == nil {
		panic("adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// create newrebot
func NewRobot(adapterName string) (adapter RoBot, err error) {
	if adapter, ok := adapters[adapterName]; ok {
		return adapter, nil
	} else {
		err = fmt.Errorf("unknow adapter name%s", adapterName)
	}
	return nil, err
}
