package iteration

import "fmt"

func ExampleRepeatBasic() {
	fmt.Println(Repeat("a"))
	// Output: aaaaa
}

func ExampleRepeatParameterized() {
	fmt.Println(Repeat("a", 4))
	// Output: aaaa
}
