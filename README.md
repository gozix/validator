# GoZix Validator

[documentation-img]: https://img.shields.io/badge/godoc-reference-blue.svg?color=24B898&style=for-the-badge&logo=go&logoColor=ffffff
[documentation-url]: https://pkg.go.dev/github.com/gozix/validator/v3
[license-img]: https://img.shields.io/github/license/gozix/validator.svg?style=for-the-badge
[license-url]: https://github.com/gozix/validator/blob/master/LICENSE
[release-img]: https://img.shields.io/github/tag/gozix/validator.svg?label=release&color=24B898&logo=github&style=for-the-badge
[release-url]: https://github.com/gozix/validator/releases/latest
[build-status-img]: https://img.shields.io/github/actions/workflow/status/gozix/validator/go.yml?logo=github&style=for-the-badge
[build-status-url]: https://github.com/gozix/validator/actions
[go-report-img]: https://img.shields.io/badge/go%20report-A%2B-green?style=for-the-badge
[go-report-url]: https://goreportcard.com/report/github.com/gozix/validator
[code-coverage-img]: https://img.shields.io/codecov/c/github/gozix/validator.svg?style=for-the-badge&logo=codecov
[code-coverage-url]: https://codecov.io/gh/gozix/validator

[![License][license-img]][license-url]
[![Documentation][documentation-img]][documentation-url]

[![Release][release-img]][release-url]
[![Build Status][build-status-img]][build-status-url]
[![Go Report Card][go-report-img]][go-report-url]
[![Code Coverage][code-coverage-img]][code-coverage-url]

The bundle provide a validator integration to GoZix application.

## Installation

```shell
go get github.com/gozix/validator/v3
```

## Dependencies

* [universal-translator](https://github.com/gozix/universal-translator)

## Built-in Tags

| Symbol                     | Value                                | Description                                                  | 
|----------------------------|--------------------------------------|--------------------------------------------------------------|
| TagConfigurator            | validator.configurator               | Tag to mark specific configurator of validator.Validate      |
| TagEnLocaleConfigurator    | validator.configurator.en-locale     | Tag to mark en locale configurator of validator.Validate     |
| TagJsonTagNameConfigurator | validator.configurator.json-tag-name | Tag to mark json tag name configurator of validator.Validate |

## Documentation

You can find documentation on [pkg.go.dev][documentation-url] and read source code if needed.

## Questions

If you have any questions, feel free to create an issue.
