package main

// import "greeting"

// import (
// 	"calc"
// 	"fmt"
// )

// import (
// 	"fmt"
// 	"keyboard"
// 	"log"
// )

import (
	"dates"
	"fmt"
)

// func main() {
// 	greeting.Hello()
// 	greeting.Hi()
// }

// func main() {
// 	fmt.Println(calc.Add(1, 2))
// 	fmt.Println(calc.Substract(7, 3))
// }

// func main() {
// 	fmt.Print("Enter a temprature in Fahrenheit: ")
// 	fahrenheit, err := keyboard.GetFloat()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	celsius := (fahrenheit - 32) * 5 / 9
// 	fmt.Printf("%0.2f degrees Celsius\n", celsius)
// }

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with follow-up in", days+dates.DaysInWeek, "days")
}
