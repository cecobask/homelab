package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"slices"
)

type Client struct {
	client  *http.Client
	baseURL string
	token   string
	logger  *slog.Logger
}

type Device struct {
	ID       string   `json:"nodeId"`
	Hostname string   `json:"hostname"`
	Tags     []string `json:"tags"`
}

type Devices []Device

type DeviceFilters struct {
	ids       []string
	hostnames []string
	tags      []string
}

const (
	endpointListDevices  = "/tailnet/%s/devices"
	endpointDeleteDevice = "/device/%s"
)

func NewClient(baseURL string, logger *slog.Logger) *Client {
	return &Client{
		client:  http.DefaultClient,
		baseURL: baseURL,
		token:   os.Getenv("TAILSCALE_TOKEN"),
		logger:  logger.With(slog.String("scope", "tailscale")),
	}
}

func NewDeviceFilters(ids, hostnames, tags []string) *DeviceFilters {
	return &DeviceFilters{
		ids:       ids,
		hostnames: hostnames,
		tags:      tags,
	}
}

func (ts *Client) ListDevices(ctx context.Context, tailnetName string, filters DeviceFilters) (Devices, error) {
	req, err := ts.request(ctx, http.MethodGet, fmt.Sprintf(endpointListDevices, tailnetName), nil)
	if err != nil {
		return nil, err
	}
	resp, err := ts.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list devices: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		ts.logger.Error("failed to list devices", slog.String("status", resp.Status))
		return nil, fmt.Errorf("received unhealthy http response status code")
	}
	var result struct {
		Devices Devices `json:"devices"`
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal devices: %v", err)
	}
	return result.Devices.filter(filters, ts.logger), nil
}

func (ts *Client) DeleteDevice(ctx context.Context, id string) error {
	ts.logger.Debug("deleting device", slog.String("id", id))
	req, err := ts.request(ctx, http.MethodDelete, fmt.Sprintf(endpointDeleteDevice, id), nil)
	if err != nil {
		return err
	}
	resp, err := ts.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete device: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		ts.logger.Error("failed to delete device", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	ts.logger.Info("deleted device", slog.String("id", id))
	return nil
}

func (ts *Client) DeleteDevices(ctx context.Context, ids []string) error {
	for _, id := range ids {
		if err := ts.DeleteDevice(ctx, id); err != nil {
			return err
		}
	}
	return nil
}

func (ts *Client) request(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, ts.baseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", ts.token))
	return req, nil
}

func (df DeviceFilters) isZero() bool {
	return len(df.ids) == 0 && len(df.hostnames) == 0 && len(df.tags) == 0
}

func (d Devices) filter(filters DeviceFilters, logger *slog.Logger) Devices {
	logger.Debug(fmt.Sprintf("found %d devices before filtering", len(d)))
	if filters.isZero() {
		logger.Info("skipping device filtering", slog.Any("result", d))
		return d
	}
	logger.Debug("applying device filters", slog.Group("filters",
		slog.Any("ids", filters.ids),
		slog.Any("hostnames", filters.hostnames),
		slog.Any("tags", filters.tags),
	))
	var filtered Devices
	for _, device := range d {
		matrix := map[string]bool{
			"ids":       true,
			"hostnames": true,
			"tags":      true,
		}
		if len(filters.ids) > 0 {
			matrix["ids"] = slices.Contains(filters.ids, device.ID)
		}
		if len(filters.hostnames) > 0 {
			matrix["hostnames"] = slices.Contains(filters.hostnames, device.Hostname)
		}
		if len(filters.tags) > 0 {
			matrix["tags"] = containsAll(filters.tags, device.Tags)
		}
		if matrix["ids"] && matrix["hostnames"] && matrix["tags"] {
			filtered = append(filtered, device)
		}
	}
	message := fmt.Sprintf("found %d devices after filtering", len(filtered))
	logger.Info(message, slog.Any("result", filtered))
	return filtered
}

func (d Devices) GetIDs() []string {
	ids := make([]string, len(d))
	for i, device := range d {
		ids[i] = device.ID
	}
	return ids
}

func containsAll(a, b []string) bool {
	matrix := make(map[string]bool, len(b))
	for _, item := range b {
		matrix[item] = true
	}
	for _, item := range a {
		if !matrix[item] {
			return false
		}
	}
	return true
}
