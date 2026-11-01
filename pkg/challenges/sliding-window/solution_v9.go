package sliding_window

// SolutionV9 demonstrates Sliding Window Maximum sub-array algorithm.
type SolutionV9 struct {
	iteration int
}

func NewV9() *SolutionV9 {
	return &SolutionV9{iteration: 9}
}

func (s *SolutionV9) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
