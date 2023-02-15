FROM alpine:latest



# Working directory
WORKDIR /app


# Copy the source from the current directory to the Working Directory inside the container
COPY main .


CMD [ "./main" ]


# Expose port 8080 to the outside world
EXPOSE 8080


