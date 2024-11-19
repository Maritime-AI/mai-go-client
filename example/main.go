package main

import (
	"fmt"
	"os"

	"github.com/Maritime-AI/mai-go-client/client"
)

func main() {
	// NewDefaultClient returns a new default authenticator client
	host := "mai.10ure.com"

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

	cli := client.NewClient(host, apiToken)

	resp, err := cli.GetConversation(sessionID, orgRefID)
	if err != nil {
		panic(err)
	}

	fmt.Println("messages", resp)
}
