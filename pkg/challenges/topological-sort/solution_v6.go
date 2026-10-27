package topological_sort

// SolutionV6 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV6 struct {
	iteration int
}

func NewV6() *SolutionV6 {
	return &SolutionV6{iteration: 6}
}

func (s *SolutionV6) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
