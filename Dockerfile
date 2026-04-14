# ==========================================
# ETAPA 1: Construcción (Builder)
# ==========================================
# Usamos la versión de Go que arreglamos (1.25) en su versión "alpine" (ligera)
FROM golang:1.25-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos primero los archivos de dependencias (para aprovechar la caché de Docker)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el resto del código fuente
COPY . .

# Compilamos la aplicación. 
# CGO_ENABLED=0 asegura que el binario no dependa de librerías de C del sistema, haciéndolo 100% portable.
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./api/main.go

# ==========================================
# ETAPA 2: Ejecución (Runner)
# ==========================================
# Usamos una imagen de Alpine limpia y vacía
FROM alpine:latest

# Añadimos certificados de seguridad (necesarios para que Go pueda hacer peticiones HTTPS si lo necesitas)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiamos SOLO el binario compilado desde la Etapa 1
COPY --from=builder /app/server .

ENV APP_ENV=production

# Exponemos el puerto que usará nuestra API (Render/Local)
EXPOSE 8080

# Comando para ejecutar el servidor
CMD ["./server"]