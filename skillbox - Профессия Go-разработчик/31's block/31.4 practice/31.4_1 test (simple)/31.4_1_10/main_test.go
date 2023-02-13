package main

import "testing"

func TestAddV1(t *testing.T) {
	exp := 5
	x := 2
	y := 3
	res := add(x, y)
	if res != exp {
		t.Fail()
	}

	y = 2
	res = add(x, y)
	if res != exp {
		t.Logf("Сумма чисел %d и %d не равна %d.\n", x, y, exp)
		t.Fail()
	}
}

func TestAddV2(t *testing.T) {
	for _, val := range []struct {
		exp int
		x   int
		y   int
	}{
		{
			exp: 5,
			x:   2,
			y:   3,
		},
		{
			exp: 6,
			x:   3,
			y:   4,
		},
		{
			exp: 12,
			x:   5,
			y:   7,
		},
		{
			exp: 15,
			x:   4,
			y:   10,
		},
	} {
		res := add(val.x, val.y)
		if res != val.exp {
			t.Logf("Сумма чисел %d и %d не равна %d.\n", val.x, val.y, val.exp)
			t.Fail()
		}
	}
}
