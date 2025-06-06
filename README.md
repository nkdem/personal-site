# Personal Site
[![Build Project](https://github.com/nkdem/personal-site/actions/workflows/docker-build-push.yaml/badge.svg)](https://github.com/nkdem/personal-site/actions/workflows/docker-build-push.yaml)

The source code to my (unfinished) [website](https://nkdem.net).

## Technology Stack
*   **Frontend:**
    *   [Eleventy](https://www.11ty.dev/) (Static Site Generator)
    *   [Tailwind CSS](https://tailwindcss.com/) (CSS Framework)
    *   [Nunjucks](https://mozilla.github.io/nunjucks/) (Template Engine)
*   **Containerization:**
    *   [Docker](https://www.docker.com/) (Containerization Platform)
    *   [Caddy](https://caddyserver.com/) (Web Server for Static Files)
*   **CI/CD & GitOps:**
    *   [Argo CD](https://argo-cd.readthedocs.io/) (GitOps Continuous Delivery)
    *   [GitHub Container Registry (GHCR)](https://github.com/features/packages) (Docker Image Registry)
    *   [GitHub Actions](https://github.com/features/actions) 
        * On new commits to `main` branch:
            * Build Docker image
            * Push Docker image to [GHCR](https://github.com/nkdem/personal-site/pkgs/container/personal-site)
            * Update Kubernetes deployment with new image (Argo CD then syncs)
            * Verify that the deployment is healthy
*   **Kubernetes & Infrastructure:**
    *   [Kubernetes (k3s)](https://k3s.io/) (Lightweight Kubernetes Distribution)
    *   [Hetzner Cloud](https://www.hetzner.com/cloud) (Cloud Infrastructure Provider)
    *   [Traefik](https://traefik.io/) (Ingress Controller - built into k3s)
    *   [Cert-Manager](https://cert-manager.io/) (Automated TLS Certificate Management)
    *   [Let's Encrypt](https://letsencrypt.org/) (Free SSL Certificates)

Running on a Hetzner VPS with K3s. (See [repo for my k8 infra](https://github.com/nkdem/k8-infra))
