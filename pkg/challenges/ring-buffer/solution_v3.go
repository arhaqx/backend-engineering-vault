package ring_buffer

// SolutionV3 demonstrates Lock-free circular ring buffer with atomic pointers.
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
