package network

import (
	"bytes"
	"html/template"
)

const (
	CALICO           = "calico"
	FLANNEL          = "flannel"
	defaultInterface = "eth.*|em.*"
	defaultPodCIDR   = "100.64.0.0/10"
)

type MetaData struct {
	Interface string
	PodCIDR   string
}

// Network is CNI interface
type Network interface {
	// if template is "" using default template
	Manifests(template string) string
	// return cni template file
	Template() string
}

func NewNetwork(t string, metadata MetaData) Network {
	switch t {
	case CALICO:
		return &Calico{metadata: metadata}
	case FLANNEL:
		return &Flannel{metadata: metadata}
	default:
		return &Calico{metadata: metadata}
	}
}

func render(data MetaData, temp string) string {
	var b bytes.Buffer
	t := template.Must(template.New("net").Parse(temp))
	t.Execute(&b, &data)
	return b.String()
}
