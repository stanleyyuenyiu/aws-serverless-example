import API  from '@aws-amplify/api';
import Auth  from '@aws-amplify/auth';
// retrieve temporary AWS credentials and sign requests
API.configure({
        endpoints: window.__initialState.awsConfig.API.endpoints
});
Auth.configure(window.__initialState.awsConfig.Auth);

let click = false
async function callAPI(){
	let apiName = window.__initialState.awsConfig.API.endpoints[0].name;
	let path = '/'+apiName; 
	let myInit = { // OPTIONAL
	    headers: {}, // OPTIONAL
	    response: true
	}
	return await API.get(apiName, path, myInit).then(response => {
	    try{
	    //	var obj = JSON.parse(response)
	    	var obj = response.data
	    	var str = "Your IP is:" +  obj.IP[0];
	    	str += "<br/>Proxy IP is:" +  obj.IP[1];
	    	document.getElementById("Result").innerHTML = str;
	    }catch(e){
	    	console.log(e);
	    	document.getElementById("Result").innerHTML = "Fail to fetch IP, please retry";
	    }
	    click = false;
	}).catch(error => {
		document.getElementById("Result").innerHTML = "Fail to fetch IP, please retry";
	    console.log(error.response)
	    click = false;
	});
}

callAPI();


const btn = document.getElementById('Retry');

Retry.addEventListener('click', (evt) => {
	if(click) return;
	document.getElementById("Result").innerHTML = "Loading your IP...";
	click = true;
	callAPI();
});