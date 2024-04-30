package official

type APIRequest struct {
	Messages []api_message `json:"messages"`
	Stream   bool          `json:"stream"`
	Model    string        `json:"model"`
}

type api_message struct {
	Role    string      `json:"role"`
	Content interface{} `json:"content"`
}
