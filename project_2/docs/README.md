# Project 2 documentation

## Setup bucket and IA labelling

https://codelabs.developers.google.com/codelabs/cloud-picadaily-lab1?hl=en#0https://codelabs.developers.google.com/codelabs/cloud-picadaily-lab1?hl=en#0

```shell
# Setup project
export GOOGLE_CLOUD_PROJECT=XXXX

# Enable APIs
gcloud services enable vision.googleapis.com
gcloud services enable cloudfunctions.googleapis.com

# Create & config bucket
export BUCKET_PICTURES=uploaded-pictures-${GOOGLE_CLOUD_PROJECT}
gsutil mb -l us-east1 gs://${BUCKET_PICTURES}
gsutil uniformbucketlevelaccess set on gs://${BUCKET_PICTURES}
```

## Deploy function

```shell
export BUCKET_PICTURE=uploaded-pictures-${GOOGLE_CLOUD_PROJECT}

cd pic-analysis

gcloud functions deploy pic-analysis \
--region us-east1 \
--entry-point VisionAnalysis \
--runtime=go120 \
--gen2 \
--trigger-bucket $BUCKET_PICTURE \
--memory=1Gi
```

## Deploy thumbnails container

Build & publish app to registry

```shell
gcloud builds submit --tag gcr.io/$GOOGLE_CLOUD_PROJECT/thumnail-service
```

Config CLI

```shell
gcloud config set run/region us-east1
gcloud config set run/platform managed
```

## Setup pubsub

Create topic

```
TOPIC_NAME=cloudstorage-cloudrun-topic
gcloud pubsub topics create $TOPIC_NAME
```

Setup notification on file creation

```shell
BUCKET_PICTURES=uploaded-pictures-$GOOGLE_CLOUD_PROJECT
gsutil notification create -t $TOPIC_NAME -f json gs://$BUCKET_PICTURES
```

Setup service account

```shell
SERVICE_ACCOUNT=$TOPIC_NAME-sa
gcloud iam service-accounts create $SERVICE_ACCOUNT \
     --display-name "Cloud Run Pub/Sub Invoker"

gcloud run services add-iam-policy-binding $SERVICE_NAME \
   --member=serviceAccount:$SERVICE_ACCOUNT@$GOOGLE_CLOUD_PROJECT.iam.gserviceaccount.com \
   --role=roles/run.invoker

gcloud pubsub subscriptions create $TOPIC_NAME-subscription --topic $TOPIC_NAME \
   --push-endpoint=$SERVICE_URL \
   --push-auth-service-account=$SERVICE_ACCOUNT@$GOOGLE_CLOUD_PROJECT.iam.gserviceaccount.com   
```


