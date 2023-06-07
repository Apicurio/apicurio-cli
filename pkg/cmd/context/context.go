package context

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/create"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/delete"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/list"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/unset"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/use"
	registryUse "github.com/apicurio/apicurio-cli/pkg/cmd/registry/use"

	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

// NewContextCmd creates a new command to manage service contexts
func NewContextCmd(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "context",
		Short:   f.Localizer.MustLocalize("context.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("context.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("context.cmd.example"),
		Args:    cobra.NoArgs,
	}

	// The implementation of `rhoas service-registry use` command has been aliased here as `rhoas context set-service-registry`
	registryUseCmd := registryUse.NewUseCommand(f)
	registryUseCmd.Use = "set-service-registry"
	registryUseCmd.Example = f.Localizer.MustLocalize("context.setRegistry.cmd.example")

	cmd.AddCommand(
		use.NewUseCommand(f),
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		unset.NewUnsetCommand(f),

		// reused sub-commands
		registryUseCmd,
	)
	return cmd
}
