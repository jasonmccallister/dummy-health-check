# Dummy Healthcheck

There are a lot of situations where you need a dummy healthcheck to always pass, especially when creating new environments. This project is a simple Go application that runs a healthcheck server as a container. It allows you to pass flags to configure the endpoint and port for the health check but also allows you to define and override those with environment variables.
