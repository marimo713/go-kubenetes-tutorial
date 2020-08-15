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
docker push marimoex/go_kube
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

## Deploymentのデプロイ
```
kubectl create -f kubernetes/deployments/go-deployment.yaml
kubectl get deployments
kubectl get pods

kubectl create -f kubernetes/services/go-loadbarancer.yaml

kubectl delete -f kubernetes/deployments/go-deployment.yaml
kubectl delete -f kubernetes/services/go-loadbarancer.yaml 
```
http://localhost:8080/pingへアクセス

## readness liveness
```
kokinoMacBook-puro:go-kubenetes-tutorial koki$ kubectl get pods
NAME                   READY   STATUS    RESTARTS   AGE
app-7b6bbb9cfb-8jqk9   1/1     Running   0          63s
app-7b6bbb9cfb-thjjn   1/1     Running   0          63s
app-7b6bbb9cfb-tsvjq   1/1     Running   0          63s
```
http://localhost:8080/changestatusへアクセス

```
kokinoMacBook-puro:go-kubenetes-tutorial koki$ kubectl get pods
NAME                   READY   STATUS    RESTARTS   AGE
app-7b6bbb9cfb-8jqk9   1/1     Running   1          79s
app-7b6bbb9cfb-thjjn   1/1     Running   0          79s
app-7b6bbb9cfb-tsvjq   1/1     Running   0          79s
```
podが1つリスタートされているのがわかる

## リリース

```
docker build -t marimoex/go_kube:1.0.0 -f docker/Dockerfile .
docker push marimoex/go_kube:1.0.0
```

main.goファイルを書き換えて
```
docker build -t marimoex/go_kube:2.0.0 -f docker/Dockerfile .
docker push marimoex/go_kube:2.0.0
```

```
kubectl apply -f kubernetes/deployments/go-deployment-v1.yaml
kubectl get deployments
kubectl get replicasets
kubectl get pods

kubectl create -f kubernetes/services/go-loadbarancer.yaml
```
変更前にアクセス

```
kubectl apply -f kubernetes/deployments/go-deployment-v2.yaml
kubectl rollout status deployment app

kubectl get deployments
kubectl get replicasets
kubectl get pods
```

アクセスすると変更後の情報に変わる

```
kubectl rollout undo deployment app
```
アクセスすると戻っている

```

kubectl delete -f kubernetes/deployments/go-deployment-v2.yaml
kubectl delete -f kubernetes/services/go-loadbarancer.yaml
```


## Istio
### Getting start
最新の手順は以下を参照  
https://istio.io/docs/setup/getting-started/ 
#### ダウンロード 
istioのダウンロード
```
curl -L https://istio.io/downloadIstio | sh -
```
ダウンロードしたフォルダに移動
```
cd istio-1.4.5
```
PATHの追加
```
export PATH=$PWD/bin:$PATH
```
#### インストール
demoプロファイルでインストール  
※本番用では使用してはいけない   
```
istioctl manifest apply -f kubernetes/istio/istio.yaml
```
インストールの確認
```
$ kubectl get svc -n istio-system
NAME                     TYPE           CLUSTER-IP       EXTERNAL-IP     PORT(S)                                                                                                                                      AGE
grafana                  ClusterIP      172.21.211.123   <none>          3000/TCP                                                                                                                                     2m
istio-citadel            ClusterIP      172.21.177.222   <none>          8060/TCP,15014/TCP                                                                                                                           2m
istio-egressgateway      ClusterIP      172.21.113.24    <none>          80/TCP,443/TCP,15443/TCP                                                                                                                     2m
istio-galley             ClusterIP      172.21.132.247   <none>          443/TCP,15014/TCP,9901/TCP                                                                                                                   2m
istio-ingressgateway     LoadBalancer   172.21.144.254   52.116.22.242   15020:31831/TCP,80:31380/TCP,443:31390/TCP,31400:31400/TCP,15029:30318/TCP,15030:32645/TCP,15031:31933/TCP,15032:31188/TCP,15443:30838/TCP   2m
istio-pilot              ClusterIP      172.21.105.205   <none>          15010/TCP,15011/TCP,8080/TCP,15014/TCP                                                                                                       2m
istio-policy             ClusterIP      172.21.14.236    <none>          9091/TCP,15004/TCP,15014/TCP                                                                                                                 2m
istio-sidecar-injector   ClusterIP      172.21.155.47    <none>          443/TCP,15014/TCP                                                                                                                            2m
istio-telemetry          ClusterIP      172.21.196.79    <none>          9091/TCP,15004/TCP,15014/TCP,42422/TCP                                                                                                       2m
jaeger-agent             ClusterIP      None             <none>          5775/UDP,6831/UDP,6832/UDP                                                                                                                   2m
jaeger-collector         ClusterIP      172.21.135.51    <none>          14267/TCP,14268/TCP                                                                                                                          2m
jaeger-query             ClusterIP      172.21.26.187    <none>          16686/TCP                                                                                                                                    2m
kiali                    ClusterIP      172.21.155.201   <none>          20001/TCP                                                                                                                                    2m
prometheus               ClusterIP      172.21.63.159    <none>          9090/TCP                                                                                                                                     2m
tracing                  ClusterIP      172.21.2.245     <none>          80/TCP                                                                                                                                       2m
zipkin                   ClusterIP      172.21.182.245   <none>          9411/TCP 
```
podが全てRunningになっていることを確認する
```
$ kubectl get pods -n istio-system
NAME                                                           READY   STATUS      RESTARTS   AGE
grafana-f8467cc6-rbjlg                                         1/1     Running     0          1m
istio-citadel-78df5b548f-g5cpw                                 1/1     Running     0          1m
istio-egressgateway-78569df5c4-zwtb5                           1/1     Running     0          1m
istio-galley-74d5f764fc-q7nrk                                  1/1     Running     0          1m
istio-ingressgateway-7ddcfd665c-dmtqz                          1/1     Running     0          1m
istio-pilot-f479bbf5c-qwr28                                    1/1     Running     0          1m
istio-policy-6fccc5c868-xhblv                                  1/1     Running     2          1m
istio-sidecar-injector-78499d85b8-x44m6                        1/1     Running     0          1m
istio-telemetry-78b96c6cb6-ldm9q                               1/1     Running     2          1m
istio-tracing-69b5f778b7-s2zvw                                 1/1     Running     0          1m
kiali-99f7467dc-6rvwp                                          1/1     Running     0          1m
prometheus-67cdb66cbb-9w2hm  
```
#### スタート
アプリの起動
```
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -n default -f kubernetes/deployments/go-deployment.yaml
$ kubectl apply -n default -f kubernetes/services/go-clusterip.yaml
```
Dashbord
```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-rc5/aio/deploy/recommended.yaml
kubectl proxy &
kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
```
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login

Jaerger
```
istioctl dashboard jaeger
```
Kiali
```
istioctl dashboard kiali
```
Grafana
http://localhost:3000/dashboard/db/istio-mesh-dashboard
 ```
 kubectl -n istio-system port-forward $(kubectl -n istio-system get pod -l app=grafana -o jsonpath='{.items[0].metadata.name}') 3000:3000 &
 ```


