package firsttry

import (
	"errors"
	"fmt"
)

func foo() {
	fmt.Println("foo")
}

type Face interface {
	Bar() (string, error)
	Baz() (string, error)
}

type Mine struct{}

func (m *Mine) Bar() (string, error) {
	return "bar", nil
}

func (m *Mine) Baz() (string, error) {
	return "", errors.New("baz returns error")
}

func Foolish(f Face) error {
	str, err := f.Bar()
	if err != nil {
		return err
	}

	if str == "hello world" {
		fmt.Println("I got hello world")
	} else if str == "bar" {
		return errors.New("I got bar")
	} else {
		return errors.New(fmt.Sprintf("What I got %s", str))
	}

	str, err = f.Baz()
	if err != nil {
		return err
	}

	if str == "" {
		fmt.Println("I got empty string")
	} else if str == "u r fooled" {
		fmt.Println("I got a dumpass")
	} else {
		return errors.New(fmt.Sprintf("What I got %s", str))
	}
	return nil
}
