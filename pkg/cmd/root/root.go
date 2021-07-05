package root

import (
	"flag"

	"github.com/apicurio/apicurio-cli/pkg/cmd/registry"

	"github.com/apicurio/apicurio-cli/internal/build"
	"github.com/apicurio/apicurio-cli/pkg/localize"

	"github.com/apicurio/apicurio-cli/pkg/cmd/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(f *factory.Factory, version string) *cobra.Command {

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           f.Localizer.MustLocalize("root.cmd.use"),
		Short:         f.Localizer.MustLocalize("root.cmd.shortDescription"),
		Long:          f.Localizer.MustLocalize("root.cmd.longDescription"),
		Example:       f.Localizer.MustLocalize("root.cmd.example"),
	}
	fs := cmd.PersistentFlags()

	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	fs.BoolVarP(&help, "help", "h", false, f.Localizer.MustLocalize("root.cmd.flag.help.description"))

	cmd.Version = version

	cmd.SetVersionTemplate(f.Localizer.MustLocalize("version.cmd.outputText", localize.NewEntry("Version", build.Version)))
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	cmd.AddCommand(registry.NewServiceRegistryCommand(f))

	return cmd
}
