# Sistema de Atención Bancaria (Colas)

Este proyecto es una simulación de un sistema de atención bancaria que utiliza una cola de clientes basada en prioridades. Los clientes son atendidos en función de su prioridad y el orden de llegada. Los niveles de prioridad incluyen: **Gerencial**, **VIP**, **Normal** y **Básico**.

---

## Tabla de Contenidos

- [Arquitectura](#arquitectura)
- [Requerimientos](#requerimientos)
- [Cómo ejecutar la aplicación](#cómo-ejecutar-la-aplicación)
- [Docker](#docker)
  - [Construir y ejecutar con Docker](#construir-y-ejecutar-con-docker)
  - [Usar Docker Compose](#usar-docker-compose)
- [Endpoints](#endpoints)
  - [POST /clients](#post-clients)
  - [POST /clients/next](#post-clientsnext)
  - [GET /queue](#get-queue)
  - [GET /history](#get-history)
- [Pruebas](#pruebas)

---

## Arquitectura

El proyecto se basa en el uso de **Clean Architecture** y **SOLID**, para que sea mantenible y escalable. La estructura de carpetas es la siguiente:
```
.
├── cmd                     # Punto de entrada de la aplicación
│   └── http
│       └── main.go         # Inicia el servidor HTTP, se hace uso de fiber
├── internal
│   ├── adapter             # Adaptadores (Handlers y Routers)
│   │   ├── handler         # Definición de Handlers
│   │   └── router          # Configuración de rutas, se hace uso de fiber para el routing
│   ├── core                # Lógica de negocio (Dominios, servicios y puertos)
│   │   ├── domain          # Definición de las estructuras de Go (Client, Queue)
│   │   ├── port            # Definición de interfaces
│   │   └── service         # Implementaciones de la lógica de negocio
│   └── util                # Utilidades y helpers que usa la aplicación
├── Dockerfile              # Dockerfile
├── docker-compose.yml      # Configuración de Docker Compose
├── go.mod                  # Dependencias del proyecto
└── go.sum                  # Hash de dependencias
```

Enlace de referencia sobre arquitectura hexagonal: [Building RESTful API with Hexagonal Architecture in Go](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)

### Componentes principales:
- **Domain**: Define los modelos o mejor dicho en Go las estructuras/struct centrales del sistema, como `Client` y `Queue`.
- **Service**: Son la Implementaciones de la lógica de negocio que gestiona la cola de atención y el historial de clientes atendidos.
- **Adapter**: Contiene los Handlers HTTP y el router que expone los endpoints de la aplicación.

---

## Requerimientos

- **Go** versión 1.22.4
- **Docker** creación de contenedores
- **Docker Compose** orquestador de contenedores
- **Testify** Testing

---

## Cómo ejecutar la aplicación (local sin Docker)

1. **Compilar el proyecto**:

   ```bash
   go build -o bank-queue-system ./cmd/http
   ```

   ```bash
   go build -o bank-queue-system ./cmd/http

2. **Ejecutar la aplicación**:

   ```bash
   ./bank-queue-system
   ```
  
La aplicación estará disponible en http://localhost:8084 (Si se prefiere modificar el puerto ir al DockerCompose)

---

## Docker

1. **Construir la imagen de Docker:**
   ```bash
   docker build -t bank-queue-system .
   ```
2. **Ejecutar la imagen de Docker:**
   ```bash
   docker run -p 8080:8080 bank-queue-system
   ```

## Usar Docker Compose

1. **Levantar el contenedor con Docker Compose:**
   ```bash
   docker-compose up -d
   ```
2. **Detener el contenedor:**
   ```bash
   docker-compose stop
   ```

---

## Endpoints

### POST /clients

Agrega un cliente a la cola. Los clientes son ordenados por **prioridad** y, si las prioridades son iguales, por el **orden de llegada**, este orden es definido por un timer establecido en las utilidades.

- **URL**: `/clients`
- **Método HTTP**: `POST`
- **Cuerpo de la solicitud**:
  ```json
  {
    "ID": "1",
    "Name": "Carlos Pérez",
    "Priority": 1
  }
  ```
- **Respuesta exitosa:**
  - Código: 201 Created
- **Cuerpo:**
  - Cliente Carlos Pérez (Normal) añadido a la cola

  
### POST /clients/next
Procesa al siguiente cliente en la cola basado en la prioridad. Si dos clientes tienen la misma prioridad, se respeta el orden de llegada.

URL: /clients/next
Método HTTP: POST
Descripción: Retira al cliente con mayor prioridad de la cola y lo atiende.
Respuesta exitosa:
Código: 200 OK
Cuerpo:
plaintext
Copy code
Cliente Carlos Pérez (Normal) ha sido atendido (esperó 2m10s)
GET /queue
Obtiene el estado actual de la cola. Muestra los clientes pendientes en el orden en que serán atendidos.

URL: /queue
Método HTTP: GET
Descripción: Devuelve la lista de clientes que están actualmente en la cola, ordenados por prioridad y llegada.
Respuesta exitosa:
Código: 200 OK
Cuerpo:
json
Copy code
[
  {
    "ID": "1",
    "Name": "Carlos Pérez",
    "Priority": 1,
    "ArrivalTime": "2024-09-20T09:55:00Z"
  },
  {
    "ID": "2",
    "Name": "Ana López",
    "Priority": 0,
    "ArrivalTime": "2024-09-20T09:50:00Z"
  }
]
GET /history
Obtiene el historial de clientes que ya han sido atendidos. Este endpoint muestra la lista de clientes que ya pasaron por la cola y fueron procesados.

URL: /history
Método HTTP: GET
Descripción: Devuelve un historial en formato JSON de los clientes que ya han sido atendidos.
Respuesta exitosa:
Código: 200 OK
Cuerpo:
json
Copy code
[
  {
    "ID": "1",
    "Name": "Carlos Pérez",
    "Priority": 2,
    "ArrivalTime": "2024-09-20T09:55:00Z",
    "WaitTime": "2m10s"
  }
]


Pruebas

El proyecto incluye pruebas unitarias e integradas utilizando Testify. Para ejecutar las pruebas, usa el siguiente comando:

bash
Copy code
go test ./...
Esto ejecutará todas las pruebas del proyecto y te dará un resumen de los resultados.