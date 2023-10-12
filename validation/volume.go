/*
   Copyright 2020 The Compose Specification Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package validation

import (
	"fmt"

	"github.com/compose-spec/compose-go/tree"
)

func checkVolume(value any, p tree.Path) error {
	v, ok := value.(map[string]any)
	if !ok {
		return fmt.Errorf("expected volume, got %s", value)
	}

	err := checkExternal(v, p)
	if err != nil {
		return err
	}
	return nil
}
