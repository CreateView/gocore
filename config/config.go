package config

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-envconfig"
)

func Load(cfg interface{}, fields ...*validation.FieldRules) error {
	if err := load(cfg, nil); err != nil {
		return errors.Wrapf(err, "failed to load config")
	}

	if err := validate(cfg, fields...); err != nil {
		return errors.Wrapf(err, "failed to validate config")
	}

	return nil
}

func validate(cfg interface{}, fields ...*validation.FieldRules) error {
	return validation.ValidateStruct(
		cfg,
		fields...,
	)
}

func load(cfg interface{}, l envconfig.Lookuper) error {
	if l == nil {
		if err := envconfig.Process(context.TODO(), cfg); err != nil {
			return errors.Wrapf(err, "process from environment variables")
		}
	} else {
		if err := envconfig.ProcessWith(context.TODO(), cfg, l); err != nil {
			return errors.Wrapf(err, "process with lookuper from environment variables")
		}
	}

	return nil
}
