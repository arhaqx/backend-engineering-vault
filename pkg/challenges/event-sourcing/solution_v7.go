package event_sourcing

// SolutionV7 demonstrates Event Store append-only ledger with snapshot support.
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
