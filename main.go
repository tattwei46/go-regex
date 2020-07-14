package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Test1: ", TestSpecial("!@#$%"))
	fmt.Println("Test2: ", TestSpecial("1234"))
	fmt.Println("Test3: ", TestLowercase("abcd!@#$"))
	fmt.Println("Test4: ", TestLowercase("ABCD1234"))
	fmt.Println("Test5: ", TestUppercase("ABCD1234"))
	fmt.Println("Test6: ", TestUppercase("abcd!@#$"))
	fmt.Println("Test7: ", TestNumeric("abcd!@#$"))
	fmt.Println("Test8: ", TestNumeric("1ABCD!@##"))
}

func TestSpecial(s string) bool {
	var pattern = "[ !\"#$%&'()*+,\\-./:;<=>?@[\\\\\\]^_`{|}~]"
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return matched
}

func TestNumeric(s string) bool {
	var pattern = "[0-9]"
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return matched
}

func TestUppercase(s string) bool {
	var pattern = "[A-Z]"
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return matched
}

func TestLowercase(s string) bool {
	var pattern = "[a-z]"
	matched, err := regexp.Match(pattern, []byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return matched
}
