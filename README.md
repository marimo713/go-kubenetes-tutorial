# go-kubenetes-tutorial

## Docker Hubにログイン
```
docker login
```

## Dockerイメージのビルド
```
docker build -t marimoex/go_kube -f docker/Dockerfile .
```

## Dockerイメージのpush
```
docker build -t marimoex/go_kube -f docker/Dockerfile .
```

## podのデプロイ
```
kubectl create -f kubernetes/pods/go-pod.yaml
kubectl get pods

kubectl port-forward go-kube 8080:8080
kubectl logs -f go-kube

kubectl delete -f kubernetes/pods/go-pod.yaml
kubectl get pods
```

## Serviceのデプロイ
```
kubectl create -f kubernetes/services/go-nodeport.yaml
kubectl create -f kubernetes/pods/go-pod.yaml
```
http://localhost:31000/pingへアクセス
