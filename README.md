Please select either one of the deployment step of the application as below

# Auto Deploy & Installation

## prerequisite
Install aws-cli 
https://docs.aws.amazon.com/en_us/cli/latest/userguide/cli-chap-install.html

Install jq
https://stedolan.github.io/jq/download/

Configurate default aws profile
https://docs.aws.amazon.com/en_us/cli/latest/userguide/cli-chap-configure.html

## How to use
Navigate to the project clone path, run below command, it will use your default aws profile to execute
```
sh run.sh
```
Optional: 
To execute with other aws profile, replace below "[my_profile_name]" with your aws profile
```
sh run.sh [my_profile_name]
```


# Manual Deploy & Installation

1. Clone the project from https://github.com/stanleyyuenyiu/aws-serverless-example
2. Export the cloned path var, replace {git cloned path} to the path where you clone the project
```
export GITClonedPath={git cloned path}
```
3. Create a S3 bucket and Upload the lambda source code to the S3 bucket
```
export AWS_DEFAULT_REGION=ap-southeast-1
export S3Bucket=backendbucket-api
aws s3 mb s3://$S3Bucket
aws s3 cp $GITClonedPath/backend/ipdetection/ipdetection.zip s3://$S3Bucket/ipdetection.zip
```
4. Create cloudformation stack
```
aws cloudformation create-stack --stack-name backend --template-body file://$GITClonedPath/cf/backend.json --capabilities CAPABILITY_IAM --parameters ParameterKey=S3Bucket,ParameterValue=$S3Bucket
```
5. Wait cloudformation stack finish
```
aws cloudformation wait stack-create-complete --stack-name backend
```
6. Grep the output of cloudformation stack
```
aws cloudformation describe-stacks --stack-name backend --query Stacks[0].Outputs
```
7. Expected output
```
[
    {
        "OutputKey": "ApiBaseUrl",
        "OutputValue": "https://14ej5gtezk.execute-api.ap-southeast-1.amazonaws.com/v1"
    },
    {
        "OutputKey": "Region",
        "OutputValue": "ap-southeast-1"
    },
    {
        "OutputKey": "CognitoId",
        "OutputValue": "ap-southeast-1:546214db-60ab-40c4-9a02-cdb4bdfb22ac"
    },
    {
        "OutputKey": "ApiId",
        "OutputValue": "14ej5gtezk"
    },
    {
        "OutputKey": "ApiPath",
        "OutputValue": "ip"
    }
]
```
8.Open frontend configuration file -> $GITClonedPath/frontend/dist/config.js, update {ApiBaseUrl},{Region},{CognitoId} to "OutputValue" from above corresponding “OutputKey”
```javascript
var awsConfig = {
	API:{
		endpoints: [
            {
                name: "ips",
                endpoint: "{ApiBaseUrl}",
                region: "{Region}"
            },
        ]
	},
	Auth:{
		identityPoolId: "{CognitoId}", 
         	region: "{Region}"
	}
}
```
9. Create cloudformation stack for frontend
```
aws cloudformation create-stack --stack-name frontend --template-body file://$GITClonedPath/cf/frontend.json --capabilities CAPABILITY_IAM 
```
10. Wait cloudformation stack finish
```
aws cloudformation wait stack-create-complete --stack-name frontend
```
11. Grep the output of cloudformation stack
```
aws cloudformation describe-stacks --stack-name frontend --query Stacks[0].Outputs
```
12. Expected output
```
[
    {
        "OutputKey": "BucketUrlForOAIVerify",
        "OutputValue": "https://s3-ap-southeast-1.amazonaws.com/frontend-frontendbucket-9zdf7nn66utp/index.html"
    },
    {
        "OutputKey": "BucketName",
        "OutputValue": "frontend-frontendbucket-9zdf7nn66utp"
    },
    {
        "OutputKey": "CloudFrontUrl",
        "OutputValue": "https://ddcbus9dh2gvn.cloudfront.net/index.html"
    },
    {
        "OutputKey": "CloudFrontId",
        "OutputValue": "EDK6SJ05GH29V"
    }
]
```
13. Upload files to S3, update {BucketName} to "OutputValue" from above corresponding “OutputKey”
```
export S3BucketFrontend={BucketName}
aws s3 cp $GITClonedPath/frontend/dist/index.html s3://$S3BucketFrontend/index.html
aws s3 cp $GITClonedPath/frontend/dist/config.js s3://$S3BucketFrontend/config.js
aws s3 cp $GITClonedPath/frontend/dist/main.bundle.js s3://$S3BucketFrontend/main.bundle.js
```
14.	Open browser and enter {CloudFrontUrl} to verify the application, open browser and enter {BucketUrlForOAIVerify} to verify the OAI restrict access of the S3
