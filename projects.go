package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const ProjectsEndpoint = "projects"

type Project struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Color        int    `json:"color"`
	ParentId     int    `json:"parent_id"`
	Order        int    `json:"order"`
	CommentCount int    `json:"comment_count"`
	Shared       bool   `json:"shared"`
	Favorite     bool   `json:"favorite"`
	InboxProject bool   `json:"inbox_project"`
	TeamInbox    bool   `json:"team_inbox"`
	SyncId       int    `json:"sync_id"`
	Url          string `json:"url"`
}

type Collaborator struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (t *Todoist) GetProjects(ctx context.Context) (projects []Project, err error) {
	projects = make([]Project, 0)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint, nil, nil, &projects)

	return
}

type AddProjectParams map[string]interface{}

func (p *AddProjectParams) WithName(name string) *AddProjectParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *AddProjectParams) WithParentId(parentId int) *AddProjectParams {
	if parentId != 0 {
		(*p)["parent_id"] = parentId
	}

	return p
}

func (p *AddProjectParams) WithColor(color int) *AddProjectParams {
	if color != 0 {
		(*p)["color"] = color
	}

	return p
}

func (p *AddProjectParams) WithFavorite(favorite bool) *AddProjectParams {
	(*p)["favorite"] = favorite
	return p
}

func (t *Todoist) AddProject(ctx context.Context, params *AddProjectParams) (project *Project, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	project = new(Project)
	err = t.request(ctx, http.MethodPost, ProjectsEndpoint, nil, bytes.NewBuffer(payload), project)

	return
}

func (t *Todoist) GetProject(ctx context.Context, projectId int) (project *Project, err error) {
	project = new(Project)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, nil, project)

	return
}

type UpdateProjectParams map[string]interface{}

func (p *UpdateProjectParams) WithName(name string) *UpdateProjectParams {
	if name != "" {
		(*p)["name"] = name
	}

	return p
}

func (p *UpdateProjectParams) WithColor(color int) *UpdateProjectParams {
	if color != 0 {
		(*p)["color"] = color
	}

	return p
}

func (p *UpdateProjectParams) WithFavorite(favorite bool) *UpdateProjectParams {
	(*p)["favorite"] = favorite
	return p
}

func (t *Todoist) UpdateProject(ctx context.Context, projectId int, params *UpdateProjectParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	return t.request(ctx, http.MethodPost, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, bytes.NewBuffer(payload), nil)
}

func (t *Todoist) DeleteProject(ctx context.Context, projectId int) (err error) {
	return t.request(ctx, http.MethodDelete, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, nil, nil)
}

func (t *Todoist) GetCollaborators(ctx context.Context, projectId int) (collaborators []Collaborator, err error) {
	collaborators = make([]Collaborator, 0)
	err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+strconv.Itoa(projectId)+"/collaborators", nil, nil, &collaborators)

	return
}
