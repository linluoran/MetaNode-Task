// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.3

package types

type UserCreateReq struct {
	Mobile   string `form:"mobile"`
	Nickname string `form:"nickname"`
}

type UserCreateResp struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
	Flag    bool   `json:"flag"`
}

type UserInfoReq struct {
	ID int64 `form:"userId"`
}

type UserInfoResp struct {
	Code    int64  `json:"code"`
	Data    string `json:'data'`
	Message string `json:"msg"`
}

type UserUpdateReq struct {
	ID       int64  `form:"userId"`
	Nickname string `form:"nickname"`
	Mobile   string `form:"mobile"`
}

type UserUpdateResp struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
	Flag    bool   `json:"flag"`
}
