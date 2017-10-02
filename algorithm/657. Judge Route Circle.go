package algorithm

import (
	"fmt"
)

func judgeCircle(moves string) bool {
	var locX, locY int
	var moveX, moveY int
	for _, move := range moves {
		switch string(move) {
		case "U":
			moveX = 0
			moveY = 1
		case "D":
			moveX = 0
			moveY = -1
		case "L":
			moveX = -1
			moveY = 0
		case "R":
			moveX = 1
			moveY = 0
		}
		locX += moveX
		locY += moveY
	}
	if locX == 0 && locY == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(judgeCircle("LL"))
}
