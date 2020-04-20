package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/sarulabs/di/v2"

	gzUT "github.com/gozix/universal-translator/v2"
)

type (
	// Bundle implements the glue.Bundle interface.
	Bundle struct{}

	// Validate is type alias of validator.Validate.
	Validate = validator.Validate

	// Configurator is register default translations func interface.
	Configurator = func(*validator.Validate) error

	// ValidationErrors is type alias of validator.ValidationErrors.
	ValidationErrors = validator.ValidationErrors

	// ValidationErrorsTranslations is type alias of validator.ValidationErrorsTranslations.
	ValidationErrorsTranslations = validator.ValidationErrorsTranslations
)

const (
	// BundleName is default definition name.
	BundleName = "validator"

	// DefEnLocaleConfigurator is english translations configurator.
	DefEnLocaleConfigurator = "validator.configurator.en-locale"

	// DefJsonTagNameConfigurator is json tag name configurator.
	DefJsonTagNameConfigurator = "validator.configurator.json-tag-name"

	// TagConfigurator is tag to mark specific validations.
	TagConfigurator = "validator.configurator"
)

// NewBundle create bundle instance.
func NewBundle() *Bundle {
	return &Bundle{}
}

// Key implements the glue.Bundle interface.
func (b *Bundle) Name() string {
	return BundleName
}

// Build implements the glue.Bundle interface.
func (b *Bundle) Build(builder *di.Builder) error {
	return builder.Add(
		di.Def{
			Name: BundleName,
			Build: func(ctn di.Container) (_ interface{}, err error) {
				var configurators = make([]Configurator, 0, 4)
				for name, def := range ctn.Definitions() {
					for _, tag := range def.Tags {
						if TagConfigurator != tag.Name {
							continue
						}

						var configurator Configurator
						if err = ctn.Fill(name, &configurator); err != nil {
							return nil, err
						}

						configurators = append(configurators, configurator)
					}
				}

				var validate = validator.New()
				for _, configurator := range configurators {
					if err = configurator(validate); err != nil {
						return nil, err
					}
				}

				return validate, nil
			},
		},
		di.Def{
			Name: DefEnLocaleConfigurator,
			Tags: []di.Tag{{
				Name: TagConfigurator,
			}},
			Build: func(ctn di.Container) (_ interface{}, err error) {
				var translator *gzUT.UniversalTranslator
				if err = ctn.Fill(gzUT.BundleName, &translator); err != nil {
					return nil, err
				}

				return func(v *validator.Validate) error {
					var t, founded = translator.GetTranslator("en")
					if !founded {
						return nil
					}

					return en.RegisterDefaultTranslations(v, t)
				}, nil
			},
		},
		di.Def{
			Name: DefJsonTagNameConfigurator,
			Tags: []di.Tag{{
				Name: TagConfigurator,
			}},
			Build: func(ctn di.Container) (interface{}, error) {
				return func(v *validator.Validate) error {
					v.RegisterTagNameFunc(func(field reflect.StructField) string {
						var name = strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
						if name == "-" {
							return ""
						}

						return name
					})

					return nil
				}, nil
			},
		},
	)
}

// DependsOn implements the glue.DependsOn interface.
func (b *Bundle) DependsOn() []string {
	return []string{gzUT.BundleName}
}
