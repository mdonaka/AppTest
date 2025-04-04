#!/bin/bash

DB_NAME="${1:-/data/spices.db}"
CSV_FILE="spices.csv"
TABLE_NAME="spices"
TEMP_TABLE="tmp"

if [[ ! -f "$CSV_FILE" ]]; then
    echo "$CSV_FILE is not found."
    exit 1
fi

sqlite3 "$DB_NAME" <<EOF

# 一時的にCSVからテーブルを作成
DROP TABLE IF EXISTS $TEMP_TABLE;
CREATE TABLE $TEMP_TABLE (
    name TEXT,
    alias TEXT,
    taste TEXT,
    flavor TEXT,
    family TEXT,
    origin TEXT
);
.mode csv
.import $CSV_FILE $TEMP_TABLE


# IDを追加して最終的なテーブルを作成
DROP TABLE IF EXISTS $TABLE_NAME;
CREATE TABLE $TABLE_NAME (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    alias TEXT,
    taste TEXT,
    flavor TEXT,
    family TEXT,
    origin TEXT
);

INSERT INTO $TABLE_NAME (name, alias, taste, flavor, family, origin)
SELECT name, alias, taste, flavor, family, origin FROM $TEMP_TABLE;
DROP TABLE $TEMP_TABLE;

SELECT * FROM $TABLE_NAME;
EOF

echo "Imported data from $CSV_FILE to $DB_NAME"
