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