package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rawDice := os.Args[1:]

	dice, err := parseDice(rawDice)
	if err != nil {
		fmt.Println(err)
		return
	}

	rerollDice(dice)

	fmt.Println(dice)
}

// todo(benCoomes): handle CSV and other delimiters?
func parseDice(rawDice []string) ([]int, error) {
	dice := make([]int, len(rawDice))
	for i, rawDie := range rawDice {
		if rawDie == "X" || rawDie == "x" {
			dice[i] = 0
			continue
		}

		val, err := strconv.Atoi(rawDie)
		if err != nil {
			// Todo: wrap error?
			return nil, err
		}

		if val > 6 {
			return nil, errors.New("dice values greater than 6 are not supported")
		}

		dice[i] = val
	}

	return dice, nil
}

func rerollDice(dice []int) {
	rsrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rsrc)

	for i, die := range dice {
		if die == 0 {
			roll := r.Intn(6) + 1
			dice[i] = roll
		}
	}
}
