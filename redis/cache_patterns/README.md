# Redis Cache Patterns Practice

This project demonstrates common Redis caching patterns and protection mechanisms in Go.

## Learning Objectives

### 1. Cache-Aside Pattern (缓存旁路模式)
The most common caching pattern:
- **Read**: Check cache first → if miss, read from DB → write to cache
- **Write**: Write to DB → invalidate cache

### 2. Cache Penetration Prevention (防止缓存穿透)
When querying non-existent data repeatedly:
- **Solution**: Cache null values with short TTL
- **Implementation**: Store `product:null:{id}` key when product not found

### 3. Cache Breakdown Prevention (防止缓存击穿)
When a hot key expires and multiple requests hit DB simultaneously:
- **Solution**: Use mutex locks per key
- **Implementation**: Lock before DB query, double-check cache after lock

### 4. Cache Avalanche Prevention (防止缓存雪崩)
When many keys expire at the same time:
- **Solution**: Use random TTL (base + random offset)
- **Implementation**: TTL = 30min + random(0-10min)

## Project Structure

```
cache_patterns/
├── main.go                 # Application entry point
├── models/
│   └── product.go         # Product model
├── store/
│   └── db.go              # Mock database (simulates DB with latency)
├── cache/
│   └── cache.go           # Cache service with all protection mechanisms
├── handlers/
│   └── product_handler.go # HTTP handlers
├── database/
│   └── redis.go           # Redis connection
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Features

- ✅ Cache-Aside pattern implementation
- ✅ Null cache for non-existent products (prevents cache penetration)
- ✅ Per-key mutex locks (prevents cache breakdown)
- ✅ Random TTL (prevents cache avalanche)
- ✅ Cache hit/miss statistics
- ✅ RESTful API with Gin framework

## API Endpoints

### Get Product by ID
```bash
# First request - cache miss (reads from DB)
curl http://localhost:8080/products/1

# Second request - cache hit (reads from Redis)
curl http://localhost:8080/products/1
```

### List All Products
```bash
curl http://localhost:8080/products
```

### Invalidate Cache
```bash
curl -X DELETE http://localhost:8080/products/1/cache
```

### Health Check
```bash
curl http://localhost:8080/health
```

## Running Locally

### Prerequisites
- Go 1.25+
- Redis running on localhost:6379

### Start Redis
```bash
# Using Docker
docker-compose up -d redis

# Or use existing Redis instance
```

### Run Application
```bash
# Using Make
make run

# Or directly
cd redis/cache_patterns
go run main.go
```

## Running with Docker

### Start Everything
```bash
docker-compose up -d
```

### View Logs
```bash
docker-compose logs -f app
```

### Stop Everything
```bash
docker-compose down
```

## Testing Cache Patterns

### 1. Test Cache-Aside Pattern
```bash
# First request (cache miss)
curl http://localhost:8080/products/1
# Response: "from_cache": false

# Second request (cache hit)
curl http://localhost:8080/products/1
# Response: "from_cache": true
```

### 2. Test Cache Penetration Prevention
```bash
# Query non-existent product multiple times
curl http://localhost:8080/products/999
# First: DB query, then null cache stored
# Subsequent: Returns immediately from null cache
```

### 3. Test Cache Breakdown Prevention
```bash
# Invalidate cache first
curl -X DELETE http://localhost:8080/products/1/cache

# Send multiple concurrent requests
for i in {1..10}; do
  curl http://localhost:8080/products/1 &
done
# Only one request hits DB, others wait for cache
```

### 4. Test Cache Avalanche Prevention
```bash
# Check Redis TTL - should be random
docker exec -it cache_patterns_redis redis-cli TTL product:1
# TTL will vary between 1800-2400 seconds (30-40 minutes)
```

## Response Format

### Success Response
```json
{
  "product": {
    "id": 1,
    "name": "Laptop",
    "price": 999.99,
    "stock": 10
  },
  "from_cache": true,
  "message": "✅ Data retrieved from Redis cache (fast!)"
}
```

### Cache Miss Response
```json
{
  "product": {
    "id": 1,
    "name": "Laptop",
    "price": 999.99,
    "stock": 10
  },
  "from_cache": false,
  "message": "⏳ Data retrieved from database and cached for next time"
}
```

## Key Implementation Details

### Cache Keys
- Product cache: `product:{id}`
- Null cache: `product:null:{id}` (TTL: 5 minutes)

### TTL Strategy
- Base TTL: 30 minutes
- Random offset: 0-10 minutes
- Total TTL: 30-40 minutes (prevents avalanche)

### Lock Strategy
- Per-key mutex locks
- Double-check pattern after acquiring lock
- Prevents multiple DB queries for same key

## Notes

- Mock database simulates 50-200ms query latency
- First request always hits DB (cache miss)
- Subsequent requests use cache (much faster)
- Null values cached for 5 minutes to prevent penetration
- Random TTL prevents all keys expiring simultaneously

