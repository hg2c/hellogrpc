1. 部署 hellogrpc

   kubectl apply -f hellogrpc.yaml
   
2. 查看 hellogrpc-client 日志，现在流量会平均分配到两个版本，即 Nihao 和 Hello 会交替显示

   kubectl logs -f $(kubectl get pods -l app=hellogrpc-client -o 'jsonpath={.items[0].metadata.name}') hellogrpc-client
   
3. 把流量全部切换到 v1, client 现在日志输出全部为 hello

   kubectl apply -f hellogrpc-v1.yaml

4. 把流量全部切换到 v2, client 现在日志输出全部为 nihao

   kubectl apply -f hellogrpc-v2.yaml
