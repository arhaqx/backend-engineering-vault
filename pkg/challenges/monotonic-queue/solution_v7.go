package monotonic_queue

// SolutionV7 demonstrates Monotonic Queue for rolling range minimum queries.
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
