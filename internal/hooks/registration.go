package hooks

func initHooks(h *Hooks) {
	eventBatcher := NewEventBatcher()
	h.registerBeforeRequestHook(eventBatcher)
	h.registerAfterSuccessHook(eventBatcher)

	globalEventBatcher = eventBatcher
}

var globalEventBatcher *EventBatcher

func FlushEvents() error {
	if globalEventBatcher == nil {
		return nil
	}
	return globalEventBatcher.Flush()
}
