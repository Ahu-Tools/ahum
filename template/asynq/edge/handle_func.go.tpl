func Handle{{.TaskName}}(ctx context.Context, t *asynq.Task) error {
	var p {{.TaskName}}Payload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}
	log.Printf("{{.TaskName}} executed!")
	// {{.TaskName}} code ...
	return nil
}