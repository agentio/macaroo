package create

import (
	"fmt"

	"github.com/agentio/macaroo/internal/generate"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create SECRET",
		Short: "Create a macaroon using a secret.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			nonce, err := generate.Nonce()
			if err != nil {
				return err
			}
			key := []byte(args[0])
			message := []byte(nonce)
			signature := generate.HMAC(key, message)
			fmt.Fprintf(cmd.OutOrStdout(), "%s:%s\n", message, signature)
			return nil
		},
	}
	return cmd
}
