package entities

import (
	"time"

	"github.com/langgenius/dify-plugin-daemon/internal/types/entities/plugin_entities"
)

type (
	PluginRuntime struct {
		State     PluginRuntimeState                `json:"state"`
		Config    plugin_entities.PluginDeclaration `json:"config"`
		Connector PluginConnector                   `json:"-"`
	}

	PluginRuntimeInterface interface {
		PluginRuntimeTimeLifeInterface
		PluginRuntimeSessionIOInterface
	}

	PluginRuntimeTimeLifeInterface interface {
		InitEnvironment() error
		StartPlugin() error
		Stopped() bool
		Stop()
		Configuration() *plugin_entities.PluginDeclaration
		RuntimeState() *PluginRuntimeState
		Wait() (<-chan bool, error)
	}

	PluginRuntimeSessionIOInterface interface {
		Listen(session_id string) *BytesIOListener
		Write(session_id string, data []byte)
	}
)

func (r *PluginRuntime) Stopped() bool {
	return r.State.Status == PLUGIN_RUNTIME_STATUS_STOPPED
}

func (r *PluginRuntime) Stop() {
	r.State.Status = PLUGIN_RUNTIME_STATUS_STOPPED
}

func (r *PluginRuntime) Configuration() *plugin_entities.PluginDeclaration {
	return &r.Config
}

func (r *PluginRuntime) RuntimeState() *PluginRuntimeState {
	return &r.State
}

type PluginRuntimeState struct {
	Restarts     int        `json:"restarts"`
	Status       string     `json:"status"`
	RelativePath string     `json:"relative_path"`
	ActiveAt     *time.Time `json:"active_at"`
	StoppedAt    *time.Time `json:"stopped_at"`
	Verified     bool       `json:"verified"`
}

const (
	PLUGIN_RUNTIME_STATUS_ACTIVE     = "active"
	PLUGIN_RUNTIME_STATUS_LAUNCHING  = "launching"
	PLUGIN_RUNTIME_STATUS_STOPPED    = "stopped"
	PLUGIN_RUNTIME_STATUS_RESTARTING = "restarting"
	PLUGIN_RUNTIME_STATUS_PENDING    = "pending"
)

type PluginConnector interface {
	OnMessage(func([]byte))
	Read([]byte) int
	Write([]byte) int
}
