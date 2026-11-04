package monotonic_queue

// SolutionV11 demonstrates Monotonic Queue for rolling range minimum queries.
type SolutionV11 struct {
	iteration int
}

func NewV11() *SolutionV11 {
	return &SolutionV11{iteration: 11}
}

func (s *SolutionV11) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
