package api

/*
 *
 */
type SecretList struct {
	Secret string `json:"secret"`
}

/*
 *
 */
type SecretRenew struct {
	Secret string `json:"secret"`
}

type AgentAdd struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UID
	Version     string `json:"version"`     // Agent version
	Secret      string `json:"secret"`      // Secret
	State       string `json:"state"`       // Operational state
}

type AgentAddRsp struct {
	Count  int    `json:"count"`  // Entries affected
	Status string `json:"status"` // Status of operation
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

type Agent struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UID
	Version     string `json:"version"`     // Agent version
	Secret      string `json:"secret"`      // Secret
	State       string `json:"state"`       // Operational state
}

type Source struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	Path        string `json:"path"`        // Source path URL
	AgentID     string `json:"agentId"`     // Agent UID
}

type SourceExt struct {
	Source
	ID string `json:"id"` // Source UID
}

type SourceID struct {
	Alias  string `json:"alias"`        // User assigned alias, human readable
	ID     string `json:"id,omitempty"` // Source UID
	New    bool   `json:"new"`          // Source is new
	Status string `json:"status"`       // Status
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

type AgentResponseList struct {
	Status string  `json:"status"`
	Data   []Agent `json:"data,omitempty"`
}

type SourceResponseList struct {
	Status string      `json:"status"`
	Data   []SourceExt `json:"data,omitempty"`
}

type SourceResponseAdd struct {
	Status string     `json:"status"`
	Data   []SourceID `json:"data,omitempty"`
}
