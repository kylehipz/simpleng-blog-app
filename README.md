# Simpleng Blog App

## Purpose
The purpose of this document is to provide information about **_Simpleng Blog App_**, what it is, the functional requirements, non-functional requirements, technology stack to be used, database design, and system architecture.

## What is Simpleng Blog App?
**Simpleng Blog App** is a simple application that can handle blog posts, just like medium, twitter, facebook. The real purpose for creating this application is to learn new technologies and practice creating software architecture that is scalable and cost efficient

## Functional Requirements
- A user should be able to register to the platform
- A user should be able to create/update/delete a blog
- A user should be able to follow another user
- A user should have a _**home feed**_, which contains the blogs of users that the user follow
## Non-functional Requirements
- 50 million users create/update/delete blogs
- 100 million users read their home feed
- Fast response times
- Real-time updates
## Database Design
[![Database Design](https://app.eraser.io/workspace/uFj0XfkrIEXLeRA9qaac/preview?elements=-wHSyNFibqALHvVjcz9u_Q&type=embed)](https://app.eraser.io/workspace/uFj0XfkrIEXLeRA9qaac?elements=-wHSyNFibqALHvVjcz9u_Q)

## High-level Architecture
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/5f9c58fd-497d-4575-ab42-aac9a830374a)

## Integration of Services
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/f4c36ce9-a1fe-415e-ba5d-4349e3853559)

## Technology Stack
- Golang
- Postgresql
- Docker
- Elasticache - Redis
- Terraform
- Terragrunt
- AWS ECS
- AWS Fargate
- AWS SNS
- AWS SQS
