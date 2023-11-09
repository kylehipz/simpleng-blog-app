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
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/997072f7-046b-4c07-af76-bb86cd24c468)
### User Blog Feed
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/793fa1e5-9149-4856-b2ad-31a72cf2d6d2)
### Scaling Blog Reads
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/93e7f9cc-964f-4ae5-b7cf-3f29b34ac02b)
![image](https://github.com/kylehipz/simpleng-blog-app/assets/31435847/a1340181-28f4-4678-be70-e0bfd9a02e3d)

### Scaling Blog Writes
### Realtime Updates (Blog Feed)
### Realtime Updates (Blog reactions/comments)

## Technology Stack
1. Go
2. PostgreSQL
3. Docker
4. Kubernetes
5. Docker Swarm
6. Nginx
7. Amazon ELB
8. Amazon EC2/ECS/EKS
9. Apache Kafka
10. RabbitMQ
11. Redis
12. Amazon S3
13. Amazon SQS
14. ElasticSearch
15. Terraform
