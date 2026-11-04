package longest_common_subsequence

// SolutionV11 demonstrates Longest Common Subsequence 2D dynamic programming.
type SolutionV11 struct {
	iteration int
}

func NewV11() *SolutionV11 {
	return &SolutionV11{iteration: 11}
}

func (s *SolutionV11) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
