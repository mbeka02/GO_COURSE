package main

import "fmt"



func main(){
	animal:=struct {
		species string
		gender string
	
	}{
		species:"homo sapien",
		gender:"male",
	}
	
  fmt.Println(animal)

}