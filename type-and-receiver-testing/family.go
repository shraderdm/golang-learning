package main

import (
	"fmt"
	"os"
)

type family []string

func (f family) print() {
	for _, member := range f {
		fmt.Println(member)
	}
}

func (f family) newMember() string {
	f = append(f, os.Args[1])
	fmt.Println(os.Args[1])
	return os.Args[1]
}

func newFamily() family {
	members := family{}
	firstName := []string{"Papa", "Mommy", "Ethan", "Elsie", "Sam", "New Puppy"}
	lastName := "Shrader"

	for _, first := range firstName {
		members = append(members, first+" "+lastName)
	}
	return members
}
