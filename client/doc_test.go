// Copyright (c) 2020-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package client_test

import (
	"context"
	"fmt"
	"log"

	"github.com/mattermost/mattermost/server/public/model"

	"github.com/mattermost/mattermost-plugin-tech-develop/client"
)

func Example() {
	ctx := context.Background()

	client4 := model.NewAPIv4Client("http://localhost:8065")
	_, _, err := client4.Login(context.Background(), "test@example.com", "testtest")
	if err != nil {
		log.Fatal(err)
	}

	c, err := client.New(client4)
	if err != nil {
		log.Fatal(err)
	}

	playbookRunID := "h4n3h7s1qjf5pkis4dn6cuxgwa"
	playbookRun, err := c.PlaybookRuns.Get(ctx, playbookRunID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Playbook Run Name: %s\n", playbookRun.Name)
}
