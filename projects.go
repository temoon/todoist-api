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
	Id           int    `json:"id,omitempty"`
	Name         string `json:"name"`
	Color        int    `json:"color,omitempty"`
	ParentId     int    `json:"parent_id,omitempty"`
	Order        int    `json:"order,omitempty"`
	CommentCount int    `json:"comment_count,omitempty"`
	Shared       bool   `json:"shared,omitempty"`
	Favorite     bool   `json:"favorite,omitempty"`
	InboxProject bool   `json:"inbox_project,omitempty"`
	TeamInbox    bool   `json:"team_inbox,omitempty"`
	SyncId       int    `json:"sync_id,omitempty"`
	Url          string `json:"url,omitempty"`
}

type Collaborator struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (t *Todoist) GetProjects(ctx context.Context) (projects []Project, err error) {
	projects = make([]Project, 0)
	if err = t.request(ctx, http.MethodGet, ProjectsEndpoint, nil, nil, &projects); err != nil {
		return
	}

	return
}

func (t *Todoist) AddProjects(ctx context.Context, name string, parentId int, color int, favorite bool) (project *Project, err error) {
	params := map[string]interface{}{
		"name":     name,
		"favorite": favorite,
	}

	if parentId != 0 {
		params["parent_id"] = parentId
	}

	if color != 0 {
		params["color"] = color
	}

	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	project = new(Project)
	if err = t.request(ctx, http.MethodPost, ProjectsEndpoint, nil, bytes.NewBuffer(payload), project); err != nil {
		return
	}

	return
}

func (t *Todoist) GetProject(ctx context.Context, projectId int) (project *Project, err error) {
	project = new(Project)
	if err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, nil, project); err != nil {
		return
	}

	return
}

func (t *Todoist) UpdateProject(ctx context.Context, projectId int, name string, color int, favorite bool) (err error) {
	params := map[string]interface{}{
		"name":     name,
		"favorite": favorite,
	}

	if color != 0 {
		params["color"] = color
	}

	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	if err = t.request(ctx, http.MethodPost, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, bytes.NewBuffer(payload), nil); err != nil {
		return
	}

	return
}

func (t *Todoist) DeleteProject(ctx context.Context, projectId int) (err error) {
	if err = t.request(ctx, http.MethodDelete, ProjectsEndpoint+"/"+strconv.Itoa(projectId), nil, nil, nil); err != nil {
		return
	}

	return
}

func (t *Todoist) GetCollaborators(ctx context.Context, projectId int) (collaborators []Collaborator, err error) {
	collaborators = make([]Collaborator, 0)
	if err = t.request(ctx, http.MethodGet, ProjectsEndpoint+"/"+strconv.Itoa(projectId)+"/collaborators", nil, nil, &collaborators); err != nil {
		return
	}

	return
}
