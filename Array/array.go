package main

import "fmt"

func main() {
	var data int
	var index int
	var input string

	slice := []int{}


for { 
	fmt.Println("Choose operations")
	fmt.Println("Insert")
	fmt.Println("Delete")
	fmt.Println("View")
	fmt.Println("Exit")
	fmt.Println("___________")

	fmt.Scanln(&input)

	if input=="insert"{
		fmt.Println("Enter data to add")
		fmt.Scanln(&data)

		slice=append(slice, data)
	}else if input=="delete"{

		fmt.Println("Enter the index to delete")
		fmt.Scanln(&index)

		if index>len(slice){
			fmt.Println("Enter a valid index")
		}

		slice=append(slice[:index],slice[index+1:]...)

	}else if input=="view"{

		fmt.Println("Array are ",slice)

	}else{
		return
	}


	}

}