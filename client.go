package client

import (
	"bytes"
	"context"
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
	defaultProtocol = "https"
	defaultHost     = "mai.10ure.com"
	basePath        = "api/v1"
	defaultTimeout  = 60 * time.Second
)

// Option is a function that configures the client
// It is used to set optional parameters for the client
type Option func(c *Client)

// WithTimeout sets the timeout for the client
func WithTimeout(t time.Duration) Option {
	return func(c *Client) {
		c.timeout = t
	}
}

// WithHost sets the host for the client
func WithHost(h string) Option {
	return func(c *Client) {
		c.host = h
	}
}

// WithProtocol sets the protocol for the client
func WithProtocol(p string) Option {
	return func(c *Client) {
		c.protocol = p
	}
}

type Client struct {
	host     string
	protocol string
	apiToken string
	timeout  time.Duration
}

// NewClient returns a new maritime ai client
func NewClient(apiToken string, opts ...Option) *Client {
	c := &Client{
		protocol: defaultProtocol,
		host:     defaultHost,
		apiToken: apiToken,
		timeout:  defaultTimeout,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) PostPartnerCrew(ctx context.Context, params models.CrewParams) error {

	url := fmt.Sprintf("%s://%s/api/v1/partners/crew", c.protocol, c.host)

	data, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to marshal crew params: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Authorization", c.apiToken)
	req.Header.Set("Content-Type", "application/json")

	httpCli := http.Client{
		Timeout: c.timeout,
	}

	resp, err := httpCli.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send http request: %w", err)
	}

	var respData []byte
	if resp.Body != nil {
		defer resp.Body.Close()

		respData, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}
	}

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("expected 204 status code got %d: %s", resp.StatusCode, string(respData))
	}

	return nil
}

// PostChatMessage sends a chat message to the oceo maritime ai api
func (c *Client) PostChatMessage(orgRefID string, sessionID *string,
	message string) (*models.ChatMessageResponse, error) {

	url := fmt.Sprintf("%s://%s/api/v1/partners/organizations/%s/converse", c.protocol, c.host, orgRefID)

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
		Timeout: c.timeout,
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

// GetConversation retrieves the conversation for a given session ID and organization reference ID
func (c *Client) GetConversation(sessionID, orgRefID string) (*models.ChatMessageResponse, error) {
	transport := client.New(c.host, basePath, []string{c.protocol})

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
	WithTimeout(c.timeout)

	resp, err := client.GetPartnersOrganizationsIDConversationSessionID(params, auth)
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}

// PostUserChatMessage sends a user chat message to the oceo maritime ai api
func (c *Client) PostUserChatMessage(userID string, sessionID *string,
	message string) (*models.ChatMessageResponse, error) {

	url := fmt.Sprintf("%s://%s/api/v1/partners/users/%s/converse", c.protocol, c.host, userID)

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
		Timeout: c.timeout,
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

// GetUserConversation retrieves the conversation for a given session ID and user reference ID
func (c *Client) GetUserConversation(sessionID, userID string) (*models.ChatMessageResponse, error) {
	url := fmt.Sprintf("%s://%s/api/v1/partners/users/%s/conversation/%s", c.protocol, c.host, userID, sessionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	req.Header.Set("Authorization", c.apiToken)
	req.Header.Set("Content-Type", "application/json")

	httpCli := http.Client{
		Timeout: c.timeout,
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
