package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary .env file
	envContent := `DATABASE_URL="postgres://user:password@localhost:5432/dbname"
API_KEY=your_api_key_here
DEBUG=true`

	file, err := os.CreateTemp("", ".env")
	if err != nil {
		t.Fatalf("Error creating temporary file: %v", err)
	}
	defer os.Remove(file.Name()) // Clean up the file after the test

	_, err = file.Write([]byte(envContent))
	if err != nil {
		t.Fatalf("Error writing to temporary file: %v", err)
	}
	file.Close()

	// Load the .env file
	err = Load(file.Name())
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	// Check if the environment variables are set correctly
	tests := []struct {
		key   string
		value string
	}{
		{"DATABASE_URL", "postgres://user:password@localhost:5432/dbname"},
		{"API_KEY", "your_api_key_here"},
		{"DEBUG", "true"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			value := os.Getenv(tt.key)
			if value != tt.value {
				t.Errorf("Expected %s to be %s, got %s", tt.key, tt.value, value)
			}
		})
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		line  string
		key   string
		value string
	}{
		{`DATABASE_URL="postgres://user:password@localhost:5432/dbname"`, "DATABASE_URL", "postgres://user:password@localhost:5432/dbname"},
		{`API_KEY=your_api_key_here`, "API_KEY", "your_api_key_here"},
		{`DEBUG=true`, "DEBUG", "true"},
		{`EMPTY_KEY=`, "EMPTY_KEY", ""},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			os.Clearenv()
			if err := parseLine(tt.line); err != nil {
				t.Fatalf("Error parsing line: %v", err)
			}

			value := os.Getenv(tt.key)
			if value != tt.value {
				t.Errorf("Expected %s to be %s, got %s", tt.key, tt.value, value)
			}
		})
	}
}
