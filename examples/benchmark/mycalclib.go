package benchmark

//START OMIT
type BigStruct struct {
	data [1024 * 1000]string // at what size are pointers faster?
}

func (bs *BigStruct) DoCalculationByReference() int {
	return 42
}

func (bs BigStruct) DoCalculationByValue() int {
	return 1984
}

// END OMIT
