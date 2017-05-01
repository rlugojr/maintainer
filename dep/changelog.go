// Copyright © 2017 Maintainer Authors
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

package dep

import "errors"

const (
	nameGitHubChangelogGenerator string = "github_changelog_generator"
)

// GitHubChangelogGenerator is the type for github_changelog_generator.
// See https://github.com/skywinder/Github-Changelog-Generator.
type GitHubChangelogGenerator struct {
	Dependency
	isInstalled bool
}

// NewGitHubChangelogGenerator returns a new GitHubChangelogGenerator instance.
func NewGitHubChangelogGenerator() *GitHubChangelogGenerator {
	return &GitHubChangelogGenerator{
		isInstalled: false,
	}
}

// Name implements Dependency.Name().
func (g *GitHubChangelogGenerator) Name() string {
	return nameGitHubChangelogGenerator
}

// IsInstalled implements Dependency.IsInstalled().
func (g *GitHubChangelogGenerator) IsInstalled() (bool, error) {
	return false, errors.New("Not Implemented")
}

// Install implements Dependency.Install().
func (g *GitHubChangelogGenerator) Install() error {
	return errors.New("Not Implemented")
}