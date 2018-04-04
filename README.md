# upproof
Simple application to monitor uptime of nodes on a distributed network.

## /
URL: `http://example.com/`

METHOD: `GET`

Display the latest nodes as defined by seenThreshold. 

```
{
  "active": [
    {
      "nodeID": "e4327043c4746c998a758db868e6d6a1ff9ddf0a",
      "lastSeen": 1522871547654
    },
    {
      "nodeID": "a14cc24085d7a9ea43ec3712d72b1dbb9ad3d2c0",
      "lastSeen": 1522871547912
    },
    ...
  ]
}
```


## /api/ping
URL: `http://example.com/api/ping/<nodeID>`

METHOD: `POST`

Allow the node to ping with a body.

```
{
  "active": [
    {
      "nodeID": "e4327043c4746c998a758db868e6d6a1ff9ddf0a",
      "lastSeen": 1522871547654
    },
    {
      "nodeID": "a14cc24085d7a9ea43ec3712d72b1dbb9ad3d2c0",
      "lastSeen": 1522871547912
    },
    ...
  ]
}
```

