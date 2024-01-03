package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/evgenivanovi/gophkeeper/api/http/auth"
	"github.com/evgenivanovi/gophkeeper/api/http/authapi"
	"github.com/evgenivanovi/gophkeeper/api/http/common"
	httpctl "github.com/evgenivanovi/gophkeeper/internal/server/http"
	authuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/auth"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/core"
	authmd "github.com/evgenivanovi/gophkeeper/internal/shared/md/auth"
	errx "github.com/evgenivanovi/gpl/err"
	"github.com/evgenivanovi/gpl/stdx/net/http/headers"
)

/* __________________________________________________ */

type SignUpController struct {
	decoder func(io.Reader) *json.Decoder
	usecase authuc.SignupUsecase
}

func ProvideSignUpController(
	usecase authuc.SignupUsecase,
) *SignUpController {
	decoder := func(reader io.Reader) *json.Decoder {
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		return decoder
	}
	return &SignUpController{
		decoder: decoder,
		usecase: usecase,
	}
}

func (c *SignUpController) Handle(
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

func (c *SignUpController) buildRequest(
	request *http.Request,
) (*authuc.SignUpRequest, error) {

	var requestModel authapi.SignupRequest

	err := c.
		decoder(request.Body).
		Decode(&requestModel)

	if err != nil {
		return nil, err
	}

	return &authuc.SignUpRequest{
		Payload: authuc.SignUpRequestPayload{
			Credentials: authmd.CredentialsModel{
				Username: requestModel.Username,
				Password: requestModel.Password,
			},
		},
	}, nil

}

/* __________________________________________________ */

func (c *SignUpController) onSuccess(
	response authuc.SignUpResponse, writer http.ResponseWriter, request *http.Request,
) {

	responseModel := authapi.SigninResponse{
		Payload: authapi.SigninResponsePayload{
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

func (c *SignUpController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *SignUpController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *SignUpController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	code := errx.ErrorCode(err)

	if code == core.ErrorExistsCode {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	if code == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
