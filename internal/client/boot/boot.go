package boot

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	authgrpcapi "github.com/evgenivanovi/gophkeeper/api/pb/auth"
	secretgrpcapi "github.com/evgenivanovi/gophkeeper/api/pb/secret/public"
	authcmd "github.com/evgenivanovi/gophkeeper/internal/client/cmd/auth"
	secretcmd "github.com/evgenivanovi/gophkeeper/internal/client/cmd/secret"
	"github.com/evgenivanovi/gophkeeper/internal/client/common"
	userdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/auth"
	tokendm "github.com/evgenivanovi/gophkeeper/internal/client/domain/auth/token"
	cfgdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/config"
	fsdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/fs"
	secretdm "github.com/evgenivanovi/gophkeeper/internal/client/domain/secret"
	authgrpc "github.com/evgenivanovi/gophkeeper/internal/client/grpc/auth"
	secretgrpc "github.com/evgenivanovi/gophkeeper/internal/client/grpc/secret"
	authuc "github.com/evgenivanovi/gophkeeper/internal/client/usecase/auth"
	configuc "github.com/evgenivanovi/gophkeeper/internal/client/usecase/config"
	secretuc "github.com/evgenivanovi/gophkeeper/internal/client/usecase/secret"
	secretshareddm "github.com/evgenivanovi/gophkeeper/internal/shared/domain/secret"
	"github.com/evgenivanovi/gpl/std"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Options common.Options

	RootCMD *cobra.Command

	SigninCMD  *cobra.Command
	SignupCMD  *cobra.Command
	SignoutCMD *cobra.Command

	GetSecretArguments = secretcmd.GetSecretArg{}
	GetSecretCommand   *cobra.Command

	CreateSecretArguments       = secretcmd.CreateSecretArg{}
	CreateBinarySecretArguments = secretcmd.CreateBinarySecretArg{}

	CreateSecretCommand            *cobra.Command
	CreateTextSecretCommand        *cobra.Command
	CreateBinarySecretCommand      *cobra.Command
	CreateCredentialsSecretCommand *cobra.Command
	CreateCardSecretCommand        *cobra.Command

	FileManager fsdm.Manager

	ConfigParser  cfgdm.Parser
	ConfigReader  cfgdm.Reader
	ConfigWriter  cfgdm.Writer
	ConfigManager cfgdm.Manager

	TokenProvider tokendm.TokenProvider

	AuthHTTPClient *resty.Client
	AuthGRPCClient authgrpcapi.AuthAPIClient
	AuthAPI        userdm.AuthAPI

	AddUserUsecase      configuc.AddUserUsecase
	SetUserUsecase      configuc.SetUserUsecase
	CreateConfigUsecase configuc.CreateConfigUsecase

	SecretContentEncoder            secretshareddm.SecretContentEncoder
	SecretContentDecoder            secretshareddm.SecretContentDecoder
	TextSecretContentDecoder        secretshareddm.TextSecretContentDecoder
	BinarySecretContentDecoder      secretshareddm.BinarySecretContentDecoder
	CredentialsSecretContentDecoder secretshareddm.CredentialsSecretContentDecoder
	CardSecretContentDecoder        secretshareddm.CardSecretContentDecoder

	SecretEncoderDecoder      secretshareddm.SecretEncoderDecoder
	OwnedSecretEncoderDecoder secretshareddm.OwnedSecretEncoderDecoder

	SecretSeekingGRPCClient secretgrpcapi.SecretSeekingAPIClient
	SecretSeekerAPI         secretdm.SeekerAPI

	SecretManagementGRPCClient secretgrpcapi.SecretManagementAPIClient
	SecretManagementAPI        secretdm.SecretManagementAPI

	SigninUsecase authuc.SigninUsecase
	SignupUsecase authuc.SignupUsecase

	GetSecretUsecase           secretuc.GetSecretUsecase
	CreateDecodedSecretUsecase secretuc.CreateDecodedSecretUsecase
	CreateEncodedSecretUsecase secretuc.CreateEncodedSecretUsecase
)

func init() {
	initOptions()
	initDependencies()
	initUsecases()
	initCommands()
	initArguments()
}

func initOptions() {

	options, err := common.NewOptions()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	Options = options

}

func initDependencies() {

	FileManager = fsdm.ProvideManagerService()

	ConfigParser = cfgdm.ProvideYamlParser()
	ConfigReader = cfgdm.ProvideReaderService(ConfigParser)
	ConfigWriter = cfgdm.ProvideWriterService(ConfigParser)
	ConfigManager = cfgdm.ProvideManagerService(ConfigReader, ConfigWriter)

	TokenProvider = tokendm.ProvideTokenProviderService(ConfigReader)

	initHTTP()
	initGRPC()

	AuthAPI = authgrpc.ProvideAPIService(AuthGRPCClient)
	SecretSeekerAPI = secretgrpc.ProvideAPISeekingService(TokenProvider, SecretSeekingGRPCClient)
	SecretManagementAPI = secretgrpc.ProvideAPIManagementService(TokenProvider, SecretManagementGRPCClient)

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

}

func initUsecases() {

	AddUserUsecase = configuc.ProvideAddUserUsecaseService(
		FileManager, ConfigManager,
	)

	SetUserUsecase = configuc.ProvideSetUserUsecaseService(
		ConfigManager,
	)

	CreateConfigUsecase = configuc.ProvideCreateConfigUsecaseService(
		FileManager,
	)

	SigninUsecase = authuc.ProvideSigninUsecaseService(
		AuthAPI, AddUserUsecase, SetUserUsecase, CreateConfigUsecase, ConfigManager,
	)

	SignupUsecase = authuc.ProvideSignupUsecaseService(
		AuthAPI, AddUserUsecase, SetUserUsecase, CreateConfigUsecase, ConfigManager,
	)

	GetSecretUsecase = secretuc.ProvideGetSecretUsecaseService(
		SecretSeekerAPI,
		ConfigManager,
		SecretEncoderDecoder,
	)

	CreateDecodedSecretUsecase = secretuc.ProvideCreateDecodedSecretUsecaseService(
		SecretManagementAPI,
		ConfigManager,
	)

	CreateEncodedSecretUsecase = secretuc.ProvideCreateEncodedSecretUsecaseService(
		SecretManagementAPI,
		ConfigManager,
	)

}

func initHTTP() {

	AuthHTTPClient = resty.
		New().
		SetBaseURL(Options.Address)

}

func initGRPC() {

	conn, err := grpc.Dial(
		Options.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		fmt.Println(
			fmt.Errorf(
				"error in creating GRPC connection: %w",
				err,
			),
		)
		os.Exit(1)
	}

	AuthGRPCClient = authgrpcapi.NewAuthAPIClient(conn)
	SecretManagementGRPCClient = secretgrpcapi.NewSecretManagementAPIClient(conn)
	SecretSeekingGRPCClient = secretgrpcapi.NewSecretSeekingAPIClient(conn)

}

func initCommands() {

	RootCMD = &cobra.Command{
		Use: ApplicationName,
	}

	SigninCMD = &cobra.Command{
		Use: "signin",
		Run: func(cmd *cobra.Command, args []string) {

			mode := authcmd.ProvideSigninModel(
				Options, SigninUsecase,
			)

			program := tea.NewProgram(mode)

			_, err := program.Run()
			if err != nil {
				return
			}

		},
	}
	RootCMD.AddCommand(SigninCMD)

	SignupCMD = &cobra.Command{
		Use: "signup",
		Run: func(cmd *cobra.Command, args []string) {

			mode := authcmd.ProvideSignupModel(
				Options, SignupUsecase,
			)

			program := tea.NewProgram(mode)

			_, err := program.Run()
			if err != nil {
				return
			}

		},
	}
	RootCMD.AddCommand(SignupCMD)

	SignoutCMD = &cobra.Command{
		Use: "signout",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("signout called")
		},
	}
	RootCMD.AddCommand(SignoutCMD)

	GetSecretCommand = &cobra.Command{
		Use: "get",
		Run: func(cmd *cobra.Command, args []string) {

			command := secretcmd.ProvideGetCommand(
				Options, GetSecretArguments, GetSecretUsecase,
			)

			res, err := command.Run()

			if err != nil {
				output := strings.Builder{}
				output.WriteString(err.Error())
				output.WriteString(std.NL)
				fmt.Println(output.String())
			} else {
				output := strings.Builder{}
				output.WriteString(res)
				fmt.Println(output.String())
			}

		},
	}
	RootCMD.AddCommand(GetSecretCommand)

	CreateSecretCommand = &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {},
	}
	RootCMD.AddCommand(CreateSecretCommand)

	CreateTextSecretCommand = &cobra.Command{
		Use: "text",
		Run: func(cmd *cobra.Command, args []string) {

			mode := secretcmd.ProvideCreateTextModel(
				Options, CreateSecretArguments, CreateDecodedSecretUsecase,
			)

			program := tea.NewProgram(mode)

			_, err := program.Run()
			if err != nil {
				return
			}

		},
	}
	CreateSecretCommand.AddCommand(CreateTextSecretCommand)

	CreateBinarySecretCommand = &cobra.Command{
		Use: "binary",
		Run: func(cmd *cobra.Command, args []string) {

			command := secretcmd.ProvideCreateBinaryCommand(
				Options, CreateBinarySecretArguments, CreateDecodedSecretUsecase,
			)

			err := command.Run()

			if err != nil {
				output := strings.Builder{}
				output.WriteString(err.Error())
				output.WriteString(std.NL)
				fmt.Println(output.String())
			} else {
				output := strings.Builder{}
				output.WriteString("Binary succeed!")
				output.WriteString(std.NL)
				fmt.Println(output.String())
			}

		},
	}
	CreateSecretCommand.AddCommand(CreateBinarySecretCommand)

	CreateCredentialsSecretCommand = &cobra.Command{
		Use: "credentials",
		Run: func(cmd *cobra.Command, args []string) {

			mode := secretcmd.ProvideCreateCredentialsModel(
				Options, CreateSecretArguments, CreateDecodedSecretUsecase,
			)

			program := tea.NewProgram(mode)

			_, err := program.Run()
			if err != nil {
				return
			}

		},
	}
	CreateSecretCommand.AddCommand(CreateCredentialsSecretCommand)

	CreateCardSecretCommand = &cobra.Command{
		Use: "card",
		Run: func(cmd *cobra.Command, args []string) {

			mode := secretcmd.ProvideCreateCardModel(
				Options, CreateSecretArguments, CreateDecodedSecretUsecase,
			)

			program := tea.NewProgram(mode)

			_, err := program.Run()
			if err != nil {
				return
			}

		},
	}
	CreateSecretCommand.AddCommand(CreateCardSecretCommand)

}

func initArguments() {
	initGetSecretCommandArguments()
	initCreateSecretCommandArguments()
	initCreateBinarySecretCommandArguments()
}

func initGetSecretCommandArguments() {

	args := &GetSecretArguments
	cmd := GetSecretCommand

	cmd.PersistentFlags().StringVarP(
		&args.Name,
		"name",
		"",
		"",
		"",
	)

	err := cmd.MarkPersistentFlagRequired("name")
	if err != nil {
		return
	}

	err = viper.BindPFlag("name", cmd.PersistentFlags().Lookup("name"))
	if err != nil {
		return
	}

}

func initCreateSecretCommandArguments() {

	args := &CreateSecretArguments
	cmd := CreateSecretCommand

	cmd.PersistentFlags().StringVarP(
		&args.Name,
		"name",
		"",
		"",
		"",
	)

	err := cmd.MarkPersistentFlagRequired("name")
	if err != nil {
		return
	}

	err = viper.BindPFlag("name", cmd.PersistentFlags().Lookup("name"))
	if err != nil {
		return
	}

}

func initCreateBinarySecretCommandArguments() {

	args := &CreateBinarySecretArguments
	cmd := CreateBinarySecretCommand

	cmd.PersistentFlags().StringVarP(
		&args.Name,
		"name",
		"",
		"",
		"",
	)

	err := cmd.MarkPersistentFlagRequired("name")
	if err != nil {
		return
	}

	err = viper.BindPFlag("name", cmd.PersistentFlags().Lookup("name"))
	if err != nil {
		return
	}

	cmd.PersistentFlags().StringVarP(
		&args.Path,
		"path",
		"",
		"",
		"",
	)

	err = cmd.MarkPersistentFlagRequired("path")
	if err != nil {
		return
	}

	err = viper.BindPFlag("path", cmd.PersistentFlags().Lookup("path"))
	if err != nil {
		return
	}

}
