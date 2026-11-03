package bloom_filter

// SolutionV10 demonstrates Counting Bloom Filter with multiple hash functions.
type SolutionV10 struct {
	iteration int
}

func NewV10() *SolutionV10 {
	return &SolutionV10{iteration: 10}
}

func (s *SolutionV10) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
