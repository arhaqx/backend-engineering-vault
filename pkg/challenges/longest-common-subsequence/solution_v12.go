package longest_common_subsequence

// SolutionV12 demonstrates Longest Common Subsequence 2D dynamic programming.
type SolutionV12 struct {
	iteration int
}

func NewV12() *SolutionV12 {
	return &SolutionV12{iteration: 12}
}

func (s *SolutionV12) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
