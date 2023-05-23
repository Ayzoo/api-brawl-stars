# The Brawl Stars API

## Descripción

The Brawl Star API is a RESTful API based on the supercell video game Brawl Stars. You will access data about characters, images.

## Requisitos previos

- Go (v1.17 o superior)
- Fiber (v2)

## Instalación

1. Clona este repositorio: `git clone https://github.com/itsbeensolong/api-brawl-stars.git`
2. Navega hasta el directorio del proyecto: `cd api-brawl-star`
4. Configura las variables de entorno en el archivo `.env` según tus necesidades.
5. Inicia el servidor: `go run .`

## Uso

Asegúrate de haber iniciado el servidor antes de realizar cualquier solicitud a la API. La API estará disponible en `http://localhost:8000`.

### Endpoints

- `GET /users`: Obtiene todos los usuarios.
- `GET /users/:id`: Obtiene un usuario específico por su ID.
- `POST /users`: Crea un nuevo usuario.
- `PUT /users/:id`: Actualiza un usuario existente.
- `DELETE /users/:id`: Elimina un usuario existente.

### Ejemplo de solicitud

# API de ejemplo

Esta es una API de ejemplo que proporciona funcionalidades para administrar usuarios.

## Descripción

La API de ejemplo está diseñada para permitir a los desarrolladores crear, leer, actualizar y eliminar usuarios en un sistema. Utiliza los métodos HTTP estándar y respuestas en formato JSON.

## Requisitos previos

- Node.js (v12 o superior)
- MongoDB (v4 o superior)

## Instalación

1. Clona este repositorio: `git clone https://github.com/tu-usuario/api-ejemplo.git`
2. Navega hasta el directorio del proyecto: `cd api-ejemplo`
3. Instala las dependencias: `npm install`
4. Configura las variables de entorno en el archivo `.env` según tus necesidades.
5. Inicia el servidor: `npm start`

## Uso

Asegúrate de haber iniciado el servidor antes de realizar cualquier solicitud a la API. La API estará disponible en `http://localhost:3000`.

### Autenticación

La API utiliza autenticación basada en tokens. Para autenticarte, debes incluir un encabezado `Authorization` en tus solicitudes con el valor `Bearer <token>`, donde `<token>` es el token de acceso válido.

### Endpoints

- `GET /users`: Obtiene todos los usuarios.
- `GET /users/:id`: Obtiene un usuario específico por su ID.
- `POST /users`: Crea un nuevo usuario.
- `PUT /users/:id`: Actualiza un usuario existente.
- `DELETE /users/:id`: Elimina un usuario existente.

### Ejemplo de solicitud

