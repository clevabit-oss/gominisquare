package container

import (
	"net"
	"os"
	"time"
)

type Status int

const (
	Creating Status = iota
	Created
	Resolved
	Installed
	Starting
	Running
	Pausing
	Paused
	Stopping
	Stopped
	Downloading
	Updating
	Failed
)

func (s Status) String() string {
	switch s {
	case Created:
		return "created"
	case Resolved:
		return "resolved"
	case Installed:
		return "installed"
	case Starting:
		return "starting"
	case Running:
		return "running"
	case Pausing:
		return "pausing"
	case Paused:
		return "paused"
	case Stopping:
		return "stopping"
	case Stopped:
		return "stopped"
	case Downloading:
		return "downloading"
	case Updating:
		return "updating"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}

type Mount struct {
	Source      string
	Destination string
	ReadOnly    bool
	Merge       bool
}

type Config struct {
	IpAddress4 net.Addr
	IpAddress6 net.Addr
	Interface  string
	Root       string
	Mounts     []Mount
}

type Container interface {
	ID() string

	Status() Status

	PID() uint

	Created() time.Time

	Config() Config

	Create() error

	Resolve() error

	Start() error

	Stop() error

	Pause() error

	Resume() error

	Update() error

	Destroy() error

	Signal(signal os.Signal) error
}
