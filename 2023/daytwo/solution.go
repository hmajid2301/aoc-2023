package daytwo

import (
	"fmt"
	"strconv"
	"strings"
)

var MAX_RED_CUBES = 12
var MAX_BLUE_CUBES = 14
var MAX_GREEN_CUBES = 13

func CalculatePossibleGames(results string) int {

	lines := strings.Split(results, "\n")
	var total int
	for _, line := range lines {

		game := strings.Split(line, ":")
		gameNum := game[0]

		num := strings.ReplaceAll(gameNum, "Game ", "")
		fmt.Println("NUM", num)
		num_, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}

		turns := strings.Split(game[1], ";")
		impossible := false
		for _, turn := range turns {
			cubes := strings.Split(turn, ",")
			for _, cube := range cubes {
				v, err := strconv.Atoi(strings.Split(cube, " ")[1])
				if err != nil {
					panic(err)
				}

				if strings.Contains(cube, "blue") {
					if v > MAX_BLUE_CUBES {
						impossible = true
						break
					}
				} else if strings.Contains(cube, "red") {
					if v > MAX_RED_CUBES {
						impossible = true
						break
					}
				} else if strings.Contains(cube, "green") {
					if v > MAX_GREEN_CUBES {
						impossible = true
						break
					}
				}
			}
			if impossible {
				break
			}
		}

		if !impossible {
			total += num_
		}
	}

	return total
}
