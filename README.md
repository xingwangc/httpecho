httpecho is a simple http service developed by Go, which just echo the URL you requested and can be used to test the network.

The simple application is provided for some network testing. Here I used it test if the network of k8s
is working rightly, include service dns, the port exposed, etc. So here also provide a simple Dockerfile to
build the image into a docker image. And a k8s configuration files pair(rc/svc fil) also provided.

# Build application

since the base image of Dockefile is using centos7, so you should build the application using linux env

	GOOS=linux GOARCH=amd64 go build .

The `httpecho` file will be generated in current directory.

## Test you application

To start the application, you should specify a port to listen, default will use 8080 when not provided. Also if
the port specified is not in the range of [1024:49151), it will use 8080 too.

	./httpecho 8000
	curl localhost:8000/hello

It will reponse:

	Hello, "/hello"

# Build docker image

Build and tag the docker image with your repository. (I set up a public repo xingwangc/httpecho in docker hub, you can directly pull from there)

	docker build -t xingwangc/httpecho .

Then push your image to your repository.

	docker push xingwangc/httpecho

## Test your image

	docker run -d -p 8000:8000 xingwangc/httpecho ./httpecho 8000

After the container is up, test is same with application.

# To verify the network of k8s cluster

Copy the rc/svc configuration files pair into your k8s cluster. And then run `kubectl create -f filename` to start the replication controller and service(you need to change the image to you private repo first, if you are using self build image or do not want to use mine).

## Verify the node port to expose outside the cluster

I use the port: *30000*(you can use anyone in 30000~32767) to expose the service outside the cluster. So choose one of any your minion nodes ip to check if the node port is work.

	curl *any minion nodes's IP*:30000/hello

in my case is `curl 192.168.1.220:30000/hello`

will response:

	Hello, "/hello"

## Use the domain name to verify the cluster dns

As defined in the svc yaml file, the name of service is "httpecho-svc", so the domain name of the service is also "httpecho-svc"(full path is "httpecho-svc.default.svc.cluster.local")

So use the domain name to access service in the pod of cluster should also work if the dns of cluster is working.

	# kubectl exec -it httpecho-3lpbc curl httpecho-svc:8000/test

should response:

	Hello, "/hello"
