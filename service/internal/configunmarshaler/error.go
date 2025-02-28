// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configunmarshaler // import "go.opentelemetry.io/collector/service/internal/configunmarshaler"

import (
	"fmt"
	"reflect"

	"go.opentelemetry.io/collector/config"
)

func errorUnknownType(component string, id config.ComponentID, factories []reflect.Value) error {
	return fmt.Errorf("unknown %s type: %q for id: %q (valid values: %v)", component, id.Type(), id, factories)
}

func errorUnmarshalError(component string, id config.ComponentID, err error) error {
	return fmt.Errorf("error reading %s configuration for %q: %w", component, id, err)
}
