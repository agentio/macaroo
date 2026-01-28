package verify

import (
	"log"
	"strings"

	"github.com/agentio/macaroo/internal/generate"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify MACAROON SECRET",
		Short: "Verify a macaroon using the original secret used to create it.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			macaroon := args[0]
			secret := args[1]
			parts := strings.Split(macaroon, ":")
			for i := range parts[:len(parts)-1] {
				message := strings.Join(parts[0:i+1], ":")
				signature := generate.HMAC([]byte(secret), []byte(message))
				log.Printf("hashing %s -> %s", message, signature)
				secret = signature
			}
			if secret == parts[len(parts)-1] {
				log.Printf("VERIFIED")
			} else {
				log.Printf("FAILED")
			}
			return nil
		},
	}
	return cmd
}
