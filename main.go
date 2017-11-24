package main

import (
	"fmt"

	gameplate "github.com/tannineo/2048-cli/gameplate"
)

const (
	LOGO = `
██████╗  ██████╗ ██╗  ██╗ █████╗        █████╗██╗     ████╗ 
╚════██╗██╔═████╗██║  ██║██╔══██╗      ██╔═══╝██║      ██╔╝ 
 █████╔╝██║██╔██║███████║╚█████╔╝█████╗██║    ██║      ██║  
██╔═══╝ ████╔╝██║╚════██║██╔══██╗╚════╝██║    ██║      ██║  
███████╗╚██████╔╝     ██║╚█████╔╝       █████╗██████╗ ████╗ 
╚══════╝ ╚═════╝      ╚═╝ ╚════╝        ╚════╝╚═════╝ ╚═══╝ 

See https://github.com/tannineo/2048-cli
Input 'h' for help. Only read the first char.
`
	BASIC_INSTRUCTION = `
4: play or restart a 2048 game in classic 4*4 girds
H, h: help message (what you're reading)
Q, q: quit the game`
)

func main() {

	fmt.Println(LOGO)

	inGame := false
	var gp gameplate.Gameplate

	for {
		var inputCmd string
		if _, err := fmt.Scanln(&inputCmd); err != nil {
			fmt.Println(err)
			return
		}
		inputRunes := []rune(inputCmd)
		switch inputRunes[0] {
		// byebye
		case 'q':
			fallthrough
		case 'Q':
			fmt.Println("Byebye~")
			return
		// 帮助
		case 'h':
			fallthrough
		case 'H':
			fmt.Println(BASIC_INSTRUCTION)
			if inGame {
				fmt.Println(gp.Rules())
			}

		// grid44
		case '4':
			inGame = true
			gp = &gameplate.Grid44{}
			gp.NewGame()
			fmt.Println(gp.Print())

		default:
			if inGame {
				if v, ok := gp.AvailableMoves()[inputRunes[0]]; ok {
					if gp.Move(v) {
						gp.GenerateNewCells()
					}
					fmt.Println(gp.Print())
					if gp.IsGameOver() {
						fmt.Println("~ Game Over ~")
					}
				} else {
					// invalid input
					fmt.Println("Invalid input.")
				}
			} else {
				// invalid input
				fmt.Println("Invalid input.")
			}
		}
	}

}
