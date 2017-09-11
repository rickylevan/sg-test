// map implementation in go. map from int->int, unassigned 0
// two methods, assign and get.

package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

const size = 4

type list struct {
	key int
	val int
	nxt *list
}

type nap struct {
	data [size]*list
}

// return a number between 0 and size-1
func (*nap) hash(n int) uint {
	str := []byte(strconv.Itoa(n))
	h := sha256.Sum256(str)
	return uint(h[0]) % size
}

func (n *nap) get(key int) int {
	spot := n.hash(key)
	l := n.data[spot]

	for {
		if l == nil {
			return 0
		}
		if l.key == key {
			return l.val
		}
		l = l.nxt
	}
}

func (n *nap) assign(key int, val int) {
	spot := n.hash(key)
	l := n.data[spot]

	if l == nil {
		n.data[spot] = &list{
			key: key,
			val: val,
			nxt: nil,
		}
		return
	}
	for {
		if l.key == key {
			l.val = val
			break
		} else if l.nxt == nil {
			l.nxt = &list{
				key: key,
				val: val,
				nxt: nil,
			}
			break
		}
		l = l.nxt
	}
}

func step(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	out := 0
	for i := n; i >= 1; i-- {
		out += step(n - i)
	}
	return out
}

type pair struct {
	x int
	y int
}

var cache map[pair]int

func pathGrid(X, Y int) int {

	if X == 1 || Y == 1 {
		return 1
	}

	if cache == nil {
		cache = make(map[pair]int)
	}

	if val, ok := cache[pair{X, Y}]; ok {
		return val
	}

	out := pathGrid(X-1, Y) + pathGrid(X, Y-1)

	cache[pair{X, Y}] = out

	return out

}

// return all permutations of a string
func sp(prior, rest string) []string {

	if rest == "" {
		return []string{prior}
	}
	var out []string
	for i := 0; i < len(rest); i++ {
		sub := sp(prior+string(rest[i]),
			rest[0:i]+rest[i+1:len(rest)])
		out = append(out, sub...)
	}

	return out
}

func main() {

	fmt.Println(sp("", "abcde"))

}
