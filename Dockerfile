FROM golang:1.22

WORKDIR /app
COPY /grinha-de-backend-2024-q1-demode /app/grinha-de-backend-2024-q1-demode
COPY /server-cert.pem /app/server-cert.pem
COPY /server.key /app/server.key

EXPOSE 8085

CMD ["./grinha-de-backend-2024-q1-demode"]
