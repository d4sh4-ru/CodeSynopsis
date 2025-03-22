package markdown

import (
	"fmt"
	"os"
	"strings"

	"github.com/d4sh4-ru/CodeSynopsis/internal/fileutils"
	"github.com/d4sh4-ru/CodeSynopsis/internal/tree"
	"github.com/d4sh4-ru/CodeSynopsis/internal/utils"
)

// GenerateMarkdown — основная функция, вызывающая вспомогательные
func GenerateMarkdown(selectedItems map[string]bool) {
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return
	}

	// Получаем все файлы проекта, игнорируя ненужные
	projectPaths := fileutils.GetProjectStructure(rootDir)

	// Получаем файлы, выбранные пользователем
	files := fileutils.GetSelectedFiles(selectedItems)

	// Генерируем Markdown
	markdownContent := buildMarkdownContent(files, projectPaths)

	// Записываем Markdown в файл
	writeMarkdownToFile(markdownContent)
}

// BuildMarkdownContent — формирует содержимое Markdown-файла
func buildMarkdownContent(files, projectPaths []string) string {
	var content strings.Builder

	// Добавляем структуру проекта
	treeStructure := tree.BuildTreeStructure(projectPaths)
	content.WriteString("# Структура проекта\n\n")
	content.WriteString("```\n" + treeStructure + "```\n\n")

	// Добавляем содержимое выбранных файлов
	for _, file := range files {
		content.WriteString(fmt.Sprintf("## %s\n\n", file))
		language := utils.DetectLanguage(file)
		data, err := os.ReadFile(file)
		if err != nil {
			content.WriteString(fmt.Sprintf("Error reading file: %v\n", err))
		} else {
			content.WriteString(fmt.Sprintf("```%s\n", language))
			content.Write(data)
			content.WriteString("\n```\n")
		}
		content.WriteString("\n\n---\n\n")
	}

	return content.String()
}

// WriteMarkdownToFile — записывает Markdown в output.md
func writeMarkdownToFile(content string) {
	err := os.WriteFile("output.md", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing output.md: %v\n", err)
	} else {
		fmt.Println("Document generated successfully: output.md")
	}
}
