package container

import (
	"github.com/satori/uuid"
	"os"
	"time"
)

type namespacesPtrs struct {
	cpu     string
	cpuset  string
	devices string
	freezer string
	net     string
	pid     string
}

type linuxContainer struct {
	id     string
	config Config
	status Status
	pid    uint
}

func NewLinuxContainer(config Config) (Container, error) {
	id := uuid.NewV4()
	container := &linuxContainer{
		id:     id.String(),
		config: config,
		status: Creating,
		pid:    -1,
	}

	return container, nil
}

func (c *linuxContainer) ID() string {
	return c.id
}

func (c *linuxContainer) Status() Status {
	return c.status
}

func (c *linuxContainer) PID() uint {
	return c.pid
}

func (c *linuxContainer) Created() time.Time {
	panic("implement me")
}

func (c *linuxContainer) Config() Config {
	return c.config
}

func (c *linuxContainer) Create() error {
	if err := prepareNamespaces(c); err != nil {
		return err
	}

	if err := assignNamespace(c); err != nil {
		return err
	}
	return nil
}

func (c *linuxContainer) Resolve() error {
	return resolveContainerConstraints(c)
}

func (c *linuxContainer) Start() error {
	return startContainer(c)
}

func (c *linuxContainer) Stop() error {
	return stopContainer(c)
}

func (c *linuxContainer) Pause() error {
	return freezeContainer(c)
}

func (c *linuxContainer) Resume() error {
	return unfreezeContainer(c)
}

func (c *linuxContainer) Update() error {
	panic("implement me")
}

func (c *linuxContainer) Destroy() error {
	return destroyContainer(c)
}

func (c *linuxContainer) Signal(signal os.Signal) error {
	panic("implement me")
}

func prepareNamespaces(container *linuxContainer) {

}
