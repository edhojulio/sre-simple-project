# Use an official Node.js runtime as a parent image
FROM node:14

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install app dependencies
RUN npm install

# Copy the local application code to the container
COPY . .

# Expose port 3000 to the outside world
EXPOSE 3000

# Define the command to run the application
CMD ["node", "app.js"]
