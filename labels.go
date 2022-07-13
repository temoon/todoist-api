package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const LabelsEndpoint = "labels"

type Label struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Color    int    `json:"color"`
	Order    int    `json:"order"`
	Favorite bool   `json:"favorite"`
}

// region GetLabels

func (t *Todoist) GetLabels(ctx context.Context) (labels []Label, err error) {
	labels = make([]Label, 0)
	err = t.request(ctx, http.MethodGet, LabelsEndpoint, nil, nil, &labels)

	return
}

// endregion

// region AddLabel

type AddLabelParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddLabelParams() *AddLabelParams {
	params := make(AddLabelParams)
	return &params
}

func (p *AddLabelParams) WithName(name string) *AddLabelParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddLabelParams) WithOrder(order int) *AddLabelParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *AddLabelParams) WithColor(color int) *AddLabelParams {
	if color != 0 {
		(*p)["color"] = color
	}

	return p
}

func (p *AddLabelParams) WithFavorite(favorite bool) *AddLabelParams {
	(*p)["favorite"] = favorite
	return p
}

func (t *Todoist) AddLabel(ctx context.Context, params *AddLabelParams) (label *Label, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	label = new(Label)
	err = t.request(ctx, http.MethodPost, LabelsEndpoint, nil, bytes.NewBuffer(payload), label)

	return
}

// endregion

// region GetLabel

func (t *Todoist) GetLabel(ctx context.Context, labelId int) (label *Label, err error) {
	label = new(Label)
	err = t.request(ctx, http.MethodGet, LabelsEndpoint+"/"+strconv.Itoa(labelId), nil, nil, label)

	return
}

// endregion

// region UpdateLabel

type UpdateLabelParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateLabelParams() *UpdateLabelParams {
	params := make(UpdateLabelParams)
	return &params
}

func (p *UpdateLabelParams) WithName(name string) *UpdateLabelParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *UpdateLabelParams) WithOrder(order int) *UpdateLabelParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *UpdateLabelParams) WithColor(color int) *UpdateLabelParams {
	if color != 0 {
		(*p)["color"] = color
	}

	return p
}

func (p *UpdateLabelParams) WithFavorite(favorite bool) *UpdateLabelParams {
	(*p)["favorite"] = favorite
	return p
}

func (t *Todoist) UpdateLabel(ctx context.Context, labelId int, params *UpdateLabelParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	return t.request(ctx, http.MethodPost, LabelsEndpoint+"/"+strconv.Itoa(labelId), nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region DeleteLabel

func (t *Todoist) DeleteLabel(ctx context.Context, labelId int) (err error) {
	return t.request(ctx, http.MethodDelete, LabelsEndpoint+"/"+strconv.Itoa(labelId), nil, nil, nil)
}

// endregion
