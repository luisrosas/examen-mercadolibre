# Examen Mercadolibre

Desarrollo de un API para detectar si un ADN es mutante o no

### Stack tecnologíco
* Go (v1.17 o superior)
* Mysql

### Comandos
Para instalar el proyecto se debe ejecutar el comando:
```
make install
```
Solo es necesario si este se va a ejecutar en la maquina local.

Para compilar el proyecto se debe ejecutar el comando:
```
make build
```
Esto generará una carpeta llamada `build` en la raíz del mismo, la cal contendrá el archivo ejecutable.

Para ejecutar los tests se debe ejecutar el comando:

```
make tests
```

Para ejecutar los tests incluido el reporte de cobertura de los mismo se debe ejecutar el comando:
```
make coverage
```
Esto generará una carpeta llamada `coverage` en la raíz del proyecto, la cual contendrá el archivo `coverage.html` que se podrá abrir en el navegador para ver la covertura de los mismos.
