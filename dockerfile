# Используем официальный образ PostgreSQL
FROM postgres:latest

# Копирование скриптов SQL в контейнер
COPY ./sql_scripts /docker-entrypoint-initdb.d/

# docker-compose up --build