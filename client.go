package client

import (
	"github.com/Maritime-AI/mai-go-client/aiapi"
	"github.com/Maritime-AI/mai-go-client/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	protocol = "https"
	basePath = "api/v1"
	host     = "mai.10ure.com"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}

func (c *Client) PostChatMessage(orgRefID string, sessionID *string,
	message string) (*models.ChatMessageResponse, error) {
	transport := client.New(host, basePath, []string{protocol})

	// Initialize the formats registry
	formats := strfmt.Default

	// Initialize the client
	client := aiapi.New(transport, formats)

	auth := runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
		return req.SetHeaderParam("Authorization", c.apiToken)
	})

	params := aiapi.NewPostPartnersOrganizationsIDConverseParams().WithBody(&models.ChatMessageParams{
		Message:   message,
		SessionID: sessionID,
	}).
		WithID(orgRefID)

	resp, err := client.PostPartnersOrganizationsIDConverse(params, auth)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

func (c *Client) GetConversation(sessionID, orgRefID string) (*models.ChatMessageResponse, error) {
	transport := client.New(host, basePath, []string{protocol})

	// Initialize the formats registry
	formats := strfmt.Default

	// Initialize the client
	client := aiapi.New(transport, formats)

	auth := runtime.ClientAuthInfoWriterFunc(func(req runtime.ClientRequest, _ strfmt.Registry) error {
		return req.SetHeaderParam("Authorization", c.apiToken)
	})

	params := aiapi.NewGetPartnersOrganizationsIDConversationSessionIDParams().
		WithSessionID(sessionID).
		WithID(orgRefID)

	resp, err := client.GetPartnersOrganizationsIDConversationSessionID(params, auth)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
