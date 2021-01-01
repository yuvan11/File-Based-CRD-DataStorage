package main

import (
	crd "File-Based-CRD/src/CRDFunctions"
	"fmt"
	"os"
)


func main() {

	crd.CreateFile()

	for {

		var Option int

		fmt.Println("Choose an outcome from the list below")
		fmt.Println("1.Create")
		fmt.Println("2.Read")
		fmt.Println("3.Delete")
		fmt.Println("4.Display All")
		fmt.Println("5.Exit")
		fmt.Print(">")
		fmt.Scanf("%d", &Option)
		switch Option {
		case 1:
			crd.Create()
		case 2:
			crd.Read()
		case 3:
			crd.Delete()
		case 4:
			crd.DisplayAll()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Choose the correct option between 1 - 5")

		}
	}

}
