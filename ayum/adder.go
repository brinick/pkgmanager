package ayum

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type repoer interface {
	Filename() string
	fmt.Stringer
}

type rpmRepoAdder interface {
	AddRemoteRepos([]repoer) error
}

type rpmRepoAdd struct {
	basedir string
}

// AddRemoteRepos configures the ayum installation with the provided remote repositories
func (r *rpmRepoAdd) AddRemoteRepos(repos []repoer) error {
	for _, repo := range repos {
		repoConf := filepath.Join(r.basedir, "ayum/etc/yum.repos.d", repo.Filename())
		if err := ioutil.WriteFile(repoConf, []byte(repo.String()), 0774); err != nil {
			return fmt.Errorf("could not configure remote repo %s (%w)", repo.Filename(), err)
		}
	}

	return nil
}
