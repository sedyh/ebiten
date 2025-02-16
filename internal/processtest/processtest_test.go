// Copyright 2021 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !android && !ios && !js

package processtest_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	exec "golang.org/x/sys/execabs"
)

func TestPrograms(t *testing.T) {
	dir := "testdata"
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range ents {
		if e.IsDir() {
			continue
		}
		n := e.Name()
		if !strings.HasSuffix(n, ".go") {
			continue
		}

		t.Run(n, func(t *testing.T) {
			cmd := exec.Command("go", "run", filepath.Join(dir, n))
			stderr := &bytes.Buffer{}
			cmd.Stderr = stderr
			if err := cmd.Run(); err != nil {
				t.Errorf("%v\n%s", err, stderr)
			}
		})
	}
}
