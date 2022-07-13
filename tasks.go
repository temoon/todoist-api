package todoist

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

const TasksEndpoint = "tasks"

type Task struct {
	Id           int    `json:"id"`
	ProjectId    int    `json:"project_id"`
	SectionId    int    `json:"section_id"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	Completed    bool   `json:"completed"`
	LabelIds     []int  `json:"label_ids"`
	ParentId     int    `json:"parent_id"`
	Order        int    `json:"order"`
	Priority     int    `json:"priority"`
	Due          Due    `json:"due"`
	Url          string `json:"url"`
	CommentCount int    `json:"comment_count"`
	Assignee     int    `json:"assignee"`
	Assigner     int    `json:"assigner"`
}

type Due struct {
	String    string `json:"string"`
	Date      string `json:"date"`
	Recurring bool   `json:"recurring"`
	Datetime  string `json:"datetime"`
	Timezone  string `json:"timezone"`
}

// region GetTasks

type GetTasksParams map[string]string

//goland:noinspection GoUnusedExportedFunction
func MakeGetTasksParams() *GetTasksParams {
	params := make(GetTasksParams)
	return &params
}

func (p *GetTasksParams) WithProjectId(projectId int) *GetTasksParams {
	if projectId != 0 {
		(*p)["project_id"] = strconv.Itoa(projectId)
	}

	return p
}

func (p *GetTasksParams) WithSectionId(sectionId int) *GetTasksParams {
	if sectionId != 0 {
		(*p)["section_id"] = strconv.Itoa(sectionId)
	}

	return p
}

func (p *GetTasksParams) WithLabelId(labelId int) *GetTasksParams {
	if labelId != 0 {
		(*p)["label_id"] = strconv.Itoa(labelId)
	}

	return p
}

func (p *GetTasksParams) WithFilter(filter string) *GetTasksParams {
	if filter != "" {
		(*p)["filter"] = filter
	}

	return p
}

func (p *GetTasksParams) WithLang(lang string) *GetTasksParams {
	if lang != "" {
		(*p)["lang"] = lang
	}

	return p
}

func (p *GetTasksParams) WithIds(ids []int) *GetTasksParams {
	if ids != nil && len(ids) != 0 {
		value := strings.Builder{}
		value.WriteString(strconv.Itoa(ids[0]))
		for i := 1; i < len(ids); i++ {
			value.WriteByte(',')
			value.WriteString(strconv.Itoa(ids[i]))
		}

		(*p)["ids"] = value.String()
	}

	return p
}

func (t *Todoist) GetTasks(ctx context.Context, params *GetTasksParams) (tasks []Task, err error) {
	tasks = make([]Task, 0)
	err = t.request(ctx, http.MethodGet, TasksEndpoint, *params, nil, &tasks)

	return
}

// endregion

// region AddTask

type AddTaskParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeAddTaskParams() *AddTaskParams {
	params := make(AddTaskParams)
	return &params
}

func (p *AddTaskParams) WithContent(content string) *AddTaskParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *AddTaskParams) WithDescription(description string) *AddTaskParams {
	if description != "" {
		(*p)["description"] = description
	}

	return p
}

func (p *AddTaskParams) WithProjectId(projectId int) *AddTaskParams {
	if projectId != 0 {
		(*p)["project_id"] = projectId
	}

	return p
}

func (p *AddTaskParams) WithSectionId(sectionId int) *AddTaskParams {
	if sectionId != 0 {
		(*p)["section_id"] = sectionId
	}

	return p
}

func (p *AddTaskParams) WithParentId(parentId int) *AddTaskParams {
	if parentId != 0 {
		(*p)["parent_id"] = parentId
	}

	return p
}

func (p *AddTaskParams) WithOrder(order int) *AddTaskParams {
	if order != 0 {
		(*p)["order"] = order
	}

	return p
}

func (p *AddTaskParams) WithLabelIds(labelIds []int) *AddTaskParams {
	if labelIds != nil && len(labelIds) != 0 {
		(*p)["label_ids"] = labelIds
	}

	return p
}

func (p *AddTaskParams) WithPriority(priority int) *AddTaskParams {
	if priority != 0 {
		(*p)["priority"] = priority
	}

	return p
}

func (p *AddTaskParams) WithDueString(dueString string) *AddTaskParams {
	if dueString != "" {
		(*p)["due_string"] = dueString
	}

	return p
}

func (p *AddTaskParams) WithDueDate(dueDate string) *AddTaskParams {
	if dueDate != "" {
		(*p)["due_date"] = dueDate
	}

	return p
}

func (p *AddTaskParams) WithDueDatetime(dueDatetime string) *AddTaskParams {
	if dueDatetime != "" {
		(*p)["due_datetime"] = dueDatetime
	}

	return p
}

func (p *AddTaskParams) WithDueLang(dueLang string) *AddTaskParams {
	if dueLang != "" {
		(*p)["due_lang"] = dueLang
	}

	return p
}

func (p *AddTaskParams) WithAssignee(assignee int) *AddTaskParams {
	if assignee != 0 {
		(*p)["assignee"] = assignee
	}

	return p
}

func (t *Todoist) AddTask(ctx context.Context, params *AddTaskParams) (task *Task, err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	task = new(Task)
	err = t.request(ctx, http.MethodPost, TasksEndpoint, nil, bytes.NewBuffer(payload), task)

	return
}

// endregion

// region GetTask

func (t *Todoist) GetTask(ctx context.Context, taskId int) (task *Task, err error) {
	task = new(Task)
	err = t.request(ctx, http.MethodGet, TasksEndpoint+"/"+strconv.Itoa(taskId), nil, nil, task)

	return
}

// endregion

// region UpdateTask

type UpdateTaskParams map[string]interface{}

//goland:noinspection GoUnusedExportedFunction
func MakeUpdateTaskParams() *UpdateTaskParams {
	params := make(UpdateTaskParams)
	return &params
}

func (p *UpdateTaskParams) WithContent(content string) *UpdateTaskParams {
	if content != "" {
		(*p)["content"] = content
	}

	return p
}

func (p *UpdateTaskParams) WithDescription(description string) *UpdateTaskParams {
	if description != "" {
		(*p)["description"] = description
	}

	return p
}

func (p *UpdateTaskParams) WithLabelIds(labelIds []int) *UpdateTaskParams {
	if labelIds != nil && len(labelIds) != 0 {
		(*p)["label_ids"] = labelIds
	}

	return p
}

func (p *UpdateTaskParams) WithPriority(priority int) *UpdateTaskParams {
	if priority != 0 {
		(*p)["priority"] = priority
	}

	return p
}

func (p *UpdateTaskParams) WithDueString(dueString string) *UpdateTaskParams {
	if dueString != "" {
		(*p)["due_string"] = dueString
	}

	return p
}

func (p *UpdateTaskParams) WithDueDate(dueDate string) *UpdateTaskParams {
	if dueDate != "" {
		(*p)["due_date"] = dueDate
	}

	return p
}

func (p *UpdateTaskParams) WithDueDatetime(dueDatetime string) *UpdateTaskParams {
	if dueDatetime != "" {
		(*p)["due_datetime"] = dueDatetime
	}

	return p
}

func (p *UpdateTaskParams) WithDueLang(dueLang string) *UpdateTaskParams {
	if dueLang != "" {
		(*p)["due_lang"] = dueLang
	}

	return p
}

func (p *UpdateTaskParams) WithAssignee(assignee int) *UpdateTaskParams {
	if assignee != 0 {
		(*p)["assignee"] = assignee
	}

	return p
}

func (t *Todoist) UpdateTask(ctx context.Context, taskId int, params *UpdateTaskParams) (err error) {
	var payload []byte
	if payload, err = json.Marshal(params); err != nil {
		return
	}

	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+strconv.Itoa(taskId), nil, bytes.NewBuffer(payload), nil)
}

// endregion

// region CloseTask

func (t *Todoist) CloseTask(ctx context.Context, taskId int) (err error) {
	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+strconv.Itoa(taskId)+"/close", nil, nil, nil)
}

// endregion

// region ReopenTask

func (t *Todoist) ReopenTask(ctx context.Context, taskId int) (err error) {
	return t.request(ctx, http.MethodPost, TasksEndpoint+"/"+strconv.Itoa(taskId)+"/reopen", nil, nil, nil)
}

// endregion

// region DeleteTask

func (t *Todoist) DeleteTask(ctx context.Context, taskId int) (err error) {
	return t.request(ctx, http.MethodDelete, TasksEndpoint+"/"+strconv.Itoa(taskId), nil, nil, nil)
}

// endregion
