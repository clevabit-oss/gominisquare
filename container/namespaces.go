package container

import (
	"os"
	"path/filepath"
)

const namespaceBaseDir = "/sys/fs/cgroup"

type namespace string

const (
	cpuNamespace     namespace = "cpu"
	cpusetNamespace  namespace = "cpuset"
	devicesNamespace namespace = "devices"
	freezerNamespace namespace = "freezer"
	memoryNamespace  namespace = "memory"
	networkNamespace namespace = "net_cls"
	pidNamespace     namespace = "pids"
)

func namespaceDir(namespace namespace) string {
	return filepath.Join(namespaceBaseDir, string(namespace))
}

func newPidNamespace(name string) (uintptr, error) {
	ns := namespaceDir(pidNamespace)

	if _, err := os.Stat(ns); os.IsExist(err) {
		panic("not yet implemented") // TODO
	}

	if err := os.Mkdir(ns, os.ModePerm); err != nil {
		return -1, err
	}

	pidmax := filepath.Join(ns, "pids.max")
	f, err := os.Open(pidmax)
	if err != nil {
		return -1, err
	}

	if _, err := f.WriteString("1"); err != nil {
		return -1, err
	}


}
