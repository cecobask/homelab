package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) CreateVolume(ctx context.Context, params CreateVolumeParams) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, params.url(c.baseURL), params.body())
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var r response
		json.NewDecoder(res.Body).Decode(&r)
		return fmt.Errorf("expected response status code %d, got %d; message: %s", http.StatusOK, res.StatusCode, r.Message)
	}
	slog.Info("volume created", "volume", params.Filename)
	return nil
}

func (c *Client) EnsureVolume(ctx context.Context, params CreateVolumeParams) error {
	_, err := c.GetVolume(ctx, GetVolumeParams{
		Node:    params.Node,
		Storage: params.Storage,
		Volume:  params.Filename,
	})
	if err == nil {
		slog.Warn("volume already exists, skipping creation", "volume", params.Filename)
		return nil
	}
	var nfe *NotFoundError
	if errors.As(err, &nfe) {
		return c.CreateVolume(ctx, params)
	}
	return err
}

func (c *Client) GetVolume(ctx context.Context, params GetVolumeParams) (*Volume, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, params.url(c.baseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()
	var r response
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		if strings.Contains(r.Message, "Failed to find logical volume") {
			return nil, &NotFoundError{Message: r.Message}
		}
		return nil, fmt.Errorf("expected response status code %d, got %d; message: %s", http.StatusOK, res.StatusCode, r.Message)
	}
	return &r.Data, nil
}

func (c *Client) DeleteVolume(ctx context.Context, params DeleteVolumeParams) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, params.url(c.baseURL), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	res, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		var r response
		json.NewDecoder(res.Body).Decode(&r)
		return fmt.Errorf("expected response status code %d, got %d; message: %s", http.StatusOK, res.StatusCode, r.Message)
	}
	slog.Info("volume deleted", "volume", params.Volume)
	return nil
}

type Volume struct {
	Format string `json:"format"`
	Path   string `json:"path"`
	Size   int    `json:"size"`
	Used   int    `json:"used"`
}

type CreateVolumeParams struct {
	Filename string
	Format   string
	Node     string
	Size     string
	Storage  string
	VMID     int
}

func (r CreateVolumeParams) url(base string) string {
	return fmt.Sprintf("%s/nodes/%s/storage/%s/content", base, r.Node, r.Storage)
}

func (r CreateVolumeParams) body() io.Reader {
	v := url.Values{}
	v.Set("filename", r.Filename)
	v.Set("format", r.Format)
	v.Set("node", r.Node)
	v.Set("size", r.Size)
	v.Set("storage", r.Storage)
	v.Set("vmid", strconv.Itoa(r.VMID))
	return strings.NewReader(v.Encode())
}

type GetVolumeParams struct {
	Node    string
	Storage string
	Volume  string
}

func (r GetVolumeParams) url(base string) string {
	return fmt.Sprintf("%s/nodes/%s/storage/%s/content/%s", base, r.Node, r.Storage, r.Volume)
}

type DeleteVolumeParams struct {
	Node    string
	Storage string
	Volume  string
}

func (r DeleteVolumeParams) url(base string) string {
	return fmt.Sprintf("%s/nodes/%s/storage/%s/content/%s", base, r.Node, r.Storage, r.Volume)
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}
