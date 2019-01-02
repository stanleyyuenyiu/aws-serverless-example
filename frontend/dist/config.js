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