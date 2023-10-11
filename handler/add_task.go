package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/hitoshi-w/go-lang/entity"
	"github.com/hitoshi-w/go-lang/store"
)

type AddTask struct {
	Store     *store.TaskStore
	Validator *validator.Validate
}

// http.ResponseWriter（HTTPレスポンスを書き込むためのインターフェース）
// *http.Request（HTTPリクエストを表す構造体）
func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// リクエストからコンテキスト（context.Context）を取得。コンテキストは、リクエストのライフサイクルを追跡し、リクエスト間で値を共有するために使用
	ctx := r.Context()

	var b struct {
		Title string `json:"title" validate:"required"`
	}
	// リクエストボディからJSONデータをb構造体にデコード
	// JSONデータが不正な場合はエラー
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)

		return
	}

	t := &entity.Task{
		Title:   b.Title,
		Status:  entity.TaskStatusToDo,
		Created: time.Now(),
	}
	id, err := store.Tasks.Add(t)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)

		return
	}

	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: id}

	RespondJSON(ctx, w, rsp, http.StatusOK)
}
