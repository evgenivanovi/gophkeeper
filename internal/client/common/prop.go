package common

import (
	"context"
	"errors"

	"github.com/evgenivanovi/gpl/cfg"
)

/* __________________________________________________ */

type Options struct {
	Config  string
	Address string
}

func NewOptions() (Options, error) {
	options := Options{}

	/* __________________________________________________ */

	ConfigProp := cfg.NewProperty(ConfigLocationPropertyName)
	ConfigProp.BindOne(cfg.NewEnvSource(ConfigLocationEnvName))
	ConfigProp.BindOne(cfg.NewArgSource(ConfigLocationArgName))
	ConfigProp.BindOne(cfg.NewValueSource(MustBuildConfigPath()))

	ConfigValue, err := ConfigProp.CalcElse(cfg.FirstStringNotEmptyElse())
	if err != nil {
		return NewEmptyOptions(), err
	}
	options.Config = ConfigValue

	/* __________________________________________________ */

	AddressProp := cfg.NewProperty(AddressPropertyName)
	AddressProp.BindOne(cfg.NewEnvSource(AddressEnvName))

	AddressValue, err := AddressProp.CalcElse(cfg.FirstStringNotEmptyElse())
	if err != nil {
		return NewEmptyOptions(), err
	}
	options.Address = AddressValue

	/* __________________________________________________ */

	return options, nil
}

func NewEmptyOptions() Options {
	return Options{}
}

/* __________________________________________________ */

type optionsContextKey string

const optionsCtxKey optionsContextKey = "ctx.options"

/* __________________________________________________ */

func OptionsFromCtx(ctx context.Context) (Options, error) {
	if value, ok := ctx.Value(optionsCtxKey).(Options); ok {
		return value, nil
	}
	return NewEmptyOptions(), errors.New("no options found")
}

func MustOptionsFromCtx(ctx context.Context) Options {
	options, err := OptionsFromCtx(ctx)
	if err != nil {
		panic(err)
	}
	return options
}

func OptionsWithCtx(ctx context.Context, options Options) context.Context {
	value, err := OptionsFromCtx(ctx)
	if err != nil && value == options {
		return ctx
	}
	return context.WithValue(ctx, optionsCtxKey, options)
}

/* __________________________________________________ */
