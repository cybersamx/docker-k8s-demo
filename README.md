# docker-k8s-demo

This project contains sample code for a docker and kubernetes workshop I am conducting.

## Setup

### The Application

1. Download and install the [go compiler](https://go.dev/dl/).
1. Build and run the application.

   ```shell
   $ go run main.go
   ```

1. Test the application.

   ```shell
   $ curl localhost:8080
   ```

#### Using Docker Compose

1. Run the application using docker compose.

   ```shell
   $ docker-compose up
   ```
   
1. To tear down the container.

   ```shell
   $ docker-compose down
   ```

### Docker Container

1. Download and install [Rancher Desktop](https://rancherdesktop.io/).
1. Launch Rancher Desktop.
1. Build a Docker image of the app.

   ```shell
   $ docker build -t hello .
   $ docker image ls  # We should see the hello image
   ```

1. Run the app as a container.

   ```shell
   $ docker run -d -p 8080:8080 --name hello hello  # Run the container
   $ docker ps  # List what docker containers are running locally
   $ docker stop hello
   $ docker rm hello
   ```

### Push the Image to a Container Registry

1. If you haven't done so, register an account on <https://hub.docker.com> and sign in.
1. Create a repository and name it `hello`. Let's assume that the username (or namespace) is `cybersamx`.

   ```shell
   $ docker login --username=cybersamx
   ```

1. Tag the image to match the repository you created on docker hub.

   ```shell
   # docker image tag <local_image_name> <remote_image_name>
   $ docker image tag hello cybersamx/hello
   ```

1. Push the image to a (public) container registry.

   ```shell
   $ docker image push cybersamx/hello
   ```

### Run the App in Kubernetes (k8s)

1. Install [kubectl](https://kubernetes.io/docs/tasks/tools/).
1. Launch Rancher Desktop, select Preferences, check `Enable Kubernetes`.
1. Switch k8s context. **Note: Always a good idea to check your k8s context. You don't want to alter anything while still connected to a production k8s cluster.**

   ```shell
   $ kubectl config use-context rancher-desktop
   $ kubectl config current-context
   ```

1. Deploy the app.

   ```shell
   $ kubectl apply -f deploy.yaml
   $ kubectl get deploy
   NAME    READY   UP-TO-DATE   AVAILABLE   AGE
   hello   3/3     3            3           8s
   ```

   Run `curl localhost:8080` but we will get `Connection refused`. This is because we have not exposed the tcp port.

1. Run the service.

   ```shell
   $ kubectl apply -f service.yaml
   $ curl "$(kubectl get svc hello-service -o json | jq -r '.status.loadBalancer.ingress[0].ip'):8080"
   ```

## Reference

* [Docker: Docs](https://docs.docker.com/)
* [Kubernetes: Docs](https://kubernetes.io/docs/home/)
