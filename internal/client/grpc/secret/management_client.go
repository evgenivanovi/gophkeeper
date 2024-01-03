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

type APIManagementService struct {
	tokenProvider token.TokenProvider
	client        secretapi.SecretManagementAPIClient
}

func ProvideAPIManagementService(
	tokenProvider token.TokenProvider,
	client secretapi.SecretManagementAPIClient,
) *APIManagementService {
	return &APIManagementService{
		tokenProvider: tokenProvider,
		client:        client,
	}
}

func (a *APIManagementService) CreateDecoded(
	ctx context.Context, user string, data secretdm.DecodedSecretData,
) (secretdm.EncodedSecret, error) {

	ctx, err := a.buildContext(ctx, user)
	if err != nil {
		return secretdm.NewEmptyEncodedSecret(), err
	}

	req := a.buildCreateDecodedRequest(data)

	res, err := a.client.CreateDecoded(ctx, req)
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

func (a *APIManagementService) buildCreateDecodedRequest(
	data secretdm.DecodedSecretData,
) *secretapi.CreateDecodedSecretRequest {

	pbdata := secretgrpc.FromDecodedSecretDataModel(
		secretsharedmd.FromDecodedSecretData(data),
	)

	return &secretapi.CreateDecodedSecretRequest{
		Payload: &secretapi.CreateDecodedSecretRequest_Payload{
			Data: pbdata,
		},
	}

}

func (a *APIManagementService) CreateEncoded(
	ctx context.Context, user string, data secretdm.EncodedSecretData,
) (secretdm.EncodedSecret, error) {

	ctx, err := a.buildContext(ctx, user)
	if err != nil {
		return secretdm.NewEmptyEncodedSecret(), err
	}

	req := a.buildCreateEncodedRequest(data)

	res, err := a.client.CreateEncoded(ctx, req)
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

func (a *APIManagementService) buildCreateEncodedRequest(
	data secretdm.EncodedSecretData,
) *secretapi.CreateEncodedSecretRequest {

	pbdata := secretgrpc.FromEncodedSecretDataModel(
		secretsharedmd.FromEncodedSecretData(data),
	)

	return &secretapi.CreateEncodedSecretRequest{
		Payload: &secretapi.CreateEncodedSecretRequest_Payload{
			Data: pbdata,
		},
	}

}

func (a *APIManagementService) buildContext(
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
