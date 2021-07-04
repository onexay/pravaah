package messaging

const (
	MSG_CONNECT_REQ = iota // Connection REQ message
	MSG_CONNECT_RSP        // Connection RSP message
	MSG_CAPABILITY         // Capabilities message
	MSG_SYNC               // Sync message
)

type ReqMsg struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type RspMsg struct {
	Type   int    `json:"type"`
	Data   string `json:"data"`
	Status string `json:"status"`
}

type ConnectReqMsg struct {
	Version string `json:"version"`
	Secret  string `json:"secret"`
}

type ConnectRspMsg struct {
	Version string `json:"version"`
	ID      string `json:"uuid"`
}

type CapabilitiesMsg struct {
	Version string `json:"version"`
	Secret  string `json:"secret"`
}

type SyncMsg struct {
	Secret string `json:"secret"`
}
