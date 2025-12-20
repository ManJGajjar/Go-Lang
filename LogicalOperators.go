package main

import "fmt"

func LogicalOperators() {
	var p int = 10
	var q int = 5

	if (p!=q) && (p>q) {
		fmt.Println("p is not equal to q and p is greater than q")
	}
	if (p==10) || (q==10) {
		fmt.Println("Either p is 10 or q is 10")
	}
	if(!(p==q)){ 
        fmt.Println("True")
    }
}