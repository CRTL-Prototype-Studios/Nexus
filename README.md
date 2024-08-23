# Nexus

## Table of Contents
1. [Introduction](#introduction)
2. [Deployment](#deployment)
3. [Authentication](#authentication)
4. [Rate Limiting](#rate-limiting)
5. [Blog Posts](#blog-posts)
6. [Comments](#comments)
7. [File Management](#file-management)
8. [Photo Management](#photo-management)
9. [Albums](#albums)
10. [Error Handling](#error-handling)

## Introduction

Welcome to the Nexus API documentation. This API provides endpoints for managing blog posts, comments, files, and photos. It's designed to be used as a backend for a content management system.

## Deployment

### Build It Yourself
You can clone this GitHub repository and build the Docker Image yourself.

### Official Docker Image
You can also deploy your instance of Nexus with the official Docker Image `type32/nexus:latest`.

Deploying an instance of Nexus requires the following Environment Variables:

```dotenv
DB_HOST=nexus.example.com
DB_USER=exampleuser
DB_PASSWORD=examplepassword
DB_NAME=example
DB_PORT=114514
MINIO_ENDPOINT=0.0.0.1:114514
MINIO_ACCESS_KEY=your-minio-access-key
MINIO_SECRET_KEY=your-minio-secret-key

# This option will forcefully migrate the schema to your database every time you start this Go server.
ENFORCE_SCHEMA_MIGRATION=true

# JWT secret key for authentication
SECRET_JWT_KEY=your-very-secure-secret-key-here

# Rate limiting configuration
RATE_LIMIT_PER_SECOND=10
RATE_LIMIT_BURST=30
```

## Authentication

Nexus now supports user authentication using JSON Web Tokens (JWT).

### Sign Up
- **POST** `/signup`
- **Body**:
  ```json
  {
    "username": "newuser",
    "password": "securepassword"
  }
  ```
- **Response**: Returns a success message

### Sign In
- **POST** `/signin`
- **Body**:
  ```json
  {
    "username": "existinguser",
    "password": "correctpassword"
  }
  ```
- **Response**: Returns a JWT token to be used for authenticated requests

### Using Authentication
For authenticated endpoints, include the JWT token in the Authorization header:
```
Authorization: Bearer your_jwt_token_here
```

## Rate Limiting

Rate limiting is implemented to prevent abuse of the API. The default configuration allows:
- 10 requests per second
- Burst of up to 30 requests

These values can be adjusted using the `RATE_LIMIT_PER_SECOND` and `RATE_LIMIT_BURST` environment variables.

## Blog Posts

### Create a Blog Post
- **POST** `/api/v1/blog`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "title": "Your Blog Post Title",
    "content": "Your blog post content goes here",
    "coverID": 123  // Optional: ID of the cover photo
  }
  ```
- **Response**: Returns the created blog post object

### Get All Blog Posts
- **GET** `/api/v1/blog`
- **Query Parameters**:
  - `page` (optional): Page number for pagination (default: 1)
  - `pageSize` (optional): Number of items per page (default: 10)
- **Response**: Returns an array of blog post objects

### Get a Specific Blog Post
- **GET** `/api/v1/blog/:id`
- **Response**: Returns the specified blog post object

### Update a Blog Post
- **PUT** `/api/v1/blog/:id`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "title": "Updated Title",
    "content": "Updated content",
    "coverID": 456  // Optional: New cover photo ID
  }
  ```
- **Response**: Returns the updated blog post object

### Delete a Blog Post
- **DELETE** `/api/v1/blog/:id`
- **Authentication**: Required
- **Response**: Returns a success message

## Comments

### Add a Comment
- **POST** `/api/v1/comments`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "content": "Your comment here",
    "blogPostID": 123  // ID of the blog post
  }
  ```
- **Response**: Returns the created comment object

### Get Comments
- **GET** `/api/v1/comments`
- **Query Parameters**:
  - `blogPostID`: ID of the blog post
- **Response**: Returns an array of comment objects

### Update a Comment
- **PUT** `/api/v1/comments/:id`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "content": "Updated comment content"
  }
  ```
- **Response**: Returns the updated comment object

### Delete a Comment
- **DELETE** `/api/v1/comments/:id`
- **Authentication**: Required
- **Response**: Returns a success message

## File Management

### Upload a File
- **POST** `/api/v1/files`
- **Authentication**: Required
- **Form Data**:
  - `file`: The file to upload
  - `path` (optional): The directory path to store the file (default: root directory)
  - `isDirectory` (optional): Set to "true" if creating a directory (default: "false")
- **Response**: Returns the file object

### List Files
- **GET** `/api/v1/files`
- **Query Parameters**:
  - `path` (optional): The directory path to list files from (default: root directory)
- **Response**: Returns an array of file objects in the specified directory

### Get File or Directory Contents
- **GET** `/api/v1/files/dir/*path`
- **Response**:
  - If path is a file: Returns the file object
  - If path is a directory: Returns an array of file objects in the directory

### Update File Metadata
- **PUT** `/api/v1/files/:id`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "name": "Updated file name"
  }
  ```
- **Response**: Returns the updated file object

### Delete a File
- **DELETE** `/api/v1/files/:id`
- **Authentication**: Required
- **Response**: Returns a success message

### Create a Directory
- **POST** `/api/v1/directories`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "name": "New Directory Name",
    "path": "/parent/directory/path"
  }
  ```
- **Response**: Returns the created directory object

## Photo Management

### Create a Photo
- **POST** `/api/v1/photos`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "title": "Photo Title",
    "description": "Photo description",
    "fileID": 123,  // ID of the associated file
    "width": 1920,
    "height": 1080
  }
  ```
- **Response**: Returns the created photo object

### Get All Photos
- **GET** `/api/v1/photos`
- **Query Parameters**:
  - `page` (optional): Page number for pagination (default: 1)
  - `pageSize` (optional): Number of items per page (default: 10)
- **Response**: Returns an array of photo objects

### Get a Specific Photo
- **GET** `/api/v1/photos/:id`
- **Response**: Returns the specified photo object

### Update a Photo
- **PUT** `/api/v1/photos/:id`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "title": "Updated Title",
    "description": "Updated description",
    "fileID": 456,  // Optional: New associated file ID
    "width": 3840,
    "height": 2160
  }
  ```
- **Response**: Returns the updated photo object

### Delete a Photo
- **DELETE** `/api/v1/photos/:id`
- **Authentication**: Required
- **Response**: Returns a success message

## Albums

### Create an Album
- **POST** `/api/v1/albums`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "name": "My New Album"
  }
  ```
- **Response**: Returns the created album object

### Get All Albums
- **GET** `/api/v1/albums`
- **Query Parameters**:
  - `page` (optional): Page number for pagination (default: 1)
  - `pageSize` (optional): Number of items per page (default: 10)
- **Response**: Returns an array of album objects

### Get a Specific Album
- **GET** `/api/v1/albums/:id`
- **Response**: Returns the specified album object with associated photos

### Update an Album
- **PUT** `/api/v1/albums/:id`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "name": "Updated Album Name"
  }
  ```
- **Response**: Returns the updated album object

### Delete an Album
- **DELETE** `/api/v1/albums/:id`
- **Authentication**: Required
- **Response**: Returns a success message

### Add a Photo to an Album
- **POST** `/api/v1/albums/:id/photos`
- **Authentication**: Required
- **Body**:
  ```json
  {
    "photoID": 123
  }
  ```
- **Response**: Returns a success message

### Remove a Photo from an Album
- **DELETE** `/api/v1/albums/:id/photos/:photoID`
- **Authentication**: Required
- **Response**: Returns a success message

## Error Handling

All endpoints will return appropriate HTTP status codes:

- 200: Successful operation
- 201: Successful creation
- 400: Bad request (e.g., invalid input)
- 401: Unauthorized (authentication required)
- 403: Forbidden (insufficient permissions)
- 404: Resource not found
- 429: Too Many Requests (rate limit exceeded)
- 500: Internal server error

Error responses will include a JSON object with an "error" field describing the issue.

## Rate Limiting

[Note: Add rate limiting details here when implemented]