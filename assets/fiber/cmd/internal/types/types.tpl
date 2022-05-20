package types

type (
	CaptchaRsp struct {
		CaptchaId     string `json:"captchaId"`
		PicPath       string `json:"picPath"`
		CaptchaLength int    `json:"captchaLength""`
	}
)

type (
	AddUserReq struct {
		PhoneNum  string `json:"phoneNum" validate:"len=11"`
		Username  string `json:"userName" validate:"required"`
		NickName  string `json:"nickName" validate:"required"`
		HeaderImg string `json:"headerImg" validate:"required"`
	}
)

type (
	User struct {
		Id        string `json:"id"`
		PhoneNum  string `json:"phoneNum"`
		Username  string `json:"userName"`
		NickName  string `json:"nickName"`
		HeaderImg string `json:"headerImg"`
	}
	DetailReq struct {
		Id string `json:"id" validate:"required"`
	}
	DetailRsp struct {
		UserInfo *User `json:"userInfo"`
	}
)

type (
	LoginReq struct {
		PhoneNum  string `json:"phoneNum" validate:"len=11"`
		Password  string `json:"password" validate:"required"`
		Captcha   string `json:"captcha" validate:"len=6"`
		CaptchaId string `json:"captchaId" validate:"required"`
	}
	LoginRsp struct {
		JWTToken string `json:"jwtToken"`
		UserInfo *User  `json:"userInfo"`
	}
)

type (
	ListOrderReq struct {
		Id        string `json:"id"`
		StartTime string `json:"startTime" validate:"datetime"`
		EndTime   string `json:"endTime" validate:"datetime"`
		Page      uint32 `json:"page"`
		Size      uint32 `json:"size"`
	}
	ListOrderRsp struct {
		Id         string `json:"id"`
		CreateTime string `json:"createTime"`
		TotalCount string `json:"totalCount"`
		TotalSum   string `json:"totalSum"`
		Items      string `json:"items"`
	}
)
