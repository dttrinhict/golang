package pattern

import (
	"github.com/docker/docker/api/types"
)

type IDocker interface {
	DockerClient()
	GetContainersByStatus(status string) (containers []types.Container, err error)
	GetAllContainers() (containers []types.Container, err error)
	GetRunningContainerById(containerID string) (containers []types.Container, err error)
	GetContainerById(containerID string) (containers []types.Container, err error)
	GetRunningContainerByName(containerName string) (containers []types.Container, err error)
	GetContainerByName(containerName string) (containers []types.Container, err error)
	GetRunningContainers() (containers []types.Container, err error)
	GetStoppedContainers() (containers []types.Container, err error)
	StopContainerByID(containerID string) (err error)
	StopContainerByName(containerName string) (err error)
	StartContainerByID(containerID string) (err error)
	StartContainerByName(containerName string) (err error)
	Close() (err error)
}
