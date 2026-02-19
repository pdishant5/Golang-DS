package main

import (
	"fmt"

	"authenticator/user"
	"authenticator/usermap"
)

func main() {
	userMap := usermap.New()

	u1 := user.New("pdishant5", "DishantP@12345")
	err := userMap.Add(u1)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	u2 := user.New("pdishant5", "DishantP@12345")
	err = userMap.Add(u2)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	u3 := user.New("pdishant55", "DishantP@12345")
	err = userMap.Add(u3)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	err = userMap.Authenticate("pdishant5", "DishantP@12345")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	err = userMap.Authenticate("pdishant5", "DishantP@1234")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	err = userMap.Authenticate("dishant555", "DishantP@12345")
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
