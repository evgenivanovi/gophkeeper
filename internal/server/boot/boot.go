package boot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evgenivanovi/gophkeeper/api/http/authapi"
	secretprivateapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/private"
	secretpublicapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/public"
	sessiondm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/session"
	"github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/token"
	userdm "github.com/evgenivanovi/gophkeeper/internal/server/domain/auth/user"
	secretdm "github.com/evgenivanovi/gophkeeper/internal/server/domain/secret"
	grpcauth "github.com/evgenivanovi/gophkeeper/internal/server/grpc/auth"
	grpcsecret "github.com/evgenivanovi/gophkeeper/internal/server/grpc/secret"
	httpctl "github.com/evgenivanovi/gophkeeper/internal/server/http"
	authctl "github.com/evgenivanovi/gophkeeper/internal/server/http/auth"
	secretpg "github.com/evgenivanovi/gophkeeper/internal/server/postgres/secret"
	sessionpg "github.com/evgenivanovi/gophkeeper/internal/server/postgres/session"
	userpg "github.com/evgenivanovi/gophkeeper/internal/server/postgres/user"
	authuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/auth"
	secretuc "github.com/evgenivanovi/gophkeeper/internal/server/usecase/secret"
	"github.com/evgenivanovi/gophkeeper/internal/shared/domain/auth/user"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	auth2 "github.com/evgenivanovi/gophkeeper/internal/shared/util/auth"
	"github.com/evgenivanovi/gophkeeper/pkg/grpc"
	"github.com/evgenivanovi/gpl/fw"
	"github.com/evgenivanovi/gpl/goose"
	"github.com/evgenivanovi/gpl/meta"
	"github.com/evgenivanovi/gpl/pg"
	jwtxgrpc "github.com/evgenivanovi/gpl/stdx/jwtx/grpc"
	slogx "github.com/evgenivanovi/gpl/stdx/log/slog"
	"github.com/evgenivanovi/gpl/stdx/mw"
	"github.com/evgenivanovi/gpl/vcs"
	"github.com/go-chi/chi/v5"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
)

/* __________________________________________________ */

var (
	Props  Properties
	VCS    vcs.VCS
	App    meta.App
	Config fw.Configuration

	Router chi.Router

	Datasource             *pg.Datasource
	PostgresReadRequester  *pg.ReadRequester
	PostgresWriteRequester *pg.WriteRequester
	PostgresTransactor     *pg.TrxRequester

	AuthReadRepository  userdm.AuthReadRepository
	AuthWriteRepository userdm.AuthWriteRepository
	AuthRepository      userdm.AuthRepository
	PasswordManager     userdm.PasswordManager
	AuthManager         user.AuthManager

	TokenManager token.Manager

	SessionReadRepository  sessiondm.ReadRepository
	SessionWriteRepository sessiondm.WriteRepository
	SessionRepository      sessiondm.Repository
	SessionIDGenerator     sessiondm.IDGenerator
	SessionManager         sessiondm.Manager

	SigninUsecase authuc.SigninUsecase
	SignupUsecase authuc.SignupUsecase

	SecretReadRepository  secretdm.ReadRepository
	SecretWriteRepository secretdm.WriteRepository
	SecretRepository      secretdm.Repository

	SecretContentEncoder            secretshareddm.SecretContentEncoder
	SecretContentDecoder            secretshareddm.SecretContentDecoder
	TextSecretContentDecoder        secretshareddm.TextSecretContentDecoder
	BinarySecretContentDecoder      secretshareddm.BinarySecretContentDecoder
	CredentialsSecretContentDecoder secretshareddm.CredentialsSecretContentDecoder
	CardSecretContentDecoder        secretshareddm.CardSecretContentDecoder

	SecretEncoderDecoder      secretshareddm.SecretEncoderDecoder
	OwnedSecretEncoderDecoder secretshareddm.OwnedSecretEncoderDecoder

	SecretManager secretdm.Manager
	SecretSeeker  secretdm.Seeker

	CreateDecodedSecretUsecase secretuc.CreateDecodedSecretUsecase
	CreateEncodedSecretUsecase secretuc.CreateEncodedSecretUsecase

	GetDecodedByNameSecretUsecase secretuc.GetDecodedByNameUsecase
	GetEncodedByNameSecretUsecase secretuc.GetEncodedByNameUsecase
)

func init() {
	Props = ProvideProperties()
	Router = chi.NewRouter()
}

/* __________________________________________________ */

// Boot ...
func Boot() {
	bootApp()
	bootPGDatasource()
	bootPGInfrastructure()
	bootPGMigrations()
	bootAuth()
	bootHTTPRouter()
	bootGRPCRouter()
	bootServer()
	bootPrint()
}

func bootApp() {

	VCS = vcs.NewVCS()
	vcs.Read(&VCS)

	App = meta.NewAppWithOps(
		meta.WithAppName(AppName),
		meta.WithAppVersion(AppVersion),
	)

	Config = *fw.NewConfiguration()

	Config.App.Settings = *fw.NewServerSettings(
		fw.WithHttp(Props.HTTPEnabled),
		fw.WithHttpPort(Props.HTTPPort),

		fw.WithHttps(Props.HTTPSEnabled),
		fw.WithHttpsPort(Props.HTTPSPort),

		fw.WithGrpc(Props.GRPCEnabled),
		fw.WithGrpcPort(Props.GRPCPort),
	)

}

func bootPGDatasource() {

	dsn, err := pg.NewDatasource(
		context.Background(),
		Props.DSNPostgres,
		*pg.NewConnectionSettings(),
	)

	if err != nil {
		panic(err)
	}

	Datasource = dsn

}

func bootPGInfrastructure() {
	PostgresReadRequester = pg.ProvideReadRequester(Datasource.Pool)
	PostgresWriteRequester = pg.ProvideWriteRequester(Datasource.Pool)
	PostgresTransactor = pg.ProvideTrxRequester(Datasource.Pool)
}

func bootPGMigrations() {

	slogx.Log().Debug(
		"Starting PostgreSQL migrations.",
	)

	goose.MigrateUp("./migrations", "postgres", Props.DSNPostgres)

	slogx.Log().Debug(
		"Finished PostgreSQL migrations.",
	)

}

func bootAuth() {

	AuthReadRepository = userpg.ProvidePGReadRepositoryService(
		*PostgresReadRequester,
	)
	AuthWriteRepository = userpg.ProvidePGWriteRepositoryService(
		*PostgresWriteRequester,
	)
	AuthRepository = userdm.ProvideAuthRepositoryService(
		AuthReadRepository, AuthWriteRepository,
	)

	PasswordManager = userdm.ProvidePasswordManagerService()

	TokenManager = token.ProvideManagerService(
		*auth2.NewTokenSettings(
			auth2.WithAccessSecret(Props.JWTAccessTokenSecretKey),
			auth2.WithAccessExpiration(Props.JWTAccessTokenExpirationTime),

			auth2.WithRefreshSecret(Props.JWTRefreshTokenSecretKey),
			auth2.WithRefreshExpiration(Props.JWTRefreshTokenExpirationTime),
		),
	)

	SessionReadRepository = sessionpg.ProvidePGReadRepositoryService(
		*PostgresReadRequester,
	)
	SessionWriteRepository = sessionpg.ProvidePGWriteRepositoryService(
		*PostgresWriteRequester,
	)
	SessionRepository = sessiondm.ProvideRepositoryService(
		SessionReadRepository, SessionWriteRepository,
	)

	SessionIDGenerator = sessiondm.ProvideIDGeneratorService()

	SessionManager = sessiondm.ProvideManagerService(
		SessionIDGenerator, SessionRepository,
	)

	AuthManager = userdm.ProvideAuthManagerService(
		PostgresTransactor,
		AuthRepository,
		PasswordManager,
		TokenManager,
		SessionManager,
	)

	SigninUsecase = authuc.ProvideSigninUsecaseService(
		PostgresTransactor, AuthManager,
	)
	SignupUsecase = authuc.ProvideSignupUsecaseService(
		PostgresTransactor, AuthManager,
	)

	SecretReadRepository = secretpg.ProvidePGReadRepositoryService(
		*PostgresReadRequester,
	)
	SecretWriteRepository = secretpg.ProvidePGWriteRepositoryService(
		*PostgresWriteRequester,
	)
	SecretRepository = secretdm.ProvideRepositoryService(
		SecretReadRepository, SecretWriteRepository,
	)

	SecretContentEncoder = secretshareddm.ProvideGobSecretContentEncoder()
	SecretContentDecoder = secretshareddm.ProvideGobSecretContentDecoder()
	TextSecretContentDecoder = secretshareddm.ProvideGobTextSecretContentDecoder()
	BinarySecretContentDecoder = secretshareddm.ProvideGobBinarySecretContentDecoder()
	CredentialsSecretContentDecoder = secretshareddm.ProvideGobCredentialsSecretContentDecoder()
	CardSecretContentDecoder = secretshareddm.ProvideGobCardSecretContentDecoder()

	SecretEncoderDecoder = secretshareddm.ProvideSecretEncoderDecoderService(
		SecretContentEncoder, SecretContentDecoder,
	)

	OwnedSecretEncoderDecoder = secretshareddm.ProvideOwnedSecretEncoderDecoderService(
		SecretContentEncoder, SecretContentDecoder,
	)

	SecretManager = secretdm.ProvideManagerService(
		SecretRepository,
		SecretContentEncoder,
		SecretContentDecoder,
	)

	SecretSeeker = secretdm.ProvideSeekerService(
		SecretRepository,
		SecretContentEncoder,
		SecretContentDecoder,
	)

	CreateDecodedSecretUsecase = secretuc.ProvideCreateDecodedSecretUsecaseService(
		SecretManager,
	)
	CreateEncodedSecretUsecase = secretuc.ProvideCreateEncodedSecretUsecaseService(
		SecretManager,
	)

	GetDecodedByNameSecretUsecase = secretuc.ProvideGetDecodedByNameUsecaseService(
		SecretSeeker,
		OwnedSecretEncoderDecoder,
	)
	GetEncodedByNameSecretUsecase = secretuc.ProvideGetEncodedByNameUsecaseService(
		SecretSeeker,
		OwnedSecretEncoderDecoder,
	)

}

func bootHTTPRouter() {

	Router.NotFound(defaultErrorHandler)
	Router.MethodNotAllowed(defaultErrorHandler)

	Router.Post(
		authapi.SigninEndpoint.String(),
		mw.Conveyor(
			httpctl.AsHandler(
				authctl.ProvideSignInController(
					SigninUsecase,
				),
			),
		).ServeHTTP,
	)

	Router.Post(
		authapi.SignupEndpoint.String(),
		mw.Conveyor(
			httpctl.AsHandler(
				authctl.ProvideSignUpController(
					SignupUsecase,
				),
			),
		).ServeHTTP,
	)

	Config.WithHTTPHandler(Router)

}

func defaultErrorHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusBadRequest)
	_, _ = writer.Write(nil)
}

func bootGRPCRouter() {

	methods := make(map[string]bool)
	methods[secretprivateapi.InternalSecretManagementAPI_CreateDecoded_FullMethodName] = false
	methods[secretprivateapi.InternalSecretManagementAPI_CreateEncoded_FullMethodName] = false
	methods[secretpublicapi.SecretManagementAPI_CreateDecoded_FullMethodName] = true
	methods[secretpublicapi.SecretManagementAPI_CreateEncoded_FullMethodName] = true
	methods[secretpublicapi.SecretSeekingAPI_GetByName_FullMethodName] = true

	authMW := jwtxgrpc.NewUnary(
		jwtxgrpc.WithKey(auth2.KeyProvider(Props.JWTAccessTokenSecretKey)),
		jwtxgrpc.WithMethod(auth2.MethodProvider()),
		jwtxgrpc.WithClaims(auth2.ClaimsProvider()),
		jwtxgrpc.WithExtractor(auth2.GRPCExtractorProvider()),
		jwtxgrpc.WithVerifier(auth2.GRPCVerifierProvider()),
	)

	Config.
		WithGRPCReflection(
			true,
		).
		WithGRPCStreamMW(
			grpcrecovery.StreamServerInterceptor(),
		).
		WithGRPCUnaryMW(
			grpc.UnaryServerProtected(methods, authMW),
			grpcrecovery.UnaryServerInterceptor(),
		).
		WithGRPCServices(
			grpcauth.ProvideAuthAPI(
				SigninUsecase, SignupUsecase,
			),
			grpcsecret.ProvideInternalSecretManagementAPI(
				CreateDecodedSecretUsecase,
				CreateEncodedSecretUsecase,
			),
			grpcsecret.ProvideSecretManagementAPI(
				CreateDecodedSecretUsecase,
				CreateEncodedSecretUsecase,
			),
			grpcsecret.ProvideSecretSeekingAPI(
				GetEncodedByNameSecretUsecase,
			),
		)

}

func bootServer() {

	if err := fw.RunServer(&Config); err != nil {
		slogx.Log().Debug("boot server failed", slogx.ErrAttr(err))
	}

}

func bootPrint() {

	fmt.Println("App name:", App.Name)
	fmt.Println("App version:", App.Version)
	fmt.Println("Build date:", VCS.Time)
	fmt.Println("Build commit:", VCS.Revision)

	if Config.App.Settings.HttpEnabled() {
		slogx.Log().Debug("HTTP is enabled.")
		slogx.Log().Debug(
			"Calculated HTTP server options: " + Config.App.Settings.HttpAddress(),
		)
	}

	if Config.App.Settings.HttpsEnabled() {
		slogx.Log().Debug("HTTPs is enabled.")
		slogx.Log().Debug(
			"Calculated HTTPs server options: " + Config.App.Settings.HttpsAddress(),
		)
	}

	if Config.App.Settings.GrpcEnabled() {
		slogx.Log().Debug("gRPC is enabled.")
		slogx.Log().Debug(
			"Calculated gRPC server options: " + Config.App.Settings.GrpcAddress(),
		)
	}

}

/* __________________________________________________ */
