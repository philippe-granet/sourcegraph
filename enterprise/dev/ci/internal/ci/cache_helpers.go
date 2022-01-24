package ci

import "github.com/sourcegraph/sourcegraph/enterprise/dev/ci/internal/buildkite"

func cacheYarn() buildkite.StepOpt {
	return buildkite.Cache(&buildkite.CacheOptions{
		ID:          "node_modules",
		Key:         "cache-node_modules-{{ checksum 'yarn.lock' }}",
		RestoreKeys: []string{"cache-node_modules-{{ checksum 'yarn.lock' }}"},
		Paths:       []string{"node_modules", "client/extension-api/node_modules", "client/eslint-plugin-sourcegraph/node_modules"},
	})
}
