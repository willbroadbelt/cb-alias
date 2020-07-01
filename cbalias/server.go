package cbalias

import (
	"github.com/go-git/go-git"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	RepoPath = "tmp"
	ConfigName = "products.yml"
)

type Products map[string]map[string]Version

type Version struct {
	Release string `yaml:release,omitempty`
	Stable  string `yaml:stable,omitempty`
}

func ParseYaml(data []byte) (Products, error) {
	var products Products
	err := yaml.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetConfigRepo() error {
	log.Printf("git clone https://github.com/willbroadbelt/cb-alias.git")

	_, err := git.PlainClone(RepoPath, false, &git.CloneOptions{
		URL:               "https://github.com/willbroadbelt/cb-alias.git",
		Progress:          os.Stdout,
	})

	log.Printf("PlainCLone err %v", err)

	if err == git.ErrRepositoryAlreadyExists {
		return pullConfigRepo()
	}
	return err
}

func pullConfigRepo() error {
	r, err := git.PlainOpen(RepoPath)
	if err != nil {
		return err
	}
	
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	
	log.Printf("Pulling products repo")
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err == git.NoErrAlreadyUpToDate {
		log.Printf("%v", err)
		return nil
	}
	return err
}
