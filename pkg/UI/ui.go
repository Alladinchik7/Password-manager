package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin": // Unix-like systems
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		// Альтернативный способ для неподдерживаемых систем
		print("\033[2J\033[H") // ANSI escape codes
	}
}
func ShowSuccess(message string) {
	fmt.Printf("%s%s✓ Success:%s%s\n", colorReset, colorGreen, colorReset, message)
}

func ShowError(message error) {
	fmt.Printf("%s%sErorr:%s%v\n", colorReset, colorRed, colorReset, message)
}

func ShowInfo(message string) {
	fmt.Printf("%s%s→ Info:%s%s\n", colorReset, colorYellow, colorReset, message)
}

func WaitForEnter() {
	fmt.Print("Press Enter to continue...")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// Убираем символы новой строки
	input = strings.TrimSpace(input)
	// Игнорируем все что было введено, важно только нажатие Enter
}
