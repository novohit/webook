# webook

## etcdctl


```bash
cd ./internal/config
# unix
etcdctl put "/webook" "$(cat dev.yaml)"
# powershell
etcdctl put "/webook" "$(Get-Content -Raw -Path 'dev.yaml')"
```