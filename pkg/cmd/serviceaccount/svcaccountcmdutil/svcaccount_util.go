package svcaccountcmdutil

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/serviceaccount/svcaccountcmdutil/credentials"
)

var (
	CredentialsOutputFormats = []string{credentials.EnvFormat, credentials.JSONFormat, credentials.PropertiesFormat, credentials.SecretFormat, credentials.JavaPropertiesFormat}
)
