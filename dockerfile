# Используем официальный образ PostgreSQL
FROM postgres:latest

# Установка переменных окружения для PostgreSQL
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB mydatabase

# Копирование скриптов SQL в контейнер
COPY ./sql_scripts /docker-entrypoint-initdb.d/

# docker-compose up --build