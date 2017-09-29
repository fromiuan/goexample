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
	RusTag = "rus"
)

var (
	RusMenu = map[int]string{
		1: "фруктовый салат ",
		2: "филе  помидоров ",
		3: "с помидорами ",
		4: "пельмени",
	}

	RusPrice = map[int]float64{
		1: 24.09,
		2: 150.32,
		3: 15,
		4: 25.09,
	}

	RussiaLang = &Russia{}
)

type Russia struct {
	lock     *sync.RWMutex
	Desk     int
	Lang     string
	DishMenu map[int]int
}

func NewRussia() *Russia {
	return &Russia{
		lock:     new(sync.RWMutex),
		Lang:     RusTag,
		DishMenu: make(map[int]int),
	}
}

func (c *Russia) Welcome(desk int) {
	call := fmt.Sprintf("Добро пожаловать,  стол номер  %d гостей  , извините,  нужно  что - нибудь  ? ", desk)
	fmt.Println(call)
	c.Desk = desk
}

func (c *Russia) Name(key int) string {
	if menu, ok := RusMenu[key]; ok {
		return menu
	}
	return "меню  не существует "
}

func (c *Russia) GetPrice(key int) float64 {
	if price, ok := RusPrice[key]; ok {
		return price
	}
	return 0.00
}

func (c *Russia) SetDish(key, number int) bool {
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

func (c *Russia) DelDish(key, number int) bool {
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

func (c *Russia) updatedish() {
	DeskNumber[c.Desk] = c.DishMenu
}

func (c *Russia) TotalDish() (count float64) {
	c.lock.RLocker()
	defer c.lock.RUnlock()
	for key, _ := range c.DishMenu {
		if price, ok := RusPrice[key]; ok {
			count = count + price
		}
	}
	msg := fmt.Sprintf("в общей сложности  %d  юаней ", count)
	fmt.Println(msg)
	return count
}

func (c *Russia) Pay() bool {
	if _, ok := DeskNumber[c.Desk]; ok {
		delete(DeskNumber, c.Desk)
		return true
		fmt.Println("Спасибо за ваше покровительство ")
	}
	return false
}
