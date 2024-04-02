package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/term"
)

func main() {
	var value [9]int = [9]int{}
	var activity [9]bool = [9]bool{false, false, false, false, false, false, false, true, false}
	check := 1
	for {
		rendering(value, activity)
		win, winner := win_check(value)
		if win {
			var winner_char string
			if winner == 1 {
				winner_char = "O"
			} else if winner == 2 {
				winner_char = "X"
			} else {
				winner_char = "XO"
			}
			fmt.Println("WINNER:", winner_char)
			break
		}
		activity, value, check = move(activity, value, check)
	}
}

func rendering(value [9]int, activity [9]bool) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	base_field := `
	┌──────────┐ ┌──────────┐ ┌──────────┐
	│          │ │          │ │          │
	│          │ │          │ │          │
	│          │ │          │ │          │
	│          │ │          │ │          │
	└──────────┘ └──────────┘ └──────────┘
	┌──────────┐ ┌──────────┐ ┌──────────┐
	│          │ │          │ │          │
	│          │ │          │ │          │
	│          │ │          │ │          │
	│          │ │          │ │          │
	└──────────┘ └──────────┘ └──────────┘
	┌──────────┐ ┌──────────┐ ┌──────────┐
	│          │ │          │ │          │ 
	│          │ │          │ │          │
	│          │ │          │ │          │
	│          │ │          │ │          │
	└──────────┘ └──────────┘ └──────────┘
	`
	var subfields_for_char [9]int = [9]int{
		87, 100, 113, 327, 340, 353, 568, 581, 594,
	}
	var subfields_for_frame [9]int = [9]int{
		43, 56, 69, 283, 296, 309, 523, 536, 549,
	}
	run := []rune(base_field)

	for i := 0; i < len(value); i++ {
		x := subfields_for_char[i]
		if value[i] == 1 {
			run[x] = rune('╭')
			run[x+1] = rune('╮')
			run[x+40] = rune('╰')
			run[x+41] = rune('╯')
		} else if value[i] == 2 {
			run[x] = rune('╲')
			run[x+1] = rune('╱')
			run[x+40] = rune('╱')
			run[x+41] = rune('╲')
		}

		if activity[i] {
			x := subfields_for_frame[i]
			run[x] = rune('┌')
			run[x+9] = rune('┐')
			if i > 5 {
				run[x+121] = rune('└')
				run[x+130] = rune('┘')
			} else {
				run[x+120] = rune('└')
				run[x+129] = rune('┘')
			}
		}
	}
	fmt.Println(string(run))
}

func move(activity [9]bool, value [9]int, check int) ([9]bool, [9]int, int) {
	t, _ := term.Open("/dev/tty")
	term.CBreakMode(t)
	var key []byte = make([]byte, 1)
	t.Read(key)

	var head int
	for i := 0; i < len(activity); i++ {
		if activity[i] {
			head = i
			break
		}
	}

	result_value := value
	fmt.Println(key[0], "---", key)
	if (key[0] == 'a' || key[0] == 68) && head > 0 {
		head = head - 1
	} else if (key[0] == 'd' || key[0] == 67) && head < 8 {
		head = head + 1
	} else if (key[0] == 'w' || key[0] == 65) && head > 2 {
		head = head - 3
	} else if (key[0] == 's' || key[0] == 66) && head < 6 {
		head = head + 3
	} else if key[0] == ' ' || key[0] == 10 {
		for i := 0; i < 9; i++ {
			result_value[i] = value[i]
			if i == head && value[i] == 0 {
				if check == 1 {
					result_value[i] = 1
					check = 2
				} else {
					result_value[i] = 2
					check = 1
				}
			}
		}
	}
	var result_activity [9]bool
	for i := 0; i < 9; i++ {
		result_activity[i] = false
		if i == head {
			result_activity[i] = true
		}
	}
	t.Close()
	return result_activity, result_value, check
}

func win_check(value [9]int) (bool, int) {
	win := false
	winner := 0

	if value[0] == value[1] && value[1] == value[2] && value[0] != 0 {
		win = true
		winner = value[0]
	} else if value[3] == value[4] && value[4] == value[5] && value[3] != 0 {
		win = true
		winner = value[3]
	} else if value[6] == value[7] && value[7] == value[8] && value[6] != 0 {
		win = true
		winner = value[6]
	} else if value[0] == value[3] && value[3] == value[6] && value[0] != 0 {
		win = true
		winner = value[0]
	} else if value[1] == value[4] && value[4] == value[7] && value[1] != 0 {
		win = true
		winner = value[1]
	} else if value[2] == value[5] && value[5] == value[8] && value[2] != 0 {
		win = true
		winner = value[2]
	} else if value[0] == value[4] && value[4] == value[8] && value[0] != 0 {
		win = true
		winner = value[0]
	} else if value[2] == value[4] && value[4] == value[6] && value[2] != 0 {
		win = true
		winner = value[0]
	} else {
		for i := 0; i < len(value); i++ {
			win = true
			if value[i] == 0 {
				win = false
				break
			}
		}
	}
	return win, winner
}
