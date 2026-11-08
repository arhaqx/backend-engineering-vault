package lfu_cache

// SolutionV13 demonstrates Least Frequently Used (LFU) cache eviction engine.
type SolutionV13 struct {
	iteration int
}

func NewV13() *SolutionV13 {
	return &SolutionV13{iteration: 13}
}

func (s *SolutionV13) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
