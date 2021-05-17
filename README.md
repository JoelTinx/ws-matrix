# How to use

## Route: "/rotate" (Method: POST)

Example of the body for the request:
```json
{
    "data": [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
}
```

Example of response of the request:
```json
{
    "status": true,
    "message": "The array was rotated successfully",
    "data": [
        [3, 6, 8 ],
        [2, 5, 8 ],
        [1, 4, 7 ]
    ]
}
```