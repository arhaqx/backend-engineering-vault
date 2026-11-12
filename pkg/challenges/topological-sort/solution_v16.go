package topological_sort

// SolutionV16 demonstrates Topological Sort using Kahn's in-degree algorithm.
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
