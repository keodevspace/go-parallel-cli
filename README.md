Go Parallel CLI 🚀
==================

A high-performance Command Line Interface (CLI) built in Go to demonstrate advanced concurrency patterns and efficient API consumption.

📌 Overview
-----------

This project was developed to explore the core strengths of the Go programming language, specifically its first-class support for concurrency. It implements a **Worker Pool pattern** to process data from a public API in parallel using synchronized Goroutines.

🛠 Technologies & Concepts
--------------------------

*   **Go 1.21+**: Leveraging the latest features of the language.
    
*   **Goroutines & Channels**: Implementing the "Fan-out" pattern for parallel processing.
    
*   **Sync WaitGroups**: Ensuring proper synchronization and lifecycle management of threads.
    
*   **Implicit Interfaces**: Decoupling the API client for better testability and maintenance.
    
*   **Pointers & Structs**: Optimized memory management and data modeling.
    
*   **Docker**: Containerized using **Multi-stage builds** to ensure a minimal footprint (~15MB image).
    

🏗 Project Structure
--------------------

Following the **Standard Go Project Layout**:

  go-parallel-cli/  
  ├── cmd/cli/          # Application entry point (main.go)  ├── internal/         # Private application code  
  │   ├── api/          # HTTP Client and Data Models  
  │   └── processor/    # Concurrency logic (Worker Pool)  
  └── Dockerfile        # Production-ready container 

🚀 How to Run
-------------

### Prerequisites

*   Go 1.21 or higher installed.
    

### Running Locally

1.  Bashgit clone https://github.com/keodevspce/go-parallel-cli.gitcd go-parallel-cli
    
2.  Bashgo run cmd/cli/main.go
    

### Running with Docker

1.  Bashdocker build -t go-parallel-cli .
    
2.  Bashdocker run go-parallel-cli
    

📝 Implementation Details
-------------------------

The CLI initializes **10 synchronized Workers**. These workers listen to a "Jobs" channel, fetch data from a public REST API, and signal completion through a sync.WaitGroup. This architecture prevents race conditions and ensures the application only exits once all tasks are processed.