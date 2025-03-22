package fileutils

import "strings"

// matchesHidden проверяет, начинается ли имя файла с точки (.), что является проверкой на скрытый файл
func matchesHidden(nameFile string) bool {
	isDotPrefix := strings.HasPrefix(nameFile, ".")
	isNotDot := nameFile != "."
	isNotDoubleDot := nameFile != ".."

	return isDotPrefix && isNotDot && isNotDoubleDot
}
