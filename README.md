# serverless-calculator

## Description

Backend of serverless-calculator project (just a regular calculator, but serverless, because why not?)

- Backend server (Google Cloud Function): https://asia-southeast2-serverless-calculator.cloudfunctions.net/serverless-calculator/calculation
- Frontend server (Firebase Hosting): https://serverless-calculator.web.app/
- Swagger documentation: https://app.swaggerhub.com/apis/RANGGAPUTRAPERTAMAPP/serverless-calculator/1.0.0

## Collaborating

- Fork and develop in your local computer, dont forget to run `. setup-pre-commit-hooks.sh` to setup pre-commit-hooks
- Open Pull Request, make sure [the Test Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/pull-request.yaml)  passed
- You can deploy manually from your branch with [Test Env Deployment Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/deploy-test.yaml)
- To make sure your Pull Request doesn't introduce significant performance degradation, run [Load Test Workflow](https://github.com/ranggarppb/serverless-calculator/actions/workflows/load-test.yaml)(artifact of last production-env-mocking server could be downloaded [here](https://github.com/ranggarppb/serverless-calculator/suites/15880973846/artifacts/906268580))

## Features
Currently the feature of latest production deployment:
- Supporting single input calculation `abs`, `neg`, `sqr`, `sqrt`, ex: `neg 3`
- Supporting two input calculation `add`, `substract`, `multiply`, `divide`
