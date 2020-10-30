# Kubernetes

#### Install K3d
```
brew install k3d
```

#### Initialize Local K3s Cluster
```
k3d cluster create cloudy-sky --api-port 6550 -p "8081:80@loadbalancer"
```
 k3d cluster create cloudy-sky \                          
--api-port 6550 \
-p "8081:80@loadbalancer"


####

