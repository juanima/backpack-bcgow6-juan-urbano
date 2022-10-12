# Ejercicio 1

## Enunciated *_Generar método PUT_*

Se solicita implementar una funcionalidad que modifique completamente una entidad. Para lograrlo, es necesario, seguir los siguientes pasos:

1. Generar un método PUT para modificar la entidad completa
2. Desde el Path enviar el ID de la entidad que se modificará
3. En caso de no existir, retornar un error 404
4. Realizar todas las validaciones (todos los campos son requeridos)


## Enunciated *_Generar método DELETE_*

Es necesario implementar una funcionalidad para eliminar una entidad. Para lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método DELETE para eliminar la entidad en base al ID
2. En caso de no existir, retornar un error 404

## Enunciated *_Generar método PATCH_*

Se requiere implementar una funcionalidad que modifique la entidad parcialmente, solo se deben modificar 2 campos:
- Si se seleccionó Productos, los campos nombre y precio.
- Si se seleccionó Usuarios, los campos apellido y edad.
- Si se seleccionó Transacciones, los campos código de transacción y monto.

Para lograrlo, es necesario, seguir los siguientes pasos:
1. Generar un método PATCH para modificar la entidad parcialmente, modificando solo 2 campo (a elección)
2. Desde el Path enviar el ID de la entidad que se modificara
3. En caso de no existir, retornar un error 404
4. Realizar las validaciones de los 2 campos a enviar

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
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   └── product.go
│   └── products
│       ├── repository.go
│       └── service.go
├── patch_products.json
├── products.json
├── script.sh
└── update_products.json

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
2022/10/05 20:18:05 1
2022/10/05 20:18:05 12
2022/10/05 20:18:05 shirt
2022/10/05 20:18:05 blue
2022/10/05 20:18:05 AC123
2022/10/05 20:18:05 12.4
2022/10/05 20:18:05 true
2022/10/05 20:18:05 December
{ID:1 Stock:12 Name:shirt Color:blue Code:AC123 Price:12.4 Publish:true Date:December}
[0x140001accc0]
[GIN] 2022/10/05 - 20:18:05 | 200 |     599.875µs |       127.0.0.1 | GET      "/filter-product?id=1&stock=12&name=shirt&color=blue&code=AC123&price=12.4&publish=true&date=December"

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

* Preparamos un archivo .json con la información de nuestro producto 

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
        -d @products.json \
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

---

## *GetAll products*

- Enpoint para acceder a la aplicación `/products/`
        
- Hacemos test 

```bash
        curl http://127.0.0.1:8080/products/ \
        -X GET \
        -H "token: 123456" \
        ; echo "" | cat -e \
```

* Output: Ya que no hay productos creados tendremos el siguiente output

```json
        [
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
        ]
```

---

## *GetOne product*

- Enpoint para acceder a la aplicación `/products/:id`
        
- Hacemos test 

```bash
        curl http://127.0.0.1:8080/products/1 \
        -X GET \
        -H "token: 123456" \
        ; echo "" | cat -e \
```

* Output: Ya que no hay productos creados tendremos el siguiente output

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
{
```

---
        
## *FilterProducto*

* Enpoint para acceder a la aplicación `/filter-product`
        * Datos del filtro `id=1&stock=2&name=beer&color=blue&code=blue&price=4.5&publish=false&date=December`
* Este enpoint requiere del token de acceso para ser consumido

* Hacemos test 

```bash
        curl "http://127.0.0.1:8080/products/filter-product?id=1&stock=2&name=beer&color=blue&code=blue&price=4.5&publish=false&date=December" 
        -X GET \
        -H "token: 123456" \
        -vvv \
        ; echo "" | cat -e \        
```

    * Output: Obtendremos los datos del usuario

```json
[
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
]
```

---

## *Update product*

* Enpoint para acceder a la aplicación `/products`

* Este enpoint requiere del token de acceso para ser consumido

* Input: Corresponde a los datos con los campos para ser actualizados 

```json
{
    "stock": 9,
    "name": "beer",
    "color": "red",
    "code": "AC888",
    "price": 4.7,
    "publish": true,
    "date": "January"
}
```

* Hacemos test 

```bash
        curl curl http://127.0.0.1:8080/products/1  \
        -X PUT \
        -H "Content-Type: application/json" \
        -H "token: 123456" \
        -d @update_products.json \
        -vvv \
```

    * Output: Obtendremos los datos del usuario

```json
{
    "Stock":9,
    "Name":"beer",
    "Color":"red",
    "Code":"blue",
    "Price":4.7,
    "Publish":true,
    "Date":"January"
}
```

---

## *Patch product*

* Enpoint para acceder a la aplicación `/products/id`

* Este enpoint requiere del token de acceso para ser consumido

* Input: Corresponde a los datos que vamos actualizar especificamente

```json
{
    "stock": 20,
}
```

* Hacemos test 

```bash

        curl curl http://127.0.0.1:8080/products/1  \
        -X PUT \
        -H "Content-Type: application/json" \
        -H "token: 123456" \
        -d @patch_products.json \
        ; echo "" | cat -e \        
```

* Output: Obtendremos los datos del producto 

```json
{
    "Id":1,
    "Stock":20,
    "Name":"beer",
    "Color":"blue",
    "Code":"blue",
    "Price":4.5,
    "Publish":false,
    "Date":"December"
}
```


---

## *Delete product*

* Enpoint para acceder a la aplicación `/products`

* Este enpoint requiere del token de acceso para ser consumido

* Hacemos test 

```bash
        curl "http://127.0.0.1:8080/products/1" 
        -X DELETE \
        -H "token: 123456" \
        -vvv \
        ; echo "" | cat -e \        
```

* Output: No tendremos respuesta de algun archivo json pero obtendremos un código de estado 204

```bash
*   Trying 127.0.0.1:8080...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> DELETE /products/1 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.79.1
> Accept: */*
> token: 123456
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 204 No Content
< Content-Type: application/json; charset=utf-8
< Date: Wed, 12 Oct 2022 22:17:27 GMT
<
* Connection #0 to host 127.0.0.1 left intact
$
```

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









2. 
