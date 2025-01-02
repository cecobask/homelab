package proxmox

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"
)

type Client struct {
	client  *http.Client
	baseURL string
	token   string
	logger  *slog.Logger
}

type Storage struct {
	ID      string
	Node    string
	Status  string
	Storage string
}

type Storages []Storage

type StorageFilters struct {
	nodes    []string
	storages []string
}

type StorageContent struct {
	ContentType string
	Volume      string
	Node        string
	Storage     string
}

type StorageContents []StorageContent

type StorageContentFilters struct {
	contentTypes []string
	volumes      []string
}
type VirtualMachine struct {
	ID     int
	Name   string
	Node   string
	Status string
	Tags   []string
}

type VirtualMachines []VirtualMachine

type VirtualMachineFilters struct {
	ids   []int
	nodes []string
	tags  []string
}

type clusterResourcesResponse struct {
	Data []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		VMID    int    `json:"vmid,omitempty"`
		Name    string `json:"name,omitempty"`
		Node    string `json:"node,omitempty"`
		Status  string `json:"status,omitempty"`
		Storage string `json:"storage,omitempty"`
		Tags    string `json:"tags,omitempty"`
	} `json:"data"`
}

type storageContentResponse struct {
	Data []struct {
		Content string `json:"content"`
		Volume  string `json:"volid"`
	} `json:"data"`
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

func NewStorageFilters(nodes, storages []string) StorageFilters {
	return StorageFilters{
		nodes:    nodes,
		storages: storages,
	}
}

func NewStorageContentFilters(contentTypes, volumes []string) StorageContentFilters {
	return StorageContentFilters{
		contentTypes: contentTypes,
		volumes:      volumes,
	}
}

func NewVirtualMachineFilters(ids []int, nodes, tags []string) VirtualMachineFilters {
	return VirtualMachineFilters{
		ids:   ids,
		nodes: nodes,
		tags:  tags,
	}
}

const (
	endpointDeleteVolume         = "/nodes/%s/storage/%s/content/%s"
	endpointDestroyVM            = "/nodes/%s/qemu/%d"
	endpointListClusterResources = "/cluster/resources"
	endpointListStorageContent   = "/nodes/%s/storage/%s/content"
	endpointStopVM               = "/nodes/%s/qemu/%d/status/stop"
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
		pm.logger.Error("failed to delete volume", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("deleted volume", slog.String("node", node), slog.String("storage", storage), slog.String("volume", volume))
	return nil
}

func (pm *Client) DeleteVolumes(ctx context.Context, storageContents StorageContents) error {
	for _, content := range storageContents {
		if err := pm.DeleteVolume(ctx, content.Node, content.Storage, content.Volume); err != nil {
			return err
		}
	}
	return nil
}

func (pm *Client) ListVMs(ctx context.Context, filters VirtualMachineFilters) (VirtualMachines, error) {
	pm.logger.Debug("listing virtual machines", slog.Group("filters", slog.Any("nodes", filters.nodes), slog.Any("tags", filters.tags)))
	req, err := pm.request(ctx, http.MethodGet, endpointListClusterResources, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("type", "vm")
	req.URL.RawQuery = query.Encode()
	resp, err := pm.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list virtual machines: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to list virtual machines", slog.String("status", resp.Status))
		return nil, fmt.Errorf("http response was not healthy")
	}
	var crr clusterResourcesResponse
	if err = json.Unmarshal(body, &crr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster resources: %v", err)
	}
	return crr.toVirtualMachines().filter(filters, pm.logger), nil
}

func (pm *Client) ListStorages(ctx context.Context, filters StorageFilters) (Storages, error) {
	pm.logger.Debug("listing cluster storages", slog.Group("filters", slog.Any("nodes", filters.nodes), slog.Any("storages", filters.storages)))
	req, err := pm.request(ctx, http.MethodGet, endpointListClusterResources, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("type", "storage")
	req.URL.RawQuery = query.Encode()
	resp, err := pm.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list cluster storages: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to list cluster storages", slog.String("status", resp.Status))
		return nil, fmt.Errorf("http response was not healthy")
	}
	var crr clusterResourcesResponse
	if err = json.Unmarshal(body, &crr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cluster resources: %v", err)
	}
	return crr.toStorages().filter(filters, pm.logger), nil
}

func (pm *Client) ListStorageContent(ctx context.Context, node, storage string, filters StorageContentFilters) (StorageContents, error) {
	pm.logger.Debug("listing storage content", slog.String("node", node), slog.String("storage", storage))
	req, err := pm.request(ctx, http.MethodGet, fmt.Sprintf(endpointListStorageContent, node, storage), nil)
	if err != nil {
		return nil, err
	}
	resp, err := pm.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to list storage content: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to list storage content", slog.String("status", resp.Status))
		return nil, fmt.Errorf("http response was not healthy")
	}
	var scr storageContentResponse
	if err = json.Unmarshal(body, &scr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal storage content: %v", err)
	}
	return scr.toStorageContents(node, storage).filter(filters, pm.logger), nil
}

func (pm *Client) ListStorageContents(ctx context.Context, storages Storages, filters StorageContentFilters) (StorageContents, error) {
	var storageContents StorageContents
	for _, storage := range storages {
		sc, err := pm.ListStorageContent(ctx, storage.Node, storage.Storage, filters)
		if err != nil {
			return nil, err
		}
		storageContents = append(storageContents, sc...)
	}
	return storageContents, nil
}

func (pm *Client) DestroyVM(ctx context.Context, node string, id int, destroyUnreferencedDisks, purge bool) error {
	pm.logger.Debug("destroying virtual machine", slog.String("node", node), slog.Int("id", id),
		slog.Bool("destroyUnreferencedDisks", destroyUnreferencedDisks),
		slog.Bool("purge", purge),
	)
	req, err := pm.request(ctx, http.MethodDelete, fmt.Sprintf(endpointDestroyVM, node, id), nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	query.Set("destroy-unreferenced-disks", boolToNumericString(destroyUnreferencedDisks))
	query.Set("purge", boolToNumericString(purge))
	req.URL.RawQuery = query.Encode()
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to destroy virtual machine: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to destroy virtual machine", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("destroyed virtual machine", slog.String("node", node), slog.Int("id", id))
	return nil
}

func (pm *Client) DestroyVMs(ctx context.Context, vms VirtualMachines, destroyUnreferencedDisks, purge bool) error {
	for _, vm := range vms {
		if err := pm.DestroyVM(ctx, vm.Node, vm.ID, destroyUnreferencedDisks, purge); err != nil {
			return err
		}
	}
	return nil
}

func (pm *Client) StopVM(ctx context.Context, node string, id int) error {
	pm.logger.Debug("stopping virtual machine", slog.String("node", node), slog.Int("id", id))
	req, err := pm.request(ctx, http.MethodPost, fmt.Sprintf(endpointStopVM, node, id), nil)
	if err != nil {
		return err
	}
	resp, err := pm.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to stop virtual machine: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		pm.logger.Error("failed to stop virtual machine", slog.String("status", resp.Status))
		return fmt.Errorf("received unhealthy http response status code")
	}
	pm.logger.Info("stopped virtual machine", slog.String("node", node), slog.Int("id", id))
	return nil
}

func (pm *Client) StopVMs(ctx context.Context, vms VirtualMachines) error {
	for _, vm := range vms {
		if err := pm.StopVM(ctx, vm.Node, vm.ID); err != nil {
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

func (crr *clusterResourcesResponse) toStorages() Storages {
	storages := make(Storages, len(crr.Data))
	for i, resource := range crr.Data {
		storages[i] = Storage{
			ID:      resource.ID,
			Node:    resource.Node,
			Status:  resource.Status,
			Storage: resource.Storage,
		}
	}
	return storages
}

func (scr *storageContentResponse) toStorageContents(node, storage string) StorageContents {
	var storageContents StorageContents
	for _, sc := range scr.Data {
		storageContents = append(storageContents, StorageContent{
			ContentType: sc.Content,
			Volume:      sc.Volume,
			Node:        node,
			Storage:     storage,
		})
	}
	return storageContents
}

func (crr *clusterResourcesResponse) toVirtualMachines() VirtualMachines {
	vms := make(VirtualMachines, len(crr.Data))
	for i, resource := range crr.Data {
		vms[i] = VirtualMachine{
			ID:     resource.VMID,
			Name:   resource.Name,
			Node:   resource.Node,
			Status: resource.Status,
			Tags:   strings.Split(resource.Tags, ";"),
		}
	}
	return vms
}

func (sf StorageFilters) isZero() bool {
	return len(sf.nodes) == 0 && len(sf.storages) == 0
}

func (s Storages) filter(filters StorageFilters, logger *slog.Logger) Storages {
	logger.Debug(fmt.Sprintf("found %d cluster storages before filtering", len(s)))
	if filters.isZero() {
		logger.Info("skipping cluster storage filtering")
		return s
	}
	logger.Debug("applying cluster storage filters", slog.Group("filters",
		slog.Any("nodes", filters.nodes),
		slog.Any("storages", filters.storages),
	))
	var filtered Storages
	for _, storage := range s {
		matrix := map[string]bool{
			"nodes":    true,
			"storages": true,
		}
		if len(filters.nodes) > 0 {
			matrix["nodes"] = slices.Contains(filters.nodes, storage.Node)
		}
		if len(filters.storages) > 0 {
			matrix["storages"] = slices.Contains(filters.storages, storage.Storage)
		}
		if matrix["nodes"] && matrix["storages"] {
			filtered = append(filtered, storage)
		}
	}
	logger.Info(fmt.Sprintf("found %d cluster storages after filtering", len(filtered)))
	return filtered
}

func (scf StorageContentFilters) isZero() bool {
	return len(scf.contentTypes) == 0 && len(scf.volumes) == 0
}

func (sc StorageContents) filter(filters StorageContentFilters, logger *slog.Logger) StorageContents {
	logger.Debug(fmt.Sprintf("found %d storage contents before filtering", len(sc)))
	if filters.isZero() {
		logger.Info("skipping storage content filtering")
		return sc
	}
	logger.Debug("applying storage content filters", slog.Group("filters",
		slog.Any("contentTypes", filters.contentTypes),
		slog.Any("volumes", filters.volumes),
	))
	var filtered StorageContents
	for _, storageContent := range sc {
		matrix := map[string]bool{
			"contentTypes": true,
			"volumes":      true,
		}
		if len(filters.contentTypes) > 0 {
			matrix["contentTypes"] = slices.Contains(filters.contentTypes, storageContent.ContentType)
		}
		if len(filters.volumes) > 0 {
			matrix["volumes"] = slices.Contains(filters.volumes, storageContent.Volume)
		}
		if matrix["contentTypes"] && matrix["volumes"] {
			filtered = append(filtered, storageContent)
		}
	}
	logger.Info(fmt.Sprintf("found %d storage contents after filtering", len(filtered)))
	return filtered
}

func (vmf VirtualMachineFilters) isZero() bool {
	return len(vmf.ids) == 0 && len(vmf.nodes) == 0 && len(vmf.tags) == 0
}

func (vms VirtualMachines) filter(filters VirtualMachineFilters, logger *slog.Logger) VirtualMachines {
	logger.Debug(fmt.Sprintf("found %d virtual machines before filtering", len(vms)))
	if filters.isZero() {
		logger.Info("skipping virtual machine filtering")
		return vms
	}
	logger.Debug("applying virtual machine filters", slog.Group("filters",
		slog.Any("storages", filters.ids),
		slog.Any("nodes", filters.nodes),
		slog.Any("tags", filters.tags),
	))
	var filtered VirtualMachines
	for _, vm := range vms {
		matrix := map[string]bool{
			"storages": true,
			"nodes":    true,
			"tags":     true,
		}
		if len(filters.ids) > 0 {
			matrix["storages"] = slices.Contains(filters.ids, vm.ID)
		}
		if len(filters.nodes) > 0 {
			matrix["nodes"] = slices.Contains(filters.nodes, vm.Node)
		}
		if len(filters.tags) > 0 {
			matrix["tags"] = containsAll(filters.tags, vm.Tags)
		}
		if matrix["storages"] && matrix["nodes"] && matrix["tags"] {
			filtered = append(filtered, vm)
		}
	}
	logger.Info(fmt.Sprintf("found %d virtual machines after filtering", len(filtered)))
	return filtered
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

func boolToNumericString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
