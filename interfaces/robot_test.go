package interfaces

import (
	"github.com/fromiuan/goexample/interfaces/lang"
	"testing"
)

func TestRebot(t *testing.T) {
	desk := 1

	l := lang.NewChinese()
	l.Welcome(desk)

	for key, name := range lang.Menu {
		t.Log(name)
		price := l.GetPrice(key)
		t.Log("价格是:", price)
	}
}
