package two_phase_commit

// SolutionV14 demonstrates 2PC Coordinator prototype for distributed transactions.
type SolutionV14 struct {
	iteration int
}

func NewV14() *SolutionV14 {
	return &SolutionV14{iteration: 14}
}

func (s *SolutionV14) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
