package api

type ListReq struct {
}

type AddReq struct {
	UUID    string `json:"uuid"`
	Alias   string `json:"alias"`
	RTPPath string `json:"rtpPath"`
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
