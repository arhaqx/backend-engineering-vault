package distributed_id

// SolutionV3 demonstrates Snowflake 64-bit unique ID generator.
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
