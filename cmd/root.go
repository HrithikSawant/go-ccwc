/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/HrithikSawant/go-ccwc/pkg/utils"
	"github.com/spf13/cobra"
)

// Flags
var filePath string
var countLines bool
var countWords bool
var countChars bool
var countBytes bool

// Function to get options based on flags
func getOptions(countBytes, countLines, countWords, countChars bool) utils.Options {
	// Default options when no flags are set
	if !countBytes && !countLines && !countWords && !countChars {
		return utils.Options{
			PrintBytes: true,
			PrintLines: true,
			PrintWords: true,
			PrintChars: true,
		}
	}

	// Options based on the flags set
	return utils.Options{
		PrintBytes: countBytes, // You can set this to true if you want byte counting as default
		PrintLines: countLines,
		PrintWords: countWords,
		PrintChars: countChars,
	}
}

// openFile opens a file and returns a reader or an error if the file doesn't exist
func openFile(path string) (*bufio.Reader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("No such file or directory: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}

	return bufio.NewReader(file), nil
}

// isInputFromPipe checks if the input is coming from a pipe (stdin)
func isInputFromPipe() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0 // Means it's from pipe
}

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "ccwc",
		Short: "A command-line utility to count characters, words, and lines in a file",
		Long: `ccwc is a CLI tool that reads input files or standard input and provides
a summary of the number of characters, words, and lines. It functions similarly
to the Unix wc command, which is commonly used for text processing.

Examples of usage:
  ccwc file.txt        		# Show lines, words, and character counts for 'file.txt'
  ccwc -l file.txt     		# Count only number of lines in 'file.txt'
  echo "Hello World" | ccwc 	# Process input from a pipe and output counts.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },

		Run: func(cmd *cobra.Command, args []string) {
			var reader *bufio.Reader
			var err error

			if filePath != "" {
				// Handle file input using the helper function
				reader, err = openFile(filePath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
			} else if len(args) > 0 {
				// Handle file input without the flag (args[0] is the file path)
				filePath = args[0]
				reader, err = openFile(filePath)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
				}
			} else if isInputFromPipe() {
				// Handle stdin input (pipe)
				reader = bufio.NewReader(os.Stdin)
			} else {
				cmd.Help()
				return
			}

			// Calculate statistics
			stats := utils.CalculateStats(reader)
			// Get options based on the flags
			options := getOptions(countBytes, countLines, countWords, countChars)

			// Format and print stats
			formattedStats := utils.FormatStats(options, stats, filePath)
			fmt.Println(formattedStats)
		},
	}
)

func init() {
	// Define flags for the count command
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the input file")
	rootCmd.Flags().BoolVarP(&countLines, "lines", "l", false, "Count only lines")
	rootCmd.Flags().BoolVarP(&countWords, "words", "w", false, "Count only words")
	rootCmd.Flags().BoolVarP(&countChars, "chars", "m", false, "Count only characters")
	rootCmd.Flags().BoolVarP(&countBytes, "bytes", "c", false, "Count only bytes")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
