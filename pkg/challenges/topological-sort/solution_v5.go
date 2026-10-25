package topological_sort

// SolutionV5 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV5 struct {
	iteration int
}

func NewV5() *SolutionV5 {
	return &SolutionV5{iteration: 5}
}

func (s *SolutionV5) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
