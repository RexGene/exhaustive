package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func each(base string, list [][]string) {
	if len(list) == 0 {
		fmt.Println(base)
		return
	}

	ls := list[0]
	for _, str := range ls {
		newBase := fmt.Sprintf("%s%s ", base, str)
		each(newBase, list[1:])
	}
}

func main() {
	dataList := [][]string{}
	re, _ := regexp.Compile("[\\s]+")
	for _, arg := range os.Args[1:] {
		arg = re.ReplaceAllString(arg, " ")
		arg = strings.TrimSpace(arg)
		dataList = append(dataList, strings.Split(arg, " "))
	}

	each("", dataList)
}
