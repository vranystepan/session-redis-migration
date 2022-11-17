# PHP session migration helper

This helper can help you with the migration of PHP session files
to the Redis database. 

## Usage

1. clone this repository
2. build the binary

    ```bash
    go build cmd/migration/main.go
    ```

3. get the help

   ```bash
   ./main -h
   ```

   ```
   Usage of main:
     -db int
       Redis DB index
     -files string
       Session files to tranfer to Redis
     -host string
       Redis host (default "localhost")
     -password string
       Redis password
     -port int
       Redis port (default 6379)
     -ttl string
       TTL for Redis keys (default "336h")
   ```
   
4. migrate the sessions

   ```bash
   go run cmd/migration/main.go -files '/tmp/bla/sess_*'
   ```
   
   > please not the quotation marks in the `files` flag. This is how
   > to prevent globbing.