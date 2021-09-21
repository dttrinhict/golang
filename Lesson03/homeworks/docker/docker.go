package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
		"github.com/docker/docker/client"
	"strings"
	"time"
)
/* Định nghĩa struct Docker
*/
type Docker struct {
	ctx context.Context
	Client *client.Client
}
/* Khởi tạo đối được Docker Client
*/
func InitDocker() (docker *Docker) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	return &Docker{
		ctx: ctx,
		Client: cli,
	}
}
/* Lấy ra danh sách tất cả các containers
*/
func (d *Docker) DockerContainerListAll() (containers []types.Container, err error){
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{All: true})
}

/* Lấy ra danh sách containers đang chạy
*/
func (d *Docker) DockerContainerListRunning() (containers []types.Container, err error){
	filter := filters.NewArgs()
	//filter.Add("status", "exited")
	filter.Add("status", "running")
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{
			Filters: filter,
		})
}
/* Start một container đang stoped
*/
func (d *Docker) DockerContainerStart(containerID string) (err error){
	return d.Client.ContainerStart(d.ctx, containerID, types.ContainerStartOptions{})
}

/* Stop container đang running
*/
func (d *Docker) DockerContainerStop(containerID string) (err error){
	duration := 30 * time.Second
	return d.Client.ContainerStop(d.ctx, containerID, &duration)
}

/* Lấy ra danh sách containers có name xác định
*/
func (d *Docker) DockerGetContainerByName(containerName string) (result []types.Container, err error) {
	containers, err := d.DockerContainerListAll()
	if err != nil {
		return nil, err
	}
	for _, v := range containers {
		for _, name := range v.Names {
			if strings.Compare(name, containerName) == 0 {
				result = append(result, v)
			}
		}
	}
	return result, err
}
/* Lấy ra danh sách các containers stoped
*/
func (d *Docker) DockerContainerListExited() (containers []types.Container, err error){
	filter := filters.NewArgs()
	filter.Add("status", "exited")
	//filter.Add("status", "running")
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{
		Filters: filter,
	})
}
/* In danh sách containers
*/
func (d *Docker) DockerPrint(containers []types.Container) {
	fmt.Printf("%64s %25s %60s %15v:%5v  %s\n","ID", "Names", "Image", "IP","Port", "Status")
	for _, c := range containers {
		ports := c.Ports
		var port types.Port
		if len(ports) > 0 {
			port = ports[0]
		}
		status := c.Status
		if strings.Contains(status, "Exited") {
			status = "stopped"
		}else{
			status = "running"
		}
		fmt.Printf("%s %25s %60s %15v:%5v  %s\n",c.ID, c.Names[0], c.Image, port.IP,port.PublicPort, status)
	}
}