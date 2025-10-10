package hooks

import "os"

func initHooks(h *Hooks) {
	if os.Getenv("POLAR_ENABLE_EVENT_BATCHING") == "true" {
		eventBatcher := NewEventBatcher()
		h.registerBeforeRequestHook(eventBatcher)
		h.registerAfterSuccessHook(eventBatcher)
		globalEventBatcher = eventBatcher
	}
}

var globalEventBatcher *EventBatcher

func FlushEvents() error {
	if globalEventBatcher == nil {
		return nil
	}
	return globalEventBatcher.Flush()
}
