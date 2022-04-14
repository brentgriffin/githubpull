package main

import (
	"fmt"
	"github.com/google/go-github/v33/github"
	"golang.org/x/net/context"
)

func main() {
	client := github.NewClient(nil)

/*	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	if err != nil {
		panic(err)
	}
	fmt.Println("repos len = ", len(repos))
	for _, repo := range repos {
		fmt.Println("Name=", *repo.Name, "|", "Desc=", *repo.Description)
	}

	// list all organizations for user "brentgriffin"
	orgs, _, err := client.Organizations.List(context.Background(), "brentgriffin", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("len=",len(orgs))
	for _, org := range orgs {
		fmt.Println("org = ", org)
	}

	fmt.Printf("orgs = ", orgs)*/

	file, dir, _, err := client.Repositories.GetContents(context.Background(), "brentgriffin", "githubpull", "/", nil)
	if err != nil {
		panic(err)
	}
	if file != nil {
		fmt.Println("file =", file)
	}
	if dir != nil {
		fmt.Println("len of dir", len(dir))
		for _, file := range dir {
			fmt.Println("file=", file.GetName())
			content, err := file.GetContent()
			if err != nil {
				panic(err)
			}
			fmt.Println("content=", content)
		}
	}
}