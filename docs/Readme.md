# Sistema de Atención Bancaria

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

El proyecto sigue los principios de **Clean Architecture** y **SOLID**, lo que asegura que sea fácil de mantener y extender. La estructura de carpetas es la siguiente:
.
├── cmd                     # Punto de entrada de la aplicación
│   └── http
│       └── main.go         # Inicia el servidor HTTP
├── internal
│   ├── adapter             # Adaptadores (Handlers y Routers)
│   │   ├── handler         # Manejadores HTTP (Endpoints)
│   │   └── router          # Configuración de rutas
│   ├── core                # Lógica de negocio (Dominios, servicios y puertos)
│   │   ├── domain          # Definición de las entidades (Client, Queue)
│   │   ├── port            # Interfaces para la capa de servicio
│   │   └── service         # Implementaciones de la lógica de negocio
│   └── util                # Utilidades compartidas (ej: manejo de tiempos)
├── Dockerfile              # Dockerfile para contenedores
├── docker-compose.yml       # Configuración de Docker Compose
├── go.mod                  # Dependencias del proyecto
└── go.sum                  # Hashes de dependencias

Enlace de referencia sobre arquitectura hexagonal: [Building RESTful API with Hexagonal Architecture in Go](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)

### Componentes principales:
- **Domain**: Define las entidades centrales del sistema, como `Client` y `Queue`, junto con las prioridades de los clientes.
- **Service**: Implementa la lógica de negocio que gestiona la cola de atención y el historial de clientes atendidos.
- **Adapter**: Contiene los manejadores HTTP y el router que expone los endpoints de la aplicación.

---

## Requerimientos

- **Go** versión 1.22.4
- **Docker** (opcional para contenedores)
- **Docker Compose** (opcional para orquestación de contenedores)
- **Testify** para las pruebas (incluido en las dependencias de Go)

---

## Cómo ejecutar la aplicación

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


Endpoints

POST /clients
Agrega un cliente a la cola. Los clientes son ordenados por prioridad y, si las prioridades son iguales, por el orden de llegada.

URL: /clients
Método HTTP: POST
Cuerpo de la solicitud:
json
Copy code
{
  "ID": "1",
  "Name": "Carlos Pérez",
  "Priority": 1
}
Respuesta exitosa:
Código: 201 Created
Cuerpo:
plaintext
Copy code
Cliente Carlos Pérez (Normal) añadido a la cola
POST /clients/next
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