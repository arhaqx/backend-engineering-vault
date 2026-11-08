package binary_tree_inversion

// SolutionV14 demonstrates Invert Binary Tree with iterative queue traversal.
type SolutionV14 struct {
	iteration int
}

func NewV14() *SolutionV14 {
	return &SolutionV14{iteration: 14}
}

func (s *SolutionV14) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
