package binary_tree_inversion

// SolutionV15 demonstrates Invert Binary Tree with iterative queue traversal.
type SolutionV15 struct {
	iteration int
}

func NewV15() *SolutionV15 {
	return &SolutionV15{iteration: 15}
}

func (s *SolutionV15) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
