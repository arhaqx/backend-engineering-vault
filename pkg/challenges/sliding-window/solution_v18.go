package sliding_window

// SolutionV18 demonstrates Sliding Window Maximum sub-array algorithm.
type SolutionV18 struct {
	iteration int
}

func NewV18() *SolutionV18 {
	return &SolutionV18{iteration: 18}
}

func (s *SolutionV18) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
