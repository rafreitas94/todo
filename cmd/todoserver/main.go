package main

import (
	"fmt"
	"net/http"
	"todo/dal"

	"github.com/labstack/echo/v4"
)

type TaskAPIResponse struct {
	Task dal.Task `json:"task"`
}

type ListTaskAPIResponse struct {
	Tasks []dal.Task `json:"tasks"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorAPIResponse struct {
	Error ErrorResponse `json:"error"`
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if err == dal.ErrorNotFound {
			c.JSON(http.StatusNotFound, ErrorAPIResponse{
				Error: ErrorResponse{
					Message: err.Error(),
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorAPIResponse{
			Error: ErrorResponse{
				Message: err.Error(),
			},
		})
	}

	dalInterface := dal.NewDataAccessLayer()
	newTask, _ := dalInterface.CreateTask(dal.CreateTaskRequest{
		Subject:     "Terminar endpoints",
		Description: "falta terminar todos os endpoints",
	})

	fmt.Println("ID:", newTask.ID)

	e.GET("/tasks/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")
		task, err := dalInterface.ReadTask(taskID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, TaskAPIResponse{
			Task: task,
		})
	})

	e.DELETE("/tasks/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")
		err := dalInterface.DeleteTask(taskID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := dalInterface.ListAllTasks(dal.ListTaskRequest{})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, ListTaskAPIResponse{
			Tasks: tasks,
		})
	})

	e.Start(":3000")
}
