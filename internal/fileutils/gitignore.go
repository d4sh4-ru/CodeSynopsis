package fileutils

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// readGitignore читает содержимое файла .gitignore и возвращает список шаблонов
func readGitignore(dir string) []string {
	gitignorePath := filepath.Join(dir, ".gitignore")
	file, err := os.Open(gitignorePath)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var patterns []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}
	return patterns
}

// matchesGitignore проверяет, соответствует ли путь шаблонам из файла .gitignore
func matchesGitignore(patterns []string, path string) bool {
	for _, pattern := range patterns {
		match, _ := filepath.Match(pattern, filepath.Base(path))
		if match {
			return true
		}
	}
	return false
}
