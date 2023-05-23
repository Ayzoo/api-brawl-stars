# The Brawl Stars API

The Brawl Star API is a RESTful API based on the supercell video game Brawl Stars. You will access data about characters, images.

## Requisitos previos

- Go (v1.17 o superior)
- Fiber (v2)

## Instalación

1. Clona este repositorio: `git clone https://github.com/itsbeensolong/api-brawl-stars.git`
2. Navega hasta el directorio del proyecto: `cd api-brawl-star`
4. Configura las variables de entorno en el archivo `.env` según tus necesidades.
5. Inicia el servidor: `go run .`


### Endpoints

- `GET /users`: .
- `GET /users/:id`: .
- `POST /users`: .
- `PUT /users/:id`: .
- `DELETE /users/:id`: .

## Uso

Asegúrate de haber iniciado el servidor antes de realizar cualquier solicitud a la API. La API estará disponible en `http://localhost:3000`.

## Errores
La API puede devolver los siguientes códigos de estado y mensajes de error:

400 Bad Request: La solicitud es inválida.
401 Unauthorized: No se proporcionó o es inválido el token de acceso.
404 Not Found: No se encontró el recurso solicitado.
500 Internal Server Error: Error interno del servidor.
Contribución
Si deseas contribuir a esta API, siéntete libre de enviar solicitudes de extracción o informar problemas en el repositorio de GitHub: https://github.com/itsbeensolong/api-brawl-stars

## Licencia
Este proyecto está licenciado bajo la Licencia MIT.

## Contacto
Para consultas o más información, puedes comunicarte conmigo en johancs.mm@gmail.com
