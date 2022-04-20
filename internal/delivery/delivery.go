package delivery

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wagaru/task/config"
	"github.com/wagaru/task/internal/errcode"
	"github.com/wagaru/task/internal/service"
)

type delivery struct {
	svc    service.Service
	engine *gin.Engine
	config *config.ServerConfig
	server *http.Server
}

func NewDelivery(svc service.Service, config *config.ServerConfig) *delivery {
	gin.SetMode(config.RunMode)
	delivery := &delivery{
		svc:    svc,
		engine: gin.New(),
		config: config,
	}
	delivery.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", delivery.config.Port),
		Handler: delivery.engine,
	}
	delivery.buildRoute()
	return delivery
}

func (d *delivery) buildRoute() {
	d.engine.GET("/tasks", d.GetTasks)
	d.engine.POST("/tasks", d.CreateTask)
	d.engine.PUT("/tasks/:id", d.UpdateTask)
	d.engine.DELETE("/tasks/:id", d.DeleteTask)
	d.engine.NoRoute(d.NoRoute)
}

func (d *delivery) Run() error {
	log.Printf("start the server in %s", d.server.Addr)
	return d.server.ListenAndServe()
}

func (d *delivery) Shutdown(ctx context.Context) error {
	return d.server.Shutdown(ctx)
}

func (d *delivery) GetTasks(c *gin.Context) {
	tasks, err := d.svc.GetTasks()
	if err != nil {
		d.ToErrorResponse(c, err)
		return
	}
	data := map[string]interface{}{
		"result": tasks,
	}
	d.ToResponse(c, http.StatusOK, data)
}

func (d *delivery) CreateTask(c *gin.Context) {
	var params CreateTaskRequest
	var err error
	if err = c.ShouldBindJSON(&params); err != nil {
		d.ToErrorResponse(c, errcode.InvalidParams)
		return
	}
	task, err := d.svc.CreateTask(params.Name)
	if err != nil {
		d.ToErrorResponse(c, err)
		return
	}
	data := map[string]interface{}{
		"result": task,
	}
	d.ToResponse(c, http.StatusCreated, data)
}

func (d *delivery) UpdateTask(c *gin.Context) {
	var params UpdateTaskRequest
	if err := c.ShouldBindJSON(&params); err != nil {
		d.ToErrorResponse(c, errcode.InvalidParams)
		return
	}
	id := c.Param("id")
	uid, _ := strconv.Atoi(id)
	task, err := d.svc.UpdateTask(uint32(uid), params.Name, *params.Status)
	if err != nil {
		d.ToErrorResponse(c, err)
		return
	}
	data := map[string]interface{}{
		"result": task,
	}
	d.ToResponse(c, http.StatusOK, data)
}

func (d *delivery) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	uid, _ := strconv.Atoi(id)
	err := d.svc.DeleteTask(uint32(uid))
	if err != nil {
		d.ToErrorResponse(c, err)
		return
	}
	d.ToResponse(c, http.StatusOK, nil)
}

func (d *delivery) NoRoute(c *gin.Context) {
	d.ToErrorResponse(c, errcode.NotFound)
}
