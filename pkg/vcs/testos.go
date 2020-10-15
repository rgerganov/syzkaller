// Copyright 2019 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package vcs

import (
	"fmt"
	"io"
)

type testos struct {
	*git
}

var _ ConfigMinimizer = new(testos)

func newTestos(dir string, opts []RepoOpt) *testos {
	return &testos{
		git: newGit(dir, nil, opts),
	}
}

func (ctx *testos) PreviousReleaseTags(commit string) ([]string, error) {
	return ctx.git.previousReleaseTags(commit, false, false)
}

func (ctx *testos) EnvForCommit(binDir, commit string, kernelConfig []byte) (*BisectEnv, error) {
	return &BisectEnv{KernelConfig: kernelConfig}, nil
}

func (ctx *testos) Minimize(original, baseline []byte, trace io.Writer,
	pred func(test []byte) (BisectResult, error)) ([]byte, error) {
	switch string(baseline) {
	case "minimize-fails":
		return nil, fmt.Errorf("minimization failure")
	case "minimize-succeeds":
		config := []byte("new-minimized-config")
		pred(config)
		return config, nil
	case "baseline-broken-build":
		return []byte("broken-build"), nil
	default:
		return original, nil
	}
}
