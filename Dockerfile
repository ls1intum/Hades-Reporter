# Stage 1: Build stage (if needed)
FROM node:lts-alpine AS base

WORKDIR /app

# 复制 package.json 和 package-lock.json
COPY package*.json ./

# 安装依赖
RUN npm install --production

# 复制源码
COPY . .

# 默认命令
CMD ["node", "src/index.js"]
