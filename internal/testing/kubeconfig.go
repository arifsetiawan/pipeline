// Copyright © 2020 Banzai Cloud
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

package testing

import (
	"io/ioutil"
	"os"
	"testing"
)

func KubeConfigFromEnv(t *testing.T) []byte {
	kubeConfigFile := os.Getenv("KUBECONFIG")
	if kubeConfigFile == "" {
		t.Skip("skipping as Kubernetes config was not provided")
	}
	kubeConfigBytes, err := ioutil.ReadFile(kubeConfigFile)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	return kubeConfigBytes
}