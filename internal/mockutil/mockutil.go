package mockutil

import (
	"context"
	"errors"

	"github.com/apicurio/apicurio-cli/pkg/core/config"
	"github.com/apicurio/apicurio-cli/pkg/core/ioutil/iostreams"
	"github.com/apicurio/apicurio-cli/pkg/core/logging"

	"github.com/apicurio/apicurio-cli/pkg/shared/connection"
	"github.com/apicurio/apicurio-cli/pkg/shared/connection/api"
	"github.com/apicurio/apicurio-cli/pkg/shared/connection/kcconnection"
)

func NewConfigMock(cfg *config.Config) config.IConfig {
	return &config.IConfigMock{
		LocationFunc: func() (string, error) {
			return ":mock_location:", nil
		},
		LoadFunc: func() (*config.Config, error) {
			return cfg, nil
		},
		SaveFunc: func(c *config.Config) error {
			cfg = c
			return nil
		},
		RemoveFunc: func() error {
			cfg = nil
			return nil
		},
	}
}

func NewConnectionMock(conn *kcconnection.Connection) connection.Connection {
	return &connection.ConnectionMock{
		RefreshTokensFunc: func(ctx context.Context) error {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return errors.New("")
			}
			if conn.Token.RefreshToken == "expired" {
				return errors.New("")
			}

			return nil
		},
		LogoutFunc: func(ctx context.Context) error {
			if conn.Token.AccessToken == "" && conn.Token.RefreshToken == "" {
				return errors.New("")
			}
			if conn.Token.AccessToken == "expired" && conn.Token.RefreshToken == "expired" {
				return errors.New("")
			}

			cfg, err := conn.Config.Load()
			if err != nil {
				return err
			}

			cfg.AccessToken = ""
			cfg.RefreshToken = ""

			return conn.Config.Save(cfg)
		},
		APIFunc: func() api.API {
			return nil
		},
	}
}

func NewLoggerMock() logging.Logger {
	io := iostreams.System()
	var logger logging.Logger
	loggerBuilder := logging.NewStdLoggerBuilder()
	loggerBuilder = loggerBuilder.Streams(io.Out, io.ErrOut)
	logger, _ = loggerBuilder.Build()
	return logger
}
