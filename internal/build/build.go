package build

import (
	"runtime/debug"
	"time"
)

type buildSource string

const (
	GithubBuildSource buildSource = "github"
)

var (
	ProductionAPIURL            = "https://api.openshift.com"
	StagingAPIURL               = "https://api.stage.openshift.com"
	ConsoleURL                  = "https://console.redhat.com"
	DefaultClientID             = "rhoas-cli-prod"
	DefaultUserAgentPrefix      = "rhoas-cli_"
	DefaultOfflineTokenClientID = "cloud-services"
	DefaultLoginTimeout         = 60 * time.Second
	OfflineTokenURL             = ConsoleURL + "/openshift/token"
	ProductionAuthURL           = "https://sso.redhat.com/auth/realms/redhat-external"
	StagingAuthURL              = "https://sso.stage.redhat.com/auth/realms/redhat-external"
)

// Define public variables here which you wish to be configurable at build time
var (
	// Version is dynamically set by the toolchain or overridden by the Makefile.
	Version = "dev"

	// Language used, can be overridden by Makefile or CI
	Language = "en"
)

// Define public variables here which you wish to be configurable at build time
var (

	// RepositoryOwner is the remote GitHub organization for the releases
	RepositoryOwner = "redhat-developer"

	// RepositoryName is the remote GitHub repository for the releases
	RepositoryName = "app-services-cli"

	// DynamicConfigURL Url used to download dynamic service constants. If empty then static service constants are  used.
	DynamicConfigURL = "https://console.redhat.com/apps/application-services/service-constants.json"

	// DefaultPageSize is the default number of items per page when using list commands
	DefaultPageSize = "10"

	// DefaultPageNumber is the default page number when using list commands
	DefaultPageNumber = "1"

	// SSORedirectPath is the default SSO redirect path
	SSORedirectPath = "sso-redhat-callback"

	// BuildSource is a unique key which indicates the infrastructure on which the binary was built
	BuildSource = "local"
)

func init() {
	if isDevBuild() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

// isDevBuild returns true if the current build is "dev" (dev build)
func isDevBuild() bool {
	return Version == "dev"
}

func init() {
	if IsDevBuild() {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}

// IsDevBuild returns true if the current build is "dev" (dev build)
func IsDevBuild() bool {
	return Version == "dev"
}
