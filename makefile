#
# Title:makefile
#
# Description:
#   'make clean' removes all core and object files
#   'make ultraclean' removes all executables
#
# Operating System/Environment:
#   Ubuntu 18.04.3 LTS (Bionic Beaver)
#   GNU Make 4.1
#
# Author:
#   G.S. Cole (guycole at gmail dot com)
#
DOCKER = docker
DARING_CHUPACABRA_FRONT_END = daring-chupacabra-fe:1
DARING_CHUPACABRA_BACK_END = daring-chupacabra-be:1
HELM = helm
KUBECTL = kubectl
MINIKUBE = minikube

be_build:
	cd chupacabra; $(DOCKER) build . -f backend.Dockerfile -t $(DARING_CHUPACABRA_BACK_END)

be_delete:
	$(KUBECTL) delete -f infra/be-deploy.yaml

be_deploy:
	$(KUBECTL) apply -f infra/be-deploy.yaml

minikube_reset:
	$(MINIKUBE) stop
	$(MINIKUBE) delete

minikube_start:
	cd infra; ./start_minikube.sh

minikube_setup:
	$(KUBECTL) apply -f infra/namespace.yaml
	$(MINIKUBE) addons enable ingress
	$(HELM) repo add stable https://charts.helm.sh/stable
	$(HELM) repo update

monitoring_delete:
	$(KUBECTL) delete -f infra/redis-dashboard.yaml -n monitoring
	$(HELM) uninstall prometheus -n monitoring

monitoring_deploy:
	$(KUBECTL) apply -f infra/redis-dashboard.yaml -n monitoring
	$(HELM) repo add prometheus-community https://prometheus-community.github.io/helm-charts
	$(HELM) upgrade --debug --install prometheus prometheus-community/kube-prometheus-stack -n monitoring --version 19.0.2 --values infra/kube-prometheus.yaml

monitoring_expose:
	$(KUBECTL) expose service prometheus-kube-prometheus-alertmanager --type=NodePort --target-port=9093 --name=prometheus-alertmanager-np --namespace=monitoring
	$(KUBECTL) expose service prometheus-kube-prometheus-prometheus --type=NodePort --target-port=9090 --name=prometheus-np --namespace=monitoring
	$(KUBECTL) expose service prometheus-grafana --type=NodePort --target-port=3000 --name=grafana-np --namespace=monitoring

redis_deploy:
	$(KUBECTL) apply -f infra/redis-secret.yaml -n chupacabra
	$(HELM) repo add bitnami https://charts.bitnami.com/bitnami
	$(HELM) upgrade --debug --install redis bitnami/redis -n chupacabra --version 15.7.6 --values infra/redis-minikube.yaml

worker_build:
	cd worker; $(DOCKER) build . -t $(DARING_CYCLOPS_WORKER)

worker_delete:
	$(KUBECTL) delete -f infra/worker-deploy.yaml -n cyclops-app

worker_deploy:
	$(KUBECTL) apply -f infra/worker-deploy.yaml -n cyclops-app
