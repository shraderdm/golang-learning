package main

import "fmt"
import "os"

func main() {
	family := []string{"Papa", "Mommy", "Ethan", "Elsie", "Sam"}
	family = append(family, "New Puppy")
	family = append(family, os.Args[1])
	for _, member := range family {
		fmt.Println(member)
	}
}
