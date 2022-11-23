package api

import (
	"fmt"
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

	// var authenticationMiddleware echo.MiddlewareFunc
	authenticationMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("executando o middleware de autenticacao")

			cookie, err := c.Cookie("session")
			if err != nil { // se erro, cookie nao esta presente
				// 401 Unauthorized
				fmt.Println("cookie de sessao nao esta presente na requisicao.")
				return c.NoContent(http.StatusUnauthorized)
			}

			sessionId := cookie.Value

			username, err := s.taskDal.AuthenticateSession(sessionId)
			if err != nil {
				// 401 Unauthorized
				fmt.Println("cookie de sessao nao corresponde a uma sessao valida.")
				return c.NoContent(http.StatusUnauthorized)
			}

			fmt.Println("usuario autenticado via sessao", username)

			return next(c)
		}
	}

	// e.Use(authenticationMiddleware)
	// separar as rotas em rotas anonimas e rotas autenticadas
	// utilizando subrouters

	tasksGroups := e.Group("/tasks", authenticationMiddleware)
	// CREATE = prefixo do grupo + a rota == /tasks
	tasksGroups.POST("", func(c echo.Context) error {
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
	tasksGroups.GET("/:task_id", func(c echo.Context) error {
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
	tasksGroups.PUT("/:task_id", func(c echo.Context) error {
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
	tasksGroups.DELETE("/:task_id", func(c echo.Context) error {
		taskID := c.Param("task_id")
		err := s.taskDal.DeleteTask(taskID)
		if err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	// LIST
	tasksGroups.GET("", func(c echo.Context) error {
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

		return c.String(http.StatusOK, "Autenticado!")
	})

	e.POST("/form-auth", func(c echo.Context) error {

		username := c.FormValue("usuario")
		password := c.FormValue("senha")

		sessionId, err := s.taskDal.AuthenticateUser(username, password)
		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		// aqui estamos autenticados
		// setar um cookie que sirva de autenticacao do usuario
		c.SetCookie(&http.Cookie{
			Name:     "session",
			Value:    sessionId,
			Domain:   "localhost",
			Expires:  time.Now().Add(2 * time.Minute),
			Secure:   true, // cookie valido somente quando utilizando HTTPS (exceto quando em localhost)
			HttpOnly: true,
		})

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

		_, err = s.taskDal.AuthenticateUser(req.Username, req.Password)
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
