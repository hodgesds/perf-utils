package perf

import "strings"

// AvailableEvents returns the list of available events.
func AvailableEvents() (map[string][]string, error) {
	events := map[string][]string{}
	rawEvents, err := fileToStrings(TracingDir + "/available_events")
	// Events are colon delimited by type so parse the type and add sub
	// events appropriately.
	if err != nil {
		return events, err
	}
	for _, rawEvent := range rawEvents {
		splits := strings.Split(rawEvent, ":")
		if len(splits) <= 1 {
			continue
		}
		eventTypeEvents, found := events[splits[0]]
		if found {
			events[splits[0]] = append(eventTypeEvents, splits[1])
			continue
		}
		events[splits[0]] = []string{splits[1]}
	}
	return events, err
}

// AvailableTracers returns the list of available tracers.
func AvailableTracers() ([]string, error) {
	return fileToStrings(TracingDir + "/available_tracers")
}

// CurrentTracer returns the current tracer.
func CurrentTracer() (string, error) {
	res, err := fileToStrings(TracingDir + "/current_tracer")
	return res[0], err
}
