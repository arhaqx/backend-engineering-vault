package two_phase_commit

// SolutionV6 demonstrates 2PC Coordinator prototype for distributed transactions.
type SolutionV6 struct {
	iteration int
}

func NewV6() *SolutionV6 {
	return &SolutionV6{iteration: 6}
}

func (s *SolutionV6) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
