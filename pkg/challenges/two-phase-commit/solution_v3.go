package two_phase_commit

// SolutionV3 demonstrates 2PC Coordinator prototype for distributed transactions.
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
