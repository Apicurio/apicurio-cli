package rule

import (
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/rule/describe"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/rule/disable"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/rule/enable"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/rule/list"
	"github.com/apicurio/apicurio-cli/pkg/cmd/registry/rule/update"
	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	"github.com/spf13/cobra"
)

func NewRuleCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rule",
		Short:   f.Localizer.MustLocalize("registry.rule.cmd.description.short"),
		Long:    f.Localizer.MustLocalize("registry.rule.cmd.description.long"),
		Example: f.Localizer.MustLocalize("registry.rule.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// add sub-commands
	cmd.AddCommand(
		enable.NewEnableCommand(f),
		list.NewListCommand(f),
		describe.NewDescribeCommand(f),
		update.NewUpdateCommand(f),
		disable.NewDisableCommand(f),
	)

	return cmd
}
