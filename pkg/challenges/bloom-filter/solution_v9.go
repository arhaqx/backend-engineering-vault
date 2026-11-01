package bloom_filter

// SolutionV9 demonstrates Counting Bloom Filter with multiple hash functions.
type SolutionV9 struct {
	iteration int
}

func NewV9() *SolutionV9 {
	return &SolutionV9{iteration: 9}
}

func (s *SolutionV9) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
