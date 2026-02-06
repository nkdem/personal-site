# Cloudflare DNS Migration (from Njal.la)

Step-by-step guide for migrating `nkdem.net` DNS from Njal.la to Cloudflare.

## 1. Create Cloudflare account

- Sign up at [dash.cloudflare.com](https://dash.cloudflare.com)
- Select the **Free** plan

## 2. Add domain to Cloudflare

- Click "Add a site" and enter `nkdem.net`
- Select the Free plan
- Note the two assigned Cloudflare nameservers (e.g. `ada.ns.cloudflare.com`, `bob.ns.cloudflare.com`)

## 3. Export current Njal.la records

- Log in to Njal.la and note all existing DNS records for `nkdem.net`
- Keep a backup in case the Cloudflare auto-scan misses anything

## 4. Verify Cloudflare auto-scan

- Cloudflare will auto-scan existing records during setup
- Confirm all records were imported correctly
- Add any missing records manually

## 5. Update nameservers at Njal.la

- In Njal.la's domain settings, replace the current nameservers with Cloudflare's assigned nameservers
- This is the point of no return — DNS will start resolving through Cloudflare once propagation completes

## 6. Wait for propagation

- Propagation can vary; check with:
  ```bash
  dig nkdem.net NS
  ```
- Wait until Cloudflare nameservers appear in the response

## 7. Configure SSL/TLS mode

- In Cloudflare dashboard → SSL/TLS → Overview
- Set encryption mode to **Flexible** (origin serves HTTP on port 80, Cloudflare terminates TLS)

## 8. Enable "Always Use HTTPS"

- In Cloudflare dashboard → SSL/TLS → Edge Certificates
- Toggle on "Always Use HTTPS"

## 9. Create API token

- Go to Cloudflare → My Profile → API Tokens → Create Token
- Use the **Edit zone DNS** template
- Scope: Zone → DNS → Edit, for `nkdem.net` only
- Copy the token

## 10. Copy Zone ID

- Go to Cloudflare dashboard → `nkdem.net` → Overview
- Copy the **Zone ID** from the right sidebar

## 11. Store secrets in Pulumi

```bash
cd infrastructure
pulumi config set --secret cloudflareAPIToken <token>
pulumi config set cloudflareZoneId <zone-id>
```
