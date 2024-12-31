# Sử dụng hình ảnh Node.js 20 trên Alpine Linux
FROM node:20-alpine 

# Thiết lập thư mục làm việc
WORKDIR /app

# Sao chép và cài đặt dependencies
COPY package.json package-lock.json ./
RUN npm ci --omit=dev

# Sao chép mã nguồn ứng dụng
COPY ./dist ./dist

# Tạo và sử dụng người dùng không phải root
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Thiết lập entrypoint
ENTRYPOINT ["node", "dist/main.js"]