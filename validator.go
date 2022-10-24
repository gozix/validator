// Copyright 2018 Sergey Novichkov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package validator

import (
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"

	"github.com/gozix/di"
	gzUT "github.com/gozix/universal-translator/v3"
)

type (
	// Bundle implements the glue.Bundle interface.
	Bundle struct{}

	// Configurator is register default translations func interface.
	Configurator = func(*validator.Validate) error
)

// BundleName is default definition name.
const BundleName = "validator"

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return &Bundle{}
}

func (b *Bundle) Name() string {
	return BundleName
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder di.Builder) error {
	return builder.Apply(
		di.Provide(b.provideUniversalTranslator, di.Constraint(0, withConfigurator())),
		di.Provide(b.provideEnLocaleConfigurator, AsConfigurator()),
		di.Provide(b.provideJsonTagNameConfigurator, AsConfigurator()),
	)
}

// DependsOn implements the glue.DependsOn interface.
func (b *Bundle) DependsOn() []string {
	return []string{
		gzUT.BundleName,
	}
}

func (b *Bundle) provideUniversalTranslator(configurators []Configurator) (_ *validator.Validate, err error) {
	var validate = validator.New()
	for _, configurator := range configurators {
		if err = configurator(validate); err != nil {
			return nil, err
		}
	}

	return validate, nil
}

func (b *Bundle) provideEnLocaleConfigurator(translator *ut.UniversalTranslator) Configurator {
	return func(v *validator.Validate) error {
		var t, founded = translator.GetTranslator("en")
		if !founded {
			return nil
		}

		return en.RegisterDefaultTranslations(v, t)
	}
}

func (b *Bundle) provideJsonTagNameConfigurator() Configurator {
	return func(v *validator.Validate) error {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			var name = strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})

		return nil
	}
}
