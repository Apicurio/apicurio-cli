// REST API exposed via the serve command.
package registry

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/factory"

	"github.com/spf13/cobra"
)

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "service-registry",
		Short:   f.Localizer.MustLocalize("registry.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("registry.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("registry.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	return cmd
}
