package sliding_window

// SolutionV2 demonstrates Sliding Window Maximum sub-array algorithm.
type SolutionV2 struct {
	iteration int
}

func NewV2() *SolutionV2 {
	return &SolutionV2{iteration: 2}
}

func (s *SolutionV2) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
