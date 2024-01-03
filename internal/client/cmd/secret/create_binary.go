package secret

import (
	"context"
	"os"

	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	"github.com/evgenivanovi/gophkeeper/internal/client/usecase/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
)

/* __________________________________________________ */

type CreateBinaryCommand struct {
	op      common.Options
	args    CreateBinarySecretArg
	usecase secret.CreateDecodedSecretUsecase
}

func ProvideCreateBinaryCommand(
	op common.Options,
	args CreateBinarySecretArg,
	usecase secret.CreateDecodedSecretUsecase,
) *CreateBinaryCommand {
	return &CreateBinaryCommand{
		op:      op,
		args:    args,
		usecase: usecase,
	}
}

func (cmd *CreateBinaryCommand) Run() error {

	ctx := context.Background()
	ctx = common.OptionsWithCtx(ctx, cmd.op)

	bytes, err := os.ReadFile(cmd.args.Path)
	if err != nil {
		panic(err)
	}

	data := secretsharedmd.DecodedSecretDataModel{
		Name: cmd.args.Name,
		Type: "BINARY",
		Content: &secretsharedmd.BinarySecretContentModel{
			Bytes: bytes,
		},
	}

	err = cmd.usecase.Execute(ctx, data)
	if err != nil {
		return err
	}

	return nil

}

/* __________________________________________________ */
