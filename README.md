# Device Configuration Microservice
Build Status: ![](https://github.com/UdamLiyanage/device-configuration-service/workflows/Go/badge.svg)
***
Device Configuration Microservice for IoT Platform. This service is responsible for the following features:
*Create configurations
*Read configurations
*Update configurations
*Delete configurations

***
## Document Structure for Configuration - Revision 1
Below is the structure of the JSON document that holds user data
```
{
	"device_token": "Device Token",
	"device_serial": "Device Serial",
	"configurations_type": "email",
	"configuration": {
		"to": "receiver@email.com",
		"from": "sender@email.com",
		"subject": "Email Subject",
		"body": "Test Body for Email"
	}
}
```
