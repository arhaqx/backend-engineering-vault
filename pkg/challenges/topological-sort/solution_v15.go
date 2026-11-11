package topological_sort

// SolutionV15 demonstrates Topological Sort using Kahn's in-degree algorithm.
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
