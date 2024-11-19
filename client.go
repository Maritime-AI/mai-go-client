package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Maritime-AI/mai-go-client/aiapi"
	"github.com/Maritime-AI/mai-go-client/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

const (
	protocol = "http"
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

	url := fmt.Sprintf("https://%s/api/v1/partners/organizations/%s/converse", host, orgRefID)
	fmt.Print("somshie")
	data, err := json.Marshal(&models.ChatMessageParams{
		Message:   message,
		SessionID: sessionID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal chat message params: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Authorization", c.apiToken)
	req.Header.Set("Content-Type", "application/json")

	httpCli := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %w", err)
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var respData models.ChatMessageResponse
	if err = json.Unmarshal(bs, &respData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &respData, nil
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
