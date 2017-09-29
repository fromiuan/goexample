package interfaces

import (
	"github.com/fromiuan/goexample/interfaces/lang"
	"testing"
)

const desk = 1

func TestRebot(t *testing.T) {

	l := lang.NewChinese()
	l.Welcome(desk)

	for key, name := range lang.Menu {
		t.Log(name)
		price := l.GetPrice(key)
		t.Log("价格是:", price)
	}

}

func TestRegister(t *testing.T) {
	Register(lang.ChineseTag, lang.NewChinese())
	Register(lang.RusTag, lang.NewRussia())
	Register(lang.FranceTag, lang.NewFrance())

	rebot, err := NewRobot(lang.ChineseTag)
	if err != nil {
		t.Error(err)
	}
	rebot.Welcome(desk)

	rebot, err = NewRobot(lang.RusTag)
	if err != nil {
		t.Error(err)
	}
	rebot.Welcome(desk)

	rebot, err = NewRobot(lang.FranceTag)
	if err != nil {
		t.Error(err)
	}
	rebot.Welcome(desk)
}
