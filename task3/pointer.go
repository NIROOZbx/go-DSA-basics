package main

import "fmt"

type Student struct {
	Name string
	Age  int
}



func main() {
	s1 := Student{Name: "nirooz", Age: 54}
	s2 :=&Student{Name: "harry",Age: 21}
	s1.Name="johnn"
	s2.Age=23
	fmt.Println(s1,s2)


}
