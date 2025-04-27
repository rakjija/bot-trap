FROM node:22-alpine AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build
