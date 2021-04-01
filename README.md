# Atomic
This is a simple POC that shows how build an polyglot microservices based infrastructure with docker and kubernetes.

The name **atomic** stands for: little apps (atoms) working together without matters his language. And follows the UNIX philosophy: "
The Unix philosophy emphasizes building simple, short, clear, modular, and extensible code that can be easily maintained and repurposed by developers other than its creators. The Unix philosophy favors composability as opposed to monolithic design." according [wikipedia](https://en.wikipedia.org/wiki/Unix) or "do one thing and do it well"

## Requirements
For linux you need:
 - docker
 - docker-compose
 - docker-machine
 - minikube
 - kubctl

## Run
### Docker
After install minikube and runing execute:
```
$ eval $(minikube docker-env)
```
This will be bind the docker environment from minikube to your host

Then you got to the root of this project and run:

```
$ docker-compose build
```
This will compile and install the images from the project into de minikube VM.

### Minikube
Now we will install [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) and the [Services](https://kubernetes.io/docs/concepts/services-networking/service/) to minikube.

We going to run:
```
$ minikube addons enable ingress
```
This gonna enable the Ingress plugin

Then we gonna up the services into minikube:
```
$ kubectl apply -f services.yaml
```

And finally configure Ingress to able the services:
```
$ kubectl apply -f ingress.yaml
```

Display atomic ingress information
```
$ kubectl describe ingress atomic-ingress
```