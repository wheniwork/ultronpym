package cmd

import "github.com/spf13/cobra"

var manifestFile string
var signatureFile string
var publicKey string

// Commands initializes all commands
func Commands() *cobra.Command {
	// RootCmd handles the default commands
	rootCmd := &cobra.Command{
		Use:   "ultron",
		Short: "Ultron handles listing, signing, and verifying manifests",
		Long: `Ultron makes creating, signing, and verifying manifests a breeze
                made with love from the team at When I Work`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVarP(&manifestFile, "manifest", "m", ".manifest.yml", "manifest file location")
	rootCmd.PersistentFlags().StringVarP(&signatureFile, "signature", "s", ".manifest.sig", "signature file location")
	rootCmd.PersistentFlags().StringVarP(&publicKey, "key", "k", "", "contents of the armored public keyring, defaults to gpg agent")

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(signCmd)
	rootCmd.AddCommand(verifyCmd)

	return rootCmd
}
