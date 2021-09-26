# golang_tp

Inicialmente se crea el archivo main y el mod mediante comandos. Luego se define a utilizar una carpeta services donde estará alojado el archivo string y string_test, siendo estos el "cerebro" de la aplicación y el main el "corazón".
En el archivo string.go se define la estructura de tipo Chain, además de realizarse los respectivos import, y en la función ValidChain se recibirá una cadena de texto y se devolverá un booleano junto con un error para especificar la validación.
En esta función se analizan posibles errores y lo que se debería devolver en cada caso, también se usa la funcion "strconv.Atoi" para convertir texto a número, que en caso de fallar devuelve un error.
En la función ModChain, se recibe una cadena de texto(ya validada) y un booleano, mientras que esta devuelve un puntero a Chain o error. En caso de que el booleano este en true se asignan los valores al puntero a Chain (r) y se devuelve el "objeto" mencionado con anterioridad.
Caso contrario y que el boolean llegue en false tras la validación, se devuelve la estructura de la siguiente manera: {"","", 0}, para poder utilizarla en los test.
Luego de realizar algunos ajustes se logró llegar al 80% de cobertura, lo cual es suficiente para analizar posibles errores en la función principal (ValidChain).

