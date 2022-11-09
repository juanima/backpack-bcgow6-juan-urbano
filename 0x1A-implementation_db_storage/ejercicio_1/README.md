# Test

## Ejercicio 1 - Implementar GetByName()
Desarrollar un método en el repositorio que permita hacer búsquedas de un producto por nombre. Para lograrlo se deberá:

-  Diseñar interfaz “Repository” en la que exista un método GetByName() que reciba por parámetro un string y retorna una estructura del tipo Product.

-  Implementar el método de forma que con el string recibido lo use para buscar en la DB por el campo “name”.



## Ejercicio 2 - Replicar Store()
Tomar el ejemplo visto en la clase y diseñar el método Store():
Puede tomar de ejemplo la definición del método Store visto en clase para incorporarlo en la interfaz.
Implementar el método Store.



## Ejercicio 3 - Ejecutar Store()
Diseñar un Test que permita ejecutar Store y comprobar la correcta inserción del registro en la tabla.

## Ejercicio 4 - Ejecutar GetByName()
Diseñar un Test que permita ejecutar GetByName y comprobar que retorne el registro buscado en caso de que exista. 

# Run

* Import database from sql script file

```bash
$ mysql -uroot -p < db/db.sql
Enter password:
$
```

* Run web server form our API

```bash
CO000PFW3GXDTQ9:ejercicio_1 juurbano$ go run cmd/server/main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/products/         --> github.com/bootcamp-go/storage/cmd/server/handler.(*Product).Store.func1 (3 handlers)
[GIN-debug] GET    /api/v1/products/         --> github.com/bootcamp-go/storage/cmd/server/handler.(*Product).GetByName.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080

. . .
```


