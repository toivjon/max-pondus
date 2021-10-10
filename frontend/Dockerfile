# syntax=docker/dockerfile:1
FROM node:14.17.6-alpine as build

# Create a directory for the application.
WORKDIR /maxpondus

# Download project dependencies.
COPY package.json .
RUN npm install

# Copy the project files and build the application.
COPY . .
RUN npm run build

# Tell Docker to start NGINX server to serve the application.
FROM nginx:1.21.3-alpine
COPY --from=build /maxpondus/dist/ /usr/share/nginx/html

# Tell Docker about the port used by the application.
EXPOSE 80