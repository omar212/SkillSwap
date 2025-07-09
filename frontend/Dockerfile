# Use official Node image
FROM node:20

# Set working directory
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy app code
COPY . .

# Expose dev server port
EXPOSE 3000

# Enable file change watching in Docker
ENV CHOKIDAR_USEPOLLING=true

# Start Next.js dev server
CMD ["npm", "run", "dev"]
