package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// We will just specify the type of the params in all of the request types
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// We will just specify the result and errors in all the response types
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
