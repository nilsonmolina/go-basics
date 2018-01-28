package main

import (
	"fmt"
	"log"
)

/* HANDS ON # 4 -----------------------------------
1. starting with this code, use the sqrt.Error struct as a value of type error.
If you would like, use these numbers:
	- lat "50.2289 N"
	- long "99.4656 W"

type sqrtError struct {
	lat  string
	long string
	err  error
}
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}
func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
	}
	return 42, nil
}
----------------------------------- */

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		//e := errors.New("input was less than 0")
		e := fmt.Errorf("input {%v} is less than 0 ", f)
		return 0, sqrtError{lat: "50.2289 N", long: "99.4656 W", err: e}
	}
	return 42, nil
}
