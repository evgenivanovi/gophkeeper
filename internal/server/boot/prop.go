package boot

import (
	"flag"
	"time"

	"github.com/evgenivanovi/gpl/cfg"
	"github.com/evgenivanovi/gpl/std"
)

/* __________________________________________________ */

// Properties ...
type Properties struct {
	/* __________________________________________________ */
	// Configuration related properties
	/* __________________________________________________ */
	ConfigLocation         string
	ConfigLocationFn       func() string
	ConfigLocationProperty cfg.Property

	/* __________________________________________________ */
	// Server related properties
	/* __________________________________________________ */
	HTTPEnabled         bool
	HTTPEnabledProperty cfg.Property
	HTTPPort            int
	HTTPPortProperty    cfg.Property

	HTTPSEnabled         bool
	HTTPSEnabledProperty cfg.Property
	HTTPSPort            int
	HTTPSPortProperty    cfg.Property

	GRPCEnabled         bool
	GRPCEnabledProperty cfg.Property
	GRPCPort            int
	GRPCPortProperty    cfg.Property

	/* __________________________________________________ */
	// Database related properties
	/* __________________________________________________ */
	DSNPostgres         string
	DSNPostgresProperty cfg.Property

	/* __________________________________________________ */
	// Auth related properties
	/* __________________________________________________ */
	JWTAccessTokenSecretKey              string
	JWTAccessTokenSecretKeyProperty      cfg.Property
	JWTAccessTokenExpirationTime         time.Duration
	JWTAccessTokenExpirationTimeProperty cfg.Property

	JWTRefreshTokenSecretKey              string
	JWTRefreshTokenSecretKeyProperty      cfg.Property
	JWTRefreshTokenExpirationTime         time.Duration
	JWTRefreshTokenExpirationTimeProperty cfg.Property
}

// ProvideProperties ...
func ProvideProperties() Properties {

	props := new(Properties)

	props.ConfigLocationProperty = cfg.NewProperty(ConfigLocationPropertyName)
	props.ConfigLocationProperty.BindOne(cfg.NewEnvSource(ConfigLocationEnvName))
	props.ConfigLocationProperty.BindOne(cfg.NewArgSource(ConfigLocationArgName))

	props.ConfigLocationFn = props.ConfigLocationProperty.CalcFn(cfg.FirstStringNotEmpty(std.Empty))

	props.HTTPEnabledProperty = cfg.NewProperty(HTTPEnabledPropertyName)
	props.HTTPEnabledProperty.BindOne(cfg.NewEnvSource(HTTPEnabledEnvName))
	props.HTTPEnabledProperty.BindOne(cfg.NewArgSource(HTTPEnabledArgName))
	props.HTTPEnabledProperty.BindOne(cfg.NewJSONFileSourceWithPath(HTTPEnabledCfgName, props.ConfigLocationFn))

	props.HTTPPortProperty = cfg.NewProperty(HTTPPortPropertyName)
	props.HTTPPortProperty.BindOne(cfg.NewEnvSource(HTTPPortEnvName))
	props.HTTPPortProperty.BindOne(cfg.NewArgSource(HTTPPortArgName))
	props.HTTPPortProperty.BindOne(cfg.NewJSONFileSourceWithPath(HTTPPortCfgName, props.ConfigLocationFn))

	props.HTTPSEnabledProperty = cfg.NewProperty(HTTPSEnabledPropertyName)
	props.HTTPSEnabledProperty.BindOne(cfg.NewEnvSource(HTTPSEnabledEnvName))
	props.HTTPSEnabledProperty.BindOne(cfg.NewArgSource(HTTPSEnabledArgName))
	props.HTTPSEnabledProperty.BindOne(cfg.NewJSONFileSourceWithPath(HTTPSEnabledCfgName, props.ConfigLocationFn))

	props.HTTPSPortProperty = cfg.NewProperty(HTTPSPortPropertyName)
	props.HTTPSPortProperty.BindOne(cfg.NewEnvSource(HTTPSPortEnvName))
	props.HTTPSPortProperty.BindOne(cfg.NewArgSource(HTTPSPortArgName))
	props.HTTPSPortProperty.BindOne(cfg.NewJSONFileSourceWithPath(HTTPSPortCfgName, props.ConfigLocationFn))

	props.GRPCEnabledProperty = cfg.NewProperty(GRPCEnabledPropertyName)
	props.GRPCEnabledProperty.BindOne(cfg.NewEnvSource(GRPCEnabledEnvName))
	props.GRPCEnabledProperty.BindOne(cfg.NewArgSource(GRPCEnabledArgName))
	props.GRPCEnabledProperty.BindOne(cfg.NewJSONFileSourceWithPath(GRPCEnabledCfgName, props.ConfigLocationFn))

	props.GRPCPortProperty = cfg.NewProperty(GRPCPortPropertyName)
	props.GRPCPortProperty.BindOne(cfg.NewEnvSource(GRPCPortEnvName))
	props.GRPCPortProperty.BindOne(cfg.NewArgSource(GRPCPortArgName))
	props.GRPCPortProperty.BindOne(cfg.NewJSONFileSourceWithPath(GRPCPortCfgName, props.ConfigLocationFn))

	props.DSNPostgresProperty = cfg.NewProperty(DSNPostgresName)
	props.DSNPostgresProperty.BindOne(cfg.NewEnvSource(DSNPostgresEnvName))
	props.DSNPostgresProperty.BindOne(cfg.NewArgSource(DSNPostgresArgName))
	props.DSNPostgresProperty.BindOne(cfg.NewJSONFileSourceWithPath(DSNPostgresCfgName, props.ConfigLocationFn))

	props.JWTAccessTokenSecretKeyProperty = cfg.NewProperty(JWTAccessTokenSecretKeyName)
	props.JWTAccessTokenSecretKeyProperty.BindOne(cfg.NewEnvSource(JWTAccessTokenSecretKeyEnvName))
	props.JWTAccessTokenSecretKeyProperty.BindOne(cfg.NewArgSource(JWTAccessTokenSecretKeyArgName))

	props.JWTAccessTokenExpirationTimeProperty = cfg.NewProperty(JWTAccessTokenExpirationTimeName)
	props.JWTAccessTokenExpirationTimeProperty.BindOne(cfg.NewEnvSource(JWTAccessTokenExpirationTimeEnvName))
	props.JWTAccessTokenExpirationTimeProperty.BindOne(cfg.NewArgSource(JWTAccessTokenExpirationTimeArgName))

	props.JWTRefreshTokenSecretKeyProperty = cfg.NewProperty(JWTRefreshTokenSecretKeyName)
	props.JWTRefreshTokenSecretKeyProperty.BindOne(cfg.NewEnvSource(JWTRefreshTokenSecretKeyEnvName))
	props.JWTRefreshTokenSecretKeyProperty.BindOne(cfg.NewArgSource(JWTRefreshTokenSecretKeyArgName))

	props.JWTRefreshTokenExpirationTimeProperty = cfg.NewProperty(JWTRefreshTokenExpirationTimeName)
	props.JWTRefreshTokenExpirationTimeProperty.BindOne(cfg.NewEnvSource(JWTRefreshTokenExpirationTimeEnvName))
	props.JWTRefreshTokenExpirationTimeProperty.BindOne(cfg.NewArgSource(JWTRefreshTokenExpirationTimeArgName))

	flag.Parse()

	props.HTTPEnabled = props.
		HTTPEnabledProperty.
		CalcValue(cfg.FirstBoolOr(HTTPEnabledDefaultValue)).
		GetBool()

	props.HTTPPort = props.
		HTTPPortProperty.
		CalcValue(cfg.FirstInt(HTTPPortDefaultValue)).
		GetInt()

	props.HTTPSEnabled = props.
		HTTPSEnabledProperty.
		CalcValue(cfg.FirstBoolOr(HTTPSEnabledDefaultValue)).
		GetBool()

	props.HTTPSPort = props.
		HTTPSPortProperty.
		CalcValue(cfg.FirstInt(HTTPSPortDefaultValue)).
		GetInt()

	props.GRPCEnabled = props.
		GRPCEnabledProperty.
		CalcValue(cfg.FirstBoolOr(GRPCEnabledDefaultValue)).
		GetBool()

	props.GRPCPort = props.
		GRPCPortProperty.
		CalcValue(cfg.FirstInt(GRPCPortDefaultValue)).
		GetInt()

	props.DSNPostgres = props.
		DSNPostgresProperty.
		Calc(cfg.FirstStringNotEmptyThrow())

	props.JWTAccessTokenSecretKey = props.
		JWTAccessTokenSecretKeyProperty.
		Calc(cfg.FirstStringNotEmpty(JWTAccessTokenSecretKeyDefaultValue))

	props.JWTAccessTokenExpirationTime = props.
		JWTAccessTokenExpirationTimeProperty.
		CalcValue(cfg.FirstDurationOr(JWTAccessTokenExpirationTimeDefaultValue)).
		GetDuration()

	props.JWTRefreshTokenSecretKey = props.
		JWTRefreshTokenSecretKeyProperty.
		Calc(cfg.FirstStringNotEmpty(JWTRefreshTokenSecretKeyDefaultValue))

	props.JWTRefreshTokenExpirationTime = props.
		JWTRefreshTokenExpirationTimeProperty.
		CalcValue(cfg.FirstDurationOr(JWTRefreshTokenExpirationTimeDefaultValue)).
		GetDuration()

	return *props

}

/* __________________________________________________ */
