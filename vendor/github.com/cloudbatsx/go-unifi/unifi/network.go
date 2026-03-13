package unifi

import (
	"context"
	"encoding/json"
	"fmt"
)

func (dst *Network) MarshalJSON() ([]byte, error) {
	type Alias Network
	aux := &struct {
		*Alias

		WANEgressQOS *emptyStringInt `json:"wan_egress_qos,omitempty"`
	}{
		Alias: (*Alias)(dst),
	}

	if dst.Purpose == "wan" {
		// only send QOS when this is a WAN network
		v := emptyStringInt(dst.WANEgressQOS)
		aux.WANEgressQOS = &v
	}

	b, err := json.Marshal(aux)
	return b, err
}

func (c *Client) DeleteNetwork(ctx context.Context, site, id, name string) error {
	err := c.do(ctx, "DELETE", fmt.Sprintf("s/%s/rest/networkconf/%s", site, id), struct {
		Name string `json:"name"`
	}{
		Name: name,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ListNetwork(ctx context.Context, site string) ([]Network, error) {
	return c.listNetwork(ctx, site)
}

func (c *Client) GetNetwork(ctx context.Context, site, id string) (*Network, error) {
	return c.getNetwork(ctx, site, id)
}

func (c *Client) CreateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data json.RawMessage `json:"data"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/rest/networkconf", site), d, &respBody)
	if err != nil {
		return nil, err
	}

	var network Network
	// try as array first (standard UniFi response)
	var dataArray []Network
	if jsonErr := json.Unmarshal(respBody.Data, &dataArray); jsonErr == nil {
		if len(dataArray) != 1 {
			return nil, &NotFoundError{}
		}
		network = dataArray[0]
	} else {
		// try as single object
		if jsonErr := json.Unmarshal(respBody.Data, &network); jsonErr != nil {
			return nil, fmt.Errorf("unable to unmarshal network response data: %w", jsonErr)
		}
	}

	return &network, nil
}

func (c *Client) UpdateNetwork(ctx context.Context, site string, d *Network) (*Network, error) {
	var respBody struct {
		Meta meta            `json:"meta"`
		Data json.RawMessage `json:"data"`
	}

	err := c.do(ctx, "PUT", fmt.Sprintf("s/%s/rest/networkconf/%s", site, d.ID), d, &respBody)
	if err != nil {
		return nil, err
	}

	var network Network
	// try as array first (standard UniFi response)
	var dataArray []Network
	if jsonErr := json.Unmarshal(respBody.Data, &dataArray); jsonErr == nil {
		if len(dataArray) != 1 {
			return nil, &NotFoundError{}
		}
		network = dataArray[0]
	} else {
		// try as single object
		if jsonErr := json.Unmarshal(respBody.Data, &network); jsonErr != nil {
			return nil, fmt.Errorf("unable to unmarshal network response data: %w", jsonErr)
		}
	}

	return &network, nil
}
