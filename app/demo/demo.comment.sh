#start
ENV=local sh tool.sh demo start
ENV=dev NACOS_ADDR=127.0.0.1 NACOS_PORT=8848 NACOS_NAMESPACE_ID=sre-test NACOS_GROUP=SRE NACOS_DATA_IDS=demo.yaml;registry.yaml;auth.yaml sh tool.sh demo start