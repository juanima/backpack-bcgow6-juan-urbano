#  Práctica clase 1 - Go Testing

## Ejercicio 1
*¿Cuáles son las diferencias entre White Box y Black Box?*

| White Box | Black Box |
| --------- | --------- |
| Diseño o funcionamiento interno, es desconocida por quien ejecuta la prueba | quien ejecuta la prueba conoce y tiene visibilidad sobre el código.|
| las pruebas se ejecuta desde la perspectiva del usuario final | las pruebas no se ejecutan desde una perspectiva de usuario final. |
| probar la funcionalidad del código, evaluando las respuestas y reacciones (comportamiento) | probar el flujo, seguridad y estructura del código  |
| Se aplica para:<br> * Test funcionales<br> * Test No funcionales<br> * Test de regresion | Se aplica para:<br> * Test Unitarios<br> * Test de integración |

## Ejercicio 2
*¿Qué es un test funcional?*

* Son las pruebas que se basan en la entrada y salida del software, con el objetivo final de comprobar que la respuesta del software ante cada escenario, coincida exactamente con el resultado esperado.
* Son test que validan el comportamiento funcional del software.

## Ejercicio 3
*¿Qué es un Test de Integración?*

* Son aquellos test que prueban la comunicación entre distintos componentes o capas de la aplicación.

> El objetivo es comprobar que todos aquellos bloques de código que fueron probados de forma unitaria, interactúen y se comuniquen entre sí generando los resultados esperados.

## Ejercicio 4
Indicar las dimensiones de calidad prioritarias en MELI.

Tomados apartir del estandar *_ISO/IEC 25010_* se definen 8 dimensiones de los cuales para MeLi son relevantes y no resta mayor importancia a los demás

* Funcionalidad (_Functionality_)
        * Funcionalidad de nuestro sistema, respondiendo a ¿cómo funciona?, ¿cómo está hecho?, ¿cómo se apropia de los requerimientos?

* Rendimiento (_Performance_)
        * Capacidad para utilizar los recursos, respondiendo a ¿cómo se comporta a traves del tiempo?

* Fiabilidad (_Reliability_)
        * La madurez con la que nuestro sistema cuenta, la capacidad que tiene nuestro sistema de recuperarse ante diversos errores, tolerancia a fallos ante ambientes ostiles

* Seguridad (_Security_)
        * Un sistema tiene que ser integral

* Mantenibilidad (_Maintainability_)
        * Enfocado a la parte del código, los sistemas sufren cambios constantemente se tiene que adaptar al entorno para no perder su identidad
        * En medio de los `refractors` y los equipos trabajando sobre el mismo código alteran los estado del programa, el sistema debe mantener la caracteristica de la mantenibilidad
        * Nuestro sistema trabajarse de manera modular, con una resposabilidad única por modulo facilitando la reusabilidad y el mantenimientok
        * Los sistemas están sujetos a recibir cambios externos y nuestro sistema debe ser capaz de recibirlos sin romperse

