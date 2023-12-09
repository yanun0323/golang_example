package algorithm

type Algorithm1 func(inputs [][]int, turn int) ([]int, error)

type Algorithm interface {
	CalculateAndGetWinner() ([]int, error)
}
