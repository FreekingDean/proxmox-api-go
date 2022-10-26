package delete

import (
	"github.com/Telmate/proxmox-api-go/cli"
	"github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/spf13/cobra"
)

var (
	delete_snapshotCmd = &cobra.Command{
		Use:   "snapshot GUESTID SNAPSHOTNAME",
		Short: "Deletes the Speciefied snapshot",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id := cli.ValidateIntIDset(args, "GuestID")
			snapName := cli.RequiredIDset(args, 1, "SnapshotName")
			c := cli.NewClient()
			_, err = c.DeleteSnapshot(proxmox.NewVmRef(id), snapName)
			if err != nil {
				return
			}
			cli.PrintItemDeleted(deleteCmd.OutOrStdout(), snapName, "Snapshot")
			return
		},
	}
)

func init() {
	deleteCmd.AddCommand(delete_snapshotCmd)
}