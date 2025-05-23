/*
Copyright 2017 The Kubernetes Authors.

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

package azure

import (
	"path"
	"runtime"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/stretchr/testify/assert"
)

func TestGetCloudConfiguration(t *testing.T) {
	tests := map[string]struct {
		cloudName string
		expected  cloud.Configuration
	}{
		"AzureChinaCloud":   {"AzureChinaCloud", cloud.AzureChina},
		"AzurePublicCloud":  {"", cloud.AzurePublic},
		"AzureUSGovernment": {"AzureUSGovernmentCloud", cloud.AzureGovernment},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cloudCfg, err := getCloudConfiguration(test.cloudName)
			if err != nil {
				t.Errorf("got unexpected err %v", err)
			}
			if cloudCfg.ActiveDirectoryAuthorityHost != test.expected.ActiveDirectoryAuthorityHost {
				t.Errorf("got %v, want %v", cloudCfg, test.expected)
			}
		})
	}
}

func TestOverrideConfiguration(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := path.Join(path.Dir(filename), "fixtures/config_test.json")
	cfg, err := getConfig(configFile, "subscription-override", "rg-override", "", "aad-endpoint-override")
	if err != nil {
		t.Errorf("got unexpected err %v", err)
	}
	assert.Equal(t, cfg.SubscriptionID, "subscription-override")
	assert.Equal(t, cfg.ResourceGroup, "rg-override")
	assert.Equal(t, cfg.ActiveDirectoryAuthorityHost, "aad-endpoint-override")
}
