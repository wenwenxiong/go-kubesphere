package kubesphere

import (
	"fmt"
	"context"
	"github.com/go-openapi/strfmt"
)
type OpenpitrixService service

type AppCategory struct {
	CategoryId *string `json:"category_id,omitempty"`
	AppTotal *int `json:"app_total,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Icon *string `json:"icon,omitempty"`
	CreateTime *string `json:"create_time,omitempty"`
	Owner *string `json:"owner,omitempty"`
	UpdateTime *string `json:"update_time,omitempty"`

}

type AppCategoryList struct {
	TotalCount *int `json:"total_count,omitempty"`
	Items []*AppCategory
}

type Message struct {
	Message *string `json:"message,omitempty"`
}

func (a *AppCategory) GetCategoryId() string {
	if a == nil || a.CategoryId == nil {
		return ""
	}
	return *a.CategoryId
}

func (a *AppCategory) GetAppTotal() int {
	if a == nil || a.AppTotal == nil {
		return 0
	}
	return *a.AppTotal
}
func (a *AppCategory) GetName() string {
	if a == nil || a.Name == nil {
		return ""
	}
	return *a.Name
}
func (a *AppCategory) GetDescription() string {
	if a == nil || a.Description == nil {
		return ""
	}
	return *a.Description
}
func (a *AppCategory) GetLocale() string {
	if a == nil || a.Locale == nil {
		return ""
	}
	return *a.Locale
}
func (a *AppCategory) GetIcon() string {
	if a == nil || a.Icon == nil {
		return ""
	}
	return *a.Icon
}
func (a *AppCategory) GetCreateTime() string {
	if a == nil || a.CreateTime == nil {
		return ""
	}
	return *a.CreateTime
}

func (a *AppCategory) GetOwner() string {
	if a == nil || a.Owner == nil {
		return ""
	}
	return *a.Owner
}

func (a *AppCategory) GetUpdateTime() string {
	if a == nil || a.UpdateTime == nil {
		return ""
	}
	return *a.UpdateTime
}

func (a *AppCategory) SetName(name *string )  {
	a.Name = name
}

func (a *AppCategory) SetLocale(locale *string )  {
	a.Locale = locale
}

func (a *AppCategory) SetDescription(description *string )  {
	a.Description = description
}

func (s *OpenpitrixService) CreateAppCategory(ctx context.Context, categoryReq *AppCategory, token *AccessToken) (*AppCategory, *Response, error) {
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/categories")
	a := new(AppCategory)

	req, err := s.client.NewRequest("POST", u, categoryReq)

	if err != nil {
		return a, nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) GetAppCategory(ctx context.Context, categoryReq *AppCategory, token *AccessToken) ( *AppCategoryList, *Response, error) {

	u := fmt.Sprintf("/kapis/openpitrix.io/v1/categories")
	if categoryReq != nil && categoryReq.Name != nil {
		categoryName := categoryReq.Name
		u = u + fmt.Sprintf("?conditions=keyword=%s", *categoryName)
	}

	a := new(AppCategoryList)

	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return a, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) UpdateAppCategory(ctx context.Context, categoryReq *AppCategory, token *AccessToken) ( *Message, *Response, error) {
	categoryId := categoryReq.CategoryId
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/categories/%s", *categoryId)
	a := new (Message)
	req, err := s.client.NewRequest("PATCH", u, categoryReq)

	if err != nil {
		return a, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) DeleteAppCategory(ctx context.Context, categoryReq *AppCategory, token *AccessToken) ( *Message, *Response, error) {
	categoryId := categoryReq.CategoryId
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/categories/%s", *categoryId)
	a := new (Message)
	req, err := s.client.NewRequest("DELETE", u, nil)

	if err != nil {
		return a, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

type App struct {
	AppId *string `json:"app_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Isv *string `json:"isv,omitempty"`
	VersionId *string `json:"version_id,omitempty"`
	VersionType *string `json:"version_type,omitempty"`
	VersionName *string `json:"version_name,omitempty"`
	VersionPackage strfmt.Base64  `json:"version_package,omitempty"`
	UpdateTime *string `json:"update_time,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Description *string `json:"description,omitempty"`
	ClusterTotal *int `json:"cluster_total,omitempty"`
	Active *bool `json:"active,omitempty"`
	PackageName *string `json:"package_name,omitempty"`
	CategoryId *string `json:"category_id,omitempty"`
}

func (a *App) GetCategoryId() string {
	if a == nil || a.CategoryId == nil {
		return ""
	}
	return *a.CategoryId
}

func (a *App) SetName(name *string )  {
	a.Name = name
}

func (a *App) SetDescription(description *string )  {
	a.Description = description
}

func (a *App) SetCategoryId(categoryId *string )  {
	a.CategoryId = categoryId
}

type AppList struct {
	TotalCount *int `json:"total_count,omitempty"`
	Items []*App
}

func (s *OpenpitrixService) GetApp(ctx context.Context, appReq *App, token *AccessToken) ( *AppList, *Response, error) {
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/apps")
	if appReq != nil && appReq.Name != nil {
		appName := appReq.Name
		u = u + fmt.Sprintf("?conditions=status=active,keyword=%s", *appName)
	}

	a := new (AppList)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return a, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) CreateApp(ctx context.Context, appReq *App, token *AccessToken) (  *App, *Response, error) {

	u := fmt.Sprintf("/kapis/openpitrix.io/v1/apps")

	a := new (App)
	req, err := s.client.NewRequest("POST", u, appReq)

	if err != nil {
		return a, nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) UpdateApp(ctx context.Context, appReq *App, token *AccessToken) (  *Message, *Response, error) {
	appId := appReq.AppId
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/apps/%s", *appId)

	a := new (Message)
	req, err := s.client.NewRequest("PATCH", u, appReq)

	if err != nil {
		return a, nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}

func (s *OpenpitrixService) DeleteApp(ctx context.Context, appReq *App, token *AccessToken) (  *Message, *Response, error) {
	appId := appReq.AppId
	u := fmt.Sprintf("/kapis/openpitrix.io/v1/apps/%s", *appId)

	a := new (Message)
	req, err := s.client.NewRequest("DELETE", u, nil)

	if err != nil {
		return a, nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.GetAccesstoken())

	resp, err := s.client.Do(ctx, req, a)

	if err != nil {
		return a, resp, err
	}
	return a, resp, nil
}