package handle

import (
	"fmt"
	password "password-manager/internal/Password"
	pmanage "password-manager/internal/Password/PasswordManager"
	ui "password-manager/pkg/UI"
	input "password-manager/pkg/UI/Input"
	output "password-manager/pkg/UI/Output"
	"strconv"
	"time"
)

func HandlePasswordGeneration(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	input, err := input.ReadUserInput("Please enter lenght password:")
	if err != nil {
		return err
	}

	len, err := strconv.Atoi(input)
	if err != nil {
		return err
	}

	pass, err := pm.GeneratePassword(len)
	if err != nil {
		return err
	}

	ui.ShowSuccess("Password generate successfully:")
	fmt.Println("Password: ", pass)
	ui.WaitForEnter()

	return nil
}

func HandlePasswordAdd(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	inputServis, err := input.ReadUserInput("Please enter service name: ")
	if err != nil {
		return err
	}

	inputPass, err := input.ReadUserInput("Enter new password (Or Enter for generate): ")
	if err != nil {
		return err
	}

	if inputPass != "" {
		err := pm.CheckPasswordStrength(inputPass)
		if err != nil {
			return err
		}
	} else {
		err := HandlePasswordGeneration(pm)
		if err != nil {
			return err
		}
	}

	inputCategory, err := input.ReadUserInput("Please enter category: ")
	if err != nil {
		return err
	}

	password := password.Password{
		Name:     inputServis,
		Value:    inputPass,
		Category: inputCategory,
	}
	if err := pm.SavePassword(password); err != nil {
		return err
	}

	ui.ShowSuccess("Password saved successfully")
	ui.WaitForEnter()

	return nil
}

func HandlePasswordStats(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	stats, err := pm.GetPasswordStats()
	if err != nil {
		return err
	}

	// Выводим статистику в читаемом виде
	ui.ShowSuccess("Total statistics:\n")
	fmt.Printf("⚡ Total passwords: %d\n", stats["total_passwords"])

	fmt.Printf("\n📂 Distribution by categories:\n")
	if categories, ok := stats["categories"].(map[string]int); ok {
		for category, count := range categories {
			fmt.Printf("   • %-15s: %d\n", category, count)
		}
	}

	if oldestDate, ok := stats["oldest_password_date"].(time.Time); ok {
		fmt.Printf("\n🕒 Time characteristics:\n")
		fmt.Printf("   • Oldest: %s\n", oldestDate.Format("2006-01-02"))
		if newestDate, ok := stats["newest_password_date"].(time.Time); ok {
			fmt.Printf("   • Newest: %s\n", newestDate.Format("2006-01-02"))
		}
	}

	ui.WaitForEnter()

	return nil
}

func HandleFindDuplication(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	// Ищем дубликаты
	duplicates, err := pm.FindDuplicatePasswords()
	if err != nil {
		return err
	}

	if len(duplicates) == 0 {
		fmt.Println("Duplicates not found")
	} else {
		fmt.Printf("\nFound duplicates:\n")
		for password, services := range duplicates {
			fmt.Printf("\nPassword '%s' is used in the following services:\n", password)
			for _, service := range services {
				fmt.Printf("- %s\n", service)
			}
		}
	}

	ui.WaitForEnter()

	return nil
}

func HandlePasswordDelete(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	service, err := input.ReadUserInput("Enter service name: ")
	if err != nil {
		return err
	}

	if err := pm.DeletePassword(service); err != nil {
		return err
	}
	ui.ShowSuccess("Password delete successfully")
	ui.WaitForEnter()

	return nil
}

func HandlePasswordSearch(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	service, err := input.ReadUserInput("Enter service name: ")
	if err != nil {
		return err
	}

	pass, err := pm.GetPassword(service)
	if err != nil {
		return err
	}

	output.ShowPasswordDetails(pass)
	ui.WaitForEnter()

	return nil
}

func HandlePasswordUpdate(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()

	service, err := input.ReadUserInput("Enter service name: ")
	if err != nil {
		return err
	}

	inputPass, err := input.ReadUserInput("Enter new password (Or Enter for generate): ")
	if err != nil {
		return err
	}

	if inputPass != "" {
		err := pm.CheckPasswordStrength(inputPass)
		if err != nil {
			return err
		}
	} else {
		err := HandlePasswordGeneration(pm)
		if err != nil {
			return err
		}
	}

	if err := pm.UpdatePassword(service, inputPass); err != nil {
		return err
	}

	ui.ShowSuccess("Password saved successfully")
	ui.WaitForEnter()

	return nil
}

func HandleExitAndSave(pm *pmanage.PasswordManager) error {
	ui.ClearScreen()
	ui.ShowInfo("Going save passwords...")

	if err := pm.SaveToFile(); err != nil {
		return err
	}

	ui.ShowSuccess("Passwords saved to file")

	fmt.Println("Bye, bye...")

	return nil
}
