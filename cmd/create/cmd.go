package create

import (
	"encoding/base64"
	"fmt"

	macaroonsv1 "github.com/agentio/macaroo/genproto/agent.io/macaroons/v1"
	"github.com/agentio/macaroo/internal/generate"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create IDENTIFIER KEY",
		Short: "Create a macaroon using an identifier (nonce) and a secret key.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			key := []byte(args[1])
			m := &macaroonsv1.Macaroon{
				Id: id,
			}
			b, err := proto.Marshal(m)
			if err != nil {
				return err
			}
			signature := generate.HMAC(key, b)
			m.Signature = signature
			b, err = proto.Marshal(m)
			if err != nil {
				return err
			}
			macaroon := base64.RawURLEncoding.EncodeToString(b)
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", macaroon)
			return nil
		},
	}
	return cmd
}
