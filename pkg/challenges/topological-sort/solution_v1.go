package topological_sort

// SolutionV1 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV1 struct {
	iteration int
}

func NewV1() *SolutionV1 {
	return &SolutionV1{iteration: 1}
}

func (s *SolutionV1) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
