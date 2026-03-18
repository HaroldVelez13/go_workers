# 🧠 Go Event-Driven Chat System (LLM + SSE)

## 🚀 Overview
Sistema de chat distribuido tipo ChatGPT construido en Go usando arquitectura **event-driven**.

## 🧩 Stack
- Go (microservicios)
- NATS (broker de eventos)
- PostgreSQL (persistencia)
- Docker (orquestación)
- SSE (streaming en tiempo real)

## 🏗️ Arquitectura

Cliente → Gateway (SSE) → NATS → chat-service → NATS → llm-service → NATS → Gateway → Cliente

## 📦 Servicios

### chat-service
- Recibe mensajes vía HTTP
- Persiste en PostgreSQL
- Publica `message.created.v1`
- Reconstruye mensajes desde chunks

### llm-service
- Consume eventos
- Genera respuestas
- Emite `response.chunk.generated.v1` (streaming)

### gateway
- Maneja conexiones SSE
- Consume eventos de chunks
- Envía streaming al cliente

### shared
- Eventos tipados
- BaseEvent (traceID, userID)
- Cliente NATS reutilizable

## 🔄 Flujo

1. Cliente envía mensaje
2. chat-service lo guarda y publica evento
3. llm-service genera respuesta en chunks
4. gateway transmite en tiempo real (SSE)
5. chat-service reconstruye y guarda mensaje final

## 🔥 Features

- Event-driven architecture
- Streaming tipo ChatGPT
- SSE en tiempo real
- Trazabilidad con traceID
- Persistencia completa

## ⚠️ Pendientes

- Cancelación de generación (context)
- Escalabilidad SSE (Redis pub/sub)
- Integración con LLM real (OpenAI/local)
- Seguridad (JWT/Auth)

## ▶️ Run

```bash
docker compose up --build
```

## 🧪 Test SSE

```bash
curl -N http://localhost:8080/chat/stream?chat_id=1
```

## 📌 Autor
Harold Vélez