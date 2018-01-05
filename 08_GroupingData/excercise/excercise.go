package excercise

import (
	"fmt"
)

//First FUNC
func First() {
	fmt.Println("ciao")
}

//TestArray FUNC
func TestArray() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

//TestSlices FUNC
func TestSlices() {
	//COMPOSITE LITERAL
	// x := type {values}
	x := []int{4, 5, 7, 8, 42} //A SLICE allows to group together VALUES of the SAME TYPE
	fmt.Println(x)
	fmt.Println("Lunghezza dello slice: ", len(x))
	fmt.Println("Primo Elemento dello Slice ", x[0])
	for i, v := range x {
		fmt.Println(i, v)
	}

}

//SlicingSlice FUNC
func SlicingSlice() {
	//COMPOSITE LITERAL
	// x := type {values}
	x := []int{4, 5, 7, 8, 42} //A SLICE allows to group together VALUES of the SAME TYPE
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	x = append(x, 100, 102)
	fmt.Println(x)

	y := []int{300, 400, 500}

	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // Remove elements from a SLICE
	fmt.Println(x)

}
