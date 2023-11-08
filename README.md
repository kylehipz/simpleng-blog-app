# Simpleng Blog Application

## Functional Requirements
1. A user should be able to write his/her own posts
2. A post should be able to contain images/attachments
3. A user should be able to follow other users
4. A user should be able to see posts from the users that he/she follwos
5. A post should be flagged if it has inappropriate content (violence, sex, ...etc)
6. A user should be able to react and comment to a post
7. A user should be able to search for a post
8. A user should be able to search for other users

## Estimations
1. 200m users
2. 50m posts per day

## Non-functional Requirements
1. The system should be highly available
2. Latency should be low

## API endpoints
1. POST /follow - following a user
2. POST /posts - creating a post
3. GET /feed - getting the blog feed of the user, should contain blogs of the users that the user follows
4. GET /posts/:id - getting a specific post
5. POST /posts/:id/reaction - reacting to a post
6. POST /posts/:id/comment - commenting to a post

## Database Design

## System Design
### Full, High-level Design
### Scaling Blog Reads
### Scaling Blog Writes
### Realtime Updates

## Technology Stack
1. NodeJS
2. TypeScript
3. NestJS
4. PNPM
5. PostgreSQL
6. Docker
7. Kubernetes
8. Docker Swarm
9. Nginx
10. Amazon ELB
11. Amazon EC2/ECS/EKS
12. Apache Kafka
13. RabbitMQ
14. Redis
15. Amazon S3
16. Amazon SQS
17. ElasticSearch
18. Terraform
