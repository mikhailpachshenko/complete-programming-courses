package main

import "testing"

func TestAdd(t *testing.T) {
	exp := 5
	x := 3
	y := 2
	res := add(x, y)
	if res != exp {
		t.Fail()
	}

	exp = 10
	x = 4
	y = 4
	res = add(x, y)
	if res != exp {
		t.Logf("Сумма чисел %d и %d не равна %d.", x, y, exp)
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
			exp: 5,
			x:   2,
			y:   3,
		},
		{
			exp: 6,
			x:   3,
			y:   3,
		},
		{
			exp: 6,
		},
		{
			exp: 7,
			y:   6,
		},
	} {
		res := add(v.x, v.y)
		if res != v.exp {
			t.Logf("Сумма чисел %d и %d не равна %d.\n", v.x, v.y, v.exp)
			t.Fail()
		}
	}
}
