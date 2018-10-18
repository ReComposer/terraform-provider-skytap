package skytap

// ptrToStr returns a string value for the passed string pointer.
// It returns the empty string if the pointer is nil.
func ptrToStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// strToPtr returns a pointer to the passed string.
func strToPtr(s string) *string {
	return &s
}

// ptrToInt returns an int value for the passed int pointer.
// It returns 0 if the pointer is nil.
func ptrToInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// intToPtr returns a pointer to the passed int.
func intToPtr(i int) *int {
	return &i
}