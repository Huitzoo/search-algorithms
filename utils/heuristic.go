package utils

import "math"

func CalculateDistance(
	goal []int,
	current []int,
) float64 {

	return math.Abs(float64(goal[0]-current[0])) + math.Abs(float64(goal[1]-current[1]))
}
