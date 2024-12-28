package tailscale

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cecobask/homelab/automation/pkg/logger"
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
	baseURL              = "https://api.tailscale.com/api/v2/"
	listDevicesEndpoint  = "tailnet/-/devices"
	deleteDeviceEndpoint = "device/%s"
)

func NewClient() *Client {
	return &Client{
		client:  &http.Client{},
		baseURL: baseURL,
		token:   os.Getenv("TAILSCALE_TOKEN"),
		logger:  logger.NewLogger(os.Stdout).With(slog.String("scope", "tailscale")),
	}
}

func NewDeviceFilters(ids, hostnames, tags []string) *DeviceFilters {
	return &DeviceFilters{
		ids:       ids,
		hostnames: hostnames,
		tags:      tags,
	}
}

func (ts *Client) GetDevices(ctx context.Context, filters DeviceFilters) (Devices, error) {
	req, err := ts.request(ctx, http.MethodGet, listDevicesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	resp, err := ts.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get devices: status code is %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
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
	req, err := ts.request(ctx, http.MethodDelete, fmt.Sprintf(deleteDeviceEndpoint, id), nil)
	if err != nil {
		return err
	}
	resp, err := ts.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete device: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete device: status code is %d", resp.StatusCode)
	}
	return nil
}

func (ts *Client) DeleteDevices(ctx context.Context, ids []string) error {
	for _, id := range ids {
		if err := ts.DeleteDevice(ctx, id); err != nil {
			return err
		}
		ts.logger.Info("deleted device", slog.String("id", id))
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
