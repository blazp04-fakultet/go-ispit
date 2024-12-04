package main

import (
	"blazperic/ispit/zadatci"
	"fmt"
)

func main() {
	fmt.Print("\n------------------[ZADATAK 1]----------------- \n")

	// ZADATAK 1
	var proizvodi map[string]int = map[string]int{
		"banana":  3,
		"telefon": 425,
		"kruska":  7,
		"mlijeko": 4,
		"mobitel": 140,
		"laptop":  648,
		"jabuka":  5,
		"tablet":  695,
	}
	zadatci.Zadatak1(proizvodi)

	fmt.Print("\n------------------[ZADATAK 2]----------------- \n")

	// ZADATAK 2

	var employees = []zadatci.Employee{
		{
			Name:     "Blaž Perić",
			Position: "Student",
			Tasks: []string{
				"Položiti ispit iz GoLanga",
				"Predati papire za dekanovu nagradu",
				"Predati papire za stipendiju",
			},
		},
		{
			Name:     "Marko Markovic",
			Position: "Dizajner",
			Tasks: []string{
				"Izraditi prototip za novu aplikaiciju",
				"Uraditi mockup za novu aplikaciju",
				"Izraditi ikonicu za GO projekt",
			},
		},
	}
	employees[0].PrintAllTasks()

	employees[0].AddTask("Početi učiti KDM")
	employees[0].AddTask("Vježbati za kolokvi iz baza podataka")
	employees[0].AddTask("Okupit se sa ekipom za hackatjon")

	employees[0].PrintAllTasks()
	employees[0].CompleateTask()
	employees[0].PrintAllTasks()

	employees[1].PrintAllTasks()
	employees[1].AddTask("Otići na zornicu")
	employees[1].CompleateTask()

	employees[1].PrintAllTasks()

}
