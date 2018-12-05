package utils

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type flag struct {
	cmd *cobra.Command
}

func NewFlagBuilder(parent_cmd, cmd *cobra.Command) *flag {
	if parent_cmd != nil {
		parent_cmd.AddCommand(cmd)
	}
	return &flag{cmd: cmd}
}

func (f flag) viperName(name string) string {
	return strings.Replace(f.cmd.CommandPath(), " ", ".", -1) + "." + name
}

func (f *flag) Bool(p *bool, name, shorthand string, value bool, usage string) {
	f.cmd.Flags().BoolVarP(p, name, shorthand, value, usage)
	vName := f.viperName(name)
	viper.BindPFlag(vName, f.cmd.Flags().Lookup(name))
}

func (f *flag) String(p *string, name, shorthand, value, usage string) {
	f.cmd.Flags().StringVarP(p, name, shorthand, value, usage)
	vName := f.viperName(name)
	viper.BindPFlag(vName, f.cmd.Flags().Lookup(name))
}
