package sliding_window

// SolutionV13 demonstrates Sliding Window Maximum sub-array algorithm.
type SolutionV13 struct {
	iteration int
}

func NewV13() *SolutionV13 {
	return &SolutionV13{iteration: 13}
}

func (s *SolutionV13) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
