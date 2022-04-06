package root

import (
	"flag"

	"github.com/redhat-developer/app-services-cli/pkg/cmd/completion"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/login"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/logout"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/registry"
	"github.com/redhat-developer/app-services-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(f *factory.Factory, version string) *cobra.Command {

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           "apicr",
		Short:         "apicurio service registry cli",
		Long:          "",
		Example:       "",
	}
	fs := cmd.PersistentFlags()

	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	fs.BoolVarP(&help, "help", "h", false, "Prints help information")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	cmd.AddCommand(login.NewLoginCmd(f))
	cmd.AddCommand(logout.NewLogoutCommand(f))
	cmd.AddCommand(completion.NewCompletionCommand(f))

	// Plugin command
	cmd.AddCommand(registry.NewServiceRegistryCommand(f))

	return cmd
}
