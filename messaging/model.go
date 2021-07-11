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
	MSG_STATUS_OK      = 1
	MSG_STATUS_ERROR   = 2
	MSG_STATUS_INVALID = 3

	MSG_CONNECT_REQ = 1 // Connection REQ message
	MSG_CONNECT_RSP = 2 // Connection RSP message
	MSG_CAPABILITY  = 3 // Capabilities message
	MSG_SYNC        = 4 // Sync message
)

var MsgStatusStr []string = []string{
	"NONE",
	"OK",
	"ERROR",
	"INVALID",
}

var MsgTypeStr []string = []string{
	"NONE",
	"CONNECT_REQ",
	"CONNECT_RSP",
}

type ReqMsg struct {
	Type int    `json:"type"`
	Data string `json:"data"`
}

type RspMsg struct {
	Type   int    `json:"type"`
	Data   string `json:"data,omitempty"`
	Status int    `json:"status"`
}

type ConnectReqMsg struct {
	Alias       string `json:"alias"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Secret      string `json:"secret"`
	ID          string `json:"id"`
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
