# Password Manager 🔐

Простой и безопасный менеджер паролей с шифрованием, написанный на Go.

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![License](https://img.shields.io/badge/License-MIT-green)
![Status](https://img.shields.io/badge/Status-Stable-brightgreen)

## 📖 Описание

Password Manager - это консольное приложение для безопасного хранения и управления паролями. Все данные шифруются с использованием алгоритма AES-GCM перед сохранением в файл.

## ✨ Основные возможности

- 🔒 **Безопасное хранение** - все пароли шифруются перед сохранением
- 🎲 **Генератор паролей** - создание надежных случайных паролей
- 📊 **Статистика** - анализ ваших паролей по категориям
- 🔍 **Поиск дубликатов** - обнаружение повторно используемых паролей
- 🗂️ **Категории** - организация паролей по категориям
- 💾 **Автосохранение** - автоматическое сохранение при выходе
- 🛡️ **Проверка сложности** - валидация надежности паролей

## 🎥 Демонстрация

```mardown
==========================================
            Password Manager              
==========================================

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

==========================================
Enter: 
```

## ⚙️ Требования к системе

- **Go** 1.21 или выше
- **PostgreSQL** 12 или выше
- Операционная система: Windows, Linux, macOS

## 🚀 Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/your-username/password-manager.git
cd password-manager
```

### 2. Настройка базы данных

Создайте файл `.env` в корне проекта:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=password_manager
DB_SSLMODE=disable
```

### 3. Установка зависимостей

```bash
go mod tidy
```

### 4. Сборка и запуск

```bash
# Сборка проекта
make build

# Сборка для разработки
make dev-build

# Запуск
make run

# Тестирование
make test

# Форматирование кода
make fmt

# Очистка
make clean

# Установка
make install

# Показать все команды
make help
```

## 💻 Примеры использования

### Генерация пароля

```bash
# Выберите опцию 1 в меню
Please enter length password: 16
✓ Success: Password generate successfully:
Password: aB3#kL9$mN2@pQ1!
```

### Добавление нового пароля

```bash
# Выберите опцию 2 в меню
Please enter service name: github
Enter new password (Or Enter for generate): 
Please enter category: Development
✓ Success: Password saved successfully
```

### Просмотр статистики

```bash
# Выберите опцию 8 в меню
✓ Success: Total statistics:

⚡ Total passwords: 15

📂 Distribution by categories:
   • Development    : 8
   • Social         : 5
   • Work           : 2

🕒 Time characteristics:
   • Oldest: 2024-01-15
   • Newest: 2024-03-20
```

## 🗂️ Описание ключевых компонентов

### База данных

Использует PostgreSQL с GORM для ORM. Все пароли хранятся в зашифрованном виде.

### Безопасность

- Мастер-пароль используется для генерации ключа шифрования
- AES-256-GCM для шифрования файлов
- Валидация сложности паролей

## 🚀 Планы по развитию

- [ ] **Веб-интерфейс** - добавление web-версии
- [ ] **Мобильное приложение** - версия для iOS/Android
- [ ] **Облачная синхронизация** - синхронизация между устройствами
- [ ] **Двухфакторная аутентификация** - дополнительная безопасность
- [ ] **Импорт/экспорт** - миграция из других менеджеров
- [ ] **Плагины для браузеров** - автозаполнение паролей
- [ ] **Аудит безопасности** - регулярная проверка паролей

## 📄 Лицензия

Этот проект распространяется под лицензией MIT. Подробнее см. в файле [LICENSE](LICENSE).

---

**⚠️ Важно**: Регулярно создавайте резервные копии вашего файла паролей и храните мастер-пароль в безопасном месте!
