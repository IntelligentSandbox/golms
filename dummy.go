package main

import (
	// CLI framework
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"

	// Terminal UI and formatting
	_ "github.com/briandowns/spinner"
	_ "github.com/chzyer/readline"
	_ "github.com/fatih/color"

	// Optional: JSON parsing
	_ "github.com/tidwall/gjson"
)

func main() {
	// Empty main - just here to satisfy Go requirements
	// The underscore imports above ensure all packages are vendored
	// Delete this file after running 'go mod vendor'
}