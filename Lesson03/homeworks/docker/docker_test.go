package docker

import (
	"github.com/docker/docker/api/types"
	"testing"
)

//go test -bench .
var docker = InitDocker()
var containers = []types.Container{}
func Benchmark_DockerContainerListAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_ = docker.DockerContainerListAll()
	}
}

func BenchmarkDocker_DockerContainerListExited(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_ = docker.DockerContainerListExited()
	}
}

func BenchmarkDocker_DockerContainerListRunning(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_ = docker.DockerContainerListRunning()
	}
}

func BenchmarkDocker_DockerContainerStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = docker.DockerContainerStart("dockercontainerID")
	}
}

func BenchmarkDocker_DockerContainerStop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = docker.DockerContainerStop("dockercontainerID")
	}
}

func BenchmarkDocker_DockerGetContainerByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_,_ = docker.DockerGetContainerByName("ContainerName")
	}
}

func BenchmarkDocker_DockerPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		docker.DockerPrint(containers)
	}
}