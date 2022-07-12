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

func (t *Todoist) GetSections(ctx context.Context, projectId int) (sections []Section, err error) {
	sections = make([]Section, 0)
	if err = t.request(ctx, http.MethodGet, SectionsEndpoint, map[string]string{"project_id": strconv.Itoa(projectId)}, nil, &sections); err != nil {
		return
	}

	return
}

func (t *Todoist) AddSection(ctx context.Context, name string, projectId int, order int) (section *Section, err error) {
	params := map[string]interface{}{
		"name":       name,
		"project_id": projectId,
	}

	if order != 0 {
		params["order"] = order
	}

	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	section = new(Section)
	if err = t.request(ctx, http.MethodPost, SectionsEndpoint, nil, bytes.NewBuffer(payload), section); err != nil {
		return
	}

	return
}

func (t *Todoist) GetSection(ctx context.Context, sectionId int) (section *Section, err error) {
	section = new(Section)
	if err = t.request(ctx, http.MethodGet, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, nil, section); err != nil {
		return
	}

	return
}

func (t *Todoist) UpdateSection(ctx context.Context, sectionId int, name string) (err error) {
	params := map[string]interface{}{
		"name": name,
	}

	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	if err = t.request(ctx, http.MethodPost, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, bytes.NewBuffer(payload), nil); err != nil {
		return
	}

	return
}

func (t *Todoist) DeleteSection(ctx context.Context, sectionId int) (err error) {
	if err = t.request(ctx, http.MethodDelete, SectionsEndpoint+"/"+strconv.Itoa(sectionId), nil, nil, nil); err != nil {
		return
	}

	return
}
