package helpers

import (
	"io"
	"os"
	"strings"
)

func GetInput(path string) string {
	file, _ := os.Open(path)
	content, _ := io.ReadAll(file)

	return strings.TrimSpace(string(content))
}
