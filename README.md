# The Brawl Stars API

The Brawl Star API is a RESTful API based on the supercell video game Brawl Stars. You will access data about characters, images.

## Prerequisites

- Go (v1.17 o superior)
- Fiber (v2)

## Installation

1. Clone this repository: `git clone https://github.com/itsbeensolong/api-brawl-stars.git`
2. Navigate to the project directory: `cd api-brawl-star`
3. Start the server: `go run .`


### Endpoints

- `GET /users`: .
- `GET /users/:id`: .
- `POST /users`: .
- `PUT /users/:id`: .
- `DELETE /users/:id`: .

## Use

Asegúrate de haber iniciado el servidor antes de realizar cualquier solicitud a la API. La API estará disponible en `http://localhost:3000`.

## Errors
La API puede devolver los siguientes códigos de estado y mensajes de error:

400 Bad Request: La solicitud es inválida.
<br>
401 Unauthorized: No se proporcionó o es inválido el token de acceso.
<br>
404 Not Found: No se encontró el recurso solicitado.
<br>
500 Internal Server Error: Error interno del servidor.

## Contribution

Si deseas contribuir a esta API, siéntete libre de enviar solicitudes de extracción o informar problemas en el repositorio de GitHub: https://github.com/itsbeensolong/api-brawl-stars

## License
Este proyecto está licenciado bajo la Licencia MIT.

## Contact
Para consultas o más información, puedes comunicarte conmigo en johancs.mm@gmail.com
