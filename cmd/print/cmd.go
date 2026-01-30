package print

import (
	"encoding/base64"
	"log"

	macaroonsv1 "github.com/agentio/macaroo/genproto/agent.io/macaroons/v1"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "print MACAROON",
		Short: "Display the contents of a macaroon.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			macaroon := args[0]
			b, err := base64.RawURLEncoding.DecodeString(macaroon)
			if err != nil {
				return err
			}
			var m macaroonsv1.Macaroon
			err = proto.Unmarshal(b, &m)
			if err != nil {
				return err
			}
			b3, err := protojson.MarshalOptions{
				Indent:    "  ",
				Multiline: true,
			}.Marshal(&m)
			log.Printf("%s", string(b3))
			return nil
		},
	}
	return cmd
}
