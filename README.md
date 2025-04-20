# Erase una vez 4

Aplicación Golang utilizada en los ejercicios del libro [Érase una vez Kubernetes](https://leanpub.com/erase-una-vez-kubernetes).

Traduction: [Inglés](README_en.md)

## Descripción

La aplicación consulta el número de Pods existentes en el Namespace e imprime un mensaje con esta información en la consola.

Es un ejemplo sencillo utilizado en el capítulo *Role Base Access Control* para mostrar cómo asignarle permisos a un Pod.

## Variables de entorno

El funcionamiento de la aplicación puede ser modificado a través de variables de entorno:

|Variable de entorno|Descripción|Valor por defecto|
|-------------------|-----------|-----------------|
|`NAMESPACE` | Modifica el namespace donde serán consulados los Pods.      | `default` |
|`SLEEP_TIME`| Modifica el intervalo de tiempo entre mensajes. En segundos.| 5 |

## Despliegue en Kubernetes

Crear namespace `admin`.

```sh
kubectl apply -f kubernetes/namespace.yaml

namespace/admin created
```

Crear los objetos ServiceAccount, Role, RoleBinding y Deployment.

```sh
kubectl apply -f kubernetes/

deployment.apps/app created
namespace/admin unchanged
role.rbac.authorization.k8s.io/developer created
rolebinding.rbac.authorization.k8s.io/developer created
serviceaccount/developer created
```

Consulte los logs de la App desplegada.

```sh
kubectl -n admin logs -l app=rbac --follow

Existen 0 pods en el namespace default
Existen 0 pods en el namespace default
Existen 0 pods en el namespace default
```
