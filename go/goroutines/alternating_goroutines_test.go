package goroutines

func Example_alternativeGoroutines() {
	alternatingGoroutinesABBA()
	alternatingGoroutinesABAB()
	alternatingGoroutinesABAB_Buffered()
	// Output:
	// ABBAABBAAB
	// ABABABABAB
	// ABABABABAB
}
