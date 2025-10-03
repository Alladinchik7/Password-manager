package pmanage

import (
	"crypto/rand"
	"fmt"
	"math/big"
	password "password-manager/internal/Password"
)

func (pm *PasswordManager) GeneratePassword(lenght int) (string, error) {
	if lenght < 8 {
		return "", fmt.Errorf("lenght for new password short")
	}

	characterSets := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*;:?"
	password := make([]byte, lenght)

	for i := 0; i < lenght; i++ {
		randNum, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSets))))
		if err != nil {
			return "", err
		}

		password[i] = characterSets[randNum.Int64()]
	}

	return string(password), nil
}

func (pm *PasswordManager) ListPassword() ([]password.Password, error) {
	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return []password.Password{}, err
	}

	return passwords, nil
}

func (pm *PasswordManager) FindDuplicatePasswords() (map[string][]string, error) {
	kash := make(map[string][]string)

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return nil, err
	}

	for _, v := range passwords {
		kash[v.Value] = append(kash[v.Value], v.Name)
	}

	result := make(map[string][]string)
	for pass, names := range kash {
		if len(names) > 1 {
			result[pass] = names
		}
	}

	return result, nil
}

func (pm *PasswordManager) ListCategories() ([]string, error) {
	categories := make(map[string]bool)

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return []string{}, err
	}

	for _, v := range passwords {
		if v.Value != "" {
			categories[v.Category] = true
		}
	}

	var result []string
	for categiry := range categories {
		result = append(result, categiry)
	}

	return result, nil
}
