package cmd

import (
	"os"

	"github.com/nullsploit01/cc-calculator/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cccalc 'expression'",
	Short: "CCCalc is a CLI calculator capable of parsing and evaluating complex arithmetic expressions.",
	Long: `CCCalc is a command-line calculator designed to efficiently evaluate arithmetic expressions including nested operations and standard mathematical functions.
			For example, you can use CCCalc to calculate expressions like:
				cccalc "5 + 2 * (3 - 1)"
				cccalc "10 / 2 + 3 * (2 + 1)"

			The calculator supports basic operations such as addition (+), subtraction (-), multiplication (*), and division (/). 
			It is designed to handle errors such as division by zero gracefully and ensures that the order of operations is correctly followed.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.PrintErr("Please provide an expression")
			cmd.Usage()
			return
		}
		calculator := internal.NewCalculator(args[0])
		result := calculator.Calculate()

		cmd.Println(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
