package extend

import (
	"fmt"
	"strings"

	"github.com/agentio/macaroo/internal/generate"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend MACAROON CONSTRAINT",
		Short: "Extend a macaroon by adding a constraint.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			macaroon := args[0]
			constraint := args[1]

			parts := strings.Split(macaroon, ":")

			key := parts[len(parts)-1]

			newparts := append(parts[0:len(parts)-1], constraint)
			message := strings.Join(newparts, ":")

			signature := generate.HMAC([]byte(key), []byte(message))
			fmt.Fprintf(cmd.OutOrStdout(), "%s:%s\n", message, signature)

			return nil
		},
	}
	return cmd
}
