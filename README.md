# Personal Site
[![Build Project](https://github.com/nkdem/personal-site/actions/workflows/docker-build-push.yaml/badge.svg)](https://github.com/nkdem/personal-site/actions/workflows/docker-build-push.yaml)

The source code to my [website](https://nkdem.net).

## Technology Stack
*   **Frontend:**
    *   [Eleventy](https://www.11ty.dev/) (Static Site Generator)
    *   [Tailwind CSS](https://tailwindcss.com/) (CSS Framework)
    *   [Nunjucks](https://mozilla.github.io/nunjucks/) (Template Engine)
*   **Containerization:**
    *   [Docker](https://www.docker.com/) (Containerization Platform)
    *   [Caddy](https://caddyserver.com/) (Web Server)
*   **CI/CD:**
    *   [GitHub Actions](https://github.com/features/actions) — build, push to [GHCR](https://github.com/nkdem/personal-site/pkgs/container/personal-site), deploy via Tailscale SSH
*   **Infrastructure:**
    *   [Hetzner Cloud](https://www.hetzner.com/cloud) (VPS)
    *   [Pulumi](https://www.pulumi.com/) (Infrastructure as Code)
    *   [Cloudflare](https://www.cloudflare.com/) (DNS, TLS termination)
    *   [Tailscale](https://tailscale.com/) (SSH access)
