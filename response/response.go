package response

type RegisterReturn struct {
	ID        uint   `json:"id"`
	Username  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
	Token     string `json:"token"`
}
