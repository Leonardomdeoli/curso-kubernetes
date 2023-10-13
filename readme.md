- Gerando imagem
````sh
docker build -t leonardomdeoli/hello-go-serve .
````

- Executar a imagem
````sh
docker run --rm -p 80:80 leonardomdeoli/hello-go-serve 
````

- Enviando docker hub
````sh
docker push leonardomdeoli/hello-go-serve 
````

Criando o cluster
````sh
kind create cluster
````
````sh
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.27.3) 🖼 
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Thanks for using kind! 😊
````

Criando o pod
````sh
kubectl apply -f k8s/pod.yaml
````
````sh
pod/goserver created
````

listando o pods
````sh
 kubectl get pod
````
````sh
NAME       READY   STATUS              RESTARTS   AGE
goserver   0/1     ContainerCreating   0          29s
````

Redirecinando a porta 80 para o 8080
````sh
kubectl port-forward pod/goserver 8080:80
````
````sh
Forwarding from 127.0.0.1:8080 -> 80
Forwarding from [::1]:8080 -> 80
Handling connection for 8080
Handling connection for 8080
````

Apagando pod
````sh
kubectl delete pod goserver
````
````sh
pod "goserver" deleted
````

Criando o pod
````sh
kubectl apply -f k8s/replicateset.yaml
````
````sh
pod/goserver created
````

````sh
kubectl get pods
````
````sh
NAME             READY   STATUS    RESTARTS   AGE
goserver-drqcw   1/1     Running   0          50s
goserver-f4tlk   1/1     Running   0          50s
````

Visualizando a replicas
````sh
kubectl get replicasets
````
````sh
NAME       DESIRED   CURRENT   READY   AGE
goserver   2         2         2       2m28s
````