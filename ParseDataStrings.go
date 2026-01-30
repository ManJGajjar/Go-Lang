package main

import (
   "fmt"
   "time"
)

func PDStr() {
   v := "Thu, 05/19/11, 10:47PM"
   l := "Mon, 01/02/06, 03:04PM"
   tt, _ := time.Parse(l, v)
   fmt.Println(tt)
}