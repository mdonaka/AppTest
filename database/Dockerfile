FROM keinos/sqlite3:3.49.0 AS dev
USER root
RUN apk add --no-cache bash

FROM keinos/sqlite3:3.49.0 AS builder
USER root
COPY --chmod=744 . .
RUN sh ./import_spices.sh spices.db

FROM busybox:latest AS prod
COPY --from=builder spices.db /data/spices.db
