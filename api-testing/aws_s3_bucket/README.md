# AWS S3 bucket handler

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Cli applcation for managing AWS S3 buckets.  

## References

- [Sample HTTP Request for general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
- [Signing and authenticating REST requests](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html)

## How to use

### Preconditions

You need an 
- AWS subscription
- the `AWS_ACCESS_KEY_ID` environment variable
- the `AWS_SECRET_ACCESS_KEY` environment variable

### Build the cli application

On Windows OS:

```sh
 go build -o aws_s3_bucket_handler.exe
```

On Unix systems (MacOS or Ubuntu >=18.04/debian >=11):

```sh
 go build -o aws_s3_bucket_handler
```

### Run the cli application

```sh
# [C]reate an aws s3 bucket
aws_s3_bucket_handler bucket create -b <bucket_name> -k <AWS_ACCESS_KEY_ID> -s <AWS_SECRET_ACCESS_KEY>

# [D]elete an aws s3 bucket 
aws_s3_bucket_handler bucket delete -b <bucket_name> -k <AWS_ACCESS_KEY_ID> -s <AWS_SECRET_ACCESS_KEY>
```