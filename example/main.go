package main

import (
	"fmt"
	"os"

	"github.com/Maritime-AI/mai-go-client/aiapi"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/runtime/client"
)

func main() {
	// NewDefaultClient returns a new default authenticator client
	host := "mai.10ure.com"
	basePath := "api/v1"

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

	transport := client.New(host, basePath, []string{"http"})

	// Initialize the formats registry
	formats := strfmt.Default

	// Initialize the client
	client := aiapi.New(transport, formats)

	auth := runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
		return req.SetHeaderParam("Authorization", apiToken)
	})

	params := aiapi.NewGetPartnersOrganizationsIDConversationSessionIDParams().
		WithSessionID(sessionID).
		WithID(orgRefID)

	// Use the client...
	resp, err := client.GetPartnersOrganizationsIDConversationSessionID(params, auth)
	if err != nil {
		panic(err)
	}

	fmt.Println("messages", resp.Payload.Messages)
}
