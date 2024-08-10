# Use the official PostgreSQL image from Docker Hub
FROM postgres:latest

# Set environment variables for PostgreSQL (optional)
ENV POSTGRES_USER=phoeniciadigital
ENV POSTGRES_PASSWORD=pdsoftware
ENV POSTGRES_DB=pd_database

# Copy the init.sql file to the Docker container's entrypoint directory
COPY sql/init.sql /docker-entrypoint-initdb.d/
