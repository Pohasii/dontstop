package functions

import "testing"

type testData struct {
	incoming int
	answer   int
}

func TestMathSign(t *testing.T) {

	var testData = []testData{
		{-21, -1},
		{0, 0},
		{32423, 1},
	}

	for _, td := range testData {
		var v int
		v = MathSign(td.incoming)
		if v != td.answer {
			t.Error(
				"For", td.incoming,
				"expected", td.answer,
				"got", v,
			)
		}
	}

}
