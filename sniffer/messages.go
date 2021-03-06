package sniffer

type connSpec struct {
	dockerEndpoint string
	sshTunnelSpec  string
}

func (cs connSpec) getConnSpec() connSpec {
	return cs
}

type ConnectionRequest interface {
	getConnSpec() connSpec
}

func DirectConnectionRequest(dockerEndpoint string) ConnectionRequest {
	return connSpec{
		dockerEndpoint: dockerEndpoint,
	}
}

func TunneledConnectionRequest(dockerEndpointOnRemoteHost string, sshSpec string) ConnectionRequest {
	return connSpec{
		dockerEndpoint: dockerEndpointOnRemoteHost,
		sshTunnelSpec:  sshSpec,
	}
}

// TODO this is not needed, remove
type ConnectionResponse int

const (
	Connected ConnectionResponse = iota
	Error
)

type SshPassword string

type Container struct {
	Name string
	Id   string
}

type NetworkInterface struct {
	Name string
}
