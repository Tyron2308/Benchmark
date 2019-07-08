default:
	@echo "====> build archi <====="

postgres:
	@ echo "===> building database Postgres to register configuration test <===="
	docker-compose -f docker-kafka/compose/pg_compose.yml up -d

kafka:
	@echo "=============building Local API============="
	docker-compose -f docker-kafka/compose/kafka-single-node.yml up -d 

visualization:
	@echo "======> building visualization pipeline <======="
	./docker-compose-influxdb-grafana/jmx_http_exporter/build_image_exporter.sh
	#docker-compose -f docker-compose-influxdb-grafana/docker-compose.yml up -d  

logs:
	docker-compose -f ${ARGS} logs 

down:
	docker-compose -f ${ARGS} down 

test:

clean: down
	@echo "=============cleaning up============="
