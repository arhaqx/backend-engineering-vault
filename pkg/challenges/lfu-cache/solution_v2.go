package lfu_cache

// SolutionV2 demonstrates Least Frequently Used (LFU) cache eviction engine.
type SolutionV2 struct {
	iteration int
}

func NewV2() *SolutionV2 {
	return &SolutionV2{iteration: 2}
}

func (s *SolutionV2) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
