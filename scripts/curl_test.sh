#!/bin/bash

BASE_URL="http://localhost:8080"

echo "===> 1. Register user"
curl -X POST $BASE_URL/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "123456"
  }'

echo -e "\n\n===> 2. Login"
TOKEN=$(curl -s -X POST $BASE_URL/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }' | jq -r '.token')

echo "TOKEN=$TOKEN"

echo -e "\n===> 3. Create post"
curl -X POST $BASE_URL/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "My First Post",
    "content": "This is my first blog post"
  }'

echo -e "\n\n===> 4. Get all posts"
curl -X GET $BASE_URL/posts \
  -H "Authorization: Bearer $TOKEN"

echo -e "\n\n===> 5. Create comment"
curl -X POST $BASE_URL/comments \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "post_id": 1,
    "content": "Nice post!"
  }'

echo -e "\n\n===> 6. Get comments of post 1"
curl -X GET "$BASE_URL/comments?post_id=1" \
  -H "Authorization: Bearer $TOKEN"
