# Retail Eshop 

This project is inspired by AWS's retail container shop examples, designed to showcase various concepts of containerization on AWS. It features a sample retail store application with functionalities like product catalog, shopping and checkout.

It provides:

- A distributed component architecture in various languages and frameworks
- Utilization of a variety of different persistence backends for different components like MySQL, DynamoDB and Redis
- The ability to run in various container orchestration technologies like Docker Compose, Kubernetes etc.
- Pre-built containers image for both x86-64 and ARM64 CPU architectures
- All components instrumented for Prometheus metrics and OpenTelemetry OTLP tracing
- Support for Istio on Kubernetes
- Load generator which exercises all of the infrastructure

**This project is intended for educational purposes only and not for production use.**

![Screenshot](/docs/images/screenshot.png)

## Application Architecture

The application has been deliberately over-engineered to generate multiple de-coupled components. These components generally have different infrastructure dependencies, and may support multiple "backends" (example: Carts service supports MongoDB or DynamoDB).

![Architecture](/docs/images/KienTrucMicroServices.png)

| Component                                                                                                                   | Language | Container Image                                                             | Description                                                                 |
| --------------------------------------------------------------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| ![ui workflow](https://github.com/aws-containers/retail-store-sample-app/actions/workflows/ci-ui.yml/badge.svg)             | Java     | [Link](https://gallery.ecr.aws/aws-containers/retail-store-sample-ui)       | Aggregates API calls to the various other services and renders the HTML UI. |
| ![catalog workflow](https://github.com/aws-containers/retail-store-sample-app/actions/workflows/ci-catalog.yml/badge.svg)   | Go       | [Link](https://gallery.ecr.aws/aws-containers/retail-store-sample-catalog)  | Product catalog API                                                         |                                                   |
| ![orders workflow](https://github.com/aws-containers/retail-store-sample-app/actions/workflows/ci-orders.yml/badge.svg)     | Java     | [Link](https://gallery.ecr.aws/aws-containers/retail-store-sample-orders)   | User orders API                                                             |
| ![checkout workflow](https://github.com/aws-containers/retail-store-sample-app/actions/workflows/ci-checkout.yml/badge.svg) | Node     | [Link](https://gallery.ecr.aws/aws-containers/retail-store-sample-checkout) | API to orchestrate the checkout process                                     |
| ![assets workflow](https://github.com/aws-containers/retail-store-sample-app/actions/workflows/ci-assets.yml/badge.svg)     | Nginx    | [Link](https://gallery.ecr.aws/aws-containers/retail-store-sample-assets)   | Serves static assets like images related to the product catalog             |

## Quickstart

The following sections provide quickstart instructions for various platforms. All of these assume that you have cloned this repository locally and are using a CLI thats current directory is the root of the code repository.

### Kubernetes

This deployment method will run the application in an existing Kubernetes cluster.

Pre-requisites:

- Kubernetes cluster
- `kubectl` installed locally

Use `kubectl` to run the application:

```
kubectl apply -f https://raw.githubusercontent.com/aws-containers/retail-store-sample-app/main/dist/kubernetes/deploy.yaml
kubectl wait --for=condition=available deployments --all
```

Get the URL for the frontend load balancer like so:

```
kubectl get svc ui
```
