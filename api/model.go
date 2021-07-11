package api

type AgentAdd struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UID
	Version     string `json:"version"`     // Agent version
	Secret      string `json:"secret"`      // Secret
	State       string `json:"state"`       // Operational state
}

type AgentDel struct {
	Alias string `json:"alias"` // User assigned alias, human readable
	ID    string `json:"id"`    // UID
}

type AgentDelRsp struct {
	Count  int    `json:"count"`  // Entries affected
	Status string `json:"status"` // Status of operation
}

type AgentSources struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UUID assigned by server
	State       string `json:"state"`       // Operational state
}

type AddReq struct {
	Alias   string `json:"alias"`
	RTPPath string `json:"rtpPath"`
}

type AddRsp struct {
	UUID string `json:"uuid"`
}

type RemoveReq struct {
	UUID  string `json:"uuid"`
	Alias string `json:"alias"`
}

type StartReq struct {
	UUID  string `json:"uuid"`
	Alias string `json:"alias"`
}

type StopReq struct {
	UUID  string `json:"uuid"`
	Alias string `json:"alias"`
}
