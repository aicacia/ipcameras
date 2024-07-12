# IPCameras API

### Deploy

- `docker build -t ghcr.io/aicacia/ipcameras-api:latest .`
- `docker push ghcr.io/aicacia/ipcameras-api:latest`
- `helm upgrade ipcameras-api helm/ipcameras-api -n api --install -f values.yaml --set image.hash="$(docker inspect --format='{{index .Id}}' ghcr.io/aicacia/ipcameras-api:latest)"`

### Undeploy

- `helm delete -n api ipcameras-api`
