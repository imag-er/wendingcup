
FROM node:23-alpine as builder

WORKDIR /app

COPY package.json ./    
RUN npm install

COPY public/ public/
COPY src/ src/
COPY svelte.config.js ./
COPY vite.config.js ./
COPY index.html ./


RUN npm run build
RUN npm prune --production 

FROM nginx:latest
COPY --from=builder /app/dist/index.html /usr/share/nginx/html/index.html
COPY --from=builder /app/dist/assets /usr/share/nginx/html/assets
COPY nginx.conf /etc/nginx/nginx.conf


# 健康检查
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
    CMD curl -f http://localhost/ || exit 1

    