package pixela

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// APIBaseURL is endpoints of base.
const APIBaseURL = "https://pixe.la/v1"

const userToken = "X-USER-TOKEN"

// Client is API client for Pixela user.
type Client struct {
	UserName string
	Token    string
	client   *http.Client
}

// New creates Client object.
func New(userName, token string) *Client {
	cli := &http.Client{
		Timeout: 3 * time.Second,
	}
	return &Client{UserName: userName, Token: token, client: cli}
}

type CreateGraphOpt func(*createGraphParams) error

type GraphResult struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

func TimeZone(tz string) CreateGraphOpt {
	return func(p *createGraphParams) error {
		p.Timezone = tz
		return nil
	}
}

func (c *Client) CreateGraph(
	ctx context.Context,
	id GraphID, name, unit string, gtype GraphType,
	color GraphColor, opts ...CreateGraphOpt,
) (*GraphResult, error) {
	p := &createGraphParams{
		ID:    id,
		Name:  name,
		Unit:  unit,
		Type:  gtype,
		Color: color,
	}
	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil, err
		}
	}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	ep := fmt.Sprintf("%s/users/%s/graphs", APIBaseURL, c.UserName)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, ep, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add(userToken, c.Token)
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var gr GraphResult
	if err := json.NewDecoder(rsp.Body).Decode(&gr); err != nil {
		return nil, err
	}
	return &gr, nil
}

func (c *Client) DeleteGraph(ctx context.Context, id GraphID) (*GraphResult, error) {
	ep := fmt.Sprintf("%s/users/%s/graphs/%s", APIBaseURL, c.UserName, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(userToken, c.Token)
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var gr GraphResult
	if err := json.NewDecoder(rsp.Body).Decode(&gr); err != nil {
		return nil, err
	}
	return &gr, nil
}

func (c *Client) GetGraph(ctx context.Context, id GraphID) (*GraphDefinition, error) {
	ep := fmt.Sprintf("%s/users/%s/graphs/%s/graph-def", APIBaseURL, c.UserName, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(userToken, c.Token)
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var gd GraphDefinition
	if err := json.NewDecoder(rsp.Body).Decode(&gd); err != nil {
		return nil, err
	}
	return &gd, nil
}

func (c *Client) GetGraphs(ctx context.Context) ([]*GraphDefinition, error) {
	ep := fmt.Sprintf("%s/users/%s/graphs", APIBaseURL, c.UserName)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ep, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(userToken, c.Token)
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var gd struct {
		Graphs []*GraphDefinition `json:"graphs"`
	}
	if err := json.NewDecoder(rsp.Body).Decode(&gd); err != nil {
		return nil, err
	}
	return gd.Graphs, nil
}

func (c *Client) UpdateGraph(
	ctx context.Context,
	gd *GraphDefinition,
) (*GraphResult, error) {
	b, err := json.Marshal(gd)
	if err != nil {
		return nil, err
	}
	ep := fmt.Sprintf("%s/users/%s/graphs/%s", APIBaseURL, c.UserName, gd.ID)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, ep, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Add(userToken, c.Token)
	rsp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	var gr GraphResult
	if err := json.NewDecoder(rsp.Body).Decode(&gr); err != nil {
		return nil, err
	}
	return &gr, nil
}
