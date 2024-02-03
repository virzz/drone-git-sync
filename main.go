package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version string
)

var (
	sshKey, remote    string
	force, skipVerify bool
)

var cmd = &cobra.Command{
	Use:     "git-sync",
	Short:   "git-sync plugin",
	Version: Version,
	RunE: func(c *cobra.Command, args []string) error {
		return pluginExec()
	},
}

func init() {
	cmd.Flags().StringVar(&sshKey, "ssh-key", "", "private ssh key")
	cmd.Flags().StringVar(&remote, "remote", "", "url of the remote repo")
	cmd.Flags().BoolVar(&force, "force", false, "force push to remote")
	cmd.Flags().BoolVar(&skipVerify, "skip-verify", false, "skip ssl verification")
	viper.BindPFlag("ssh-key", cmd.Flags().Lookup("PLUGIN_SSH_KEY"))
	viper.BindPFlag("remote", cmd.Flags().Lookup("PLUGIN_REMOTE"))
	viper.BindPFlag("force", cmd.Flags().Lookup("PLUGIN_FORCE"))
	viper.BindPFlag("skip-verify", cmd.Flags().Lookup("PLUGIN_SKIP_VERIFY"))
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
