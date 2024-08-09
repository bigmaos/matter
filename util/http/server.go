package http

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mlee-msl/taskgroup"
)

type Router interface {
	CreateRouter() *gin.Engine
}

type ServerManager struct {
	// server name -> gin server
	servers map[string]*GinServer
}

type GinServer struct {
	router     *gin.Engine
	port       string
	serverName string
}

func NewServerManager() *ServerManager {
	return &ServerManager{
		servers: make(map[string]*GinServer),
	}
}

func NewGinServer(router Router, port string) *GinServer {
	return &GinServer{
		router: router.CreateRouter(),
		port:   port,
	}
}

func (m *ServerManager) RegisterServer(server *GinServer, serverName string) error {
	if serverName == "" {
		return fmt.Errorf("server name is empty")
	}
	if server == nil {
		return fmt.Errorf("server is nil")
	}
	if _, ok := m.servers[serverName]; ok {
		return fmt.Errorf("server %s already exists", serverName)
	}
	m.servers[serverName] = server
	return nil
}

func (m *ServerManager) RunServer() {
	err := m.StartServers()
	if err != nil {
		log.Fatalf("in run server:\n%v", err)
	}

}

// 并发start server
func (m *ServerManager) StartServers() error {
	var tasks []*taskgroup.Task

	taskfNO := 0
	for _, server := range m.servers {
		tasks = append(tasks,
			taskgroup.NewTask(uint32(taskfNO),
				func() (interface{}, error) {
					return nil, server.Start()
				},
				true))
		taskfNO++
	}

	tg := taskgroup.NewTaskGroup(
		taskgroup.WithWorkerNums(uint32(len(tasks))),
	)

	tg.AddTask(tasks...)
	_, err := tg.Run()
	if err != nil {
		return fmt.Errorf("start server error:%v", err)
	}
	return nil
}

func (g *GinServer) Start() error {
	svr := http.Server{
		Addr:    ":" + g.port,
		Handler: g.router,
	}

	if err := svr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
