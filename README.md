# aws-serverless-example

1. Clone the project from https://github.com/stanleyyuenyiu/aws-serverless-example
2. Export the cloned path var
```
export GITClonedPath={git cloned path}
```
3. Create a S3 bucket and Upload the lambda source code to the S3 bucket
```
export AWS_DEFAULT_REGION=ap-southeast-1
export S3Bucket=stanleyyuen-restapi
aws s3 mb s3://$S3Bucket
aws s3 cp $GITClonedPath/backend/ipdetection.zip s3://$S3Bucket/ipdetection.zip
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
                endpoint: "{ApiBaseUrl}/",
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
```aws cloudformation wait stack-create-complete --stack-name frontend
```
11. Grep the output of cloudformation stack
```
aws cloudformation describe-stacks --stack-name frontend --query Stacks[0].Outputs
```
12. Expected output
13. Upload files to S3, update {BucketName} to "OutputValue" from above corresponding “OutputKey”
```
export S3BucketFrontend={BucketName}
aws s3 mb s3://$S3BucketFrontend
aws s3 cp $GITClonedPath/frontend/index.html s3://$S3BucketFrontend/index.html
aws s3 cp $GITClonedPath/frontend/config.js s3://$S3BucketFrontend/config.js
aws s3 cp $GITClonedPath/frontend/index.html s3://$S3BucketFrontend/main.bundle.js
```
14.	Get Etag from current cloudfront distribution, update {CloudFrontId}  to "OutputValue" from above corresponding “OutputKey” at step 12
```
aws cloudfront get-distribution-config --id {CloudFrontId}
```
15.	Enable cloudfront from CLI, update {CloudFrontId} to "OutputValue" from above corresponding “OutputKey” at step 12, update {Etag} to the "Etag" value from step 14
```
export cloudfrontEtag={Etag}
aws cloudfront update-distribution --id {CloudFrontId} --distribution-config file://$GITClonedPath/distconfig-enable.json --if-match $cloudfrontEtag
```
16.	Open browser and enter {CloudFrontUrl} to verify the application, open browser and enter {BucketUrlForOAIVerify} to verify the OAI restrict access of the S3
