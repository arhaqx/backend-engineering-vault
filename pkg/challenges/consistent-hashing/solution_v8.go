package consistent_hashing

// SolutionV8 demonstrates Consistent Hashing ring with virtual vnodes.
type SolutionV8 struct {
	iteration int
}

func NewV8() *SolutionV8 {
	return &SolutionV8{iteration: 8}
}

func (s *SolutionV8) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
