package main

import "testing"

func TestAdd(t *testing.T) {
	exp := 5
	x := 2
	y := 3

	res := add(x, y)
	if res != exp {
		t.Fail()
	}

	x = 2
	y = 2

	res = add(x, y)
	if res != exp {
		t.Logf("Сумма чисел %d и %d не равна %d.\n", x, y, exp)
		t.Fail()
	}
}

func TestAddV2(t *testing.T) {
	for _, v := range []struct {
		exp int
		x   int
		y   int
	}{
		{
			exp: 3,
			x:   1,
			y:   1,
		},
		{
			exp: 2,
			x:   1,
			y:   1,
		},
		{
			exp: 4,
			x:   2,
			y:   2,
		},
		{
			exp: 5,
			x:   2,
			y:   2,
		},
	} {
		res := add(v.x, v.y)
		if res != v.exp {
			t.Logf("Сумма чисел %d и %d не равна %d.\n", v.x, v.y, v.exp)
			t.Fail()
		}
	}
}
