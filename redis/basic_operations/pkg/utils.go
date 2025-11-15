package redisops

// Helper function to repeat strings (Go doesn't have this built-in)
func Repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
