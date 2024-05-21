# **Golang Image Compression Service**

This repository contains the code for a Golang-based image compression service that integrates with AWS S3 and AWS Lambda for serverless deployment. The service reduces the file size of images while maintaining their quality, improving storage efficiency and web application performance.

## **Table of Contents**

- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Part 1: Local Server Setup](#part-1-local-server-setup)
- [Part 2: AWS Integration and Lambda Deployment](#part-2-aws-integration-and-lambda-deployment)
- [Contributing](#contributing)

## **Introduction**

This project demonstrates how to build an image compression service using Golang, AWS S3, and AWS Lambda. It covers creating a local server for image compression, fetching and storing images in AWS S3, and deploying the service as a Lambda function.

## **Features**

- Efficient image compression using Golang's `imaging` package.
- Integration with AWS S3 for storing and retrieving images.
- Serverless deployment using AWS Lambda.
- Docker support for easy deployment and scaling.

## **Getting Started**

### **Prerequisites**

- Golang installed on your machine. [Download Golang](https://golang.org/dl/)
- AWS account with access to S3 and Lambda services.
- Docker installed for creating container images. [Download Docker](https://www.docker.com/get-started)

### **Installation**

1. Clone the repository:

   ```sh
   git clone https://github.com/rabinthapa18/goImageCompressor
   cd goImageCompressor
   ```

2. Initialize the Go module:

   ```sh
   go mod init imageCompressor
   ```

3. Install dependencies:
   ```sh
   go get github.com/disintegration/imaging
   ```

## **Usage**

### **Running the Local Server**

To start the local server for testing image compression:

```sh
go run main.go
```

The server will start on `localhost:3000`. You can test the compression by sending a request with an image file and the desired width.

### **Testing Image Compression**

To test the image compression, you can use a tool like `Postman` or `ThunderClient` to send a POST request to the server:
![App Platorm](https://cdn.hashnode.com/res/hashnode/image/upload/v1716269842714/7f4ac7ea-028b-4fae-9b3c-c5157511cd85.png?auto=compress,format&format=webp)

## **Part 1: Local Server Setup**

In Part 1, we set up a local server to handle image compression:

- Create a simple local server using Golang's `net/http` package.
- Implement image compression using the `imaging` package.
- Handle HTTP requests to compress and return images.

For detailed steps, refer to the [Part 1 Blog Post](https://rabinson.hashnode.dev/building-an-image-compression-service-with-golang-aws-s3-and-aws-lambda).

## **Part 2: AWS Integration and Lambda Deployment**

In Part 2, we enhance the service with AWS integration and serverless deployment:

- Fetch images from AWS S3, compress them, and save them back to S3.
- Package the service as a Lambda function and deploy it on AWS Lambda.
- Create a Docker image for the service for easier deployment and scaling.

For detailed steps, refer to the [Part 2 Blog Post](#) (coming soon).

## **Contributing**

Contributions are welcome! If you have any ideas, suggestions, or issues, feel free to open an issue or submit a pull request.

### **Steps to Contribute**

1. Fork the repository.
2. Create your feature branch:
   ```sh
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```sh
   git commit -m 'Add your feature'
   ```
4. Push to the branch:
   ```sh
   git push origin feature/your-feature
   ```
5. Open a pull request.
