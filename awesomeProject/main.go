package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	var a = "v[0-9]\\.[0-9]\\.[0-9]-snd"
	match, _ := regexp.MatchString(a, "v1.2.1-snd")
	fmt.Println(match)

	a = "regexp:v[0-9]\\.[0-9]\\.[0-9]-snd"
	opt := strings.SplitN(a, ":", 2)

	if len(opt) != 2 {

	}
	switch strings.ToLower(opt[0]) {
	case "regexp":
		re, err := regexp.Compile(opt[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(re)

	}
}
