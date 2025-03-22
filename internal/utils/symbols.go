package utils

import (
	"path/filepath"
	"strings"
)

func GetFileSymbol(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	base := strings.ToLower(filepath.Base(filename))

	// Фильтрация по base файла
	switch base {
	case "dockerfile":
		return "🐋"
	case "makefile":
		return "🛠️"
	case "license":
		return "🔒"
	default:
		if strings.HasPrefix(base, ".env") {
			return "🔧"
		}
	}

	// Фильтрация по расширению
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff":
		return "🖼️"
	case ".txt", ".md", ".markdown", ".rst":
		return "📋"
	case ".go", ".py", ".js", ".ts", ".java", ".cpp", ".c", ".html", ".css", ".json", ".php", ".rb", ".rs", ".swift", ".kt", ".scala", ".sql":
		return "📄"
	default:
		return "📄"
	}
}
