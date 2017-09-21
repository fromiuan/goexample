package lang

import (
	"fmt"
)

// Welcome(tableNumber string)
// Name(key int) string
// GetPrice(kye int) float64
// SetDish(key, number int) bool
// DelDish(key, number int) bool
// TotalDish() string
// Pay() bool

var (
	lang = "zh_cn"
	Menu = map[int]string{
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

type Language struct {
	Lang string
}

func NewLanguage() *Language {
	return &Language{
		Lang: lang,
	}
}

func (l *Language) Welcome(name string) {
	call := fmt.Sprintf("欢迎光临,第%s号桌客人", name)
	fmt.Println(call)
}

func (l *Language) Name(key int) string {
	if menu, ok := Menu[key]; ok {
		return men
	}
	return "菜单不存在"
}

func (l *Language) GetPrice(key int) float64 {
	if price, ok := Price[key]; ok {
		return price
	}
	return 0.00
}

func (l *Language) SetDish(key, number int) bool {

}

func (l *Language) DelDish(key, number int) bool {

}

func (l *Language) TotalDish() {

}

func (l *Language) Pay() {

}
