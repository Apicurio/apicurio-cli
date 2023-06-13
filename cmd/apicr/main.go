package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/apicurio/apicurio-cli/internal/build"
	"github.com/apicurio/apicurio-cli/pkg/core/config"
	"github.com/apicurio/apicurio-cli/pkg/core/ioutil/icon"
	"github.com/apicurio/apicurio-cli/pkg/core/localize"
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

	cmdFactory.Logger.Errorf("%v\n", rootError(err, localizer))
	os.Exit(1)
}

// rootError creates the root error which is printed to the console
// it wraps the error which has been returned from subcommands with a prefix
func rootError(err error, localizer localize.Localizer) error {
	prefix := icon.ErrorPrefix()
	errMessage := err.Error()
	if prefix == icon.ErrorSymbol {
		errMessage = firstCharToUpper(errMessage)
	}

	if strings.Contains(errMessage, "\n") {
		return fmt.Errorf("%v %v\n%v", icon.ErrorPrefix(), errMessage, localizer.MustLocalize("common.log.error.verboseModeHint"))
	}
	return fmt.Errorf("%v %v. %v", icon.ErrorPrefix(), errMessage, localizer.MustLocalize("common.log.error.verboseModeHint"))
}

// Ensure that the first character in the string is capitalized
func firstCharToUpper(message string) string {
	return strings.ToUpper(message[:1]) + message[1:]
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
