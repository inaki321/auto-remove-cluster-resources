package main

import (
	"log";
	"os";
	"os/signal";
	"syscall";
	"time";

	"autocluster/k8s";
	"autocluster/server";
)

func main() {
	log.Println("--- Autocluster Application Starting ---");

	// 1. Initialize K8s ClientManager
	k8sManager, err := k8s.NewClientManager(); // main cluster manager to share throughout the app files
	if err != nil {
		log.Fatalf("Fatal: Could not initialize Kubernetes client: %v", err);
	}

	go func() {

		err := k8sManager.PollCluster(time.Second * 5); 
		if err != nil {
			log.Fatalf("Fatal: Cluster polling terminated with a critical error: %v", err);
		}
	}();

	// 3. Start the HTTP server in its own Go routine
	go func() {
		log.Println("--- Autocluster HTTP Service Starting (in background) ---");

		if err := server.StartServer(); err != nil {
			log.Fatalf("Fatal error starting HTTP server: %v", err);
		}
	}();

	log.Println("Core logic started successfully.");

	// 4. Graceful Shutdown (Wait for OS signal)
	quit := make(chan os.Signal, 1);
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM);
	<-quit;

	log.Println("--- Application Shutting Down Gracefully ---");
	log.Println("autocluster application finished.");
}