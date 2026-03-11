# Como usar

Usar la extension REST Client de Visual Studio para poder correr las peticiones desde el archivo api_test.http

Luego dentro del archivo vas a tener que loguearte como cliente o admin para poder usar los distintos metodos. Estos estan separados por Categorias segun que usaurio debes usar para loguearte.

Una ves te logueas debes poner el token que te devuelve el servidor en la variable correspondiente.

# Variables
post_id: se usa para buscar una publicacion especifica o borrarla

delete_id: se usa para borrarla una publicacion especifica

admin_token: debe completarse con el token devuelto por el pedido de logueo como admin

user_token: debe completarse con el token devuelto por el pedido de logueo como user