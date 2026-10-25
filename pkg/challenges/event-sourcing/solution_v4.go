package event_sourcing

// SolutionV4 demonstrates Event Store append-only ledger with snapshot support.
type SolutionV4 struct {
	iteration int
}

func NewV4() *SolutionV4 {
	return &SolutionV4{iteration: 4}
}

func (s *SolutionV4) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
