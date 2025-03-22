# Переменные конфигурации
APP_NAME := CodeSynopsis
VERSION := 0.1.0
BUILD_DIR := build
SRC_FILES := ./cmd/app/main.go

# Основные цели
.PHONY: all run build release clean help

all: 
	build

# Запуск приложения
run:
	go run $(SRC_FILES)

# Сборка для текущей платформы
build:  $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) $(SRC_FILES)

# Кросс-компиляция для Linux
release-linux: $(BUILD_DIR)/$(APP_NAME)-linux-amd64

$(BUILD_DIR)/$(APP_NAME)-linux-amd64: $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-linux-amd64 $(SRC_FILES)

# Кросс-компиляция для Windows
release-windows: $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe

$(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe: $(SRC_FILES)
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME)-windows-amd64.exe $(SRC_FILES)

# Полная сборка для всех платформ
release: release-linux release-windows

# Очистка
clean:
	rm -rf $(BUILD_DIR)
	rm -f *.exe

# Справка
help:
	@echo "Доступные команды:"
	@echo "  make run       - Запустить приложение"
	@echo "  make build     - Собрать для текущей платформы"
	@echo "  make release   - Собрать для Linux и Windows"
	@echo "  make clean     - Удалить все сборки"
	@echo "  make help      - Показать эту справку"

# Отображение справки по умолчанию
.DEFAULT_GOAL := help