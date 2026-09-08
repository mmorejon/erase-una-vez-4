# erase-una-vez-4

[![English](https://img.shields.io/badge/Read_in-English-blue?style=flat-square)](README.en.md)

<div align="center">

<img src="https://github.com/mmorejon/erase-una-vez-k8s/blob/main/assets/book-cover.jpg" alt="Portada Libro Érase una vez Kubernetes" width="300"/>

Aplicación Golang utilizada en los ejercicios del libro Érase una vez Kubernetes.

👇 **Consigue la edición actualizada 2026 aquí:** 👇

[![Amazon](https://img.shields.io/badge/Amazon-Comprar_Ebook-orange?style=for-the-badge&logo=amazon)](https://www.amazon.es/dp/B0HJ338RNR)
[![Amazon](https://img.shields.io/badge/Amazon-Comprar_en_Tapa_Blanda-orange?style=for-the-badge&logo=amazon)](https://www.amazon.es/dp/B0HJ3WB7ZJ)
[![LeanPub](https://img.shields.io/badge/LeanPub-Descargar_Ebook-blue?style=for-the-badge&logo=leanpub)](https://leanpub.com/erase-una-vez-kubernetes)

</div>

---

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

---

## 🤝 Comunidad y Feedback

1.  ⭐ **¿Te ha sido útil?** Dale una **estrella** al repositorio (arriba a la derecha). Nos ayuda a llegar a más ingenieros.
2.  📚 **¿Aún no tienes el libro?** Compra el libro en Amazon o Leanpub.

<div align="center">
    <a href="https://www.amazon.es/dp/B0HJ338RNR">
        <img src="https://img.shields.io/badge/Amazon-Comprar_Ebook-orange?style=for-the-badge&logo=amazon" />
    </a>
    <a href="https://www.amazon.es/dp/B0HJ3WB7ZJ">
        <img src="https://img.shields.io/badge/Amazon-Comprar_Tapa_Blanda-orange?style=for-the-badge&logo=amazon" />
    </a>
</div>
