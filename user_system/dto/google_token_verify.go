package dto

type GoogleTokenVerifyRes struct {
	Azp            string `json:"azp"`
	Aud            string `json:"aud"`
	Sub            string `json:"sub"`
	Scope          string `json:"scope"`
	Exp            string `json:"exp"`
	Expires_in     string `json:"expires_in"`
	Email          string `json:"email"`
	Email_verified string `json:"email_verified"`
	Access_type    string `json:"access_type"`
}
