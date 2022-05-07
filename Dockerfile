# Pull down the latest Postgres image from the Docker Hub
FROM postgres

# Set the environment variable for password to 'password'
ENV POSTGRES_PASSWORD password

# Create a database, let's call it 'gym-db'
ENV POSTGRES_DB gym-db

# # Use a sql dump file to create the table schema and populate it with data
# COPY gym.sql /docker-entrypoint-initdb.d/
