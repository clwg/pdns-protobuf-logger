package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// GenerateUUIDv5 creates a UUID v5 based on the provided namespace and name, and returns it as a string
func GenerateUUIDv5(name string) (string, error) {

	customNamespace, err := uuid.Parse("6ba7b810-6969-6969-6969-00c04fd430c8")
	if err != nil {
		return "", fmt.Errorf("failed to parse custom namespace UUID: %w", err)
	}
	return uuid.NewSHA1(customNamespace, []byte(name)).String(), nil
}
