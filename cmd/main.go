package main

import (
	"fmt"
	"os"
	handle "password-manager/internal/Handler"
	pmanage "password-manager/internal/Password/PasswordManager"
	"password-manager/pkg/DB/opendb"
	ui "password-manager/pkg/UI"
	input "password-manager/pkg/UI/Input"
	output "password-manager/pkg/UI/Output"
)

func main() {
	ui.ClearScreen()
	if err := opendb.Init(); err != nil {
		ui.ShowError(fmt.Errorf("❌ Database initialization failed: %v", err))
		return
	}
	ui.ShowSuccess("Database initialization ✅")

	path := "./Password-manager/cmd/manage.dat"
	pm := pmanage.NewPasswordManager(path)

	fmt.Println("=== Password Manager Initialization ===")
	fmt.Print("Enter master password: ")
	masterPassword, err := input.ReadPassword()
	if err != nil {
		ui.ShowError(err)
		return
	}

	if err := pm.SetMasterPassword(masterPassword); err != nil {
		ui.ShowError(fmt.Errorf("master error: %v", err))
		return
	}
	ui.ShowSuccess("Master password set successfully")

	if err := pm.LoadFromFile(); err != nil && !os.IsNotExist(err) {
		ui.ShowError(fmt.Errorf("file error: %v", err))
		return
	}

	ui.ShowSuccess("Password manager initialized successfully")
	ui.WaitForEnter()

	for {
		output.ShowMainMenu()
		in, err := input.ReadUserInput("Enter: ")
		if err != nil {
			ui.ShowError(err)
		}

		switch in {
		case "1":
			if err := handle.HandlePasswordGeneration(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "2":
			if err := handle.HandlePasswordAdd(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "3":
			if err := handle.HandlePasswordSearch(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "4":
			pass, err := pm.ListPassword()
			if err != nil {
				ui.ShowError(err)
				continue
			}
			output.PrintPasswordList(pass)
		case "5":
			if err := handle.HandlePasswordUpdate(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "6":
			if err := handle.HandlePasswordDelete(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "7":
			list, err := pm.ListCategories()
			if err != nil {
				ui.ShowError(err)
				continue
			}

			for _, v := range list {
				fmt.Println("Category: ", v)
			}
			ui.WaitForEnter()
		case "8":
			if err := handle.HandlePasswordStats(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "9":
			if err := handle.HandleFindDuplication(pm); err != nil {
				ui.ShowError(err)
				continue
			}
		case "0":
			if err := handle.HandleExitAndSave(pm); err != nil {
				ui.ShowError(err)
				ui.WaitForEnter()
				return
			}
			ui.ShowSuccess("Goodbye!")
			return
		default:
			ui.ShowError(fmt.Errorf("invalid choice. Please try again"))
			ui.WaitForEnter()
		}
	}
}
