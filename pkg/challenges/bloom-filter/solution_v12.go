package bloom_filter

// SolutionV12 demonstrates Counting Bloom Filter with multiple hash functions.
type SolutionV12 struct {
	iteration int
}

func NewV12() *SolutionV12 {
	return &SolutionV12{iteration: 12}
}

func (s *SolutionV12) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
