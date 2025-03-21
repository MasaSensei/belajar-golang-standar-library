package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (u UserSlice) Len() int {
	return len(u)
}

func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	users := []User{
		{Name: "Budi", Age: 22},
		{Name: "Eko", Age: 20},
		{Name: "Kurniawan", Age: 21},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
