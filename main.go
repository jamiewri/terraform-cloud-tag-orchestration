package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-tfe"
)

func main () {

	// Get tags from positional args
	t := os.Args[1:]
	ts := strings.Join(t, ",")

	// Get org and token from environment vars
	org := os.Getenv("TFC_ORGANIZATION")
	token := os.Getenv("TFC_ORGANIZATION_TOKEN")

	// Output expected behaviour
	fmt.Println("Terraform cloud orchestration with tags")
	fmt.Println("Organization:", org)
	fmt.Println("Searching for workspaces with the following tags:", ts)

	// Init go-tfe config
	config := &tfe.Config{
		Token: token,
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create context
	ctx := context.Background()

	listOptions := &tfe.ListOptions{
		PageNumber: 1,
		PageSize: 50,
	}

	workspaceListOptions := &tfe.WorkspaceListOptions{
		ListOptions: *listOptions,
		Tags: ts,
	}

	w, err := client.Workspaces.List(ctx, org, workspaceListOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check if any workspaces matched, if not, exit.
	if len(w.Items) == 0 {
		fmt.Println("No workspaces found with tags: ", ts)
		os.Exit(1)
	}

	// Output list of workspaces with matching tags
	for i := range w.Items {
		fmt.Println(w.Items[i].Name)
	}
}