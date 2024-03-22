package lsp

type DidChangeNotification struct {
	Notification
	Params DidChangeParams `json:"params"`
}

type DidChangeParams struct {
	TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []ContentChangeEvent            `json:"contentChanges"`
}

type ContentChangeEvent struct {
	Text string `json:"text"`
}
