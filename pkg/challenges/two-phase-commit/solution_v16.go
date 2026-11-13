package two_phase_commit

// SolutionV16 demonstrates 2PC Coordinator prototype for distributed transactions.
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
