package benchmark

type BigStruct struct {
	data [1024 * 100]string
}

func (bs *BigStruct) DoCalculationByReference() int {
	return 42
}

func (bs BigStruct) DoCalculationByValue() int {
	return 1984
}
