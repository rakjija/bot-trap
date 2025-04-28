FROM node:22-alpine AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM nginx:latest
COPY dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf