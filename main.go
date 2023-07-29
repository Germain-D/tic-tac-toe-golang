package main

import (
	"fmt"

	"github.com/fatih/color"
)

func affichageterrain(terrain map[int]string) {
	for i := 1; i < 10; i++ {
		if terrain[i] == " " {
			fmt.Print(i, " ")
		}
		if terrain[i] == "X" {
			color.Set(color.FgRed)
			fmt.Print(terrain[i], " ")
			color.Unset()
		}
		if terrain[i] == "O" {
			color.Set(color.FgGreen)
			fmt.Print(terrain[i], " ")
			color.Unset()

		}
		if i%3 == 0 {
			fmt.Println()
		}

	}
}

func winner(terrain map[int]string, tour int) int {
	for i := 1; i < 10; i += 3 {
		if terrain[i] == terrain[i+1] && terrain[i] == terrain[i+2] && terrain[i] != " " {
			affichageterrain(terrain)
			fmt.Println("Le joueur", terrain[i], "a gagné")
			return 10
		}
	}
	for i := 1; i < 4; i++ {
		if terrain[i] == terrain[i+3] && terrain[i] == terrain[i+6] && terrain[i] != " " {
			affichageterrain(terrain)
			fmt.Println("Le joueur", terrain[i], "a gagné")
			return 10
		}
	}
	if terrain[1] == terrain[5] && terrain[1] == terrain[9] && terrain[1] != " " {
		affichageterrain(terrain)
		fmt.Println("Le joueur", terrain[1], "a gagné")
		return 10
	}
	return tour + 1
}

func choix(terrain map[int]string) int {
	var choix int
	fmt.Println("Choisissez une case (1-9) : ")
	bonchoix := false

	for !bonchoix {
		fmt.Scanln(&choix)
		if terrain[choix] == " " {
			bonchoix = true
		} else {
			fmt.Println("Case déjà prise")
		}
	}
	return choix
}

func main() {

	tour := 0

	var terrain = map[int]string{
		1: " ",
		2: " ",
		3: " ",
		4: " ",
		5: " ",
		6: " ",
		7: " ",
		8: " ",
		9: " ",
	}

	for tour < 9 {
		affichageterrain(terrain)
		if tour%2 == 0 {
			terrain[choix(terrain)] = "X"
		} else {
			terrain[choix(terrain)] = "O"
		}
		tour = winner(terrain, tour)
	}
}
