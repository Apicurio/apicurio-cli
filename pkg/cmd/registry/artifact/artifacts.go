package artifact

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/crud/create"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/crud/delete"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/crud/get"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/crud/list"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/crud/update"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/download"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/metadata"
	migrate "github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/migrate"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/owner"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/state"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/types"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/artifact/versions"
	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewArtifactsCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "artifact",
		Short:   f.Localizer.MustLocalize("artifact.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("artifact.cmd.description.long"),
		Example: f.Localizer.MustLocalize("artifact.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		// CRUD
		create.NewCreateCommand(f),
		get.NewGetCommand(f),
		delete.NewDeleteCommand(f),
		list.NewListCommand(f),
		update.NewUpdateCommand(f),

		// Misc
		metadata.NewGetMetadataCommand(f),
		metadata.NewSetMetadataCommand(f),
		versions.NewVersionsCommand(f),
		download.NewDownloadCommand(f),
		migrate.NewExportCommand(f),
		migrate.NewImportCommand(f),
		state.NewSetStateCommand(f),
		owner.NewGetCommand(f),
		owner.NewSetCommand(f),
		types.NewGetTypesCommand(f),
	)

	return cmd
}
