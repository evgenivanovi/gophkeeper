package secret

import (
	"context"
	"strings"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/usecase/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

type GetCommand struct {
	op      common.Options
	args    GetSecretArg
	usecase secret.GetSecretUsecase
}

func ProvideGetCommand(
	op common.Options,
	args GetSecretArg,
	usecase secret.GetSecretUsecase,
) *GetCommand {
	return &GetCommand{
		op:      op,
		args:    args,
		usecase: usecase,
	}
}

func (cmd *GetCommand) Run() (string, error) {

	ctx := context.Background()
	ctx = common.OptionsWithCtx(ctx, cmd.op)

	sec, err := cmd.usecase.Execute(ctx, cmd.args.Name)
	if err != nil {
		return "", err
	}

	return cmd.format(sec), nil

}

func (cmd *GetCommand) format(
	sec secretsharedmd.DecodedSecretModel,
) string {
	var output strings.Builder

	output.WriteString("Name: " + sec.Data.Name)
	output.WriteString(std.NL)
	output.WriteString("Type: " + sec.Data.Type)
	output.WriteString(std.NL)
	output.WriteString(std.NL)

	output.WriteString("Content:")
	output.WriteString(std.NL)
	output.WriteString(sec.Data.Content.String())
	output.WriteString(std.NL)

	return output.String()
}

/* __________________________________________________ */
