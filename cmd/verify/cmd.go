package verify

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	macaroonsv1 "github.com/agentio/macaroo/genproto/agent.io/macaroons/v1"
	"github.com/agentio/macaroo/internal/generate"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify MACAROON KEY",
		Short: "Verify a macaroon using the original key used to create it.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			macaroon := args[0]
			key := []byte(args[1])
			b, err := base64.RawURLEncoding.DecodeString(macaroon)
			if err != nil {
				return err
			}
			var m macaroonsv1.Macaroon
			err = proto.Unmarshal(b, &m)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStderr(), "want %s\n", hex.EncodeToString(m.Signature))
			var signature []byte
			var m2 macaroonsv1.Macaroon
			m2.Id = m.Id
			m2.Caveats = nil
			b, err = proto.Marshal(&m2)
			if err != nil {
				return err
			}
			signature = generate.HMAC(key, b)
			fmt.Fprintf(cmd.OutOrStderr(), "sig0 %s\n", hex.EncodeToString(signature))
			key = signature
			for i := range m.Caveats {
				var m2 macaroonsv1.Macaroon
				m2.Id = m.Id
				m2.Caveats = m.Caveats[0 : i+1]
				b, err := proto.Marshal(&m2)
				if err != nil {
					return err
				}
				signature = generate.HMAC(key, b)
				fmt.Fprintf(cmd.OutOrStderr(), "sig%d %s\n", i+1, hex.EncodeToString(signature))
				key = []byte(signature)
			}
			if bytes.Equal(signature, m.Signature) {
				fmt.Fprintf(cmd.OutOrStdout(), "VERIFIED\n")
			} else {
				fmt.Fprintf(cmd.OutOrStderr(), "%s != %s\n", hex.EncodeToString(signature), hex.EncodeToString(m.Signature))
				fmt.Fprintf(cmd.OutOrStdout(), "FAILED\n")
			}
			return nil
		},
	}
	return cmd
}

//return hex.EncodeToString(signature)
