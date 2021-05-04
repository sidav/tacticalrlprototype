package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func euclideanDistance(fx, fy, tx, ty int) int {
	return abs(fx-tx)+abs(fy-ty)
}

func areCoordsNeighbouring(x1, y1, x2, y2 int, diagonalsCount bool) bool {
	if diagonalsCount {
		return abs(x1-x2) <= 1 && abs(y1-y2) <= 1
	}
	return (x1 == x2 || y1 == y2) && (abs(x1-x2) == 1 || abs(y1-y2) == 1)
}
