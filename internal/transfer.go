package internal

type Transfer struct {
	ID       uint64 `json:"id"`
	Time     string `json:"time"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Value    string `json:"value"`
	Sign
}

type Sign struct {
	Pk, Sign []byte
}
