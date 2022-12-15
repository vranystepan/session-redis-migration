# PHP session migration helper

Migration helper can help you with the migration of PHP session files
to the Redis database. TTL helper configures TTL of already existing
keys.

## Usage

1. clone this repository
2. build the binary

    ```bash
    go build cmd/migration/main.go
    ```
    
    or
    
    ```bash
    go build cmd/ttl/main.go
    ```

3. get the help

   ```bash
   ./main -h
   ```
   
4. migrate the sessions

   ```bash
   go run cmd/migration/main.go -files '/tmp/bla/sess_*'
   ```
   
   or set the TTL

   ```bash
   go run cmd/migration/main.go -ttl 336h -keys 'PHP*'
   ```
   
   > please note the apostrophes in the -files or -keys flags. This is to prevent globbing.
