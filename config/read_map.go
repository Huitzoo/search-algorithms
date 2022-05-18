package config

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"search/structs/algorithms"
	"strings"
)

var runeVariable rune = '#'

func NewStageFromBoard(path, kind string) algorithms.StageInterface {
	f, err := os.Open(path)
	fmt.Println("READ FROM: ", path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	row := 0
	scanner := bufio.NewScanner(f)
	col := 0
	blockedNodes := [][]int{}
	initialNode := []int{}
	goalNode := []int{}

	for scanner.Scan() {
		values := scanner.Text()
		for col, value := range values {
			if value == runeVariable {
				blockedNodes = append(blockedNodes, []int{col, row})
			}
		}
		if strings.Contains(values, "A") {
			initialNode = []int{strings.Index(values, "A"), row}
		}

		if strings.Contains(values, "B") {
			goalNode = []int{strings.Index(values, "B"), row}
		}
		row++
		col = len(values)
	}

	size := []int{col, row}

	return TypeOfBoard(
		size, initialNode, goalNode, blockedNodes, kind,
	)
}

func TypeOfBoard(
	size []int,
	initialState []int,
	goalState []int,
	blockedNodes [][]int,
	kind string,
) algorithms.StageInterface {
	switch kind {
	case "greedy":
		return algorithms.NewGreedyBestFirstSearch(
			size, initialState, goalState, blockedNodes,
		)
	case "breadth-lifo":
		return algorithms.NewBreadthFirstSearch(
			size, initialState, goalState, blockedNodes,
		)
	case "a_star":
		return algorithms.NewAStarSearch(
			size, initialState, goalState, blockedNodes,
		)

	default:
		return nil
	}
}
