package api

import (
	"net/http"
	"time"
	"todo/dal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
			Description: request.Description,
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

		cookie, err := c.Cookie("<nome do cookie>")
		if err != nil { // se erro, cookie nao esta presente
			// returnar um erro
		}

		// validar o cookie
		// if validateCookie(cookie) { returnar um erro }

		tasks, err := s.taskDal.ListAllTasks(dal.ListTaskRequest{})
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, ListTasksAPIResponse{
			Tasks: tasks,
		})
	})

	e.GET("/basic-auth", func(c echo.Context) error {

		username, password, ok := c.Request().BasicAuth()

		if !ok || (username != "usuario" || password != "senha") {
			c.Response().Header().Add(echo.HeaderWWWAuthenticate, `Basic realm="teste"`)
			return c.NoContent(http.StatusUnauthorized)
		}

		// aqui estamos autenticados

		// setar um cookie que sirva de autenticacao do usuario
		c.SetCookie(&http.Cookie{
			Name:     "<nome do cookie>",
			Value:    "",
			Domain:   "api.maua.br",
			Expires:  time.Now().Add(2 * time.Second),
			Secure:   true, // cookie valido somente quando utilizando HTTPS (exceto quando em localhost)
			HttpOnly: true,
		})

		return c.String(http.StatusOK, "Autenticado!")
	})

	e.POST("/form-auth", func(c echo.Context) error {

		username := c.FormValue("usuario")
		password := c.FormValue("senha")

		if username != "usuario" || password != "senha" {
			return c.String(http.StatusUnauthorized, "usuário e/ou senha incorretos")
		}

		return c.String(http.StatusOK, "Autenticado!")
	})

	e.POST("/api-auth", func(c echo.Context) error {
		type loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var req loginReq
		err := c.Bind(&req)
		if err != nil {
			return err
		}

		err = s.taskDal.AuthenticateUser(req.Username, req.Password)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, "Autenticado!")
	})

	// para requisicoes XHR feitas de uma pagina web que nao
	// esta hospedada no mesmo dominio da api.
	e.Use(middleware.CORS())
	// mais seguro.
	// meu site frontend ta em web.maua.br
	// minha api esta em api.maua.br
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"web.maua.br"},
	// }))

	return e.Start(":8080")
}
