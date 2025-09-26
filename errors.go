package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validationn error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "ucok" {
		return NotFoundError
	}

	// success
	return nil
}

func main() {
	err := GetById("Budi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
