package functions

import "fmt"

// START OMIT
func ConvertIt(arg int) string { // exported function // HL
	return convertInternal(arg)
}
func convertInternal(arg int) string { // internal, unexported function
	return fmt.Sprintf("My integer value as string: %d", arg)
}

// END OMIT
