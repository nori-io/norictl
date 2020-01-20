package utils

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type flags struct {
	cmd        *cobra.Command
	prefixFunc func(name, commandPathPrefix string) string
}

func NewFlagBuilder(parent_cmd, cmd *cobra.Command) *flags {
	if parent_cmd != nil {
		parent_cmd.AddCommand(cmd)
	}
	return &flags{cmd: cmd}
}

func (fb *flags) SetPrefixFunc(f func(name, commandPathPrefix string) string) {
	fb.prefixFunc = f
}

func (fb flags) viperName(name string) string {
	return strings.Replace(fb.cmd.CommandPath(), " ", ".", -1) + "." + name
}

func (fb *flags) Bool(f *func() bool, name, shorthand string, value bool, usage string) {
	fb.cmd.Flags().BoolP(name, shorthand, value, usage)
	vName := fb.viperName(name)
	if fb.prefixFunc != nil {
		vName = fb.prefixFunc(name, vName)
	}

	viper.BindPFlag(vName, fb.cmd.Flags().Lookup(name))

	*f = func() bool {
		return viper.GetBool(vName)
	}
}

func (fb *flags) String(f *func() string, name, shorthand, value, usage string) {
	fb.cmd.Flags().StringP(name, shorthand, value, usage)
	vName := fb.viperName(name)
	if fb.prefixFunc != nil {
		vName = fb.prefixFunc(name, vName)
	}

	viper.BindPFlag(vName, fb.cmd.Flags().Lookup(name))

	*f = func() string {
		return viper.GetString(vName)
	}
}

func (fb *flags) BoolP(f *func() bool, name, shorthand string, value bool, usage string) {
	fb.cmd.PersistentFlags().BoolP(name, shorthand, value, usage)
	vName := fb.viperName(name)
	if fb.prefixFunc != nil {
		vName = fb.prefixFunc(name, vName)
	}

	viper.BindPFlag(vName, fb.cmd.PersistentFlags().Lookup(name))

	*f = func() bool {
		return viper.GetBool(vName)
	}
}

func (fb *flags) StringP(f *func() string, name, shorthand, value, usage string) {
	fb.cmd.PersistentFlags().StringP(name, shorthand, value, usage)
	vName := fb.viperName(name)
	if fb.prefixFunc != nil {
		vName = fb.prefixFunc(name, vName)
	}

	viper.BindPFlag(vName, fb.cmd.PersistentFlags().Lookup(name))

	*f = func() string {
		return viper.GetString(vName)
	}
}
