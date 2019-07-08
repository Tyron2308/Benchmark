all:  clean build test run
buildsource:
		@echo "build source go "
		$(MAKE) -C src/ build 
testsource:
		$(MAKE) -C src/ test
cleansource:
		@echo "clean source repo"
		$(MAKE) -C src/ clean
runsource:
		@echo "run source go "
		$(MAKE) -C src/ run 
deps:
		$(MAKE) -C src/ deps 
		
up: 
		$(MAKE) -C pipelineBench/ default
buildarchi:
		@echo "build container archi "
		$(MAKE) -C pipelineBench/ postgres
		$(MAKE) -C pipelineBench/ kafka 
		$(MAKE) -C pipelineBench/ visualization 
stopkafka: 
		$(MAKE) -C pipelineBench/ ARGS=pipelineBench/docker-kafka/compose/kafka-single-node.yml down
stopvisu:
		$(MAKE) -C pipelineBench/ ARGS=pipelineBench/docker-compose-influxdb-grafana/docker-compose.yml down
stoparchi: stopkafka stopvisu
		@echo "stop container archi "
	
