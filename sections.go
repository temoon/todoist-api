package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const SectionsEndpoint = "sections"

type Section struct {
	Id        int    `json:"id"`
	ProjectId int    `json:"project_id"`
	Order     int    `json:"order"`
	Name      string `json:"name"`
}

// region GetSections

type GetSectionsParams map[string]string

//goland:noinspection GoUnusedExportedFunction
func MakeGetSectionsParams() *GetSectionsParams {
	params := make(GetSectionsParams)
	return &params
}

func (p *GetSectionsParams) WithProjectId(projectId int) *GetSectionsParams {
	if projectId != 0 {
		(*p)["project_id"] = strconv.Itoa(projectId)
	}

	return p
}

func (t *Todoist) GetSections(ctx context.Context, params *GetSectionsParams) (sections []Section, err error) {
	sections = make([]Section, 0)
	err = t.request(ctx, http.MethodGet, SectionsEndpoint, *params, nil, &sections)

	return
}

// endregion

// region AddSection

type AddSectionParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddSectionParams() *AddSectionParams {
	params := make(AddSectionParams)
	return &params
}

func (p *AddSectionParams) WithName(name string) *AddSectionParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddSectionParams) WithProjectId(projectId int) *AddSectionParams {
	if projectId != 0 {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddSectionParams) WithOrder(order int) *AddSectionParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (t *Todoist) AddSection(ctx context.Context, params *AddSectionParams) (section *Section, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	section = new(Section)
	err = t.request(ctx, http.MethodPost, SectionsEndpoint, nil, bytes.NewBuffer(payload), section)

	return
}

// endregion

// region GetSection

func (t *Todoist) GetSection(ctx context.Context, sectionId int) (section *Section, err error) {
	section = new(Section)
	err = t.request(ctx, http.MethodGet, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, nil, section)

	return
}

// endregion

// region UpdateSection

type UpdateSectionParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateSectionParams() *UpdateSectionParams {
	params := make(UpdateSectionParams)
	return &params
}

func (p *UpdateSectionParams) WithName(name string) *UpdateSectionParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (t *Todoist) UpdateSection(ctx context.Context, sectionId int, params *UpdateSectionParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	return t.request(ctx, http.MethodPost, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteSection

func (t *Todoist) DeleteSection(ctx context.Context, sectionId int) (err error) {
	return t.request(ctx, http.MethodDelete, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, nil, nil)
}

// endregion
