#!/bin/bash

# 確保 PostgreSQL 正在運行
if ! pg_isready -h localhost -p 5433 >/dev/null 2>&1; then
    echo "PostgreSQL is not running. Please start PostgreSQL first."
    exit 1
fi

# 創建測試數據庫
PGPASSWORD=postgres psql -h localhost -p 5433 -U postgres -d postgres -c "DROP DATABASE IF EXISTS bookstore_test;"
PGPASSWORD=postgres psql -h localhost -p 5433 -U postgres -d postgres -c "CREATE DATABASE bookstore_test;"

echo "Test database created successfully!"
