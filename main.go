package main

import (
	"bufio"
	"fmt"
	"github.com/Delta456/box-cli-maker"
	"github.com/fatih/color"
	"github.com/guptarohit/asciigraph"
	"log"
	"os"
	"time"
)

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
	Box := box.New(box.Config{Px: 2, Py: 2, Type: "Double", Color: "Cyan", ContentAlign: "Center", TitlePos: "Top"})
	Box.Print("Grand Go Greece ", "Grand Go Greece")
	fmt.Printf("\n(1) Start New Game\n(2) Load\n(3) Exit\n")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
	switch input {
	case 1:
		newGame()
	}

}
func displayMap() {

}
func displayStats() {

	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	color.Blue(graph)

}
func newGame() {
	var cname, cfname string
	currentTime := time.Now()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Chose from a host of city states with hthier own stenrhg or weakneses or create youre own.\n")
	fmt.Printf("Enter your character and city name. The character will be the founder and carve thier name and resence in the citys hitorty.\n First, estabish you're city with a proper name and location :\n")
	scanner.Scan()
	if err := scanner.Err(); err != nil || scanner.Text() != "" {
		log.Fatal(err)
	}
	cname = scanner.Text()
	fmt.Printf("\nThe founder of this is :\n")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	cfname = scanner.Text()
	fmt.Printf("%s was founded by %s in %d of %dth", cfname, cname, currentTime.Day(), currentTime.Month())

}
