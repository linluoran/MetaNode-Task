package types

type (
	CreateUserReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
