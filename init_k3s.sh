#!/usr/bin/env bash

/usr/local/bin/k3s-killall.sh | true
/usr/local/bin/k3s-uninstall.sh | true
curl -sfL https://get.k3s.io | sh -

#curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --no-deploy traefik" sh

sudo cat /etc/rancher/k3s/k3s.yaml > ~/.kube/local.yaml
export KUBECONFIG=~/.kube/local.yaml

sleep 30
k3s kubectl get node

docker run -d -p 5000:5000 --restart=always --name registry registry:2

curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64
sudo install skaffold /usr/local/bin/
# or
#docker run gcr.io/k8s-skaffold/skaffold:latest skaffold -h

#skaffold init --force
