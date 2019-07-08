all:  cleansource buildsource testsource buildarchi runsource

usage:
	@echo "===> deploy pipeline with : make buildarchi"
	@echo "===> get dependance for source: make deps " 
	@echo "===> build source to run test : make buildsource/runsource"
	@echo "===> just run make if you wish to run the pipeline directly"
	@echo "===> dashboard grafana should be accessible at http://<IP>:8089"
	@echo "===> dashboard adminer should be accessible at htpp://<IP>:8088"


buildsource:
		@echo "build source go "
		$(MAKE) -C src/ build 

testsource:
		$(MAKE) -C src/ test

cleansource:
		@echo "clean source repo"
		$(MAKE) -C src/ clean

runsource: buildsource
		@echo "run source go "
		$(MAKE) -C src/ run 
deps:
		$(MAKE) -C src/ deps 


cleanall: cleansource stoparchi

# ----------------------------------------------- # 
# ----------------------------------------------- # 
# ----------------------------------------------- # 
# ----------------------------------------------- # 

up: 
		$(MAKE) -C pipelineBench/ default
buildarchi:
		@echo "build container archi "
		$(MAKE) -C pipelineBench/ postgres
		$(MAKE) -C pipelineBench/ kafka 
		$(MAKE) -C pipelineBench/ visualization 

stopkafka: 
		$(MAKE) -C pipelineBench/ ARGS=docker-kafka/compose/kafka-single-node.yml down

stopvisu:
		$(MAKE) -C pipelineBench/ ARGS=docker-compose-influxdb-grafana/docker-compose.yml down
stopostgres:
		$(MAKE) -C pipelineBench/ ARGS=docker-kafka/compose/pg_compose.yml down

stoparchi: stopkafka stopvisu stopostgres
		@echo "stop container archi "	
