PGPASSWORD=${DB_PASSWORD} \
psql -U postgres -c "CREATE USER ${DB_USER};" -c "CREATE DATABASE ${DB_NAME};" -c "GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} TO ${DB_USER};" -c "ALTER DATABASE ${DB_NAME} OWNER TO ${DB_USER};"


../goose postgres "user=${DB_USER} dbname=${DB_NAME} sslmode=disable" up