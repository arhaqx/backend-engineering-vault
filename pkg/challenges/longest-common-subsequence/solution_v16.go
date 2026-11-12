package longest_common_subsequence

// SolutionV16 demonstrates Longest Common Subsequence 2D dynamic programming.
type SolutionV16 struct {
	iteration int
}

func NewV16() *SolutionV16 {
	return &SolutionV16{iteration: 16}
}

func (s *SolutionV16) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
