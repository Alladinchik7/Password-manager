package pmanage

import (
	password "password-manager/internal/Password"
	"password-manager/pkg/DB/opendb"
)

type PasswordManager struct {
	Passwords     password.DB `json:"passwords"` // Хранилище паролей
	masterKey     []byte      `json:"-"`         // Главный ключ шифрования
	filePath      string      `json:"-"`         // Путь к файлу данных
	isInitialized bool        `json:"-"`         // Флаг инициализации
}

// NewPasswordManager создает новый экземпляр PasswordManager
func NewPasswordManager(filePath string) *PasswordManager {
	return &PasswordManager{
		Passwords:     password.DB{DB: *opendb.DB},
		masterKey:     nil,
		filePath:      filePath,
		isInitialized: false,
	}
}
