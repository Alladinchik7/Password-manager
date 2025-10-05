package pmanage

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"os"
	password "password-manager/internal/Password"
	errorConst "password-manager/pkg/Error"
	ui "password-manager/pkg/UI"
)

func (pm *PasswordManager) SaveToFile() error {
	if !pm.isInitialized {
		return fmt.Errorf(errorConst.PassUninit)
	}

	var passwords []password.Password
	if err := pm.Passwords.ReadPassword(&passwords); err != nil {
		return err
	}

	data, err := json.Marshal(passwords)
	if err != nil {
		return fmt.Errorf("failed serialize password from json: %v", err)
	}

	block, err := aes.NewCipher(pm.masterKey)
	if err != nil {
		return fmt.Errorf("failed creat new blok: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed create GCM: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed generate nonce: %v", err)
	}

	// Шифруем (GCM не требует padding)
	encryptedData := gcm.Seal(nonce, nonce, data, nil)

	file, err := os.Create(pm.filePath)
	if err != nil {
		return fmt.Errorf("failed create file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(encryptedData); err != nil {
		return fmt.Errorf("failed write data: %v", err)
	}

	return nil
}

func (pm *PasswordManager) LoadFromFile() error {
	if !pm.isInitialized {
		return fmt.Errorf(errorConst.PassUninit)
	}

	// Проверяем существование и размер файла
	fileInfo, err := os.Stat(pm.filePath)
	if os.IsNotExist(err) {
		ui.ShowInfo("✓ No saved passwords found, starting fresh")
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to check file: %v", err)
	}

	// Проверяем что файл не пустой
	if fileInfo.Size() == 0 {
		ui.ShowInfo("⚠ Password file is empty, ignoring")
		return nil
	}

	// Проверяем минимальный размер для GCM
	if fileInfo.Size() < 28 { // nonce(12) + min ciphertext(16)
		return fmt.Errorf("encrypted file is too small or corrupted")
	}

	encryptedData, err := os.ReadFile(pm.filePath)
	if err != nil {
		return fmt.Errorf("failed read file: %v", err)
	}

	block, err := aes.NewCipher(pm.masterKey)
	if err != nil {
		return fmt.Errorf("failed create cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed create GCM: %v", err)
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return fmt.Errorf("encrypted data too short")
	}

	nonce, ciphertext := encryptedData[:nonceSize], encryptedData[nonceSize:]
	decryptedData, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("failed decrypt: %v", err)
	}

	var passwords []password.Password
	if err := json.Unmarshal(decryptedData, &passwords); err != nil {
		return fmt.Errorf("failed parse JSON: %v", err)
	}

	for _, v := range passwords {
		if err := pm.SavePassword(v); err != nil {
			return fmt.Errorf("failed save password: %v", err)
		}
	}

	return nil
}
