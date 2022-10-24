// Copyright 2018 Sergey Novichkov. All rights reserved.
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package validator

import "github.com/gozix/di"

// tagConfigurator is tag to mark specific validations.
const tagConfigurator = "validator.configurator"

// AsConfigurator is syntax sugar for the di container.
func AsConfigurator() di.Tags {
	return di.Tags{{
		Name: tagConfigurator,
	}}
}

func withConfigurator() di.Modifier {
	return di.WithTags(tagConfigurator)
}
