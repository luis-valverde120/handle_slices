# Manejo de Arreglo Unidimensional

Este programa permite realizar diversas operaciones sobre un arreglo unidimensional, como insertar, eliminar, buscar y ordenar elementos. Está diseñado para ejecutarse en la terminal y ofrece un menú interactivo para que el usuario seleccione las opciones deseadas.

## Funcionalidades

El programa incluye las siguientes funcionalidades:

1. **Ingresar un nuevo valor**: Permite agregar un nuevo valor al arreglo.
2. **Eliminar un valor**: Elimina un valor del arreglo según el índice especificado.
3. **Búsqueda binaria**: Busca un valor en el arreglo utilizando el algoritmo de búsqueda binaria (requiere que el arreglo esté ordenado).
4. **Búsqueda lineal**: Busca un valor en el arreglo mediante una búsqueda lineal.
5. **Ordenar arreglo**: Ordena el arreglo utilizando el algoritmo de ordenamiento por mezcla (merge sort).
6. **Imprimir arreglo**: Muestra los elementos actuales del arreglo.
7. **Salir**: Finaliza la ejecución del programa.

## Requisitos

- Go 1.21.1 o superior.

## Estructura del Proyecto

El proyecto está compuesto por los siguientes archivos:

- `main.go`: Contiene la lógica principal del programa y el menú interactivo.
- `sort.go`: Implementa las funciones de ordenamiento (`merge` y `sort`).
- `search.go`: Implementa las funciones de búsqueda (`binarySearch` y `linearSearch`).
- `go.mod`: Archivo de configuración del módulo Go.
- `.gitignore`: Define los archivos que deben ser ignorados por Git.

## Cómo Ejecutar el Programa

1. Clona este repositorio en tu máquina local:
   ```bash
   git clone https://github.com/actividad_8.git
   cd actividad8