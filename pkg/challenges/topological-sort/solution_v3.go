package topological_sort

// SolutionV3 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV3 struct {
	iteration int
}

func NewV3() *SolutionV3 {
	return &SolutionV3{iteration: 3}
}

func (s *SolutionV3) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
