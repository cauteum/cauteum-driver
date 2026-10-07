package driver

import (
	"fmt"
	"strings"
	"sync"
)

var (
	mu                  sync.RWMutex
	openers             = map[string]func() (ComputeDriver, error){}
	engineOpeners       = map[string]func() (Engine, error){}
	engineConfigOpeners = map[string]func(map[string]any) (Engine, error){}
)

// Register binds a ComputeDriver factory (called from backend init).
func Register(kind string, open func() (ComputeDriver, error)) {
	mu.Lock()
	defer mu.Unlock()
	openers[normalizeKind(kind)] = open
}

// RegisterEngine binds a Docker/Podman Engine factory.
func RegisterEngine(kind string, open func() (Engine, error)) {
	mu.Lock()
	defer mu.Unlock()
	engineOpeners[normalizeKind(kind)] = open
}

// RegisterEngineWithConfig binds a factory that consumes driver-specific
// configuration. Config-aware factories are preferred by OpenEngineWithConfig.
func RegisterEngineWithConfig(kind string, open func(map[string]any) (Engine, error)) {
	mu.Lock()
	defer mu.Unlock()
	engineConfigOpeners[normalizeKind(kind)] = open
}

// Open returns a ComputeDriver for kind: docker | podman | vm | kubernetes.
// Import driver/all (blank) once from main so backends register.
func Open(kind string) (ComputeDriver, error) {
	mu.RLock()
	f := openers[normalizeKind(kind)]
	mu.RUnlock()
	if f == nil {
		return nil, fmt.Errorf("driver: unknown kind %q (import _ \"…/driver/all\"?)", kind)
	}
	return f()
}

// OpenEngine returns a Docker/Podman Engine (ComputeDriver + host API helpers).
func OpenEngine(kind string) (Engine, error) {
	mu.RLock()
	f := engineOpeners[normalizeKind(kind)]
	mu.RUnlock()
	if f == nil {
		return nil, fmt.Errorf("driver: kind %q has no Engine API (want docker|podman; import _ \"…/driver/all\"?)", kind)
	}
	return f()
}

// OpenEngineWithConfig opens an Engine using its backend-specific options.
// Non-empty config is rejected when the backend has no config-aware factory.
func OpenEngineWithConfig(kind string, config map[string]any) (Engine, error) {
	mu.RLock()
	f := engineConfigOpeners[normalizeKind(kind)]
	mu.RUnlock()
	if f != nil {
		return f(config)
	}
	if len(config) != 0 {
		return nil, fmt.Errorf("driver: kind %q does not accept driver configuration", kind)
	}
	return OpenEngine(kind)
}

func normalizeKind(kind string) string {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "podman":
		return "podman"
	case "vm", "microvm":
		return "vm"
	case "kubernetes", "k8s":
		return "kubernetes"
	case "", "docker":
		return "docker"
	default:
		return strings.ToLower(strings.TrimSpace(kind))
	}
}
