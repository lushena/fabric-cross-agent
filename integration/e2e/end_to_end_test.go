/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package e2e

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
)

func TestE2E(t *testing.T) {
	t.Run("Base", func(t *testing.T) {
		configPath := "/opt/gopath/src/github.com/hyperledger/fabric-cross-agent/config/config_e2e.yaml"
		Run(t, config.FromFile(configPath))
	})
}
