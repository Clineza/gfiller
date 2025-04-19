/* Copyright © 2025 NAME HERE <EMAIL ADDRESS>
 */
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "",
	Long:  ``,
	Args:  cobra.RangeArgs(0, 3),
	Run:   gen,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("iterations", "i", 1, "Number of iterations")
	generateCmd.Flags().BoolP("lines", "l", false, "Print out text on new lines")
	generateCmd.Flags().StringP("text", "t", "Hello World!", "Text that will iterate")
}

func gen(cmd *cobra.Command, args []string) {
	iterations, _ := cmd.Flags().GetInt("iterations")
	if iterations <= 0 {
		fmt.Fprintf(cmd.ErrOrStderr(), "Error: iterations must be greater than 0\n")
		return
	}
	isLines, _ := cmd.Flags().GetBool("lines")
	text, _ := cmd.Flags().GetString("text")

	if len(args) >= 1 {
		fmt.Fprintf(cmd.ErrOrStderr(), "Error: you must have at least 2 arguments if not 0")
		return
	}

	if len(args) >= 2 {

	}
	//note for tomorrow: you have to check to see if the first argument is an integer or not.. then the second argument you have to check if it's a string.. that way you can guarantee the order of the arguments

	for i := 0; i < iterations; i++ {
		if isLines {
			fmt.Println(text)
		} else {
			fmt.Print(text, " ")
		}
	}
}

// gfiller generate -w 1000 "hello"
