package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spotless",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
		   examples and usage of using your application. For example:

		   Cobra is a CLI library for Go that empowers applications.
		   This application is a tool to generate the needed files
		   to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Adds all child commands to the root command 
// and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Flags and configuration settings.
	
	// Persistent flags (global in the entire app)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spotless.yaml)")

	// Local flags (run when action is called directly)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


