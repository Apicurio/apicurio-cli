package main

import (
	"fmt"
	"os"

	"github.com/apicurio/apicurio-cli/internal/build"
	"github.com/apicurio/apicurio-cli/pkg/core/config"
	"github.com/apicurio/apicurio-cli/pkg/core/localize/goi18n"

	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	"github.com/apicurio/apicurio-cli/pkg/shared/factory/defaultfactory"

	"github.com/apicurio/apicurio-cli/pkg/cmd/root"
)

func main() {
	localizer, err := goi18n.New(&goi18n.Config{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := defaultfactory.New(localizer)

	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	initConfig(cmdFactory)

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	err = rootCmd.Execute()

	if err == nil {
		return
	}
}

func initConfig(f *factory.Factory) {
	cfgFile, err := f.Config.Load()

	if cfgFile != nil {
		return
	}
	if !os.IsNotExist(err) {
		fmt.Fprintln(f.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	cfgFile = &config.Config{}
	if err := f.Config.Save(cfgFile); err != nil {
		fmt.Fprintln(f.IOStreams.ErrOut, err)
		os.Exit(1)
	}
}
