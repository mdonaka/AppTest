FROM node:22.14 AS dev


FROM node:22.14 AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build


FROM node:22.14 AS prod
COPY --from=builder /app/build/server ./build/server
COPY --from=builder /app/package*.json ./
COPY --from=builder /app/node_modules ./node_modules
CMD ["npm", "run", "start"]
