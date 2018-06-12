// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package sniffer

import (
	"sync"
)

var (
	lockDockerClientMockConnect              sync.RWMutex
	lockDockerClientMockGetContainers        sync.RWMutex
	lockDockerClientMockGetNetworkInterfaces sync.RWMutex
	lockDockerClientMockPullImage            sync.RWMutex
)

// DockerClientMock is a mock implementation of DockerClient.
//
//     func TestSomethingThatUsesDockerClient(t *testing.T) {
//
//         // make and configure a mocked DockerClient
//         mockedDockerClient := &DockerClientMock{
//             ConnectFunc: func(endpoint string) error {
// 	               panic("TODO: mock out the Connect method")
//             },
//             GetContainersFunc: func() ([]Container, error) {
// 	               panic("TODO: mock out the GetContainers method")
//             },
//             GetNetworkInterfacesFunc: func(containerId string) ([]NetworkInterface, error) {
// 	               panic("TODO: mock out the GetNetworkInterfaces method")
//             },
//             PullImageFunc: func(imageName string) error {
// 	               panic("TODO: mock out the PullImage method")
//             },
//         }
//
//         // TODO: use mockedDockerClient in code that requires DockerClient
//         //       and then make assertions.
//
//     }
type DockerClientMock struct {
	// ConnectFunc mocks the Connect method.
	ConnectFunc func(endpoint string) error

	// GetContainersFunc mocks the GetContainers method.
	GetContainersFunc func() ([]Container, error)

	// GetNetworkInterfacesFunc mocks the GetNetworkInterfaces method.
	GetNetworkInterfacesFunc func(containerId string) ([]NetworkInterface, error)

	// PullImageFunc mocks the PullImage method.
	PullImageFunc func(imageName string) error

	// calls tracks calls to the methods.
	calls struct {
		// Connect holds details about calls to the Connect method.
		Connect []struct {
			// Endpoint is the endpoint argument value.
			Endpoint string
		}
		// GetContainers holds details about calls to the GetContainers method.
		GetContainers []struct {
		}
		// GetNetworkInterfaces holds details about calls to the GetNetworkInterfaces method.
		GetNetworkInterfaces []struct {
			// ContainerId is the containerId argument value.
			ContainerId string
		}
		// PullImage holds details about calls to the PullImage method.
		PullImage []struct {
			// ImageName is the imageName argument value.
			ImageName string
		}
	}
}

// Connect calls ConnectFunc.
func (mock *DockerClientMock) Connect(endpoint string) error {
	if mock.ConnectFunc == nil {
		panic("moq: DockerClientMock.ConnectFunc is nil but DockerClient.Connect was just called")
	}
	callInfo := struct {
		Endpoint string
	}{
		Endpoint: endpoint,
	}
	lockDockerClientMockConnect.Lock()
	mock.calls.Connect = append(mock.calls.Connect, callInfo)
	lockDockerClientMockConnect.Unlock()
	return mock.ConnectFunc(endpoint)
}

// ConnectCalls gets all the calls that were made to Connect.
// Check the length with:
//     len(mockedDockerClient.ConnectCalls())
func (mock *DockerClientMock) ConnectCalls() []struct {
	Endpoint string
} {
	var calls []struct {
		Endpoint string
	}
	lockDockerClientMockConnect.RLock()
	calls = mock.calls.Connect
	lockDockerClientMockConnect.RUnlock()
	return calls
}

// GetContainers calls GetContainersFunc.
func (mock *DockerClientMock) GetContainers() ([]Container, error) {
	if mock.GetContainersFunc == nil {
		panic("moq: DockerClientMock.GetContainersFunc is nil but DockerClient.GetContainers was just called")
	}
	callInfo := struct {
	}{}
	lockDockerClientMockGetContainers.Lock()
	mock.calls.GetContainers = append(mock.calls.GetContainers, callInfo)
	lockDockerClientMockGetContainers.Unlock()
	return mock.GetContainersFunc()
}

// GetContainersCalls gets all the calls that were made to GetContainers.
// Check the length with:
//     len(mockedDockerClient.GetContainersCalls())
func (mock *DockerClientMock) GetContainersCalls() []struct {
} {
	var calls []struct {
	}
	lockDockerClientMockGetContainers.RLock()
	calls = mock.calls.GetContainers
	lockDockerClientMockGetContainers.RUnlock()
	return calls
}

// GetNetworkInterfaces calls GetNetworkInterfacesFunc.
func (mock *DockerClientMock) GetNetworkInterfaces(containerId string) ([]NetworkInterface, error) {
	if mock.GetNetworkInterfacesFunc == nil {
		panic("moq: DockerClientMock.GetNetworkInterfacesFunc is nil but DockerClient.GetNetworkInterfaces was just called")
	}
	callInfo := struct {
		ContainerId string
	}{
		ContainerId: containerId,
	}
	lockDockerClientMockGetNetworkInterfaces.Lock()
	mock.calls.GetNetworkInterfaces = append(mock.calls.GetNetworkInterfaces, callInfo)
	lockDockerClientMockGetNetworkInterfaces.Unlock()
	return mock.GetNetworkInterfacesFunc(containerId)
}

// GetNetworkInterfacesCalls gets all the calls that were made to GetNetworkInterfaces.
// Check the length with:
//     len(mockedDockerClient.GetNetworkInterfacesCalls())
func (mock *DockerClientMock) GetNetworkInterfacesCalls() []struct {
	ContainerId string
} {
	var calls []struct {
		ContainerId string
	}
	lockDockerClientMockGetNetworkInterfaces.RLock()
	calls = mock.calls.GetNetworkInterfaces
	lockDockerClientMockGetNetworkInterfaces.RUnlock()
	return calls
}

// PullImage calls PullImageFunc.
func (mock *DockerClientMock) PullImage(imageName string) error {
	if mock.PullImageFunc == nil {
		panic("moq: DockerClientMock.PullImageFunc is nil but DockerClient.PullImage was just called")
	}
	callInfo := struct {
		ImageName string
	}{
		ImageName: imageName,
	}
	lockDockerClientMockPullImage.Lock()
	mock.calls.PullImage = append(mock.calls.PullImage, callInfo)
	lockDockerClientMockPullImage.Unlock()
	return mock.PullImageFunc(imageName)
}

// PullImageCalls gets all the calls that were made to PullImage.
// Check the length with:
//     len(mockedDockerClient.PullImageCalls())
func (mock *DockerClientMock) PullImageCalls() []struct {
	ImageName string
} {
	var calls []struct {
		ImageName string
	}
	lockDockerClientMockPullImage.RLock()
	calls = mock.calls.PullImage
	lockDockerClientMockPullImage.RUnlock()
	return calls
}

var (
	lockSshClientMockClose        sync.RWMutex
	lockSshClientMockCreateTunnel sync.RWMutex
)

// SshClientMock is a mock implementation of SshClient.
//
//     func TestSomethingThatUsesSshClient(t *testing.T) {
//
//         // make and configure a mocked SshClient
//         mockedSshClient := &SshClientMock{
//             CloseFunc: func() error {
// 	               panic("TODO: mock out the Close method")
//             },
//             CreateTunnelFunc: func(remoteSpec string) (SshTunnelLocalPort, error) {
// 	               panic("TODO: mock out the CreateTunnel method")
//             },
//         }
//
//         // TODO: use mockedSshClient in code that requires SshClient
//         //       and then make assertions.
//
//     }
type SshClientMock struct {
	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// CreateTunnelFunc mocks the CreateTunnel method.
	CreateTunnelFunc func(remoteSpec string) (SshTunnelLocalPort, error)

	// calls tracks calls to the methods.
	calls struct {
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// CreateTunnel holds details about calls to the CreateTunnel method.
		CreateTunnel []struct {
			// RemoteSpec is the remoteSpec argument value.
			RemoteSpec string
		}
	}
}

// Close calls CloseFunc.
func (mock *SshClientMock) Close() error {
	if mock.CloseFunc == nil {
		panic("moq: SshClientMock.CloseFunc is nil but SshClient.Close was just called")
	}
	callInfo := struct {
	}{}
	lockSshClientMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockSshClientMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedSshClient.CloseCalls())
func (mock *SshClientMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockSshClientMockClose.RLock()
	calls = mock.calls.Close
	lockSshClientMockClose.RUnlock()
	return calls
}

// CreateTunnel calls CreateTunnelFunc.
func (mock *SshClientMock) CreateTunnel(remoteSpec string) (SshTunnelLocalPort, error) {
	if mock.CreateTunnelFunc == nil {
		panic("moq: SshClientMock.CreateTunnelFunc is nil but SshClient.CreateTunnel was just called")
	}
	callInfo := struct {
		RemoteSpec string
	}{
		RemoteSpec: remoteSpec,
	}
	lockSshClientMockCreateTunnel.Lock()
	mock.calls.CreateTunnel = append(mock.calls.CreateTunnel, callInfo)
	lockSshClientMockCreateTunnel.Unlock()
	return mock.CreateTunnelFunc(remoteSpec)
}

// CreateTunnelCalls gets all the calls that were made to CreateTunnel.
// Check the length with:
//     len(mockedSshClient.CreateTunnelCalls())
func (mock *SshClientMock) CreateTunnelCalls() []struct {
	RemoteSpec string
} {
	var calls []struct {
		RemoteSpec string
	}
	lockSshClientMockCreateTunnel.RLock()
	calls = mock.calls.CreateTunnel
	lockSshClientMockCreateTunnel.RUnlock()
	return calls
}
