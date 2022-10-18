package api

import (
	"net/http"
	"todo/dal"

	"github.com/labstack/echo/v4"
)

type CreateTaskAPIRequest struct {
	Subject     string `json:"subject"` //struct tags - anotacao de campos de struct
	Description string `json:"description"`
}

type UpdateTaskAPIRequest struct {
	Subject     string `json:"subject"` //struct tags - anotacao de campos de struct
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskAPIResponse struct {
	Task dal.Task `json:"task"`
}

type ListTasksAPIResponse struct {
	Tasks []dal.Task `json:"tasks"`
}

/*
{
  "error": {
    "message": "sunt dolore"
  }
}
*/

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorAPIResponse struct {
	Error ErrorResponse `json:"error"`
}

type Server struct {
	taskDal dal.DataAccessLayerInterface
}

func NewServer(taskDal dal.DataAccessLayerInterface) *Server { //injecao de dependencias
	return &Server{
		taskDal: taskDal,
	}
}

func (s *Server) HTTPErrorHandler(err error, c echo.Context) {
	if err == dal.ErrNotFound {
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

func (s *Server) Start(address string) error {
	e := echo.New()

	e.HTTPErrorHandler = s.HTTPErrorHandler

	// CREATE
	e.POST("/tasks", func(c echo.Context) error {
		var request CreateTaskAPIRequest

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		task, err := s.taskDal.CreateTask(dal.CreateTaskRequest{
			Subject:     request.Subject,
			Description: request.Subject,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, TaskAPIResponse{
			Task: task,
		})
	})

	// READ
	e.GET("/tasks/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")
		task, err := s.taskDal.ReadTask(taskID)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, TaskAPIResponse{
			Task: task,
		})
	})

	// UPDATE
	e.PUT("/tasks/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")

		var request UpdateTaskAPIRequest

		err := c.Bind(&request)
		if err != nil {
			return err
		}

		task, err := s.taskDal.UpdateTask(taskID, dal.UpdateTaskRequest{
			Subject:     request.Subject,
			Description: request.Description,
			Status:      request.Status,
		})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, TaskAPIResponse{
			Task: task,
		})
	})

	// DELETE
	e.DELETE("/tasks/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")
		err := s.taskDal.DeleteTask(taskID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	// LIST
	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := s.taskDal.ListAllTasks(dal.ListTaskRequest{})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, ListTasksAPIResponse{
			Tasks: tasks,
		})
	})

	return e.Start(":8080")
}
