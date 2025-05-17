package cmd

import (
	"os"

	"github.com/pass1on-ok/EmailVerification/internal/email"
	"github.com/spf13/cobra"
)

const emailFlag = "email"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ev",
	Short: "",
	Long:  `A powerful Golang-based tool for effortlessly scraping and tracking stock prices and financial ratios, with seamless data saving options.`,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flag(emailFlag).Changed {
			emailString := cmd.Flag(emailFlag).Value.String()
			err := email.VerifyEmail(emailString)
			if err != nil {
				cmd.Printf("Error: %s ❌\n", err.Error())
				return
			}
			cmd.Println("\u2705Email address is valid\u2705")
			return
		}

		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("email", "", "Email address to verify")
}
