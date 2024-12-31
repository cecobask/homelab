package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
)

type Client struct {
	client  *http.Client
	baseURL string
	token   string
	logger  *slog.Logger
}

func NewClient(baseURL string, logger *slog.Logger) *Client {
	return &Client{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		baseURL: baseURL,
		token:   fmt.Sprintf("PVEAPIToken=%s=%s", os.Getenv("PROXMOX_TOKEN_ID"), os.Getenv("PROXMOX_SECRET")),
		logger:  logger.With(slog.String("scope", "proxmox")),
	}
}

const (
	endpointDeleteVolume = "/nodes/%s/storage/%s/content/%s"
	endpointDestroyVM    = "/nodes/%s/qemu/%s"
	endpointShutdownVM   = "/nodes/%s/qemu/%s/status/shutdown"
	endpointStopVM       = "/nodes/%s/qemu/%s/status/stop"
)

func (pm *Client) DeleteVolume(ctx context.Context, node, storage, volume string) error {
	pm.logger.Debug("deleting volume", slog.String("node", node), slog.String("storage", storage), slog.String("volume", volume))
	req, err := pm.request(ctx, http.MethodDelete, fmt.Sprintf(endpointDeleteVolume, node, storage, volume), nil)
	if err != nil {
		return err
	}
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to delete volume: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to delete vm", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("deleted volume", slog.String("node", node), slog.String("storage", storage), slog.String("volume", volume))
	return nil
}

func (pm *Client) DestroyVM(ctx context.Context, node, vmid string, destroyUnreferencedDisks, purge bool) error {
	pm.logger.Debug("destroying vm", slog.String("node", node), slog.String("vmid", vmid), slog.Bool("destroyUnreferencedDisks", destroyUnreferencedDisks), slog.Bool("purge", purge))
	req, err := pm.request(ctx, http.MethodDelete, fmt.Sprintf(endpointDestroyVM, node, vmid), nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	query.Set("destroy-unreferenced-disks", boolToNumericString(destroyUnreferencedDisks))
	query.Set("purge", boolToNumericString(purge))
	req.URL.RawQuery = query.Encode()
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to destroy vm: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to destroy vm", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("destroyed vm", slog.String("node", node), slog.String("vmid", vmid), slog.Bool("destroyUnreferencedDisks", destroyUnreferencedDisks), slog.Bool("purge", purge))
	return nil
}

func (pm *Client) DestroyVMs(ctx context.Context, node string, vmids []string, destroyUnreferencedDisks, purge bool) error {
	for _, vmid := range vmids {
		if err := pm.DestroyVM(ctx, node, vmid, destroyUnreferencedDisks, purge); err != nil {
			return err
		}
	}
	return nil
}

func (pm *Client) ShutdownVM(ctx context.Context, node, vmid string) error {
	pm.logger.Debug("shutting down vm", slog.String("node", node), slog.String("vmid", vmid))
	req, err := pm.request(ctx, http.MethodPost, fmt.Sprintf(endpointShutdownVM, node, vmid), nil)
	if err != nil {
		return err
	}
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to shut down vm: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to shut down vm", slog.String("status", resp.Status))
		return fmt.Errorf("http response was not healthy")
	}
	pm.logger.Info("shut down vm", slog.String("node", node), slog.String("vmid", vmid))
	return nil
}

func (pm *Client) ShutdownVMs(ctx context.Context, node string, vmids []string) error {
	for _, vmid := range vmids {
		if err := pm.ShutdownVM(ctx, node, vmid); err != nil {
			return err
		}
	}
	return nil
}

func (pm *Client) StopVM(ctx context.Context, node, vmid string) error {
	pm.logger.Debug("stopping vm", slog.String("node", node), slog.String("vmid", vmid))
	req, err := pm.request(ctx, http.MethodPost, fmt.Sprintf(endpointStopVM, node, vmid), nil)
	if err != nil {
		return err
	}
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to stop vm: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to stop vm", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("stopped vm", slog.String("node", node), slog.String("vmid", vmid))
	return nil
}

func (pm *Client) StopVMs(ctx context.Context, node string, vmids []string) error {
	for _, vmid := range vmids {
		if err := pm.StopVM(ctx, node, vmid); err != nil {
			return err
		}
	}
	return nil
}

func (pm *Client) request(ctx context.Context, method, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, pm.baseURL+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", pm.token)
	return req, nil
}

func boolToNumericString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
