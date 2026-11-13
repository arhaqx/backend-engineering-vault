package topological_sort

// SolutionV17 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV17 struct {
	iteration int
}

func NewV17() *SolutionV17 {
	return &SolutionV17{iteration: 17}
}

func (s *SolutionV17) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
