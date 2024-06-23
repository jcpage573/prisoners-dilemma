# Storage
An outline of how prisoners-dilemma manage storage of data

## Datastore
As of right now I don't see any reason to use anything beyond simple key-value storage. I say we adopt something that implements RESP, but not Redis, fuck Redis.

We could use:
- Garnet (Microsoft, C#)
- [Valkey](https://github.com/valkey-io/valkey) (Google, Oracle, C)

or if we don't want to use RESP
- Memcached (Meta)

Valkey looks the most interesting to me, has a lot of positive reception online. And of course its interchangeable with redis and garnet.

## Implementation
As a new user, my first step is to get an API key, which will verify the identity of my prisoner.
- let's say POST `/user/{name}` returns an API key
- The hash of the key should be stored on the server at `{key hash}:{name}`
- When the user makes requests to other endpoints, to verify and identify the user, we RESP `GET {key hash}` which returns the username
- We can use the username to get other user attributes, like `{name}:games` or `{game id}:{name}` etc.

Will continue to update this doc as we go.