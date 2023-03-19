# Docker Compose Reference

Docker Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application's services. Then, with a single command, you create and start all the services from your configuration.

Using Compose is basically a three-step process.

- Define your app's environment with a Dockerfile so it can be reproduced anywhere.

- Define the services that make up your app in docker-compose.yml so they can be run together in an isolated environment.

- Run docker-compose up and Compose starts and runs your entire app.

<!-- ## Installation and documentation

## Reference -->

## Utility References

### **Debugging**

```yaml
entrypoint: ["sh", "-c", "sleep 2073600"]
command: ["sh", "-c", "sleep 2073600"]
```

this is the command that will be executed when the container is started. It is the same as the command parameter in docker run.

### **Build**

```yaml
build:
    context: ${SERVER_DIRECTORY} # build from another directory
    dockerfile: docker/dev/Dockerfile
```

- context: The path to the directory containing the Dockerfile. This path is relative to the location of the Compose file. If you specify a path to a file, Compose uses the directory that contains the file as the build context.
- dockerfile: The name of the Dockerfile (Default is path/Dockerfile).
