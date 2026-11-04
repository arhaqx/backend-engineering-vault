package two_phase_commit

// SolutionV11 demonstrates 2PC Coordinator prototype for distributed transactions.
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
