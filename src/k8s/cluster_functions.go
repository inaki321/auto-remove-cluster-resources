package k8s

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// list cluster namespaces
func (cm *ClientManager) ListNamespaces() error {
	fmt.Println("Listing Namespaces...");

	// cm.Clientset is available here
	namespaces, err := cm.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{});
	if err != nil {
		return fmt.Errorf("error listing namespaces: %w", err);
	}

	for _, ns := range namespaces.Items {
		fmt.Printf("-> %s\n", ns.Name);
	}
	fmt.Print("--------- \n");

	return nil;
}

// DeletePod
func (cm *ClientManager) DeletePod(namespace, podName string) error {
	fmt.Printf("Attempting to delete pod '%s' in namespace '%s'...\n", podName, namespace)

	// cm.Clientset is available here
	err := cm.Clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{});
	if err != nil {
		return fmt.Errorf("failed to delete pod %s/%s: %w", namespace, podName, err);
	}

	fmt.Printf("Successfully triggered deletion of pod '%s'.\n", podName);
	return nil;
}

func (cm *ClientManager) PollCluster(interval time.Duration) error {
	fmt.Printf("Starting continuous polling for cluster health (interval: %s)...\n", interval)

	for {
		// 1. Perform a simple, general check (e.g., list all namespaces).
		_, err := cm.Clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

		fmt.Println("--- Allowed Resources Configuration ---")
		for namespace, resources := range clusterResources {
			fmt.Printf("Namespace: %s allows the following resources: %v\n", namespace, resources)
			fmt.Printf("Checking which resources aren't used \n");
		}
		fmt.Println("---------------------------------------")

		// 2. Check for errors.
		if err != nil {
			fmt.Printf("Cluster check failed: API server unreachable or critical error: %v. Retrying in %s...\n", err, interval)

		} else {
			fmt.Println("Cluster check successful: API server is reachable.")
		}

		// 3. Wait for the next interval before re-checking.
		time.Sleep(interval)
	}
}
