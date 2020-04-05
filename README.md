# Cloud Run Template

## Container Deploy
```bash
gcloud builds submit --tag gcr.io/PROJECT_NAME/CONTAINER_NAME
```

## Deploy and Update
```bash
gcloud beta run deploy --image gcr.io/PROJECT_NAME/CONTAINER_NAME
        --add-cloudsql-instances INSTANCE_CONNECTION_NAME \
        --set-env-vars INSTANCE_CONNECTION_NAME=INSTANCE_CONNECTION_NAME \
        --set-env-vars DB_USER=YOUR_DB_USER \
        --set-env-vars DB_PASS=YOUR_DB_PASS \
        --set-env-vars DB_NAME=YOUR_DB

gcloud beta run services update APP_NAME \
    --add-cloudsql-instances INSTANCE_CONNECTION_NAME \
    --set-env-vars INSTANCE_CONNECTION_NAME=INSTANCE_CONNECTION_NAME \
    --set-env-vars DB_USER=YOUR_DB_USER \
    --set-env-vars DB_PASS=YOUR_DB_PASS \
    --set-env-vars DB_NAME=YOUR_DB
    --set-env-vars RELEASE_MODE=Release
```

## Cloud SQL
```bash
~/cloud_sql_proxy -instances=<PROJECT:REGION:INCETANCE>=tcp:3306
```