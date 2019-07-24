// Copyright © 2019 Banzai Cloud
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

package dns

import (
	"context"
	"encoding/json"

	"emperror.dev/errors"
	"github.com/mitchellh/mapstructure"

	"github.com/banzaicloud/pipeline/dns"
	"github.com/banzaicloud/pipeline/internal/cluster/clustersecret"
	"github.com/banzaicloud/pipeline/internal/clusterfeature"
	"github.com/banzaicloud/pipeline/internal/clusterfeature/features"
	"github.com/banzaicloud/pipeline/internal/common"
)

type ClusterSecretStore interface {
	// GetSecret gets a secret for a cluster if exists
	GetSecret(ctx context.Context, clusterID uint, name string) (clustersecret.SecretResponse, error)
}

type externalDnsFeatureSpecProcessor struct {
	secretStore ClusterSecretStore

	logger common.Logger
}

// wrapper struct for handling user inputs
type externalDnsFeatureSpec struct {
	// embedding "real" type
	Overrides dns.ExternalDnsChartValues `json:"overrides"`

	SecretName string `json:"secretName"`
}

// Process method for assembling the "values" for the helm deployment
func (p *externalDnsFeatureSpecProcessor) Process(ctx context.Context, clusterID uint, spec clusterfeature.FeatureSpec) (interface{}, error) {

	rawValues := externalDnsFeatureSpec{}
	if err := mapstructure.Decode(spec, &rawValues); err != nil {

		return nil, errors.WrapIf(err, "could not process feature spec")
	}

	secrets, err := p.secretStore.GetSecret(ctx, clusterID, rawValues.SecretName)
	if err != nil {
		return nil, errors.WrapIf(err, "failed to process feature spec secrets")
	}

	// parse secrets - aws only for the time being
	creds := awsCredentials{}
	if err := mapstructure.Decode(secrets, &creds); err != nil {

		return nil, errors.WrapIf(err, "failed to bind feature spec credentials")
	}

	// set secret values
	rawValues.Overrides.Aws.Credentials = dns.ExternalDnsAwsCredentials{
		AccessKey: creds.AccessKeyID,
		SecretKey: creds.SecretAccessKey,
	}

	values, err := json.Marshal(rawValues.Overrides)
	if err != nil {

		return nil, errors.WrapIf(err, "failed to decode values")
	}

	return values, nil
}

func NewDnsFeatureSpecProcessor(secretStore ClusterSecretStore, logger common.Logger) features.FeatureSpecProcessor {

	return &externalDnsFeatureSpecProcessor{
		logger:      logger,
		secretStore: secretStore,
	}
}

// awsCredentials helper struct for getting secret values
type awsCredentials struct {
	AccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
}