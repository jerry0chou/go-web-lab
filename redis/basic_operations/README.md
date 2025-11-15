# Redis Fundamentals with Go-Redis

This project demonstrates Redis fundamentals using the Go-Redis client library. It covers all core Redis data types, expiration, and persistence concepts.

## 📚 What You'll Learn

- **Redis Installation & Connection**: How to set up Redis and connect via redis-cli
- **Core Redis Data Types**:
  - **String**: Basic key-value operations
  - **Hash**: Field-value pairs (like objects)
  - **List**: Ordered collections
  - **Set**: Unordered unique collections
  - **Sorted Set (ZSet)**: Ordered sets with scores
- **Expiration**: TTL, EXPIRE, EXPIREAT, PERSIST
- **Persistence**: RDB snapshots and AOF logging
- **Go-Redis Client**: Basic usage patterns

## 🚀 Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.25+ (if running locally)

### Using Docker Compose (Recommended)

1. **Start Redis and the Go application:**
   ```bash
   cd redis
   docker-compose up --build
   ```

2. **The application will automatically:**
   - Start Redis server
   - Connect to Redis
   - Run all demonstration examples
   - Display output showing each Redis operation

3. **To stop:**
   ```bash
   docker-compose down
   ```

### Running Locally

1. **Install Redis:**
   ```bash
   # macOS
   brew install redis
   
   # Linux
   sudo apt-get install redis-server
   
   # Or use Docker
   docker run -d -p 6379:6379 redis:7-alpine
   ```

2. **Start Redis:**
   ```bash
   redis-server
   ```

3. **Update main.go** to use `localhost:6379` instead of `redis:6379`

4. **Install dependencies and run:**
   ```bash
   go mod tidy
   go run redis/main.go
   ```

## 🔧 Connecting via redis-cli

Once Redis is running, you can connect using the redis-cli:

```bash
# If using Docker
docker exec -it redis_server redis-cli

# If running locally
redis-cli
```

### Basic redis-cli Commands

```bash
# Test connection
PING
# Should return: PONG

# Set a key
SET mykey "Hello Redis"

# Get a key
GET mykey

# Check if key exists
EXISTS mykey

# Get all keys (use with caution in production)
KEYS *

# Delete a key
DEL mykey

# Get TTL (Time To Live)
TTL mykey

# Set expiration
EXPIRE mykey 60

# Exit
EXIT
```

## 📖 Redis Data Types Explained

### 1. String

The simplest Redis type. A key-value pair where the value is a string.

**Use Cases:**
- Caching
- Session storage
- Counters
- Simple key-value storage

**Example Operations:**
- `SET key value` - Set a value
- `GET key` - Get a value
- `INCR key` - Increment a number
- `MSET key1 val1 key2 val2` - Set multiple keys
- `MGET key1 key2` - Get multiple keys

### 2. Hash

A hash is a map of field-value pairs. Perfect for representing objects.

**Use Cases:**
- User profiles
- Product information
- Configuration objects

**Example Operations:**
- `HSET user:1000 name "Alice" email "alice@example.com"` - Set fields
- `HGET user:1000 name` - Get a field
- `HGETALL user:1000` - Get all fields
- `HINCRBY user:1000 age 1` - Increment numeric field
- `HDEL user:1000 age` - Delete a field

### 3. List

An ordered collection of strings. Elements can be added/removed from both ends.

**Use Cases:**
- Message queues
- Activity feeds
- Task lists
- Recent items

**Example Operations:**
- `LPUSH list item1 item2` - Push to left (head)
- `RPUSH list item3 item4` - Push to right (tail)
- `LRANGE list 0 -1` - Get all elements
- `LPOP list` - Pop from left
- `RPOP list` - Pop from right
- `LLEN list` - Get length

### 4. Set

An unordered collection of unique strings.

**Use Cases:**
- Tags
- Unique user IDs
- Categories
- Set operations (union, intersection, difference)

**Example Operations:**
- `SADD tags golang redis docker` - Add members
- `SMEMBERS tags` - Get all members
- `SISMEMBER tags golang` - Check membership
- `SCARD tags` - Get size
- `SREM tags docker` - Remove member
- `SINTER set1 set2` - Intersection
- `SUNION set1 set2` - Union
- `SDIFF set1 set2` - Difference

### 5. Sorted Set (ZSet)

Like a Set, but each member has a score. Members are sorted by score.

**Use Cases:**
- Leaderboards
- Rankings
- Time-series data
- Priority queues

**Example Operations:**
- `ZADD leaderboard 1000 player1 1500 player2` - Add with scores
- `ZRANGE leaderboard 0 -1` - Get range (ascending)
- `ZREVRANGE leaderboard 0 2` - Get top 3 (descending)
- `ZRANGEBYSCORE leaderboard 1000 2000` - Get by score range
- `ZSCORE leaderboard player1` - Get score
- `ZRANK leaderboard player1` - Get rank
- `ZINCRBY leaderboard 100 player1` - Increment score

## ⏱️ Expiration

Redis allows you to set expiration times on keys.

**Operations:**
- `SET key value EX 60` - Set with 60 second expiration
- `EXPIRE key 60` - Set expiration on existing key
- `EXPIREAT key timestamp` - Expire at specific timestamp
- `TTL key` - Get time to live (-1 = no expiration, -2 = key doesn't exist)
- `PERSIST key` - Remove expiration

**Use Cases:**
- Session management
- Cache invalidation
- Temporary data
- Rate limiting

## 💾 Persistence

Redis offers two persistence mechanisms:

### RDB (Redis Database Backup)
- Point-in-time snapshots
- Binary format
- Fast to load
- Configured via `save` directives in redis.conf

### AOF (Append Only File)
- Logs every write operation
- More durable
- Can be slower
- Enabled with `appendonly yes`

**Commands:**
- `SAVE` - Synchronous save (blocks)
- `BGSAVE` - Background save (non-blocking)
- `BGREWRITEAOF` - Rewrite AOF file

**In docker-compose.yml:**
```yaml
command: redis-server --appendonly yes
```
This enables AOF persistence.

## 🔌 Go-Redis Client Usage

### Connection

```go
import "github.com/redis/go-redis/v9"

rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

ctx := context.Background()
_, err := rdb.Ping(ctx).Result()
```

### Basic Operations

```go
// String
rdb.Set(ctx, "key", "value", 0)
val, _ := rdb.Get(ctx, "key").Result()

// Hash
rdb.HSet(ctx, "user:1", "name", "Alice")
name, _ := rdb.HGet(ctx, "user:1", "name").Result()

// List
rdb.LPush(ctx, "list", "item")
item, _ := rdb.LPop(ctx, "list").Result()

// Set
rdb.SAdd(ctx, "set", "member")
members, _ := rdb.SMembers(ctx, "set").Result()

// Sorted Set
rdb.ZAdd(ctx, "zset", redis.Z{Score: 1.0, Member: "member"})
members, _ := rdb.ZRange(ctx, "zset", 0, -1).Result()
```

### Expiration

```go
// Set with expiration
rdb.Set(ctx, "key", "value", 10*time.Second)

// Set expiration on existing key
rdb.Expire(ctx, "key", 10*time.Second)

// Get TTL
ttl, _ := rdb.TTL(ctx, "key").Result()
```

## 📁 Project Structure

```
redis/
├── main.go              # Main application with all examples
├── Dockerfile           # Docker image for Go app
├── docker-compose.yml   # Docker Compose configuration
└── README.md           # This file
```

## 🧪 Running Examples

The application demonstrates:

1. ✅ String operations (SET, GET, INCR, MSET, MGET)
2. ✅ Hash operations (HSET, HGET, HGETALL, HINCRBY, HDEL)
3. ✅ List operations (LPUSH, RPUSH, LRANGE, LPOP, RPOP, LLEN)
4. ✅ Set operations (SADD, SMEMBERS, SISMEMBER, SCARD, SREM, SINTER)
5. ✅ Sorted Set operations (ZADD, ZRANGE, ZREVRANGE, ZSCORE, ZRANK, ZINCRBY)
6. ✅ Expiration (SET EX, EXPIRE, EXPIREAT, TTL, PERSIST)
7. ✅ Persistence concepts (RDB, AOF)

## 🔍 Exploring Redis Data

After running the application, connect with redis-cli to explore:

```bash
docker exec -it redis_server redis-cli

# See all keys
KEYS *

# Check a specific key type
TYPE user:1000

# Get hash data
HGETALL user:1000

# Get list data
LRANGE tasks 0 -1

# Get set data
SMEMBERS tags

# Get sorted set data
ZRANGE leaderboard 0 -1 WITHSCORES
```

## 📚 Additional Resources

- [Redis Documentation](https://redis.io/docs/)
- [Go-Redis Documentation](https://redis.uptrace.dev/)
- [Redis Commands Reference](https://redis.io/commands/)

## 🐛 Troubleshooting

**Connection refused:**
- Ensure Redis is running
- Check if port 6379 is available
- Verify connection string in main.go

**Module not found:**
- Run `go mod tidy` to download dependencies

**Docker issues:**
- Ensure Docker is running
- Try `docker-compose down -v` to remove volumes and start fresh

## 📝 Notes

- This is a learning/demonstration project
- Production Redis should have authentication enabled
- Consider using connection pooling for high-traffic applications
- Monitor memory usage in production
- Configure appropriate persistence settings based on your needs

