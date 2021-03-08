package fmt

import "fmt"

func quote(input interface{}) string {
	var inputStr string
	switch input := input.(type) {
	case string:
		inputStr = input
		break
	case fmt.Stringer:
		inputStr = input.String()
		break
	default:
		inputStr = fmt.Sprintf("%v", input)
	}
	return fmt.Sprintf("%q", inputStr)
}
