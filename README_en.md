# Once Upon a Time Kubernetes

Golang application used in the exercises of the book [Once Upon a Time Kubernetes](https://leanpub.com/erase-una-vez-kubernetes).

Traducción: [Español](README.md)

## Description

The application queries the number of Pods existing in the Namespace and prints a message with this information on the console.

It is a simple example used in the chapter *Role Base Access Control* to show how to assign permissions to a Pod.

## Environment variables

The application can be modified through environment variables:

|Environment variable|Description|Default value|
|-------------------|-----------|-----------------|
|`NAMESPACE` | Modifies the namespace where the Pods will be queried.      | `default` |
|`SLEEP_TIME`| Modifies the time interval between messages. In seconds.| 5 |

## Deployment in Kubernetes

Create namespace `admin`.

```sh
kubectl apply -f kubernetes/namespace.yaml

namespace/admin created
```

Create the ServiceAccount, Role, RoleBinding and Deployment objects.

```sh
kubectl apply -f kubernetes/

deployment.apps/app created
namespace/admin unchanged
role.rbac.authorization.k8s.io/developer created
rolebinding.rbac.authorization.k8s.io/developer created
serviceaccount/developer created
```

Check the logs of the deployed App.

```sh
kubectl -n admin logs -l app=rbac --follow

Existen 0 pods en el namespace default
Existen 0 pods en el namespace default
Existen 0 pods en el namespace default
```
