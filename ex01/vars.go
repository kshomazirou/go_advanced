package main

import "fmt"

type FortyTwo struct{}

func main() {
	fmt.Printf("%v.$\n", "42 is a string")
	fmt.Printf("%v is a %T.$\n", 42, uint(42))
	fmt.Printf("%v is an %T.$\n", 42, int(42))
	fmt.Printf("%v is a %T.$\n", 42, uint8(42))
	fmt.Printf("%v is an %T.$\n", 42, int16(42))
	fmt.Printf("%v is a %T.$\n", 42, uint32(42))
	fmt.Printf("%v is an %T.$\n", 42, int64(42))
	fmt.Printf("%v is a %T.$\n", false, false)
	fmt.Printf("%v is a %T.$\n", 42, float32(42))
	fmt.Printf("%v is a %T.$\n", 42, float64(42))
	fmt.Printf("%v is a %T.$\n", complex(42, 0), complex64(complex(42, 0)))
	fmt.Printf("%v is a %T.$\n", complex(42, 21), complex128(complex(42, 21)))
	fmt.Printf("%v is a %T.$\n", FortyTwo{}, FortyTwo{})
	fmt.Printf("%v is a %T.$\n", [1]int{42}, [1]int{42})
	fmt.Printf("%v is a %T.$\n", map[string]int{"42": 42}, map[string]int{"42": 42})
	fmt.Printf("%#v is an %T.$\n", (*int)(nil), (*int)(nil))
	fmt.Printf("%#v is a %T.$\n", []int{}, []int{})
	fmt.Printf("%#v is a %T.$\n", (chan int)(nil), (chan int)(nil))
	fmt.Printf("%#v is a %T.$\n", nil, nil)
}
