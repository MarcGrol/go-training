package callingc
import "C"
import "fmt"

func main()  {
	f := C.getPerson(C.fortytwo)
	fmt.Println(int(C.bridge_int_func(f)))
}
