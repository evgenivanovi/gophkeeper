package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/evgenivanovi/gophkeeper/api/http/auth"
	httpapi "github.com/evgenivanovi/gophkeeper/api/http/authapi"
	"github.com/evgenivanovi/gophkeeper/api/http/common"
	httpctl "github.com/evgenivanovi/gophkeeper/internal/server/http"
	authuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
)

/* __________________________________________________ */

type SignInController struct {
	decoder func(io.Reader) *json.Decoder
	usecase authuc.SigninUsecase
}

func ProvideSignInController(
	usecase authuc.SigninUsecase,
) *SignInController {
	decoder := func(reader io.Reader) *json.Decoder {
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		return decoder
	}
	return &SignInController{
		decoder: decoder,
		usecase: usecase,
	}
}

func (c *SignInController) Handle(
	writer http.ResponseWriter, request *http.Request,
) {

	requestModel, requestError := c.buildRequest(request)
	if requestError != nil {
		httpctl.LogErrorRequest(requestError)
		c.onError(requestError, nil, writer, request)
		return
	}

	responseModel, responseError := c.usecase.Execute(
		request.Context(), *requestModel,
	)

	if responseError != nil {
		httpctl.LogSuccessRequest(requestModel)
		httpctl.LogErrorResponse(responseError)
		c.onError(nil, responseError, writer, request)
	} else {
		httpctl.LogSuccessRequest(requestModel)
		httpctl.LogSuccessResponse(responseModel)
		c.onSuccess(responseModel, writer, request)
	}

}

func (c *SignInController) buildRequest(
	request *http.Request,
) (*authuc.SignInRequest, error) {

	var requestModel httpapi.SigninRequest

	err := c.
		decoder(request.Body).
		Decode(&requestModel)

	if err != nil {
		return nil, err
	}

	return &authuc.SignInRequest{
		Payload: authuc.SignInRequestPayload{
			Credentials: authmd.CredentialsModel{
				Username: requestModel.Username,
				Password: requestModel.Password,
			},
		},
	}, nil

}

/* __________________________________________________ */

func (c *SignInController) onSuccess(
	response authuc.SignInResponse, writer http.ResponseWriter, request *http.Request,
) {

	responseModel := httpapi.SigninResponse{
		Payload: httpapi.SigninResponsePayload{
			Session: auth.SessionModel{
				ID: response.Payload.Session.ID,
				Tokens: auth.TokensModel{
					AccessToken: auth.AccessTokenModel{
						Token:     response.Payload.Session.Tokens.AccessToken.Token,
						ExpiresAt: response.Payload.Session.Tokens.AccessToken.ExpiresAt,
					},
					RefreshToken: auth.RefreshTokenModel{
						Token:     response.Payload.Session.Tokens.RefreshToken.Token,
						ExpiresAt: response.Payload.Session.Tokens.RefreshToken.ExpiresAt,
					},
				},
			},
			User: auth.UserModel{
				ID: response.Payload.User.ID,
				Metadata: common.MetadataModel{
					CreatedAt: response.Payload.User.Metadata.CreatedAt,
					UpdatedAt: response.Payload.User.Metadata.UpdatedAt,
					DeletedAt: response.Payload.User.Metadata.DeletedAt,
				},
			},
		},
	}

	jsonResponseModel, jsonError := json.Marshal(responseModel)
	if jsonError != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add(
		headers.ContentTypeKey.String(),
		headers.TypeApplicationJSON.String(),
	)

	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write(jsonResponseModel)

}

func (c *SignInController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *SignInController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *SignInController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	code := errx.ErrorCode(err)

	if code == core.ErrorNotFoundCode {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	if code == core.ErrorUnauthenticatedCode {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if code == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
