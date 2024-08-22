# Nexus API Documentation

## Table of Contents
1. [Introduction](#introduction)
2. [Authentication](#authentication)
3. [Blog Posts](#blog-posts)
4. [Comments](#comments)
5. [File Management](#file-management)
6. [Photo Management](#photo-management)

## Introduction

Welcome to the Nexus API documentation. This API provides endpoints for managing blog posts, comments, files, and photos. It's designed to be used as a backend for a content management system.

Base URL: `http://your-api-domain.com/api/v1`

## Authentication

[Note: Add authentication details here when implemented]

## Blog Posts

### Create a Blog Post
- **POST** `/blog`
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
- **GET** `/blog`
- **Query Parameters**:
    - `page` (optional): Page number for pagination (default: 1)
    - `pageSize` (optional): Number of items per page (default: 10)
- **Response**: Returns an array of blog post objects

### Get a Specific Blog Post
- **GET** `/blog/:id`
- **Response**: Returns the specified blog post object

### Update a Blog Post
- **PUT** `/blog/:id`
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
- **DELETE** `/blog/:id`
- **Response**: Returns a success message

## Comments

### Add a Comment
- **POST** `/comments`
- **Body**:
  ```json
  {
    "content": "Your comment here",
    "blogPostID": 123,  // ID of the blog post (use either blogPostID or postID)
    "postID": 456       // ID of the regular post
  }
  ```
- **Response**: Returns the created comment object

### Get Comments
- **GET** `/comments`
- **Query Parameters**:
    - `blogPostID` or `postID`: ID of the blog post or regular post
- **Response**: Returns an array of comment objects

### Update a Comment
- **PUT** `/comments/:id`
- **Body**:
  ```json
  {
    "content": "Updated comment content"
  }
  ```
- **Response**: Returns the updated comment object

### Delete a Comment
- **DELETE** `/comments/:id`
- **Response**: Returns a success message

## File Management

### Upload a File
- **POST** `/files`
- **Form Data**:
    - `file`: The file to upload
    - `path` (optional): The directory path to store the file (default: root directory)
    - `isDirectory` (optional): Set to "true" if creating a directory (default: "false")
- **Response**: Returns the file object

### List Files
- **GET** `/files`
- **Query Parameters**:
    - `path` (optional): The directory path to list files from (default: root directory)
- **Response**: Returns an array of file objects in the specified directory

### Get File or Directory Contents
- **GET** `/files/*path`
- **Response**:
    - If path is a file: Returns the file object
    - If path is a directory: Returns an array of file objects in the directory

### Update File Metadata
- **PUT** `/files/:id`
- **Body**:
  ```json
  {
    "name": "Updated file name"
  }
  ```
- **Response**: Returns the updated file object

### Delete a File
- **DELETE** `/files/:id`
- **Response**: Returns a success message

### Create a Directory
- **POST** `/directories`
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
- **POST** `/photos`
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
- **GET** `/photos`
- **Query Parameters**:
    - `page` (optional): Page number for pagination (default: 1)
    - `pageSize` (optional): Number of items per page (default: 10)
- **Response**: Returns an array of photo objects

### Get a Specific Photo
- **GET** `/photos/:id`
- **Response**: Returns the specified photo object

### Update a Photo
- **PUT** `/photos/:id`
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
- **DELETE** `/photos/:id`
- **Response**: Returns a success message

## Error Handling

All endpoints will return appropriate HTTP status codes:

- 200: Successful operation
- 201: Successful creation
- 400: Bad request (e.g., invalid input)
- 404: Resource not found
- 500: Internal server error

Error responses will include a JSON object with an "error" field describing the issue.

## Rate Limiting

[Note: Add rate limiting details here when implemented]

## Changelog

[Note: Add changelog entries here as you update the API]