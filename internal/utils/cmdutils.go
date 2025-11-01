package utils

import (
	"strings"
)

func PerpareTemps(temps []string) {
	for i, temp := range temps {
		parts := strings.Split(temp, "/")
		temps[i] = parts[len(parts)-1]
	}
}
