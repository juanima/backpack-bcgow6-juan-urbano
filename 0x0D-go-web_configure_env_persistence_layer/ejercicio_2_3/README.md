# Ejercicio 2 y 3

## Enunciated *_Guardar Información_*

Se debe implementar la funcionalidad para guardar la información de la petición en un archivo json, para eso se deben realizar los siguientes pasos:

1. En lugar de guardar los valores de nuestra entidad en memoria, se debe crear un archivo; los valores que se vayan agregando se guardan en él.

# Ejercicio 2

## Enunciated *_Leer Información_*

Se debe implementar la funcionalidad para leer la información requerida en la petición del archivo json generado al momento de guardar, para eso se deben realizar los siguientes pasos:

1. En lugar de leer los valores de nuestra entidad en memoria, se debe obtener del archivo generado en el punto anterior.


## Run 

- Estructura de directorios

```bash
$ tree
.
├── README.md
├── cmd
│   └── server
│       ├── handler
│       │   └── handler.go
│       └── main.go
├── create_products.json
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   └── product.go
│   └── products
│       ├── repository.go
│       └── service.go
├── patch_products.json
├── pkg
│   └── store
│       └── file.go
├── products.json
└── update_products.json

8 directories, 13 files
```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

*Window 1*

> Correr el servidor

```bash
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /filter-product           --> main.FiltrarProductsHandler (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000

...

```

*Window 0*

*  Asegurate de que en tu navegador esté corriendo

```bash
$ curl --write-out "%{http_code}\n" --silent --output /dev/null http://127.0.0.1:8080/status
200
```
>> Este endpoint `/test` está diseñado para indicarte el estado de funcionamiento del servidor web


*Window 2*


## Access to our application resources

> Consumir los servicios web usando el `curl` client http

## *Create products*

* Endpoint para crear un producto `/products/`

* Inicialmente el archivo `products.json` está vacío

```bash
$ cat products.json
[]
$
```

* Preparamos un archivo `create_products.json` con la información de nuestro producto 

* Input: Corresponde a los datos necesarios para crear un producto

```json
{
    "id": 1,
    "stock": 2,
    "name": "beer",
    "color": "blue",
    "code": "AC323",
    "price": 4.5,
    "publish": false,
    "date": "December"
}
```

* Hacemos test 

```bash
        curl curl http://127.0.0.1:8080/products/  \
        -X POST \
        -H "Content-Type: application/json" \
        -H "token: 123456" \
        -d @create_products.json \
        -vvv \
        ; echo "" | cat -e \
```

* Output: Token de acceso, creado y almacenado en nuestra applicación

```json
{
    "Id":1,
    "Stock":2,
    "Name":"beer",
    "Color":"blue",
    "Code":"blue",
    "Price":4.5,
    "Publish":false,
    "Date":"December"
}
```

* Revisar el archivo `products.json`

```bash
$ cat products.json
[
  {
    "Id": 1,
    "Stock": 20,
    "Name": "beer",
    "Color": "red",
    "Code": "red",
    "Price": 4.7,
    "Publish": true,
    "Date": ""
  }
]
$
```

*Importante*

> El proceso se repite para los siguientes _endpoints_

* Endpoint para obtener todos los productos `/products/`
* Endpoint para actualizar un producto `/products/id`
* Endpoint para actualizar un campo un producto `/products/id`

---


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```


1. Filtremos products --> ´C1TT´
2. Get on products --> ´C1TT´
2. GetAll on products --> ´C1TT´
3. Validar campos --> ´C2TM´
	1. store()
4. Validar token --> ´C2TM´
	1. store()

5. Estructura proyecto:´internal´ and ´cmd´  --> ´C2TT´ 
6. Update, delete patch -->  ´C3TM´ 
7. Configuración ENV --> i´C3TT´
8. Leer y escribir la información del JSON ´C3TT´
9. Manejo de respuestas genericas, ´struct Response´ ´C4TM´
10. Documentación ´C4TM´

