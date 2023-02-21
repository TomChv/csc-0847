####
# Build About website docker image
####
FROM node:19-alpine3.17 AS builder

# Initiliases pp working directory
WORKDIR /app

# Copy and install dependencies
COPY package.json package.json
COPY yarn.lock yarn.lock

# Install dependencies
RUN yarn install

# Copy sources
COPY . .

# Build react app
RUN yarn build

####
# Production website hosted by nginx
####
FROM node:19-alpine3.17 as APP

# Setup working directory
WORKDIR /app

# Copy bundle to serve it by nginx
COPY --from=builder /app/build /app/build

# Install serv
RUN npm install -g serve

# Expose website port
EXPOSE 3000

# Serve website
ENTRYPOINT ["serve", "-s", "/app/build"]