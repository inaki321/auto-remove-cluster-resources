package k8s

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// list cluster namespaces
func (cm *ClientManager) ListNamespaces() error {
	fmt.Println("Listing Namespaces...")

	// cm.Clientset is available here
	namespaces, err := cm.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("error listing namespaces: %w", err)
	}

	for _, ns := range namespaces.Items {
		fmt.Printf("-> %s\n", ns.Name)
	}
	fmt.Print("--------- \n")

	return nil
}

// DeletePod
func (cm *ClientManager) DeletePod(namespace, podName string) error {
	fmt.Printf("Attempting to delete pod '%s' in namespace '%s'...\n", podName, namespace)

	// cm.Clientset is available here
	err := cm.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete pod %s/%s: %w", namespace, podName, err)
	}

	fmt.Printf("Successfully triggered deletion of pod '%s'.\n", podName)
	return nil
}
