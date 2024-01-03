package secret

import (
	"context"

	secretapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/public"
	"github.com/evgenivanovi/gophkeeper/internal/client/domain/auth/token"
	secretgrpc "github.com/evgenivanovi/gophkeeper/internal/server/grpc/secret"
	secretdm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	secretsharedmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/secret"
	sharedauthutil "github.com/evgenivanovi/gophkeeper/internal/shared/util/auth"
	"google.golang.org/grpc/metadata"
)

/* __________________________________________________ */

type APISeekingService struct {
	tokenProvider token.TokenProvider
	client        secretapi.SecretSeekingAPIClient
}

func ProvideAPISeekingService(
	tokenProvider token.TokenProvider,
	client secretapi.SecretSeekingAPIClient,
) *APISeekingService {
	return &APISeekingService{
		tokenProvider: tokenProvider,
		client:        client,
	}
}

func (a *APISeekingService) GetByName(
	ctx context.Context, user string, name string,
) (secretdm.EncodedSecret, error) {

	ctx, err := a.buildContext(ctx, user)
	if err != nil {
		return secretdm.NewEmptyEncodedSecret(), err
	}

	req := a.buildGetByNameRequest(name)

	res, err := a.client.GetByName(ctx, req)
	if err != nil {
		return secretdm.NewEmptyEncodedSecret(), err
	}

	result := secretsharedmd.ToEncodedSecret(
		secretgrpc.ToEncodedSecretModel(
			res.GetPayload().GetData(),
		),
	)

	return result, nil

}

func (a *APISeekingService) buildGetByNameRequest(
	name string,
) *secretapi.GetByNameSecretRequest {

	return &secretapi.GetByNameSecretRequest{
		Payload: &secretapi.GetByNameSecretRequest_Payload{
			Name: name,
		},
	}

}

func (a *APISeekingService) buildContext(
	ctx context.Context, user string,
) (context.Context, error) {

	access, err := a.tokenProvider.ProvideAccess(ctx, user)
	if err != nil {
		return nil, err
	}

	inMD := metadata.Pairs(
		sharedauthutil.MetadataAuthKey, access.Token,
	)

	inCTX := metadata.NewOutgoingContext(
		context.Background(), inMD,
	)

	return inCTX, nil

}

/* __________________________________________________ */
