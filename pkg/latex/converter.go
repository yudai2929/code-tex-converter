package latex

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Convert(paths []string, isLstInput bool, targetDir string, basePath string) ([]string, error) {
	converted := []string{}
	for _, path := range paths {
		language := determineLanguage(path)

		resolvedPath, err := resolvePath(targetDir, path, basePath)

		if err != nil {
			return nil, err
		}

		if isLstInput {
			converted = append(converted, lstInputListing(language, resolvedPath))
			continue
		}

		content, err := os.ReadFile(path)

		if err != nil {
			return nil, err
		}
		converted = append(converted, listing(language, resolvedPath, string(content)))
	}
	return converted, nil
}

func resolvePath(targetDir, path, basePath string) (string, error) {
	relPath, err := filepath.Rel(targetDir, path)
	if err != nil {
		return "", err
	}
	if basePath == "" {
		return relPath, nil
	}
	return basePath + "/" + relPath, nil
}

func lstInputListing(language, path string) string {
	return fmt.Sprintf("\\lstinputlisting[language=%s, caption=%s]", language, path)
}

func listing(language, path, content string) string {
	begin := fmt.Sprintf("\\begin{lstlisting}[language=%s, caption=%s]", language, escapeLaTeX(path))
	end := "\\end{lstlisting}"
	return begin + "\n" + content + "\n" + end
}

func escapeLaTeX(str string) string {
	str = strings.ReplaceAll(str, "\\", "\\textbackslash ")
	str = strings.ReplaceAll(str, "{", "\\{")
	str = strings.ReplaceAll(str, "}", "\\}")
	str = strings.ReplaceAll(str, "%", "\\%")
	str = strings.ReplaceAll(str, "$", "\\$")
	str = strings.ReplaceAll(str, "#", "\\#")
	str = strings.ReplaceAll(str, "&", "\\&")
	str = strings.ReplaceAll(str, "_", "\\_")
	str = strings.ReplaceAll(str, "^", "\\textasciicircum ")
	str = strings.ReplaceAll(str, "~", "\\textasciitilde ")
	return str
}

func determineLanguage(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		return "Go"
	case ".py":
		return "Python"
	case ".js":
		return "JavaScript"
	case ".ts":
		return "TypeScript"
	case ".rs":
		return "Rust"
	default:
		return "Text"
	}
}
