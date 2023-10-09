# Sbubnom app

## Deployment

1. Copy `docker/.env-sample` to `.env`.
2. Update the environment variables in the `.env` file (if needed).
3. Build the images and run the containers:

```sh
docker compose -f docker-compose.yaml --env-file=docker/.env up -d --build
```

Test it out at [http://localhost:824](http://localhost:810). No mounted
folders. To apply changes, the image must be re-built.


