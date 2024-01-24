# [package] ServiceRunner

The *ServiceRunner* package provides all the essentials you need to develop and deploy your car module. Its features include:

- Registering your service with the System Manager and resolving all dependencies of your service, as can be defined in a *service.yaml* file
- Fetching the initial tuning state, and listening for updates to tuning state
- Setting up logging using Zerolog and adding flags for log verbosity
- Automatic (exponentially backed off) restarts of the service when it might fail

## Getting started

First, add the service runner package to your Go module:

```bash
go get github.com/VU-ASE/pkg-ServiceRunner
```

> [!TIP]
> You can install a specific version like this: `github.com/VU-ASE/pkg-ServiceRunner@v0.1.2` to improve reproducibility

Then, in your `main()` function, call the `servicerunner.Run` function. This function needs two parameters:

- Your "entrypoint" function to run. This entrypoint function should take as input a `TuningState` argument: this is the initial tuning state, fetched from the system manager. It should also have return type error
- A tuning state callback function. This function should take as input a `TuningState` argument and is called whenever a new tuning state is available

### Example: basic *main.go*
```go
package main

import (
	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/include"
	servicerunner "github.com/VU-ASE/pkg-ServiceRunner/src"
)

// Your entrypoint function
func run(service servicerunner.ResolvedService, initialTuningState *protobuf_msgs.TuningState) error {
	// all your application code can go here!
}

// Tuning state callback
func onNewTuningstate(ts *protobuf_msgs.TuningState) {
    // update your tuning!
}

// Let the servicerunner take care of everything
func main() {
	servicerunner.Run(run, onNewTuningState)
}
```

### Reaching the System Manager

As said, the ServiceRunner will register your service from a *service.yaml* file. This file should be in the root folder and should be pushed to git. In order to register the
service, the ServiceRunner needs to know how to reach the SystemManager module. This is done by reading from environment variables. You can set the environment variables as follows:

```bash
# set the SystemManager req/res endpoint
export ASE_SYSMAN_SERVER_ADDRESS=tcp://localhost:1337
# set the SystemManager pub/sub (broadcasting) endpoint
export ASE_SYSMAN_BROADCAST_ADDRESS=tcp://localhost:1338
```

### Registering your service

The ServiceRunner will register your service from a *service.yaml* file. This file should be placed in the root folder and should be pushed to git. The basic file structure is shown below:

```yaml
# Service definitions
name: serviceA
description: This is my basic service

# Define dependencies that our service depends on
dependencies:
  - service: serviceB
    output: outputA

# Define the outputs that our service exposes
outputs:
  - name: outputC
    address: tcp://*:5555
```

The servicerunner will parse this yaml file and will try to reach the SystemManager. It will register its own service and the outputs it exposes. If a service with this name is already registered and still running, the ServiceRunner will stop and will not retry.

After registering, the ServiceRunner will try to resolve all dependencies by contacting the ServiceManager. If a dependency service has not started yet, the ServiceRunner will back off and autmatically retry later. The ServiceRunner will not execute your service until all dependencies have been resolved.

### Example: using resolved dependencies and outputs

In your entrypoint function, the first argument must be of type `servicerunner.ResolvedService`. This interface provides some handy utilities to fetch the addresses of 
dependencies and outputs defined in your *service.yaml* file. Suppose your service is defined as:

```yaml
# Service definitions
name: serviceA
description: This is my basic service

# Define dependencies that our service depends on
dependencies:
  - service: serviceB
    output: outputA

# Define the outputs that our service exposes
outputs:
  - name: outputC
    address: tcp://*:5555
```

Then in your entrypoint function, you access both the address of your own output (*"outputC"*) as well as the address of your dependency (*"serviceB.outputA"*). You can do that as follows:

```go
package main

import (
	protobuf_msgs "github.com/VU-ASE/pkg-ServiceRunner/include"
	servicerunner "github.com/VU-ASE/pkg-ServiceRunner/src"
)

// Your entrypoint function
func run(service servicerunner.ResolvedService, initialTuningState *protobuf_msgs.TuningState) error {
    // Get your own output address, notice that outputC corresponds with the service.yaml outputs
	outputCaddr, err := service.GetOutputAddress("outputC")
	if err != nil {
		return err
	}

    // Get your dependency address, notice that serviceB and outputA correspond with service.yaml dependencies
    dependencyAddr, err := service.GetDependencyAddress("serviceB", "outputA")
    if err != nil {
        return err
    }
}

// Tuning state callback
func onNewTuningstate(ts *protobuf_msgs.TuningState) {
    // update your tuning!
}

// Let the servicerunner take care of everything
func main() {
	servicerunner.Run(run, onNewTuningState)
}
```

## Options and configuration

The ServiceRunner exposes the following flags that can be used to run your binary as you see fit:

| Flag                   | Default      | Description                                                                                  |
|------------------------|--------------|----------------------------------------------------------------------------------------------|
| `-debug`               | false        | Enable debug level logging                                                                   |
| `-output`              | stdout       | Path of the output file to log to                                                            |
| `-retries`             | 0            | Number of restart attempts after your service entrypoint function returns an error           |
| `-service-yaml`        | service.yaml | Path to the service.yaml file defining your service                                          |
| `-disable-live-tuning` | false        | Disable live tuning (your tuningstate callback will not be called when this flag is enabled) |

## Important

**Service name `systemmanager` is a reserved name**
When specifying your service name as `systemmanager` in *service.yaml*, the ServiceRunner will not attempt service registration, dependency resolving and tuning state fetching. (As the system manager cannot register with itself).
