package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
    Use:   "spotless [directory]",
    Short: "Spotless keeps your directories clean and organized",
    Long: `Spotless is a CLI tool for organizing your files effectively.
Use Spotless to organize a directory by letting it place files into separate 
folders based on their name, extension...`,
    Args: cobra.MaximumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        toggle, _ := cmd.Flags().GetBool("toggle")
        fmt.Println(toggle, args)
    },
}

/** Adds all child commands to the root command 
 *  and sets flags appropriately.
 */
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Local flags (run when action is called directly)
	rootCmd.Flags().BoolP("name", "n", false, "group based on the starting names")
	rootCmd.Flags().BoolP("ext", "e", false, "group based on the file extensions")
	rootCmd.Flags().BoolP("date", "d", false, "group based on the creation dates")
}