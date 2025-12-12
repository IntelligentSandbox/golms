package utils

import (
	"regexp"
	"strings"
)

// RemoveThinkTags removes <think></think> blocks from LLM responses
func RemoveThinkTags(content string) string {
	// Match <think>...</think> blocks (including newlines and any content)
	re := regexp.MustCompile(`(?s)<think>.*?</think>`)

	// Remove all think blocks
	cleaned := re.ReplaceAllString(content, "")

	// Trim leading and trailing whitespace
	cleaned = strings.TrimSpace(cleaned)

	return cleaned
}
