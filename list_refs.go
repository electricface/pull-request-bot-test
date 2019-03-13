package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	token, err := getToken()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tc := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		}))

	client := github.NewClient(tc)

	owner := "electricface"
	repo := "pull-request-bot-test"

	refs, _, err := client.Git.ListRefs(ctx, owner, repo, &github.ReferenceListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, ref := range refs {
		log.Println(ref.GetURL())
	}
}

func getToken() (string, error) {
	tokenFile := filepath.Join(os.Getenv("HOME"), ".prbot-token")
	tokenData, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}
	tokenData = bytes.TrimSpace(tokenData)
	return string(tokenData), nil
}
