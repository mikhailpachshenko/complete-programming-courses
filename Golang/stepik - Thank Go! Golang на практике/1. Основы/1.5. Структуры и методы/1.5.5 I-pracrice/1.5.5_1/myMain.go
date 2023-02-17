package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type result byte

const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

type team byte

type match struct {
	first  team
	second team
	result result
}

type rating map[team]int

type tournament []match

func (trn *tournament) calcRating() rating {
	a := rating{}
	var A, B, C, D team
	for _, v := range *trn {
		if string(v.first) == "A" {
			A = v.first
			if string(v.result) == "W" {
				a[A] += 3
			} else if string(v.result) == "L" {
				a[A] += 0
			} else if string(v.result) == "D" {
				a[A] += 1
			}
		} else if string(v.second) == "A" {
			A = v.second
			if string(v.result) == "W" {
				a[A] += 0
			} else if string(v.result) == "L" {
				a[A] += 3
			} else if string(v.result) == "D" {
				a[A] += 1
			}
		}

		if string(v.first) == "B" {
			B = v.first
			if string(v.result) == "W" {
				a[B] += 3
			} else if string(v.result) == "L" {
				a[B] += 0
			} else if string(v.result) == "D" {
				a[B] += 1
			}
		} else if string(v.second) == "B" {
			B = v.second
			if string(v.result) == "W" {
				a[B] += 0
			} else if string(v.result) == "L" {
				a[B] += 3
			} else if string(v.result) == "D" {
				a[B] += 1
			}
		}

		if string(v.first) == "C" {
			C = v.first
			if string(v.result) == "W" {
				a[C] += 3
			} else if string(v.result) == "L" {
				a[C] += 0
			} else if string(v.result) == "D" {
				a[C] += 1
			}
		} else if string(v.second) == "C" {
			C = v.second
			if string(v.second) == "W" {
				a[C] += 0
			} else if string(v.result) == "L" {
				a[C] += 3
			} else if string(v.result) == "D" {
				a[C] += 1
			}
		}

		if string(v.first) == "D" {
			D = v.first
			if string(v.result) == "W" {
				a[D] += 3
			} else if string(v.result) == "L" {
				a[D] += 0
			} else if string(v.result) == "D" {
				a[D] += 1
			}
		} else if string(v.second) == "D" {
			D = v.second
			if string(v.second) == "W" {
				a[D] += 0
			} else if string(v.result) == "L" {
				a[D] += 3
			} else if string(v.result) == "D" {
				a[D] += 1
			}
		}
	}
	return a
}

func main() {
	src := readString() // ABW DCD DAW CBL BDL ACW
	trn := parseTournament(src)

	rt := trn.calcRating()
	rt.print() // A6 B3 C1 D7
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
