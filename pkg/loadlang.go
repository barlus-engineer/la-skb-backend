package pkg

import (
	"bufio"
	"os"
	"strings"
)

func LoadLangFile(filePath string) (map[string]string, error) {
	langMap := make(map[string]string)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		langMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return langMap, nil
}
