package pattern

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"strings"
	"time"
)

var docker *Docker
var dockerClient *client.Client
var ctx = context.Background()
var duration = 3 * time.Second

type Docker struct {
	Context context.Context
	Client *client.Client
}

func NewDockerClient() *client.Client {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	if _, err = dockerClient.Ping(context.Background()); err != nil {
		panic(err.Error())
	}
	return dockerClient
}

func NewDocker() IDocker {
	if docker == nil {
		dockerClient = NewDockerClient()
		docker = &Docker{
			Context: ctx,
			Client: dockerClient,
		}
	}else{
		if dockerClient == nil {
			dockerClient = NewDockerClient()
			docker = &Docker{
				Context: ctx,
				Client: dockerClient,
			}
		}
	}
	return docker
}

func (d *Docker) DockerClient() {
	if d.Client == nil {
		dockerClient = NewDockerClient()
		d.Context = ctx
		d.Client = dockerClient
	}
}
/* Get Containers By Status
https://docs.docker.com/engine/api/v1.41/#tag/Container
status=(created|restarting|running|removing|paused|exited|dead)
*/
func (d *Docker) GetContainersByStatus(status string) (containers []types.Container, err error) {
	d.DockerClient()
	filter := filters.NewArgs()
	filter.Add("status", status)
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		Filters: filter,
	})
}

/* Get All Containers
*/
func (d *Docker) GetAllContainers() (containers []types.Container, err error) {
	d.DockerClient()
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		All: true,
	})
}

/*Get Containers By Container ID
*/
func (d *Docker) GetContainerById(containerID string) (containers []types.Container, err error) {
	allContainers, err := d.GetAllContainers()
	if err != nil {
		return nil, err
	}
	for _, c := range allContainers {
		if strings.Contains(c.ID, containerID) {
			containers = append(containers, c)
		}
	}
	return containers, err
}

/* Get Containers By Container Name
*/
func (d *Docker) GetContainerByName(containerName string) (containers []types.Container, err error) {
	allContainers, err := d.GetAllContainers()
	if err != nil {
		return nil, err
	}
	for _, c := range allContainers {
		for _, cname := range c.Names {
			if strings.Compare(TrimSlash(cname), containerName) == 0 {
				containers = append(containers, c)
			}
		}
	}
	return containers, err
}

/* Get Running Container By ID
*/
func (d *Docker) GetRunningContainerById(containerID string) (containers []types.Container, err error) {
	d.DockerClient()
	optionFilters := filters.NewArgs()
	optionFilters.Add("id", containerID)
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		Filters: optionFilters,
	})
}

/* Get Running Container By Name
*/
func (d *Docker) GetRunningContainerByName(containerName string) (containers []types.Container, err error) {
	d.DockerClient()
	optionFilters := filters.NewArgs()
	optionFilters.Add("name", containerName)
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		Filters: optionFilters,
	})
}

/* Get Running Containers
*/
func (d *Docker) GetRunningContainers() (containers []types.Container, err error) {
	d.DockerClient()
	filter := filters.NewArgs()
	filter.Add("status", "running")
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		Filters: filter,
	})
}

/* Get Running Containers
 */
func (d *Docker) GetStoppedContainers() (containers []types.Container, err error) {
	d.DockerClient()
	filter := filters.NewArgs()
	filter.Add("status", "exited")
	return d.Client.ContainerList(d.Context, types.ContainerListOptions{
		Filters: filter,
	})
}

/* Stop Running Container By ID
 */
func (d *Docker) StopContainerByID(containerID string) error {
	d.DockerClient()
	return d.Client.ContainerStop(d.Context, containerID, &duration)
}

/* Stop Running Container By Name
 */
func (d *Docker) StopContainerByName(containerName string) (err error) {
	runningContainers, err := d.GetRunningContainerByName(containerName)
	if err != nil {
		return err
	}
	return d.StopContainerByID(runningContainers[0].ID)
}

/* Start Container by ID
*/
func (d *Docker) StartContainerByID(containerID string) (err error) {
	d.DockerClient()
	return d.Client.ContainerStart(d.Context, containerID, types.ContainerStartOptions{})
}

/* Start Container By Name
 */
func (d *Docker) StartContainerByName(containerName string) (err error) {
	containers, err := d.GetAllContainers()
	if err != nil {
		return err
	}
	return d.Client.ContainerStart(d.Context, containers[0].ID, types.ContainerStartOptions{})
}

func (d *Docker) Close() (err error) {
	return d.Client.Close()
}