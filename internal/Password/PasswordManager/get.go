package pmanage

import (
	"fmt"
	password "password-manager/internal/Password"
	errorConst "password-manager/pkg/Error"
)

func (pm *PasswordManager) GetPassword(name string) (password.Password, error) {
	if !pm.isInitialized {
		return password.Password{}, fmt.Errorf(errorConst.PassUninit)
	}

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return password.Password{}, err
	}

	for _, v := range passwords {
		if v.Name == name {
			return v, nil
		}
	}

	return password.Password{}, fmt.Errorf("couldn't find the password")
}

func (pm *PasswordManager) GetPasswordsByCategory(category string) ([]password.Password, error) {
	var passwords, getPasswords []password.Password

	err := pm.Passwords.ReadPassword(&passwords)
	if err != nil {
		return []password.Password{}, err
	}

	for _, v := range passwords {
		if v.Category == category {
			getPasswords = append(getPasswords, v)
		}
	}

	return getPasswords, nil
}

func (pm *PasswordManager) GetPasswordStats() (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return nil, err
	}

	if passwords == nil {
		return nil, fmt.Errorf(errorConst.PassNotFound)
	}

	stats["total_passwords"] = len(passwords)

	categories, err := pm.ListCategories()
	if err != nil {
		return nil, err
	}

	pass, err := pm.ListPassword()
	if err != nil {
		return nil, err
	}

	countCategory := make(map[string]int)
	for _, category := range categories {
		var passwords []password.Password
		for _, p := range pass {
			if p.Category == category {
				passwords = append(passwords, p)
			}
		}
		count := len(passwords)
		countCategory[category] = count
	}
	stats["categories"] = countCategory

	minCreat := passwords[0].CreateAt
	maxCreat := passwords[0].CreateAt
	for _, v := range passwords {
		if v.CreateAt.Before(minCreat) {
			minCreat = v.CreateAt
		}
		if v.CreateAt.After(maxCreat) {
			maxCreat = v.CreateAt
		}
	}

	stats["oldest_password_date"] = minCreat
	stats["newest_password_date"] = maxCreat

	return stats, nil
}
