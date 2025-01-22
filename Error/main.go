package main

import (
	"fmt"
	"strconv"
)

func main() {
	// val := true
	// ans, err := Check(val)
	// if err != nil {
	// 	fmt.Printf("%v is greater than 5? %v\n", val, err)
	// } else {
	// 	fmt.Printf("%v is greater than 5? %v\n", val, ans)
	// }

	e := process()
	if e != nil {
		fmt.Println(e)
	}

}

// making custom error
type CodeError struct {
	Code int
}

func process() error {
	c := CodeError{
		Code: 201,
	}
	return c
}

// Implement the Error method for CustomError
func (e CodeError) Error() string {
	return fmt.Sprintf("ErrorCode:%d\n", e.Code)
}

func Check(val any) (bool, error) {
	switch v := val.(type) {
	case int:
		return (val.(int) > 5), nil
	case string:
		value := val.(string)
		num, err := strconv.Atoi(value)
		if err != nil {
			return false, err
		}
		return num > 5, nil
	default:
		fmt.Println(v)
		return false, fmt.Errorf("unknown datatype %v", v)
	}
}
