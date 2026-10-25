package sliding_window

// SolutionV5 demonstrates Sliding Window Maximum sub-array algorithm.
type SolutionV5 struct {
	iteration int
}

func NewV5() *SolutionV5 {
	return &SolutionV5{iteration: 5}
}

func (s *SolutionV5) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
