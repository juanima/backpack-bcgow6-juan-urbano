# Ejercicio 2

## Enunciated *_Ecommerce_*

Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios. Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main del programa como en las funciones.

Se necesitan las estructuras:

- Usuario: Nombre, Apellido, Correo, Productos (array de productos).
- Producto: Nombre, precio, cantidad.

Se requieren las funciones:
- Nuevo producto: recibe nombre y precio, y retorna un producto.
- Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
- Borrar productos: recibe un usuario, borra los productos del usuario.


## Run 

- Estructura de directorios

```bash
$ ls -lR
total 40
-rw-r--r--  1 juurbano  1010544492  1886 Sep 27 16:47 README.md
-rw-r--r--  1 juurbano  1010544492  3215 Sep 29 11:12 ecommerce.go
-rw-r--r--  1 juurbano  1010544492    47 Sep 28 07:21 go.mod
-rw-r--r--  1 juurbano  1010544492   201 Sep 28 07:19 main.go
drwxr-xr-x  4 juurbano  1010544492   128 Sep 29 11:03 models
-rw-r--r--  1 juurbano  1010544492   259 Sep 28 16:38 testcase.txt

./models:
total 16
-rw-r--r--  1 juurbano  1010544492   925 Sep 29 11:02 Product.go
-rw-r--r--  1 juurbano  1010544492  1258 Sep 29 08:40 User.go

$
```

- test cases
    * la primera linea es el número de testcases
    * la segunda linea es el número de usuarios para registrar
    * las siguientes líneas correspondes a los datos de los productos 
	* `nameUser` `lastName` `email` 
    * la siguiente lineas corresponden número de productos a registrar
    * las siguientes líneas correspondes a los datos de los productos 
	* `nameProduct` `price` 
    * la siguiente linea corresponde a la información del usuario y producto que quiere agregar
	* `nameUser` `lastName` `nameProduct` `price` `amount`
    * la linea final info del usuario que queremos eliminar sus productos
	* `nameUser` `lastName`

```
$ cat testcase.txt
1
-----------------------1
2
Nicolas Salva nico29@gmail.com
2
detergente 30012.00
suavizante 3400.00
Andres Ronda nico29@gmail.com
1
guantes 50212.00
Nicolas Salva vaso 360.80 6
Andres Ronda

```

### Option 1

- Ejecutar los casos de acuerdo al archivo de pruebas

```bash
$ go run . < testcase.txt
-----------------------1
0x1400010c140
Info Product:
        Name: detergente,
        Price: 30012.000000,
        Amount: 1

Info Product:
        Name: suavizante,
        Price: 3400.000000,
        Amount: 1

Info User:
        Name: Nicolas,
        Last name: Salva,
        Email: nico29@gmail.com,
        Products:
0x1400010c190
Info Product:
        Name: guantes,
        Price: 50212.000000,
        Amount: 1

Info User:
        Name: Andres,
        Last name: Ronda,
        Email: nico29@gmail.com,
        Products:
------
0x1400010c1e0
Info Product:
        Name: detergente,
        Price: 30012.000000,
        Amount: 1

Info Product:
        Name: suavizante,
        Price: 3400.000000,
        Amount: 1

Info Product:
        Name: vaso,
        Price: 360.800000,
        Amount: 6

Info User:
        Name: Nicolas,
        Last name: Salva,
        Email: nico29@gmail.com,
        Products:
0x1400010c230
Info Product:
        Name: guantes,
        Price: 50212.000000,
        Amount: 1

Info User:
        Name: Andres,
        Last name: Ronda,
        Email: nico29@gmail.com,
        Products:
------
0x1400010c280
Info Product:
        Name: detergente,
        Price: 30012.000000,
        Amount: 1

Info Product:
        Name: suavizante,
        Price: 3400.000000,
        Amount: 1

Info Product:
        Name: vaso,
        Price: 360.800000,
        Amount: 6

Info User:
        Name: Nicolas,
        Last name: Salva,
        Email: nico29@gmail.com,
        Products:
0x1400010c2d0
Info User:
        Name: Andres,
        Last name: Ronda,
        Email: nico29@gmail.com,
        Products:

$
```


## Otros

Crear un modulo, fuera del GOPATH variable

```
$ go mod init ejercicio_2
go: creating new go.mod: module example.com/m
go: to add module requirements and sums:
	go mod tidy
$
```

