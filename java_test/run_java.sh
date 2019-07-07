#export KAFKA_JMX_OPTS="-Dcom.sun.management.jmxremote=true -Dcom.sun.management.jmxremote.authenticate=false
#-Dcom.sun.management.jmxremote.ssl=false -Djava.rmi.server.hostname=jmx_exporter -Djava.net.preferIPv4Stack=true
#-Dcom.sun.management.jmxremote.rmi.port=9096 -Dcom.sun.management.jmxremote.port=9096 -Dcom.sun.management.jmxremote.local.only=true"

javac test.java 

export JAVA_OPTS="-javaagent:/Users/tyron/Desktop/Work/sncf/Benchmark/docker-compose-influxdb-grafana/jmx_http_exporter/jmx_prometheus_javaagent-0.3.0.jar=7070:/Users/tyron/Desktop/Work/sncf/Benchmark/docker-compose-influxdb-grafana/jmx_http_exporter/jmx-config.yml" java Simple
