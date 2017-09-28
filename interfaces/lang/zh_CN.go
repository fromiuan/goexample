package lang

import (
	"fmt"
	"sync"
)

// Welcome(tableNumber string)
// Name(key int) string
// GetPrice(kye int) float64
// SetDish(key, number int) bool
// DelDish(key, number int) bool
// TotalDish() float64
// Pay() bool

const (
	lang = "zh_cn"
)

var (
	DeskNumber = make(map[int]map[int]int)
	Menu       = map[int]string{
		1: "水果沙拉",
		2: "番茄牛腩",
		3: "番茄炒蛋",
		4: "水饺",
	}

	Price = map[int]float64{
		1: 24.09,
		2: 150.32,
		3: 15,
		4: 25.09,
	}
)

type Chinese struct {
	lock     *sync.RWMutex
	Desk     int
	Lang     string
	DishMenu map[int]int
}

func NewChinese() *Chinese {
	return &Chinese{
		lock:     new(sync.RWMutex),
		Lang:     lang,
		DishMenu: make(map[int]int),
	}
}

func (c *Chinese) Welcome(desk int) {
	call := fmt.Sprintf("欢迎光临,第%d号桌客人,	请问需要点儿什么？", desk)
	fmt.Println(call)
	c.Desk = desk
}

func (c *Chinese) Name(key int) string {
	if menu, ok := Menu[key]; ok {
		return menu
	}
	return "菜单不存在"
}

func (c *Chinese) GetPrice(key int) float64 {
	if price, ok := Price[key]; ok {
		return price
	}
	return 0.00
}

func (c *Chinese) SetDish(key, number int) bool {
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

func (c *Chinese) DelDish(key, number int) bool {
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

func (c *Chinese) updatedish() {
	DeskNumber[c.Desk] = c.DishMenu
}

func (c *Chinese) TotalDish() (count float64) {
	c.lock.RLocker()
	defer c.lock.RUnlock()
	for key, _ := range c.DishMenu {
		if price, ok := Price[key]; ok {
			count = count + price
		}
	}
	return count
}

func (c *Chinese) Pay() bool {
	if _, ok := DeskNumber[c.Desk]; ok {
		delete(DeskNumber, c.Desk)
		return true
	}
	return false
}
