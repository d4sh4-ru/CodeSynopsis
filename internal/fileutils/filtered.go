package fileutils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	gitignorePatterns []string
	once              sync.Once
)

// loadGitignore загружает .gitignore только один раз
func loadGitignore() {
	once.Do(func() {
		gitignorePatterns = readGitignore(".")
	})
}

// shouldIgnore проверяет, нужно ли исключить файл/папку
func shouldIgnore(name, path string) bool {
	return matchesHidden(name) || matchesGitignore(gitignorePatterns, path)
}

// GetSortedDirEntries — возвращает отсортированные файлы и папки в директории, исключая скрытые и исключенные
func GetSortedDirEntries(dir string) ([]os.DirEntry, error) {
	loadGitignore()

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var dirs, files []os.DirEntry

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())

		if shouldIgnore(entry.Name(), path) {
			continue
		}

		if entry.IsDir() {
			dirs = append(dirs, entry)
		} else {
			files = append(files, entry)
		}
	}

	return append(dirs, files...), nil
}

// GetProjectStructure — собирает пути всех файлов проекта, игнорируя скрытые и исключенные
func GetProjectStructure(rootDir string) []string {
	loadGitignore()

	var projectPaths []string

	err := filepath.WalkDir(rootDir, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if shouldIgnore(d.Name(), p) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		relPath, _ := filepath.Rel(rootDir, p)
		projectPaths = append(projectPaths, relPath)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking project root: %v\n", err)
	}

	return projectPaths
}

// GetSelectedFiles — возвращает список файлов, выбранных пользователем
func GetSelectedFiles(selectedItems map[string]bool) []string {
	loadGitignore()

	var files []string

	for path := range selectedItems {
		fileInfo, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			continue
		}

		if fileInfo.IsDir() {
			err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if !d.IsDir() && !shouldIgnore(d.Name(), p) {
					files = append(files, p)
				}
				return nil
			})
			if err != nil {
				fmt.Printf("Error walking directory %s: %v\n", path, err)
			}
		} else {
			files = append(files, path)
		}
	}

	return files
}
