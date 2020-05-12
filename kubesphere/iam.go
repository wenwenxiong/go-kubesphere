package kubesphere

import (
	"context"
	"fmt"
)

type IamService service

type IamRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (a *IamRequest) GetUsername() string {
	if a == nil || a.Username == nil {
		return ""
	}
	return *a.Username
}

func (a *IamRequest) GetPassword() string {
	if a == nil || a.Password == nil {
		return ""
	}
	return *a.Password
}

type AccessToken struct {
	Accesstoken *string `json:"access_token,omitempty"`
}

func (a *AccessToken) GetAccesstoken() string {
	if a == nil || a.Accesstoken == nil {
		return ""
	}
	return *a.Accesstoken
}


func (s *IamService) GetAccessToken(ctx context.Context, authReq *IamRequest) (*AccessToken, *Response, error) {
	u := fmt.Sprintf("/kapis/iam.kubesphere.io/v1alpha2/login")
	a := new(AccessToken)
	req, err := s.client.NewRequest("POST", u, authReq)



	if err != nil {
		return a, nil, err
	}


	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}
