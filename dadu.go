package main

import (
	"fmt"
	"math/rand"
)

const REMOVED_WHEN_DICE_TOP = 6
const MOVE_WHEN_DICE_TOP = 1

var random = []int{1, 2, 3, 4, 5, 6}

func Dadu(N, M int) []int {
	var j []int
	var players []int
	var dadu []int

	// point := 0

	for l := 1; l <= M; l++ {
		players = append(players, l)
	}

	for k := 1; k <= N; k++ {
		dadu = append(dadu, k)
	}

	throw := random[rand.Intn(len(random))]

	for i := 1; i <= throw; i++ {
		j = append(j, i)
	}

	return j
}

func main() {
	fmt.Println(Dadu(3, 4))
}
