package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

const CommentsEndpoint = "comments"

type Comment struct {
	Id         int                    `json:"id"`
	TaskId     int                    `json:"task_id"`
	ProjectId  int                    `json:"project_id"`
	Posted     string                 `json:"posted"`
	Content    string                 `json:"content"`
	Attachment map[string]interface{} `json:"attachment"`
}

type Attachment struct {
	ResourceType string `json:"resource_type"`
	FileName     string `json:"file_name"`
	FileSize     int    `json:"file_size"`
	FileType     string `json:"file_type"`
	FileUrl      string `json:"file_url"`
	UploadState  string `json:"upload_state"`
}

type ImageAttachment struct {
	Attachment

	LargeThumbnail  []interface{} `json:"tn_l"`
	MediumThumbnail []interface{} `json:"tn_m"`
	SmallThumbnail  []interface{} `json:"tn_s"`
}

type AudioAttachment struct {
	Attachment

	FileDuration int `json:"file_duration"`
}

type GetCommentsParams map[string]string

func (p *GetCommentsParams) WithProjectId(projectId int) *GetCommentsParams {
	if projectId != 0 {
		(*p)["project_id"] = strconv.Itoa(projectId)
	}

	return p
}

func (p *GetCommentsParams) WithTaskId(taskId int) *GetCommentsParams {
	if taskId != 0 {
		(*p)["task_id"] = strconv.Itoa(taskId)
	}

	return p
}

func (t *Todoist) GetComments(ctx context.Context, params *GetCommentsParams) (sections []Section, err error) {
	sections = make([]Section, 0)
	err = t.request(ctx, http.MethodGet, CommentsEndpoint, *params, nil, &sections)

	return
}

type AddCommentParams map[string]interface{}

func (p *AddCommentParams) WithTaskId(taskId int) *AddCommentParams {
	if taskId != 0 {
		(*p)["task_id"] = taskId
	}

	return p
}

func (p *AddCommentParams) WithProjectId(projectId int) *AddCommentParams {
	if projectId != 0 {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddCommentParams) WithContent(content string) *AddCommentParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *AddCommentParams) WithAttachment(attachment interface{}) *AddCommentParams {
	if attachment != nil {
		(*p)["attachment"] = attachment
	}

	return p
}

func (t *Todoist) AddComment(ctx context.Context, params *AddCommentParams) (comment *Comment, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	comment = new(Comment)
	err = t.request(ctx, http.MethodPost, CommentsEndpoint, nil, bytes.NewBuffer(payload), comment)

	return
}

func (t *Todoist) GetComment(ctx context.Context, commentId int) (comment *Comment, err error) {
	comment = new(Comment)
	err = t.request(ctx, http.MethodGet, CommentsEndpoint+"/"+strconv.Itoa(commentId), nil, nil, comment)

	return
}

type UpdateCommentParams map[string]interface{}

func (p *UpdateCommentParams) WithContent(content string) *UpdateCommentParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (t *Todoist) UpdateComment(ctx context.Context, commentId int, params *UpdateCommentParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	return t.request(ctx, http.MethodPost, CommentsEndpoint+"/"+strconv.Itoa(commentId), nil, bytes.NewBuffer(payload), nil)
}

func (t *Todoist) DeleteComment(ctx context.Context, commentId int) (err error) {
	return t.request(ctx, http.MethodDelete, CommentsEndpoint+"/"+strconv.Itoa(commentId), nil, nil, nil)
}
