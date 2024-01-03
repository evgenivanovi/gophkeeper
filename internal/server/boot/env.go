package boot

import "time"

const (
	ConfigLocationPropertyName = "Config"
	ConfigLocationEnvName      = "GOPHKEEPER__CONFIG"
	ConfigLocationArgName      = "cfg"

	HTTPEnabledPropertyName = "HTTPEnabled"
	HTTPEnabledEnvName      = "GOPHKEEPER__HTTP_ENABLED"
	HTTPEnabledArgName      = "http.enabled"
	HTTPEnabledCfgName      = "http_enabled"
	HTTPEnabledDefaultValue = false

	HTTPPortPropertyName = "HTTPPort"
	HTTPPortEnvName      = "GOPHKEEPER__HTTP_PORT"
	HTTPPortArgName      = "http.port"
	HTTPPortCfgName      = "http_port"
	HTTPPortDefaultValue = 80

	HTTPSEnabledPropertyName = "HTTPSEnabled"
	HTTPSEnabledEnvName      = "GOPHKEEPER__HTTPS_ENABLED"
	HTTPSEnabledArgName      = "https.enabled"
	HTTPSEnabledCfgName      = "https_enabled"
	HTTPSEnabledDefaultValue = false

	HTTPSPortPropertyName = "HTTPSPort"
	HTTPSPortEnvName      = "GOPHKEEPER__HTTPS_PORT"
	HTTPSPortArgName      = "https.port"
	HTTPSPortCfgName      = "https_port"
	HTTPSPortDefaultValue = 443

	GRPCEnabledPropertyName = "GRPCEnabled"
	GRPCEnabledEnvName      = "GOPHKEEPER__GRPC_ENABLED"
	GRPCEnabledArgName      = "grpc.enabled"
	GRPCEnabledCfgName      = "grpc_enabled"
	GRPCEnabledDefaultValue = true

	GRPCPortPropertyName = "GRPCPort"
	GRPCPortEnvName      = "GOPHKEEPER__GRPC_PORT"
	GRPCPortArgName      = "grpc.port"
	GRPCPortCfgName      = "grpc_port"
	GRPCPortDefaultValue = 82

	DSNPostgresName    = "DSN"
	DSNPostgresEnvName = "GOPHKEEPER__DSN"
	DSNPostgresArgName = "dsn"
	DSNPostgresCfgName = "dsn"

	JWTAccessTokenSecretKeyName         = "JWTAccessTokenSecretKey"
	JWTAccessTokenSecretKeyEnvName      = "GOPHKEEPER__JWT_ACCESS_TOKEN_SECRET_KEY"
	JWTAccessTokenSecretKeyArgName      = "jwt.token.access.secret.key"
	JWTAccessTokenSecretKeyDefaultValue = "default"

	JWTAccessTokenExpirationTimeName         = "JWTAccessTokenExpirationTime"
	JWTAccessTokenExpirationTimeEnvName      = "GOPHKEEPER__JWT_ACCESS_TOKEN_TTL"
	JWTAccessTokenExpirationTimeArgName      = "jwt.token.access.ttl"
	JWTAccessTokenExpirationTimeDefaultValue = time.Minute * 15

	JWTRefreshTokenSecretKeyName         = "JWTRefreshTokenSecretKey"
	JWTRefreshTokenSecretKeyEnvName      = "GOPHKEEPER__JWT_REFRESH_TOKEN_SECRET_KEY"
	JWTRefreshTokenSecretKeyArgName      = "jwt.token.refresh.secret.key"
	JWTRefreshTokenSecretKeyDefaultValue = "default"

	JWTRefreshTokenExpirationTimeName         = "JWTRefreshTokenExpirationTime"
	JWTRefreshTokenExpirationTimeEnvName      = "GOPHKEEPER__JWT_REFRESH_TOKEN_TTL"
	JWTRefreshTokenExpirationTimeArgName      = "jwt.token.refresh.ttl"
	JWTRefreshTokenExpirationTimeDefaultValue = time.Hour * 24 * 60
)
