BINARY_NAME=BenchTesteur

usage:
	@echo "===> deploy pipeline with : make build-archi"
	@echo "===> get dependance for source: make deps " 
	@echo "===> build source to run test : make build/run"
	@echo "===> just run make if you wish to run the pipeline directly"
	@echo "===> dashboard grafana should be accessible at http://<IP>:8089"
	@echo "===> dashboard adminer should be accessible at htpp://<IP>:8088"


all:  clean build buildarchi test run

build:
	@echo "build source go "
	$(MAKE) -C src/ build 
	cp src/$(BINARY_NAME) bin/
	rm src/$(BINARY_NAME)

test:
	$(MAKE) -C src/ test

clean:
	@echo "clean source repo"
	rm bin/$(BINARY_NAME)

run: build 
	@echo "run source go "
	./bin/$(BINARY_NAME)

deps:
		$(MAKE) -C src/ deps 

cleanall: clean stoparchi

# ----------------------------------------------- # 
# ----------------------------------------------- # 
# ----------------------------------------------- # 
# ----------------------------------------------- # 

build-archi:
		@echo "build container archi "
		$(MAKE) -C pipelineBench/ postgres
		$(MAKE) -C pipelineBench/ kafka 
		$(MAKE) -C pipelineBench/ visualization 

stop-kafka: 
		$(MAKE) -C pipelineBench/ ARGS=docker-kafka/compose/kafka-single-node.yml down

stop-visu:
		$(MAKE) -C pipelineBench/ ARGS=docker-compose-influxdb-grafana/docker-compose.yml down
stop-postgres:
		$(MAKE) -C pipelineBench/ ARGS=docker-kafka/compose/pg_compose.yml down

destroy-archi: stop-kafka stop-visu stop-postgres
		@echo "stop container archi "	
