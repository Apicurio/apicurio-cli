package factory

import (
	"github.com/apicurio/apicurio-cli/internal/config"
	"github.com/apicurio/apicurio-cli/pkg/cmd/debug"
	"github.com/apicurio/apicurio-cli/pkg/iostreams"
	"github.com/apicurio/apicurio-cli/pkg/localize"
	"github.com/apicurio/apicurio-cli/pkg/logging"
)

// New creates a new command factory
// The command factory is available to all command packages
// giving centralized access to the config

// nolint:funlen
func New(cliVersion string, localizer localize.Localizer) *Factory {
	io := iostreams.System()

	var logger logging.Logger
	cfgFile := config.NewFile()

	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(io.Out, io.ErrOut)

		debugEnabled := debug.Enabled()
		loggerBuilder = loggerBuilder.Debug(debugEnabled)

		logger, err := loggerBuilder.Build()
		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	return &Factory{
		IOStreams: io,
		Config:    cfgFile,
		Logger:    loggerFunc,
		Localizer: localizer,
	}
}
