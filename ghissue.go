package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/user"

	"bitbucket.org/mischief/libauth"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	branch = flag.String("b", "master", "Github branch")
	title = flag.String("t", "", "Title of the issue")
)

func flagFatal() {
	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.Parse()
	if flag.Lookup("-h") != nil {
		flagFatal()
	}
	if flag.NArg() != 1 {
		log.Fatal("No reponame specified")
	}
	if *title == "" { 
		log.Fatal("No title set for issue (-t <string>)")
	}
	usr, err := user.Current()
	token, err := libauth.Getuserpasswd("proto=pass service=github role=client user=%s", usr.Username)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token.Password,
	})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	createIssue(ctx, client.Issues)
}
