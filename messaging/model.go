package messaging

/*
 * Interface for messaging between agents and server
 */
type Messaging_T interface {
	// Connection establishment
	ConnectReq()
	ConnectRsp(rsp *RspMsg)

	// Dynamic feature negotiation
	FeatureSupportReq()
	FeatureSupportRsp()

	// Periodic sync
	SyncReq()
	SyncRsp()

	// Add source
	AddSourceReq(req *ReqMsg)
	AddSourceRsp()
}

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
	ID      string `json:"id"`
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

type AddSourceReqMsg struct {
	Path  string `json:"path"`
	Alias string `json:"alias"`
}

type AddSourceRspMsg struct {
}
