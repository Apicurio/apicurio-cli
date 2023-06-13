package context

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/create"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/delete"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/list"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/unset"
	"github.com/apicurio/apicurio-cli/pkg/cmd/context/use"

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

	cmd.AddCommand(
		use.NewUseCommand(f),
		list.NewListCommand(f),
		create.NewCreateCommand(f),
		delete.NewDeleteCommand(f),
		unset.NewUnsetCommand(f),
	)
	return cmd
}
