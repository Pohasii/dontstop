package being

import (
	"testing"
)

/*
import "testing"

type testData struct {
	incoming int
	answer   int
}

func TestSing(t *testing.T) {

	var testData = []testData{
		{-21, -1},
		{0, 0},
		{32423, 1},
	}

	for _, td := range testData {
		var v int
		v = Sign(td.incoming)
		if v != td.answer {
			t.Error(
				"For", td.incoming,
				"expected", td.answer,
				"got", v,
			)
		}
	}

}
*/

func TestDefaultHp(t *testing.T) {

	var testD hp = 500

	if DefaultHp != testD {
		t.Error("Expected 500, got ", DefaultHp)
	}

}

func TestDefaultLiveStatus(t *testing.T) {

	var testD int = 1

	if DefaultLiveStatus != testD {
		t.Error("Expected 1, got ", DefaultLiveStatus)
	}

}

func TestDefaultSize(t *testing.T) {

	var testD int = 15

	if DefaultSize != testD {
		t.Error("Expected 1, got ", DefaultSize)
	}

}

func TestDefaultEnergy(t *testing.T) {

	var testD int = 100

	if DefaultEnergy != testD {
		t.Error("Expected 1, got ", DefaultEnergy)
	}

}

func TestDefaultAge(t *testing.T) {

	var testD int = 0

	if DefaultAge != testD {
		t.Error("Expected 1, got ", DefaultAge)
	}

}

func TestDefaultColor(t *testing.T) {

	var testD color = color{185, 244, 66}

	if DefaultColor != testD {
		t.Error("Expected 185, 244, 66, got ", DefaultColor)
	}

}
