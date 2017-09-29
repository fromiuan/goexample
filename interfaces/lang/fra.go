package lang

import (
	"fmt"
	"sync"
)

// Welcome(tableNumber string)
// 	Name(key int) string
// 	GetPrice(kye int) float64
// 	SetDish(key, number int) bool
// 	DelDish(key, number int) bool
// 	TotalDish() float64
// 	Pay() bool

const (
	FranceTag = "fra"
)

var (
	Menu = map[int]string{
		1: "Salade de fruits",
		2: "Tomate aloyau",
		3: "Une omelette aux tomates",
		4: "Des boulettes",
	}

	FrancePrice = map[int]float64{
		1: 24.09,
		2: 150.32,
		3: 15,
		4: 25.09,
	}

	FranceLang = &France{}
)

type France struct {
	lock     *sync.RWMutex
	Desk     int
	Lang     string
	DishMenu map[int]int
}

func NewFrance() *France {
	return &France{
		lock:     new(sync.RWMutex),
		Lang:     FranceTag,
		DishMenu: make(map[int]int),
	}
}

func (c *France) Welcome(desk int) {
	call := fmt.Sprintf("Bienvenue, l'hôte%d, vous avez besoin de quelque chose?", desk)
	fmt.Println(call)
	c.Desk = desk
}

func (c *France) Name(key int) string {
	if menu, ok := Menu[key]; ok {
		return menu
	}
	return "Le menu n'existe pas"
}

func (c *France) GetPrice(key int) float64 {
	if price, ok := FrancePrice[key]; ok {
		return price
	}
	return 0.00
}

func (c *France) SetDish(key, number int) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if value, ok := c.DishMenu[key]; ok {
		c.DishMenu[key] = value + number
	} else {
		c.DishMenu[key] = number
	}
	c.updatedish()
	return true
}

func (c *France) DelDish(key, number int) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if value, ok := c.DishMenu[key]; ok {
		if value <= number {
			delete(c.DishMenu, key)
		} else {
			c.DishMenu[key] = value - key
		}
	}
	c.updatedish()
	return true
}

func (c *France) updatedish() {
	DeskNumber[c.Desk] = c.DishMenu
}

func (c *France) TotalDish() (count float64) {
	c.lock.RLocker()
	defer c.lock.RUnlock()
	for key, _ := range c.DishMenu {
		if price, ok := FrancePrice[key]; ok {
			count = count + price
		}
	}
	msg := fmt.Sprintf("Total%d $", count)
	fmt.Println(msg)
	return count
}

func (c *France) Pay() bool {
	if _, ok := DeskNumber[c.Desk]; ok {
		delete(DeskNumber, c.Desk)
		return true
		fmt.Println("Merci")
	}
	return false
}
