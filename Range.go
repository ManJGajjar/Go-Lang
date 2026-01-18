package main

import "fmt"

func RangeExample() {
   numbers := []int{0,1,2,3,4,5,6,7,8} 

   for i:= range numbers {
      fmt.Println("Slice item",i,"is",numbers[i])
   }
  
   countryCapitalMap := map[string] string {"India":"New Delhi","Italy":"Rome","Japan":"Tokyo"}
  
   for country := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",countryCapitalMap[country])
   }
  
   for country,capital := range countryCapitalMap {
      fmt.Println("Capital of",country,"is",capital)
   }
}