package extender

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

const Label = "priority.lixueduan.com"

type Extender struct {
	ClientSet kubernetes.Interface
}

func NewExtender() *Extender {
	clientset, err := NewClient()
	if err != nil {
		log.Fatalf("failed to create k8s clientset: %v", err)
	}

	return &Extender{
		ClientSet: clientset,
	}
}

// NewClient connects to an API server.
func NewClient() (kubernetes.Interface, error) {
	kubeConfig := os.Getenv("KUBECONFIG")
	if kubeConfig == "" {
		kubeConfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return nil, err
		}
	}
	client, err := kubernetes.NewForConfig(config)
	return client, err
}
