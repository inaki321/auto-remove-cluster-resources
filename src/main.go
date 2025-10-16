package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"autocluster/k8s"
	"autocluster/server"
)

func main() {
	log.Println("--- Autocluster Application Starting ---")

	k8sManager, err := k8s.NewClientManager(); // main cluster manager to share throughout the app files
	if err != nil {
		log.Fatalf("Fatal: Could not initialize Kubernetes client: %v", err);
	}

	if err := k8sManager.ListNamespaces(); err != nil {
		log.Printf("Warning: Failed to list namespaces: %v", err);
	}

	go func() {
		log.Println("--- Autocluster HTTP Service Starting (in background) ---");

		if err := server.StartServer(); err != nil {
			log.Fatalf("Fatal error starting HTTP server: %v", err);
		}
	}()

	log.Println("Core logic started successfully.");

	// 4. Graceful Shutdown (Wait for OS signal)
	quit := make(chan os.Signal, 1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	<-quit

	log.Println("--- Application Shutting Down Gracefully ---");
	log.Println("autocluster application finished.");
}
