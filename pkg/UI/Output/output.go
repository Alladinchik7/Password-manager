package output

import (
	"fmt"
	password "password-manager/internal/Password"
	ui "password-manager/pkg/UI"

	"github.com/cheynewallace/tabby"
)

func ShowMainMenu() {
	ui.ClearScreen()
	fmt.Println("==========================================")
	fmt.Println("            Password Manager              ")
	fmt.Println("==========================================")
	fmt.Println(`
		1. Generate new password
		2. Add new password
		3. Get password
		4. List all passwords
		5. Update password
		6. Delete password
		7. List categories
		8. Show password statistics
		9. Find duplicate passwords
		0. Exit
	`)
	fmt.Println("==========================================")
}

func PrintPasswordList(passwords []password.Password) {
	ui.ClearScreen()

	fmt.Println("=== Password list ===")
	t := tabby.New()
	t.AddHeader("Name", "Category", "Created", "Last Modified")
	for _, v := range passwords {
		t.AddLine(v.Name, v.Category, v.CreateAt, v.LastModified)
	}
	t.Print()

	ui.WaitForEnter()
}

func ShowPasswordDetails(password password.Password) {
	ui.ClearScreen()

	fmt.Println("=== Password details ===")
	fmt.Printf("Service: %s\n", password.Name)
	fmt.Printf("Category: %s\n", password.Category)
	fmt.Printf("Password: %s\n", password.Value)
	fmt.Println("Created: ", password.CreateAt)
	fmt.Println("Last Modified: ", password.LastModified)

	ui.WaitForEnter()
}
