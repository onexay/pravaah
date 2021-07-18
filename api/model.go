package api

/*
 *
 */
type StatusElem struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
}

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

type AgentSources struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UUID assigned by server
	State       string `json:"state"`       // Operational state
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

//-----------------------------------------------------------------------------
type AgentListResultElem struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	ID          string `json:"id"`          // UID
	Version     string `json:"version"`     // Agent version
	Secret      string `json:"secret"`      // Secret
	State       string `json:"state"`       // Operational state
}
type AgentListResponse struct {
	Status string                `json:"status"`
	Data   []AgentListResultElem `json:"data,omitempty"`
}
type AgentDeleteElem struct {
	Alias string `json:"alias"` // User assigned alias, human readable
	ID    string `json:"id"`    // UID
}
type AgentDeleteResultElem struct {
	Status      string                   `json:"status"`
	Description string                   `json:"description,omitempty"`
	Alias       string                   `json:"alias"`             // User assigned alias
	ID          string                   `json:"id"`                // Source UUID
	Sources     []SourceDeleteResultElem `json:"sources,omitempty"` // Sources from this agent
}
type AgentDeleteResponse struct {
	Status string                  `json:"status"`
	Data   []AgentDeleteResultElem `json:"data,omitempty"`
}
type SourceListResultElem struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	Path        string `json:"path"`        // Source path URL
	ID          string `json:"id"`          // Source UID
	AgentID     string `json:"agentId"`     // Agent UID
}
type SourceListResponse struct {
	Status string                 `json:"status"`
	Data   []SourceListResultElem `json:"data,omitempty"`
}
type SourceAddElem struct {
	Alias       string `json:"alias"`       // User assigned alias, human readable
	Description string `json:"description"` // Description
	Path        string `json:"path"`        // Source path URL
	AgentID     string `json:"agentId"`     // Agent UID
}
type SourceAddResultElem struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
	Alias       string `json:"alias"`        // User assigned alias, human readable
	ID          string `json:"id,omitempty"` // Source UID
}
type SourceAddResponse struct {
	Status string                `json:"status"`
	Data   []SourceAddResultElem `json:"data,omitempty"`
}
type SourceDeleteElem struct {
	Alias   string `json:"alias"`
	ID      string `json:"id"`
	AgentID string `json:"agentId"`
}
type SourceDeleteResultElem struct {
	Status      string `json:"status"`
	Description string `json:"description,omitempty"`
	Alias       string `json:"alias"` // User assigned alias
	ID          string `json:"id"`    // Source UUID
}
type SourceDeleteResponse struct {
	Status string                   `json:"status"`
	Data   []SourceDeleteResultElem `json:"data,omitempty"`
}
