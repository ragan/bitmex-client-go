package bitmex

type Error struct {
	Error Err `json:"error"`
}

type Err struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}
