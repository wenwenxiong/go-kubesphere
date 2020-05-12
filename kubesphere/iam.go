package kubesphere

import (
	"context"
	"fmt"
)

type IamService service

type IamRequest struct {
	Username     *string `json:"username,omitempty"`
	Password     *string `json:"password,omitempty"`
}

func (s *IamService) getAccessToken(ctx context.Context,  authReq *IamRequest) ( *Response, error) {
	u := fmt.Sprintf("/kapis/iam.kubesphere.io/v1alpha2/login")
	req, err := s.client.NewRequest("POST", u, authReq)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, u)
	if err != nil {
		return resp, err
	}
	return resp, nil
}