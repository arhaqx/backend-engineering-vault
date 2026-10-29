package topological_sort

// SolutionV7 demonstrates Topological Sort using Kahn's in-degree algorithm.
type SolutionV7 struct {
	iteration int
}

func NewV7() *SolutionV7 {
	return &SolutionV7{iteration: 7}
}

func (s *SolutionV7) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
