# End-to-End DevOps for Go Web Application
### CI/CD with GitHub Actions, Docker, Kubernetes (EKS), Helm & ArgoCD

This project demonstrates a production-style DevOps workflow for a Go web application, implementing complete automation from code to Kubernetes using modern DevOps and GitOps tools.

The pipeline covers:

- Application build & containerization

- Kubernetes deployment on AWS EKS

- Helm-based packaging

- CI using GitHub Actions

- CD using ArgoCD (GitOps)

## Architecture Overview

Flow:
<img width="241" height="596" alt="Untitled Diagram drawio (5)" src="https://github.com/user-attachments/assets/bab60ce8-f637-4da2-9efc-c786f9047b7a" />


## Tech Stack

- Language: Go

- Containerization: Docker (Multi-stage build)

- CI: GitHub Actions

- Container Registry: Docker Hub

- Orchestration: Kubernetes (AWS EKS)

- Ingress: NGINX Ingress Controller

- Packaging: Helm

- CD (GitOps): ArgoCD

- Cloud: AWS

### What This Project Solves

Manual Kubernetes deployments are complex and error-prone.
This project provides:

- Fully automated CI/CD pipeline

- GitOps-based deployments

- Reusable Helm-based packaging

- Zero manual kubectl apply
  
## Screenshots
<img width="1554" height="723" alt="fb4b26f4-9ac2-48e9-b542-16fe295955a5" src="https://github.com/user-attachments/assets/832a7e64-2b3e-4159-b1f0-4c80d2de474c" />

<img width="1586" height="640" alt="766cf9b3-cde7-4ceb-a1ef-63957e9fa4ee" src="https://github.com/user-attachments/assets/dea7b48f-7299-4595-abea-0f31c33b9d2e" />

<img width="1720" height="945" alt="eb65f75b-1daf-4f77-8f75-590791730cd8" src="https://github.com/user-attachments/assets/e3cfe3e8-f811-418c-a9ff-17ce755fad03" />






Step-by-step instructions and detailed explanations are available in the article :[Click here](https://anjumohan.hashnode.dev/end-to-end-devops-for-a-go-web-app-cicd-with-github-actions-kubernetes-argocd)
