# serverless-calculator

## Description

Backend of serverless-calculator project (just a regular calculator, but serverless, because why not?)

[Archived Executable](https://github.com/ranggarppb/serverless-calculator/suites/16157713024/artifacts/922573855):
- linux/amd64
- darwin/amd64
- windows/amd64

Serverless Function App:
- Backend server (Google Cloud Function): https://asia-southeast2-serverless-calculator.cloudfunctions.net/serverless-calculator/calculation
- Frontend server (Firebase Hosting): https://serverless-calculator.web.app/
- Swagger documentation: https://app.swaggerhub.com/apis/RANGGAPUTRAPERTAMAPP/serverless-calculator/1.0.0

## How To Run Executable
For now executable can be run for **console app** and **function app**. Run it with `console` / `function` argument
- To run console app, go to the executable folder and run command `./serverless-calculator console` and type the operation input
- To run function app, go to the executable folder and run command `./serverless-calculator function` and try to hit the cURL specified in [Features](#features)

## How To Compile
If you dont find your OS architecture in [Archived Executable](https://github.com/ranggarppb/serverless-calculator/suites/16157713024/artifacts/922573855), you can run command `make build` to build your own executable (already tested for `darwin/arm64`)

## How To Test In Local Without Compiling
- Testing in console: run the command `make console`
- Testing with local HTTP function `make local_function`

## Collaborating

- Fork and develop in your local computer, dont forget to run `. setup-pre-commit-hooks.sh` to setup pre-commit-hooks
- Open Pull Request, make sure [the Test Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/pull-request.yaml)  passed
- You can deploy manually from your branch with [Test Env Deployment Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/deploy-test.yaml)
- To make sure your Pull Request doesn't introduce significant performance degradation, run [Load Test Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/load-test.yaml) (artifact of last production-env-mocking server could be downloaded [here](https://github.com/ranggarppb/serverless-calculator/suites/16253957144/artifacts/928137578))

## Features
- For Console app, currently the feature of latest executable:
	- Supporting without input calculation `abs`, `neg`, `sqr`, `sqrt`, `cube`, `cubert`, ex: typing `sqr`
	- Supporting multiple input calculation `add`, `subtract`, `multiply`, `divide`, ex: typing `add 3 multiply 5`
	- The console start with `0`
- For local HTTP function app, currently the feature of latest production deployment:
	- Supporting single input calculation `abs`, `neg`, `sqr`, `sqrt`, `cube`, `cubert`
	```
	curl --location --request POST 'localhost:8080/calculation' \
		--header 'Content-Type: application/json' \
		--data-raw '{
    		"input": "cubert -2"
		}'
	```
	- Supporting multiple input calculation `add`, `subtract`, `multiply`, `divide`
	```
	curl --location --request POST 'localhost:8080/calculation' \
		--header 'Content-Type: application/json' \
		--data-raw '{
    		"input": "1 add 2 multiple 5"
		}'
	```
- For deployed HTTP function app, currently the feature of latest production deployment, its the same with local function just change URL to `https://asia-southeast2-serverless-calculator.cloudfunctions.net/serverless-calculator/calculation``

In Progress:
- Multiple input calculation using stack data structure
- Saving calculation history in Firestore
