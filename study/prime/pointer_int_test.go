package prime

import (
	"fmt"
	"testing"
)

func TestPointerInt(t *testing.T) {
	v := 5
	p := &v
	println(*p) //5
	fmt.Println(p)
	changePointer(p)
	println(*p) //3
	changePointerOther(p)
	println(*p) //3
}

func changePointer(p *int) {
	v := 3
	p = &v

}

func changePointerOther(p *int) *int {
	fmt.Println(p)
	v := 3
	p = &v
	return p
}
