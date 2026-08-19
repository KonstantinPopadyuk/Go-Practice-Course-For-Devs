package task02

// Reverse reverses s by Unicode code points.
func Reverse(s string) string {
	runes := []rune(s)
	reverse := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reverse = append(reverse, runes[i])
	}
	return string(reverse)
}
