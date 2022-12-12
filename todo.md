# ToDo

1. Change testing db to localhost for development (this is much harder than anticipated)
2. Session management/expiration
    a. idea -> give no cookie expiration so automatically expires when client shuts down.
        Keep track of client connection somehow and delete session row from db when it does.
3. Get metadata for tables working