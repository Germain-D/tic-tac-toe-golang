package main

import (
	"fmt"

	"github.com/fatih/color"

	"math/rand"
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
	fmt.Println()
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
	if terrain[3] == terrain[5] && terrain[3] == terrain[7] && terrain[3] != " " {
		affichageterrain(terrain)
		fmt.Println("Le joueur", terrain[3], "a gagné")
		return 10
	}

	return tour + 1
}

func CheckWin(terrain map[int]string, ia_state string) bool {
	for i := 0; i < 3; i++ {
		if terrain[i] == terrain[i+1] && terrain[i] == terrain[i+2] && terrain[i] != " " && terrain[i] == ia_state {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if terrain[i] == terrain[i+3] && terrain[i] == terrain[i+6] && terrain[i] != " " && terrain[i] == ia_state {
			return true
		}
	}
	if terrain[0] == terrain[4] && terrain[0] == terrain[8] && terrain[0] != " " && terrain[0] == ia_state {
		return true
	}
	if terrain[2] == terrain[4] && terrain[2] == terrain[6] && terrain[2] != " " && terrain[2] == ia_state {
		return true
	}

	return false
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

func partietype() int {
	fmt.Println("1 : Joueur vs Joueur")
	fmt.Println("2 : Joueur vs IA (easy)")
	fmt.Println("3 : Joueur vs IA (hard)")

	var choix int
	bonchoix := false

	for !bonchoix {
		fmt.Println("Choisissez un mode de jeu : ")
		fmt.Scanln(&choix)
		if choix == 1 || choix == 2 || choix == 3 {
			bonchoix = true
		} else {
			fmt.Println("Choix incorrect")
		}
	}
	return choix
}

func ia_easy(terrain map[int]string) int {
	var choix int
	bonchoix := false

	for !bonchoix {
		choix = rand.Intn(9) + 1
		if terrain[choix] == " " {
			bonchoix = true
		}
	}
	return choix
}

func availableMoves(terrain map[int]string) []int {
	var available []int
	for i := 1; i < 10; i++ {
		if terrain[i] == " " {
			available = append(available, i)
		}
	}
	return available
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimax(terrain map[int]string, ia_state string) int {
	if CheckWin(terrain, ia_state) {
		if ia_state == "X" {
			return 1
		} else {
			return -1
		}
	}
	if ia_state == "X" {
		bestValue := -1
		for _, move := range availableMoves(terrain) {
			new_terrain := make(map[int]string, len(terrain))
			for k, v := range terrain {
				new_terrain[k] = v
			}
			new_terrain[move] = "X"
			value := minimax(new_terrain, "O")
			bestValue = max(value, bestValue)
		}
		return bestValue
	} else {
		bestValue := 1
		for _, move := range availableMoves(terrain) {
			new_terrain := make(map[int]string, len(terrain))
			for k, v := range terrain {
				new_terrain[k] = v
			}
			new_terrain[move] = "O"
			value := minimax(new_terrain, "X")
			bestValue = min(value, bestValue)
		}
		return bestValue
	}
}

func bestMove(terrain map[int]string, ia_state string) int {
	available := availableMoves(terrain)

	var bestmove int
	best := -1
	fmt.Println("available")
	fmt.Println(available)

	for _, move := range available {
		new_terrain := make(map[int]string, len(terrain))
		for k, v := range terrain {
			new_terrain[k] = v
		}
		new_terrain[move] = ia_state
		value := minimax(new_terrain, ia_state)
		if value > best {
			best = value
			bestmove = move
		}
	}
	fmt.Println("bestmove")
	fmt.Println(bestmove)
	return bestmove

}

func ia_hard(terrain map[int]string, qui_commence int, ia_state string) int {
	if qui_commence == 1 {
		if ia_state == "X" {
			return bestMove(terrain, "X")
		} else {
			return bestMove(terrain, "O")
		}
	} else {
		if ia_state == "X" {
			return bestMove(terrain, "O")
		} else {
			return bestMove(terrain, "X")
		}
	}
}

func howstart() int {
	return rand.Intn(2) + 1
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

	partie := partietype()

	if partie == 1 {
		for tour < 9 {
			affichageterrain(terrain)
			if tour%2 == 0 {
				terrain[choix(terrain)] = "X"
			} else {
				terrain[choix(terrain)] = "O"
			}
			tour = winner(terrain, tour)
		}
	} else if partie == 2 {
		qui_commence := howstart()
		if qui_commence == 1 {
			fmt.Println("Le joueur commence")
		} else {
			fmt.Println("L'IA commence")
		}
		fmt.Println(qui_commence)
		for tour < 9 {
			affichageterrain(terrain)
			if tour%2 == 0 {
				if qui_commence == 1 {
					terrain[choix(terrain)] = "X"
				} else {
					terrain[ia_easy(terrain)] = "X"
				}
			} else {
				if qui_commence == 1 {
					terrain[ia_easy(terrain)] = "O"
				} else {
					terrain[choix(terrain)] = "O"
				}
			}
			tour = winner(terrain, tour)
		}

	} else {
		qui_commence := howstart()
		var ia_state string
		if qui_commence == 1 {
			fmt.Println("Le joueur commence")
			ia_state = "O"
		} else {
			fmt.Println("L'IA commence")
			ia_state = "X"
		}
		fmt.Println(qui_commence)
		for tour < 9 {
			affichageterrain(terrain)
			if tour%2 == 0 {
				if qui_commence == 1 {
					terrain[choix(terrain)] = "X"
				} else {
					terrain[ia_hard(terrain, qui_commence, ia_state)] = "X"
				}
			} else {
				if qui_commence == 1 {
					terrain[ia_hard(terrain, qui_commence, ia_state)] = "O"
				} else {
					terrain[choix(terrain)] = "O"
				}
			}
			tour = winner(terrain, tour)
		}
	}
	if tour == 9 {
		affichageterrain(terrain)
		fmt.Println("Match nul")
	}
}
