FROM node:lts-alpine
WORKDIR /app
COPY package*.json /app
RUN npm install
COPY . .
EXPOSE 8080
CMD ["npm", "run", "serve"]