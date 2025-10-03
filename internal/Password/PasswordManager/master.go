package pmanage

import (
	"fmt"
	"strings"
)

func (pm *PasswordManager) SetMasterPassword(masterPassword string) error {
	err := pm.CheckPasswordStrength(masterPassword)
	if err != nil {
		return err
	}

	bt := make([]byte, 32)
	copy(bt, []byte(masterPassword))

	pm.masterKey = bt
	pm.isInitialized = true

	return nil
}

func (pm *PasswordManager) CheckPasswordStrength(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password short")
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*", char):
			hasSpecial = true
		}
	}

	if !hasUpper || !hasLower || !hasNumber || !hasSpecial {
		return fmt.Errorf("invalid password")
	}

	return nil
}
