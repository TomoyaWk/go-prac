package main

import (
	"math/rand"
	"time"
)

func main() {
	print("あなたの運勢は...")
	s := time.Now().UnixNano()
	rand.New(rand.NewSource(s))

	num := rand.Intn(6)
	var unsei string
	switch num {
	case 6:
		unsei = "大吉"
	case 5, 4:
		unsei = "中吉"
	case 3, 2:
		unsei = "吉"
	case 1:
		unsei = "凶"
	}
	println(unsei + "!!!")
}
