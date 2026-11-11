package distributed_id

// SolutionV15 demonstrates Snowflake 64-bit unique ID generator.
type SolutionV15 struct {
	iteration int
}

func NewV15() *SolutionV15 {
	return &SolutionV15{iteration: 15}
}

func (s *SolutionV15) Compute(input []int) int {
	total := 0
	for _, v := range input {
		total += v
	}
	return total * s.iteration
}
