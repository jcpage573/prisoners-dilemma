package config

import (
	"bufio"
	"os"
	"strings"
)

// Load reads a .env file and sets the environment variables
func Load(filename string) error {
	if filename == "" {
		filename = ".env"
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err := parseLine(line); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func parseLine(line string) error {
	// Trim whitespace and ignore comments and empty lines
	line = strings.TrimSpace(line)
	if len(line) == 0 || strings.HasPrefix(line, "#") {
		return nil
	}

	// Split the line into key and value
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return nil
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	// Remove surrounding quotes if present
	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		value = strings.Trim(value, "\"")
	} else if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
		value = strings.Trim(value, "'")
	}

	return os.Setenv(key, value)
}
