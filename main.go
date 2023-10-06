package main

import "fmt"

type City struct {
	CouncilMembers []CouncilMember `json:"councilMembers"`
	CityName       string          `json:"cityName"`
	Exists         bool            `json:"exists"`
	Founder        string          `json:"founder"`
	Coins          int64           `json:"coins"`
}

func NewFounder(founder string) *City {
	return &City{Founder: founder}
}

func NewCity(cityName, founder string) *City {
	return &City{CityName: cityName}
}

type CouncilMember struct {
	Skills []string `json:"skills"`
	Name   string   `json:"name"`
	Class  string   `json:"class"`
	Age    int64    `json:"age"`
}

func main() {
	displayMenu()
}
func displayMenu() {
	var input int
	fmt.Println("---------------------------------------------Grand Go Greece-----------------------------------------------")
	fmt.Printf("\n(1) Start New Game\n(2)Load\n(3)Exit\n")
	fmt.Scanln(&input)
	switch input {
	case 1:
		newGame()
	}

}
func displayMap() {

}
func displayStats() {

}
func newGame() {
	fmt.Println("Enter your character and city name. The character will b the city's founder")

}
