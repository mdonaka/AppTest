#!/bin/bash

DB_NAME="/data/spices.db"
CSV_FILE="spices.csv"
TABLE_NAME="spices"

if [[ ! -f "$CSV_FILE" ]]; then
    echo "$CSV_FILE is not found."
    exit 1
fi

sqlite3 "$DB_NAME" <<EOF
DROP TABLE IF EXISTS $TABLE_NAME;
CREATE TABLE $TABLE_NAME (
    id INTEGER PRIMARY KEY,
    name TEXT,
    flavor TEXT,
    family TEXT
);

.mode csv
.import $CSV_FILE $TABLE_NAME

SELECT * FROM $TABLE_NAME;
EOF

echo "Imported data from $CSV_FILE to $DB_NAME"
