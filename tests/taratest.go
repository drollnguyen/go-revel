package tests

import (
	"github.com/revel/revel/testing"
)

type TaraTest struct {
	testing.TestSuite
}

func (t *TaraTest) Before() {
	println("Show hand")
}

func (t *TaraTest) TestStandUp() {
	t.Get("/")
	a := 4
	if a == 4 {

		t.AssertOk()
	} else {
		t.AssertNotFound()
	}
}

func (t *TaraTest) TestStandDown() {
	// t.Get("/")
	t.AssertOk()
	// t.AssertNotFound()
}

// func (t *TaraTest) TestRaiseYourHands() {
// 	// t.Get("/")
// 	t.AssertOk()
// 	// t.AssertContentType("text/html; charset=utf-8")
// }

func (t *TaraTest) After() {
	println("Tear down:: end show")
}
