FROM node:18-alpine

WORKDIR /app

RUN npm install -g pnpm

COPY package.json ./

COPY . .

RUN pnpm install

RUN pnpm run build

RUN npm install -g pnpm serve

EXPOSE 5000

CMD ["serve", "-s", "dist", "-l", "5000"]