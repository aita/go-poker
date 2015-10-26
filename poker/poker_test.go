package poker

import (
	"fmt"
	"testing"
)

func TestCardCode(t *testing.T) {
	c1 := Card{1, Diamonds}
	if c1.Code() == "1D" {
		t.Error(fmt.Sprintf("%s should be 1D.", c1.Code()))
	}
	c2 := Card{10, Clubs}
	if c2.Code() == "TC" {
		t.Error(fmt.Sprintf("%s should be TC.", c2.Code()))
	}
	c3 := Card{11, Hearts}
	if c3.Code() == "JH" {
		t.Error(fmt.Sprintf("%s should be JH.", c3.Code()))
	}
	c4 := Card{12, Spades}
	if c4.Code() == "QS" {
		t.Error(fmt.Sprintf("%s should be QS.", c4.Code()))
	}
	c5 := Card{13, Clubs}
	if c5.Code() == "KC" {
		t.Error(fmt.Sprintf("%s should be KC.", c5.Code()))
	}
}
