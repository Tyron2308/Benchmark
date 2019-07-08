FROM openjdk:8-jre-alpine

ENV	EXPORTER_HTTP_PORT=${JMX_EXPORTER_HTTP_PORT}
ENV REMOTE_PORT=${JMX_REMOTE_PORT}
ENV EXPORTER_RMI_PORT=${JMX_REMOTE_PORT}

EXPOSE ${EXPORTER_HTTP_PORT}
EXPOSE ${REMOTE_PORT}
EXPOSE ${EXPORTER_RMI_PORT}

RUN mkdir /jmx-exporter

WORKDIR /jmx-exporter

ADD jmx_prometheus_httpserver-0.3.1-jar-with-dependencies.jar jmx_prometheus_httpserver.jar
ADD jmx-config.yml config/jmx-config.yml

ENV JAVA_OPTS="-Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=false -Dcom.sun.management.jmxremote.port=9096 -Dcom.sun.management.jmxremote.rmi.port=9096 -Dcom.sun.management.jmxremote.local.only=true -Djava.rmi.server.hostname=jmx_exporter "

ARG HTTP_PORT 
ENTRYPOINT exec java $JAVA_OPTS -jar jmx_prometheus_httpserver.jar 5556 config/jmx-config.yml
#CMD ["5556",  "config/jmx-config.yml"]

