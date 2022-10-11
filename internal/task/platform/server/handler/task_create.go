package handler

import (
	"GoRestSimpleApi/internal/task"
	"GoRestSimpleApi/internal/task/create"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewTaskHandler(taskCreateService create.TaskCreateService) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		req := TaskRequest{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {

			task, err := task.NewTask(strconv.Itoa(req.Id), req.Name, req.Content)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			} else {

				err = taskCreateService.Save(ctx, task)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, err)
				} else {
					ctx.Status(http.StatusAccepted)
				}
			}
		}
	}
}
