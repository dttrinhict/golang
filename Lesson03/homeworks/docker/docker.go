package docker

import (
	"context"
	"github.com/docker/docker/api/types/filters"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Docker struct {
	ctx context.Context
	Client *client.Client
}

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

func (d *Docker) DockerContainerListAll() (containers []types.Container, err error){
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{All: true})
}

func (d *Docker) DockerContainerListRunning() (containers []types.Container, err error){
	filter := filters.NewArgs()
	//filter.Add("status", "exited")
	filter.Add("status", "running")
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{
			Filters: filter,
		})
}


func (d *Docker) DockerContainerListExited() (containers []types.Container, err error){
	filter := filters.NewArgs()
	filter.Add("status", "exited")
	//filter.Add("status", "running")
	return d.Client.ContainerList(d.ctx, types.ContainerListOptions{
		Filters: filter,
	})
}

func DockerExample() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
		Tty:   false,
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
}