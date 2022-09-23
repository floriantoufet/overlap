package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/floriantoufet/overlap/internal/handler"
	"github.com/floriantoufet/overlap/internal/processor"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "[CIDR1] [CIDR2]",
	Example: " 10.0.2.0/24 10.0.0.0/20",
	Short:   "IP Overlap CLI",
	Long: `IP Overlap CLI  prints to STDOUT the relation between
two CIDRs. The relations can be:
• subset: if the network of the second address is included in the first one
• superset: if the network of the second address includes the first one
• different: if the two networks are not overlapping
• same: if both address are in the same network
The program is only intended to work with IPv4 addresses.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get relations
		h := handler.NewHandler(processor.NewProcessor())
		if err := h.PrintOverlapRelation(args, os.Stdout); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.overlap.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
