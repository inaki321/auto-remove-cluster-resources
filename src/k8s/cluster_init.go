package k8s

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ClientManager holds the Kubernetes Clientset and methods to interact with the cluster.
// This struct is accessible by all files in the 'k8s' package.
type ClientManager struct {
	Clientset *kubernetes.Clientset
}

// NewClientManager initializes the Clientset and returns a pointer to the ClientManager struct.
func NewClientManager() (*ClientManager, error) {
	// --- 1. Load Kubeconfig and Flags ---
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubernetes config: %w", err)
	}

	// --- 2. Create the Clientset ---
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating clientset: %w", err)
	}

	fmt.Println("Successfully connected to the Kubernetes cluster.")

	// Return the new struct containing the initialized client
	return &ClientManager{
		Clientset: clientset,
	}, nil
}
