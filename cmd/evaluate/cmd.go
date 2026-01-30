package evaluate

import (
	"encoding/base64"
	"fmt"
	"log"

	macaroonsv1 "github.com/agentio/macaroo/genproto/agent.io/macaroons/v1"
	"github.com/google/cel-go/cel"
	"github.com/spf13/cobra"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	"google.golang.org/protobuf/proto"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "evaluate MACAROON",
		Short: "Evaluate the constraints in a macaroon.",
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
			for i := range m.Checks {
				check := m.Checks[i]
				if check.TypeUrl == "type.googleapis.com/google.api.expr.v1alpha1.CheckedExpr" {
					var ex expr.CheckedExpr
					err = proto.Unmarshal(check.Value, &ex)
					if err != nil {
						return err
					}
					checked := cel.CheckedExprToAst(&ex)
					env, err := cel.NewEnv()
					if err != nil {
						return err
					}
					prg, err := env.Program(checked)
					if err != nil {
						log.Fatalln(err)
					}
					out, _, err := prg.Eval(map[string]any{
						"name": "CEL",
					})
					if err != nil {
						log.Fatalln(err)
					}
					fmt.Println(out)
				}
			}
			return nil
		},
	}
	return cmd
}
