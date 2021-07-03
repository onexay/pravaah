package messaging

type CapabilitiesMsg struct {
	Version string `json:"version"`
	Secret  string `json:"secret"`
}

type SyncMsg struct {
	Secret string `json:"secret"`
}
