package main

import (
	"fmt"
	"os"

	client "github.com/Maritime-AI/mai-go-client"
)

func main() {
	// NewDefaultClient returns a new default authenticator client
	apiToken := os.Getenv("MAI_API_TOKEN")
	if len(apiToken) == 0 {
		panic("MAI_API_TOKEN is not set")
	}

	sessionID := os.Getenv("MAI_SESSION_ID")
	if len(sessionID) == 0 {
		panic("MAI_SESSION_ID is not set")
	}

	orgRefID := os.Getenv("ORG_REF_ID")
	if len(orgRefID) == 0 {
		panic("ORG_REF_ID is not set")
	}

	cli := client.NewClient(apiToken)

	resp, err := cli.GetConversation(sessionID, orgRefID)
	if err != nil {
		panic(err)
	}

	fmt.Println("messages", resp)

	presp, err := cli.PostChatMessage(orgRefID, &sessionID, "Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Println("messages", presp)
}
