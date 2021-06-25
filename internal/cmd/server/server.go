package server

import (
	"github.com/hetznercloud/cli/internal/hcapi2"
	"github.com/hetznercloud/cli/internal/state"
	"github.com/spf13/cobra"
)

func NewCommand(cli *state.State, client hcapi2.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "server",
		Short:                 "Manage servers",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		ListCmd.CobraCommand(cli.Context, client, cli),
		describeCmd.CobraCommand(cli.Context, client, cli),
		CreateCmd.CobraCommand(cli.Context, client, cli, cli),
		deleteCmd.CobraCommand(cli.Context, client, cli),
		newRebootCommand(cli),
		newPoweronCommand(cli),
		newPoweroffCommand(cli),
		newResetCommand(cli),
		newShutdownCommand(cli),
		newCreateImageCommand(cli),
		newResetPasswordCommand(cli),
		newEnableRescueCommand(cli),
		newDisableRescueCommand(cli),
		newAttachISOCommand(cli),
		newDetachISOCommand(cli),
		newUpdateCommand(cli),
		newChangeTypeCommand(cli),
		newRebuildCommand(cli),
		newEnableBackupCommand(cli),
		newDisableBackupCommand(cli),
		newEnableProtectionCommand(cli),
		newDisableProtectionCommand(cli),
		newSSHCommand(cli),
		labelCmds.AddCobraCommand(cli.Context, hcapi2.NewClient(cli.Client()), cli),
		labelCmds.RemoveCobraCommand(cli.Context, hcapi2.NewClient(cli.Client()), cli),
		newSetRDNSCommand(cli),
		newAttachToNetworkCommand(cli),
		newDetachFromNetworkCommand(cli),
		newChangeAliasIPsCommand(cli),
		newIPCommand(cli),
		newRequestConsoleCommand(cli),
		newMetricsCommand(cli),
	)
	return cmd
}
