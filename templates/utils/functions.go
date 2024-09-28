package utils

func eta(a int, b int) int {
	time := b - a

	return (time / 60) + 60
}
