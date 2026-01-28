package main

import (
   "fmt"
   "regexp"
)

func Regex() {
   pattern := "Hamilton"
   text := "SF-26 Lewis Hamilton"
   matched, _ := regexp.MatchString(pattern, text)
   fmt.Println(matched)
}