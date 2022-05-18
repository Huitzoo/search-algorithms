package utils

func CalculateIDStateByCoords(coords []int, columns int) int {
	//id =  n_columns*y+x
	return (columns)*coords[1] + coords[0]
}
