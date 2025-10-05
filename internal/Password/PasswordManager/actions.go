package pmanage

import (
	"fmt"
	"log"
	password "password-manager/internal/Password"
	errorConst "password-manager/pkg/Error"
	"time"
)

func (pm *PasswordManager) SavePassword(pass password.Password) error {
	if !pm.isInitialized {
		return fmt.Errorf(errorConst.PassUninit)
	}

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return err
	}

	for _, v := range passwords {
		if v.Name == pass.Name {
			return fmt.Errorf("password with that name already exists")
		}
	}

	if err := pm.Passwords.NewPassword(pass); err != nil {
		return err
	}

	return nil
}

func (pm *PasswordManager) UpdatePassword(name, newValue string) error {
	if !pm.isInitialized {
		return fmt.Errorf(errorConst.PassUninit)
	}

	pass, err := pm.GetPassword(name)
	if err != nil {
		return err
	}

	if err := pm.CheckPasswordStrength(newValue); err != nil {
		return err
	}
	log.Println("New password valid ✅")

	tx := pm.Passwords.DB.Model(pass).Where("name = ?", name).Update("value", newValue).Update("last_modified", time.Now())
	if tx.Error != nil {
		return fmt.Errorf("failed update password: %v", tx.Error)
	}

	return nil
}

func (pm *PasswordManager) DeletePassword(name string) error {
	if !pm.isInitialized {
		return fmt.Errorf(errorConst.PassUninit)
	}

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return err
	}

	flag := false
	for _, v := range passwords {
		if v.Name == name {
			flag = true
		}
	}

	if !flag {
		return fmt.Errorf("password not found")
	}

	if tx := pm.Passwords.DB.Delete(passwords, name); tx.Error != nil {
		return fmt.Errorf("failed password deleted: %v", tx.Error)
	}

	return nil
}
