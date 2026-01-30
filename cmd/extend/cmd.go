package extend

import (
	"encoding/base64"
	"fmt"

	macaroonsv1 "github.com/agentio/macaroo/genproto/agent.io/macaroons/v1"
	"github.com/agentio/macaroo/internal/generate"
	"github.com/google/cel-go/cel"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend MACAROON CONSTRAINT",
		Short: "Extend a macaroon by adding a constraint.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			macaroon := args[0]
			constraint := args[1]

			b, err := base64.RawURLEncoding.DecodeString(macaroon)
			if err != nil {
				return err
			}
			var m macaroonsv1.Macaroon
			err = proto.Unmarshal(b, &m)
			if err != nil {
				return err
			}
			env, err := cel.NewEnv()
			if err != nil {
				return err
			}
			// Check that the expression compiles and returns a String.
			ast, iss := env.Parse(`"` + constraint + `"`)
			// Report syntactic errors, if present.
			if iss.Err() != nil {
				return err
			}
			// Type-check the expression for correctness.
			checked, iss := env.Check(ast)
			checkedExpr, err := cel.AstToCheckedExpr(checked)
			if err != nil {
				return err
			}
			v, err := anypb.New(checkedExpr)
			m.Checks = append(m.Checks, v)

			k := m.Signature
			m.Signature = nil
			b2, err := proto.Marshal(&m)
			if err != nil {
				return err
			}
			signature := generate.HMAC([]byte(k), b2)
			m.Signature = signature
			b4, err := proto.Marshal(&m)
			if err != nil {
				return err
			}
			macaroon2 := base64.RawURLEncoding.EncodeToString(b4)
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", macaroon2)
			return nil

		},
	}
	return cmd
}
