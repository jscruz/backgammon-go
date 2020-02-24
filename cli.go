package main

import "fmt"

// PrintBoard prints the backgammon board in a CLI
func PrintBoard(b *Board) {
	var output string
	var redPositionUp = make([]int, 12)
	var whitePositionUp = make([]int, 12)
	var redPositionDown = make([]int, 12)
	var whitePositionDown = make([]int, 12)
	copy(redPositionUp, b.board[1][12:])
	copy(whitePositionUp, b.board[0][12:])
	copy(whitePositionDown, b.board[0][:12])
	copy(redPositionDown, b.board[1][:12])
	for i := 13; i < 25; i++ {
		output += fmt.Sprintf("%4v", i)
	}
	output += fmt.Sprintf("\n")
	for i := 12; i > 0; i-- {
		output += fmt.Sprintf("%4v", "-")
	}
	output += fmt.Sprintf("\n")

	for i := 0; i < 7; i++ {
		for j := 0; j < 12; j++ {
			if whitePositionUp[j] > 0 || redPositionUp[j] > 0 {
				if whitePositionUp[j] > 0 {
					output += fmt.Sprintf("%4v", 0)
					whitePositionUp[j]--
				}
				if redPositionUp[j] > 0 {
					output += fmt.Sprintf("%4v", "X")
					redPositionUp[j]--
				}
			} else {
				output += fmt.Sprintf("%4v", " ")
			}
		}
		output += fmt.Sprintf("\n")
	}
	output += fmt.Sprintf("   ==============================================\n")
	var lines [8]string
	for i := 0; i < 7; i++ {
		for j := 11; j >= 0; j-- {
			if whitePositionDown[j] > 0 || redPositionDown[j] > 0 {
				if whitePositionDown[j] > 0 {
					lines[i] += fmt.Sprintf("%4v", 0)
					whitePositionDown[j]--
				}
				if redPositionDown[j] > 0 {
					lines[i] += fmt.Sprintf("%4v", "X")
					redPositionDown[j]--
				}
			} else {
				lines[i] += fmt.Sprintf("%4v", " ")
			}
		}
		lines[i] += fmt.Sprintf("\n")
	}

	for i := 6; i >= 0; i-- {
		output += fmt.Sprintf("%s", lines[i])
	}
	for i := 12; i > 0; i-- {
		output += fmt.Sprintf("%4v", "-")
	}
	output += fmt.Sprintf("\n")
	for i := 12; i > 0; i-- {
		output += fmt.Sprintf("%4v", i)
	}
	output += fmt.Sprintf("\n")
	fmt.Printf(output)
}
