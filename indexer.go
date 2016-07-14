package main

import (
	"fmt"
	"github.com/cep21/xdgbasedir"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
)

// GetToken gets the access token from the configuration directory
func GetToken() (token string, err error) {
	var data []byte
	file, err := xdgbasedir.GetConfigFileLocation("github-indexer/token")
	if err != nil {
		return token, err
	}

	data, err = ioutil.ReadFile(file)
	if err != nil {
		return token, err
	}

	token = string(data)
	return token, nil
}

// GetClient returns a usable Github API client
func GetClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)
	return client
}

func main() {
	token, err := GetToken()
	if err != nil {
		log.Fatal(err)
	}

	client := GetClient(token)
	repos, _, err := client.Repositories.List("", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}
