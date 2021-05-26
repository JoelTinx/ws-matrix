# WS-MATRIX

## Running up:

```zsh
docker build -t ws-matrix .
docker run --rm -it -p 3000:3000 ws-matrix
```

or

```bash
docker-componse up
```

## Example:

### POST: **/rotate**
body content:
```json
{
    "data": [
        [1, 2, 3],
        [4, 5, 6],
        [7, 8, 9]
    ]
}
```

Example of the response:
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