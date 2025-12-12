package utils

import "testing"

func TestRemoveThinkTags(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Single think block with newlines",
			input:    "<think>\nOkay, the user is asking...\nI need to respond.\n</think>\n\nHello, cm! What's your name?",
			expected: "Hello, cm! What's your name?",
		},
		{
			name:     "Multiple think blocks",
			input:    "<think>First thought</think>Answer<think>Second thought</think>",
			expected: "Answer",
		},
		{
			name:     "Think block at the beginning",
			input:    "<think>reasoning here</think>Final answer",
			expected: "Final answer",
		},
		{
			name:     "Think block at the end",
			input:    "Final answer<think>reasoning here</think>",
			expected: "Final answer",
		},
		{
			name:     "Think block in the middle",
			input:    "Before<think>reasoning here</think>After",
			expected: "BeforeAfter",
		},
		{
			name:     "No think blocks",
			input:    "Just a normal response",
			expected: "Just a normal response",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Only think block",
			input:    "<think>Only thinking, no answer</think>",
			expected: "",
		},
		{
			name:     "Think block with special characters",
			input:    "<think>Testing with $pecial ch@rs & symbols!</think>Clean answer",
			expected: "Clean answer",
		},
		{
			name:     "Multiple lines between think blocks",
			input:    "<think>First</think>\n\nAnswer line 1\nAnswer line 2\n\n<think>Second</think>\n\nMore answer",
			expected: "Answer line 1\nAnswer line 2\n\n\n\nMore answer",
		},
		{
			name:     "Nested angle brackets inside think",
			input:    "<think>What if I write <something> here?</think>Answer",
			expected: "Answer",
		},
		{
			name:     "Extra whitespace around content",
			input:    "   <think>thinking</think>   Answer   ",
			expected: "Answer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveThinkTags(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveThinkTags() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestRemoveThinkTags_RealWorldExample(t *testing.T) {
	input := `<think>
Okay, the user is asking, "what is my name?" after saying "my name is cm." So they're confirming their name. I need to respond appropriately.

First, I should acknowledge their name. Since they just said "my name is cm," maybe I can say something like "Hello, cm! What's your name?" But wait, they already provided it. Maybe I should just confirm and offer help.

Wait, maybe they are testing if I can recognize their name. So I should respond with a friendly message confirming their name and offering assistance. Also, check if they need anything else. Keep it simple and positive. Make sure to use the correct punctuation as per the conversation history.

I think that's all. Just need to make sure the response is clear and helpful.
</think>

Hello, cm! What's your name? 😊`

	expected := "Hello, cm! What's your name? 😊"

	result := RemoveThinkTags(input)
	if result != expected {
		t.Errorf("RemoveThinkTags() real world example failed\nGot:  %q\nWant: %q", result, expected)
	}
}
