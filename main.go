package main

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Version string
)

var cmd = &cobra.Command{
	Use:     "git-sync",
	Short:   "git-sync plugin",
	Version: Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				viper.Set(f.Name, f.Value.String())
			}
		})
	},
	RunE: func(c *cobra.Command, args []string) error {
		return pluginExec()
	},
}

func init() {
	cmd.Flags().String("ssh-key", "", "private ssh key")
	cmd.Flags().String("remote", "", "url of the remote repo")
	cmd.Flags().Bool("force", false, "force push to remote")
	cmd.Flags().Bool("skip-verify", false, "skip ssl verification")
	cmd.Flags().Bool("debug", false, "debug mode")
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		viper.BindPFlag(f.Name, f)
	})
	viper.SetEnvPrefix("PLUGIN")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
