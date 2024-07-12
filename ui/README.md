# IPCameras UI

## Tools

- [nvm](https://github.com/nvm-sh/nvm#installing-and-updating)
- [pnpm](https://pnpm.io/installation)
- [tailwind css](https://tailwindcss.com/docs)
- [svelte kit](https://kit.svelte.dev/docs)
- [icons](https://lucide.dev/icons/)

## Docker/Helm

### Deploy

- `docker build -t ghcr.io/aicacia/ipcameras-ui:latest .`
- `docker push ghcr.io/aicacia/ipcameras-ui:latest`
- `helm upgrade ipcameras-ui helm/ipcameras-ui -n ui --install -f values.yaml --set image.hash="$(docker inspect --format='{{index .Id}}' ghcr.io/aicacia/ipcameras-ui:latest)"`

### Undeploy

- `helm delete -n ui ipcameras-ui`
