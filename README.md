# aws-serverless-example

1.	Clone the project from https://github.com/stanleyyuenyiu/aws-serverless-example

2.	Create a S3 bucket and Upload the lambda source code to the S3 bucket
```export AWS_DEFAULT_REGION=ap-southeast-1
export S3Bucket=stanleyyuen-restapi
aws s3 mb s3://$S3Bucket
aws s3 cp {git cloned path}/backend/ipdetection.zip s3://$S3Bucket/ipdetection.zip
```
3.	Create cloudformation stack
```
aws cloudformation create-stack --stack-name backend --template-body file://cloudformation.json --capabilities CAPABILITY_IAM --parameters ParameterKey=S3Bucket,ParameterValue=$S3Bucket
```
4.	Wait cloudformation stack finish
```
aws cloudformation wait stack-create-complete --stack-name backend
```
5.	Grep the output of cloudformation stack
```
aws cloudformation describe-stacks --stack-name backend --query Stacks[0].Outputs
```
6.	Expected output
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
7.	Open frontend configuration file -> {git cloned path}/frontend/dist/config.js, update below yellow  highlighted “OutputKey” from above corresponding “OutputValue”
```
var awsConfig = {
	API:{
		endpoints: [
            {
                name: "ips",
                endpoint: "{ApiBaseUrl}/",
                region: "{Region}"
            },
        ]
	},
	Auth:{
		identityPoolId: {CognitoId}', 
         region: {Region}'
	}
}
```
8.	Create cloudformation stack for frontend
aws cloudformation create-stack --stack-name frontend --template-body file:// {git cloned path}/cf/frontend.json --capabilities CAPABILITY_IAM 

9.	Wait cloudformation stack finish
aws cloudformation wait stack-create-complete --stack-name frontend

10.	Grep the output of cloudformation stack
aws cloudformation describe-stacks --stack-name frontend --query Stacks[0].Outputs

11.	Expected output


12.	Open folder {git cloned path}/frontend/dist/, upload files to S3, update below yellow highlighted “OutputKey” from above corresponding “OutputValue”
export S3BucketFrontend={BucketName}
aws s3 mb s3://$S3BucketFrontend
aws s3 cp {git cloned path}/frontend/index.html s3://$S3BucketFrontend/index.html
aws s3 cp {git cloned path}/frontend/config.js s3://$S3BucketFrontend/config.js
aws s3 cp {git cloned path}/frontend/index.html s3://$S3BucketFrontend/main.bundle.js

13.	Get Etag from current t cloudfront distribution, update below yellow highlighted “OutputKey” from above corresponding “OutputValue” of step 11
aws cloudfront get-distribution-config --id {CloudFrontId}

14.	Enable cloudfront from CLI, update below yellow highlighted “OutputKey” from above corresponding “OutputValue” of step 11
aws cloudfront update-distribution --id {CloudFrontId} --distribution-config file://{git cloned path}/distconfig-enable.json --if-match {CloudFrontEtag}

15.	Open browser and enter {CloudFrontUrl} to verify the application, open browser and enter {BucketUrlForOAIVerify} to verify the OAI restrict access of the S3
