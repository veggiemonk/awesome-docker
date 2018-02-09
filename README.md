
# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/veggiemonk/awesome-docker](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/veggiemonk/awesome-docker) [![Say Thanks](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg)](https://saythanks.io/to/veggiemonk)

> A curated list of Docker resources and projects
Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome][sindresorhus] and improved by these **[amazing contributors](https://github.com/veggiemonk/awesome-docker/graphs/contributors)**.

It's now a GitHub project because it's considerably easier for other people to edit, fix and expand on Docker using GitHub.
Just click [README.md][editREADME] to submit a [pull request][editREADME].
If this list is not complete, you can [contribute][editREADME] to make it so. Here is a great video tutorial to [contribute on Github](https://egghead.io/lessons/javascript-identifying-how-to-contribute-to-an-open-source-project-on-github)

***You can see the updates from [TWITTER](https://twitter.com/awesome_docker)***

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)** for tips!

#### *If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editREADME] to improve this file. Thank you!*



The creators and maintainers of this list do not receive and should not receive any form of payment to accept a change made by any contributor.
The goal of this repo is to index open-source projects, not to advertise for profit.

All the links are monitored and tested with [awesome_bot](https://github.com/dkhamsing/awesome_bot) made by [@dkhamsing](https://github.com/dkhamsing)

# Contents

- Introduction
  - [What is Docker?](#what-is-docker)
  - [Where to start?](#where-to-start)
  - [Where to start? (Windows)](#where-to-start-windows)
- [Projects](#projects)
	- [Container Operations](#container-operations)
		- [Container Composition](#container-composition)
		- [Deployment and Infrastructure](#deployment-and-infrastructure)
		- [Monitoring](#monitoring)
		- [Networking](#networking)
		- [Orchestration](#orchestration)
		- [PaaS](#paas)
		- [Reverse Proxy](#reverse-proxy)
		- [Security](#security)
		- [Service Discovery](#service-discovery)
		- [Volume Management / Data](#volume-management--data)
		- [User Interface](#user-interface)
			- [Desktop](#desktop)
			- [Terminal](#terminal)
			- [Web](#web)
	- [Docker Images](#docker-images)
		- [Base Tools](#base-tools)
		- [Builder](#builder)
		- [Dockerfile](#dockerfile)
		- [Linter](#linter)
		- [Registry](#registry)
	- [Development with Docker](#development-with-docker)
		- [API Client](#api-client)
		- [CI/CD](#cicd)
		- [Development Environment](#development-environment)
		- [Garbage Collection](#garbage-collection)
		- [Serverless](#serverless)
		- [Testing](#testing)
		- [Wrappers](#wrappers)
	- [Services based on Docker ($$$)](#services-based-on-docker-)
		- [CI/CD](#ci-services)
		- [CaaS](#caas)
		- [Monitoring](#monitoring-services)
- [Useful Resources](#useful-resources)
	- [Awesome Lists](#awesome-lists)
	- [Good Tips](#good-tips)
	- [Raspberry Pi / ARM](#raspberry-pi--arm)
	- [Security](#security-1)
  - [Videos](#videos)
- [Communities and Meetups](#communities-and-meetups)
	- [Brazilian](#brazilian)
	- [Chinese](#chinese)
	- [English](#english)
	- [Russian](#russian)

# Legend
- Abandoned :skull:
- Beta :construction:
- Monetized :heavy_dollar_sign:

# What is Docker?

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ [What is Docker](https://www.docker.com/what-docker)

# Where to start?
* [Basics – Docker, Containers, Hypervisors, CoreOS](http://etherealmind.com/basics-docker-containers-hypervisors-coreos/)
* [Dive Into Docker: From "What is Docker?" to "Hello World"](https://www.youtube.com/watch?v=XeSD17YRijk&list=PL-v3vdeWVEsXT-u0JDQZnM90feU3NE3v8) (60:25) by [@nickjanetakis][nickjanetakis]
* [Docker Curriculum](https://docker-curriculum.com): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
* [Docker Documentation](https://docs.docker.com/)
* [Docker for all - Developers, Testers, DevOps, Product Owners + Videos](https://github.com/machzqcq/docker-for-all) Docker Training Videos for all
* [Docker Jumpstart](https://github.com/odewahn/docker-jumpstart/): a quick introduction
* [Docker Toolbox](https://docs.docker.com/toolbox/overview/) :skull: - Quick setup and launch of a Docker environment on older Mac (10.10 and below) and Windows (8.1 and below) systems. On newer systems it's recommended to use the [Docker for Mac][docker-for-mac] or [Docker for Windows][docker-for-windows].
* [Docker Training](https://training.docker.com/) - Includes a free self-paced hands-on tutorial (free registration required or sign-in with DockerHub ID)
* [Katacoda](https://www.katacoda.com/): Learn Docker using Interactive Browser-Based Labs
* [Learn Docker](https://github.com/dwyl/learn-docker) Full environment set up, screenshots, step-by-step tutorial and more resources (video, articles, cheat sheets) by [@dwyl](https://github.com/dwyl)
* [Play With Docker](http://training.play-with-docker.com/) - PWD is a great way to get started with Docker from beginner to advanced users. Docker runs directly in your browser.
* [Play With Moby](http://play-with-moby.com/) - PWM is a web based Moby playground which allows you to try different components of the platform in seconds. It gives you the experience of having a free Alpine Linux Virtual Machine in the cloud where you can build and run Moby projects and even create clusters to experiment.

**Cheatsheets** by 
* [@eon01](https://github.com/eon01/DockerCheatSheet)
* [@dimonomid][docker-quick-ref] (PDF)
* [@JensPiegsa](http://docker.jens-piegsa.com) 
* [@wsargent][docker-cheat-sheet] 

# Where to start? (Windows)
* [Windows Containers Quick Start](https://docs.microsoft.com/en-us/virtualization/windowscontainers/about/index) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016
* [Build And Run Your First Docker Windows Server Container](https://blog.docker.com/2016/09/build-your-first-docker-windows-server-container/) Walkthrough installing Docker on Windows 10, building a Docker image and running a Windows container
* [Docker on Windows behind a firewall](https://toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by [@kaitoedter](https://twitter.com/kaitoedter)
* [Windows Containers and Docker: The 101](https://www.youtube.com/watch?v=N7SG2wEyQtM) :movie_camera: - A 20-minute overview, using Docker to run  PowerShell, ASP.NET Core and ASP.NET apps
* [A Comparative Study of Docker Engine on Windows Server vs Linux Platform](http://collabnix.com/a-comparative-study-of-docker-engine-on-windows-server-vs-linux-platform/) Comparing the feature sets and implementations of Docker on Windows and Linux
* [Docker with Microsoft SQL 2016 + ASP.NET](https://blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
* [Running a Legacy ASP.NET App in a Windows Container](https://blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and runnning as a Windows container
* [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](https://www.hanselman.com/blog/ExploringASPNETCoreWithDockerInBothLinuxAndWindowsContainers.aspx) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows][docker-for-windows]

----

# Projects

Moby      = open source development

Docker CE = free product release based on Moby

Docker EE = commercial product release based on Docker CE.

Docker EE is on the same code base as Docker CE, so also built from Moby, with commercial components added, such as "docker data center / universal control plane"

* [Moby](https://github.com/moby/moby)
* [Docker Images](https://hub.docker.com)
* [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
* [Docker Machine](https://github.com/docker/machine) (Machine management for a container-centric world)
* [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)
* [Docker Swarm](https://github.com/docker/swarm) (Swarm: a Docker-native clustering system)


## Container Operations

### Container Composition

* [bocker](https://github.com/p8952/bocker) (1) :skull: - Docker implemented in 100 lines of bash by [p8952](https://github.com/p8952)
* [bocker](https://github.com/icy/bocker) (2) - Write Dockerfile completely in Bash. Extensible and simple. --> Reusable by [@icy](https://github.com/icy)
* [box](https://github.com/box-builder/box) - Build Dockerfile images with a mruby DSL, includes flattening and layer manipulation
* [Capitan](https://github.com/byrnedo/capitan) - Composable docker orchestration with added scripting support by [@byrnedo](https://github.com/byrnedo).
* [crowdr](https://github.com/polonskiy/crowdr) - Tool for managing multiple Docker containers (`docker-compose` alternative) by [@polonskiy](https://github.com/polonskiy/)
* [Composerize](https://github.com/magicmark/composerize) - Convert docker run commands into docker-compose files
* [compose_plantuml](hhttps://github.com/funkwerk/compose_plantuml) -  Generate Plantuml graphs from docker-compose files by [@funkwerk](https://github.com/funkwerk)
* [docker-compose-graphviz](https://github.com/abesto/docker-compose-graphviz) - Turn a docker-compose.yml files into Graphviz .dot files by [@abesto](https://github.com/abesto)
* [docker-compose-search](https://github.com/francescou/docker-compose-search) - A search engine for Docker Compose application stacks by [@francescou](https://github.com/francescou/)
* [draw-compose](https://github.com/Alexis-benoist/draw-compose) - Utility to draw a schema of a docker compose by [@Alexis-benoist](https://github.com/Alexis-benoist)
* [elsy](https://github.com/cisco/elsy) - An opinionated, multi-language, build tool based on Docker and Docker Compose
* [habitus](https://github.com/cloud66/habitus) - A Build Flow Tool for Docker http://www.habitus.io by [@cloud66](https://github.com/cloud66)
* [Maestro](https://github.com/toscanini/maestro) :skull: - Maestro provides the ability to easily launch, orchestrate and manage mulitiple Docker containers as single unit by [@tascanini](https://github.com/toscanini)
* [percheron][percheron] :skull: - Organise your Docker containers with muscle and intelligence by [@ashmckenzie](https://github.com/ashmckenzie)
* [rocker](https://github.com/grammarly/rocker) - Extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc. by [grammarly](grammarly).
* [rocker-compose](https://github.com/grammarly/rocker-compose) - Docker composition tool with idempotency features for deploying apps composed of multiple containers. By [@grammarly](grammarly)
* [Stacker](https://github.com/stacker/stacker-cli) - Docker Compose Templates. Stacker provides an abstraction layer over Docker Compose and a better DX (developer experience).
* [Zodiac](https://github.com/CenturyLinkLabs/zodiac) :skull: - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs][CenturyLinkLabs]

### Deployment and Infrastructure

* [blackfish](https://gitlab.com/blackfish/blackfish) - a CoreOS VM to build swarm clusters for Dev & Production by [@DataMC](http://datamc.io/)
* [Centurion](https://github.com/newrelic/centurion) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [@newrelic](https://github.com/newrelic)
* [Clocker](https://github.com/brooklyncentral/clocker) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and [Brooklyn][brooklyn] for application blueprints. By [@brooklyncentral](https://github.com/brooklyncentral)
* [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker by [@ehazlett](https://github.com/ehazlett)
* [depcon](https://github.com/gondor/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes.  By [@gonodr][gondor]
* [deploy](https://github.com/ttiny/deploy) :skull: - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration by [@ttiny](https://github.com/ttiny)
* [dockit](https://github.com/humblec/dockit) :skull: - Do docker actions and Deploy gluster containers! By [@humblec](https://github.com/humblec)
* [Grafeas](https://github.com/Grafeas/Grafeas) - A common API for metadata about containers, from image and build details to security vulnerabilities. By [Grafeas](https://github.com/Grafeas)
* [Longshoreman](https://github.com/longshoreman/longshoreman) :skull: - Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By [longshoreman](https://github.com/longshoreman)

### Monitoring

* [Axibase Collector](https://github.com/axibase/atsd-use-cases/tree/master/Solutions/docker) - Axibase Collector streams performance counters, configuration changes and lifecycle events from the Docker engine(s) into Axibase Time Series Database for roll-up dashboards and integration with upstream monitoring systems.
* [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers. Created by [@Google](https://github.com/google)
* [Docker-Alertd](https://github.com/deltaskelta/docker-alertd) - Monitor and send alerts based on docker container resource usage/statistics
* [Docker-Flow-Monitor](https://github.com/vfarcic/docker-flow-monitor) - Reconfigures Prometheus when a new service is updated or deployed automatically by [@vfarcic][vfarcic]
* [Docker-Fluentd][fluentd] - Docker container to Log Other Containers' Logs. One can aggregate the logs of Docker containers running on the same host using Fluentd by [@kiyoto][kiyoto]
* [Dockerana](https://github.com/dockerana/dockerana) :skull: - packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.
* [Dynatrace](https://www.dynatrace.com/technologies/cloud-and-microservices/docker-monitoring/) - Monitor containerized applications without installing agents or modifying your Run commands
* [Glances](https://nicolargo.github.io/glances/) - A cross-platform curses-based system monitoring tool written in Python by [@nicolargo](https://github.com/nicolargo)
* [Grafana Docker Dashboard Template](https://grafana.com/dashboards/179) - A template for your Docker, Grafana and Prometheus stack [@vegasbrianc][vegasbrianc]
* [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) - InfluxDB Time series DB in combination with Grafana and cAdvisor by [@vegasbrianc][vegasbrianc]
* [LogJam](https://github.com/gocardless/logjam) - Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash) by [@gocardless](https://github.com/gocardless)
* [Logsene for Docker][spm] Monitoring of Metrics, Events and Logs implemented in Node.js. Integrated [logagent-js](https://github.com/sematext/logagent-js) to detect and parse various log formats. [@sematext][sematext]
* [Logspout](https://github.com/gliderlabs/logspout) - Log routing for Docker container logs by [@gliderlabs][gliderlabs]
* [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting) - Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].
* [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
* [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.

### Networking

* [Calico-Docker](https://www.projectcalico.org/getting-started/docker/) - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
* [Flannel](https://github.com/coreos/flannel/) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes. By [@coreos][coreos]
* [netshoot](https://github.com/nicolaka/netshoot) - The netshoot container has a powerful set of networking tools to help troubleshoot Docker networking issues by [@nicolaka](https://github.com/nicolaka)
* [Weave][weave] (The Docker network) - Weave creates a virtual network that connects Docker containers deployed across multiple hosts.

### Orchestration


* [athena](https://github.com/athena-oss/athena) - An automation platform with a plugin architecture that allows you to easily create and share services.
* [blimp](https://github.com/tubesandlube/blimp) :skull: - Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more by [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin)
* [CloudSlang](http://www.cloudslang.io/) - CloudSlang is a workflow engine to create Docker process automation
* [clusterdock](https://github.com/clusterdock/clusterdock) - Docker container orchestration to enable the testing of long-running cluster deployments
* [ContainerShip](https://github.com/containership/containership) A simple container management platform
* [Crane](https://github.com/Dataman-Cloud/crane) - Control plane based on docker built-in swarm [@Dataman-Cloud](https://github.com/Dataman-Cloud)
* [Docker Flow Swarm Listener](https://github.com/vfarcic/docker-flow-swarm-listener) :construction: - Docker Flow Swarm Listener project is to listen to Docker Swarm events and send requests when a change occurs.. By [@vfarcic][vfarcic]
* [gantryd](https://github.com/DevTable/gantryd) :skull: - A framework for easy management of docker-based components across machines by [@DevTable](https://github.com/DevTable)
* [Haven](https://github.com/codeabovelab/haven-platform) - Haven is a simplified container management platform that integrates container, application, cluster, image, and registry managements. By [@codeabovelab](https://github.com/codeabovelab)
* [Helios](https://github.com/spotify/helios) - A simple platform for deploying and managing containers across an entire fleet of servers by [@spotify](spotify)
* [Kontena](https://github.com/kontena/kontena) - Application Containers for Masses https://www.kontena.io/
* [Kubernetes](https://github.com/kubernetes/kubernetes) - Open source orchestration system for Docker containers by Google
* [ManageIQ](https://github.com/ManageIQ/manageiq) - Discover, optimize and control your hybrid IT. By [ManageIQ](https://github.com/ManageIQ)
* [Mantl](https://github.com/mantl/mantl) - Mantl is a modern platform for rapidly deploying globally distributed services
* [Marathon](https://mesosphere.github.io/marathon/docs/) - Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on"
* [Mesos](https://mesos.apache.org/documentation/latest/docker-containerizer/) - Resource/Job scheduler for containers, VM's and physical hosts [@apache](https://mesos.apache.org/)
* [Mesosphere DC/OS](https://mesosphere.com/product/) :heavy_dollar_sign: - Integrated platform for data and containers built on Apache Mesos by [@mesosphere](https://mesosphere.com)
* [Nebula](https://github.com/nebula-orchestrator) - A Docker orchestration tool designed to manage massive scale distributed clusters.
* [Nomad](https://github.com/hashicorp/nomad) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler by [@hashicorp][hashicorp]
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui) :skull: - An open-source project that makes deploying complex containerized apps as easy as Drag-and-Drop by [@CenturyLinkLabs][CenturyLinkLabs].
* [Rancher](https://github.com/rancher/rancher) - An open source project that provides a complete platform for operating Docker in production by [@rancher][rancher].
* [Swarmpit](https://github.com/swarmpit/swarmpit) - Lightweight Docker Swarm orchestration. Swarmpit provides clean way to manage your Docker Swarm cluster with various handful features such Service management, smart search, shared access and private registries.

### PaaS

* [Atlantis](https://github.com/ooyala/atlantis) :skull: - Atlantis is an Open Source PaaS for HTTP applications built on Docker and written in Go
* [Convox Rack](https://github.com/convox/rack) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
* [Dcw](https://github.com/pbertera/dcw) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels.
* [Dokku][dokku] - Docker powered mini-Heroku that helps you build and manage the lifecycle of applications (originally by [@progrium][progrium])
* [Empire](https://github.com/remind101/empire) - A PaaS built on top of Amazon EC2 Container Service (ECS)
* [Flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service
* [Jelastic](https://jelastic.com/) - An advanced PaaS for developers that simplifies clustering and complex cloud deployments with powerful web UI and usage-only pricing
* [Nanobox](https://github.com/nanobox-io/nanobox) :heavy_dollar_sign: - An application development platform that creates local environments that can then be deployed and scaled in the cloud.
* [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by [Red Hat](https://www.redhat.com/)
* [Tsuru](https://github.com/tsuru/tsuru) - Tsuru is an extensible and open source Platform as a Service software
* [Workflow](https://github.com/deis/workflow) - The open source PaaS for Kubernetes by [Deis](https://github.com/deis). Formerly Deis v1.

### Reverse Proxy

* [docker-flow-proxy](https://github.com/vfarcic/docker-flow-proxy) - Reconfigures proxy every time a new service is deployed, or when a service is scaled. By [@vfarcic][vfarcic]
* [docker-proxy](https://github.com/silarsis/docker-proxy) :skull: - Transparent proxy for docker containers, run in a docker container. By [@silarsis](https://github.com/silarsis)
* [fabio](https://github.com/fabiolb/fabio) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By [@magiconair](https://github.com/magiconair) (Frank Schroeder)
* [h2o-proxy](https://github.com/zchee/h2o-proxy) :skull: - Automated H2O reverse proxy for Docker containers. An alternative to [jwilder/nginx-proxy][nginxproxy] by [@zchee](https://github.com/zchee)
* [Let's Encrypt Nginx-proxy Companion](https://github.com/JrCs/docker-letsencrypt-nginx-proxy-companion) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By [@JrCs](https://github.com/JrCs)
* [muguet](https://github.com/mattallty/muguet) - DNS Server & Reverse proxy for Docker environments. By [@mattallty](https://github.com/mattallty)
* [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [@jwilder][jwilder]
* [Swarm Ingress Router](https://github.com/tpbowden/swarm-ingress-router) - Route DNS names to Swarm services based on labels. By [@tpbowden](https://github.com/tpbowden/)
* [Træfɪk](https://github.com/containous/traefik) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [@EmileVauge](https://github.com/emilevauge)

### Security

* [Anchor Engine](https://github.com/anchore/anchore) - Analyze images for CVE vulnerabilities and against custom security policies by [@Anchor][anchore]
* [Anchor Cloud](https://anchore.com/cloud/) :heavy_dollar_sign: - Hosted version of Anchor Engine by [@Anchor][anchore]
* [Aqua Security](https://www.aquasec.com) :heavy_dollar_sign: - Securing container-based applications from Dev to Production on any platform
* [bane](https://github.com/jessfraz/bane) - AppArmor profile generator for Docker containers by [@jessfraz][jessfraz]
* [CIS Docker Benchmark](https://github.com/dev-sec/cis-docker-benchmark) - This [InSpec][inspec] compliance profile implement the CIS Docker 1.12.0 Benchmark in an automated way to provide security best-practice tests around Docker daemon and containers in a production environment. By [@dev-sec](https://github.com/dev-sec)
* [Clair](https://github.com/coreos/clair) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [@coreos][CoreOS]
* [docker-bench-security](https://github.com/docker/docker-bench-security) - script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker][docker]
* [notary](https://github.com/theupdateframework/notary) - a server and a client for running and interacting with trusted collections. By [@TUF](https://github.com/theupdateframework)
* [oscap-docker](https://github.com/OpenSCAP/openscap) - OpenSCAP provides oscap-docker tool which is used to scan Docker containers and images. By RedHat
* [Sysdig Falco](https://github.com/draios/falco) - Sysdig Falco is an open source container security monitor. It can monitor application, container, host, and network activity and alert on unauthorized activity.
* [Sysdig Secure](https://sysdig.com/product/secure/) :heavy_dollar_sign: - Sysdig Secure addresses run-time security through behavioral monitoring and defense, and provides deep forensics based on open source Sysdig for incident response.
* [Twistlock](https://www.twistlock.com/) :heavy_dollar_sign: - Twistlock Security Suite detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.

### Service Discovery

* [Docker Grand Ambassador](https://github.com/cpuguy83/docker-grand-ambassador) :skull: - This is a fully dynamic docker link ambassador. + [Article](https://docs.docker.com/engine/admin/ambassador_pattern_linking/) by [@cpuguy83][cpuguy83]
* [docker-consul](https://github.com/gliderlabs/docker-consul) by [@progrium][progrium]
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery by [@coreOS][coreos]
* [istio](https://github.com/istio/istio) - An open platform to connect, manage, and secure microservices by [@IstioMesh](istio)
* [proxy](https://github.com/factorish/proxy) :skull: - lightweight nginx based load balancer self using service discovery provided by registrator. by [@factorish](https://github.com/factorish)
* [registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker by [@gliderlabs][gliderlabs] and [@progrium][progrium]

### Volume Management / Data

* [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume) - The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By [@blockbridge][blockbridge]
* [Convoy](https://github.com/rancher/convoy) - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher][rancher]
* [Docker Unison](https://github.com/leighmcculloch/docker-unison) A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By [@leighmcculloch](https://github.com/leighmcculloch)
* [Netshare](https://github.com/ContainX/docker-volume-netshare) Docker NFS, AWS EFS, Ceph & Samba/CIFS Volume Plugin. By [@ContainX](https://github.com/ContainX)
* [Docker Machine NFS](https://github.com/adlogix/docker-machine-nfs) Activates NFS for an existing boot2docker box created through Docker Machine on OS X.
* [Local Persist](https://github.com/CWSpear/local-persist) Specify a mountpoint for your local volumes (created via `docker volume create`) so that files will always persist and so you can mount to different directories in different containers.
* [Minio](https://github.com/jelastic-jps/minio) - S3 compatible object storage server in Docker containers
* [portworx](https://portworx.com) :heavy_dollar_sign: - Decentralized storage solution for persistent, shared and replicated volumes.
* [REX-Ray](https://github.com/thecodeteam/rexray) Vendor agnostic storage orchestration engine to provide persistent storage for Docker containers as well as Mesos frameworks and tasks. By[@Dell](https://github.com/thecodeteam)
* [quobyte](https://www.quobyte.com/) :heavy_dollar_sign: - fully fault-tolerant distributed file system with a docker volume driver

### User Interface

#### Desktop

*Native desktop applications for managing and montoring docker hosts and clusters*

* [Captain](https://getcaptain.co/) - Manage containers from the MacOSX menu bar by [@RickWong](https://github.com/rickwong)
* [DockStation](https://github.com/DockStation/dockstation) - A developer centric UI to configure, monitor, and manage services and containers [@dock_station](https://twitter.com/dock_station)
* [Dockeron](https://github.com/dockeron/dockeron) - A project built on Electron + Vue.js for Docker on desktop. [@fluency03](https://github.com/fluency03)
* [Lifeboat](https://github.com/jplhomer/lifeboat) - An easy way to launch Docker projects with a graphical interface on your Mac. [@jplhomer](https://github.com/jplhomer)

#### Terminal

* [ctop (1)](https://github.com/yadutaf/ctop) - A command line / text based Linux Containers monitoring tool that works just like you expect (Python) by [@yadutaf](https://github.com/yadutaf)
* [ctop (2)](https://github.com/bcicen/ctop) - Top-like interface for container metrics (Golang) by [@bcicen](https://github.com/bcicen/)
* [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin) - Search the Docker Registry with the Dext smart launcher.
* [dockercraft](https://github.com/docker/dockercraft) - Docker + Minecraft = Dockercraft by [@docker][docker]
* [docker.el](https://github.com/Silex/docker.el) Manage docker from Emacs by [Silex](https://github.com/Silex)
* [dockerfile-mode](https://github.com/spotify/dockerfile-mode) An emacs mode for handling Dockerfiles by [spotify][spotify]
* [Docker-mon](https://github.com/icecrime/docker-mon) :skull: - Console-based Docker monitoring by [@icecrime](https://github.com/icecrime)
* [dockersql](https://github.com/crosbymichael/dockersql) - A command line interface to query Docker using SQL by [@crosbymichael][crosbymichael]
* [dockly](https://github.com/lirantal/dockly) - An interactive shell UI for managing Docker containers by [@lirantal](https://github.com/lirantal)
* [docker-ls](https://github.com/mayflower/docker-ls) - CLI tools for browsing and manipulating docker registries by [@mayflower](https://github.com/mayflower)
* [dry](https://github.com/moncho/dry) - An interactive CLI for Docker containers by [@moncho](https://github.com/moncho)
* [DVM](https://github.com/howtowhale/dvm) - Docker version manager by [@howtowhale](https://github.com/howtowhale)
* [MultiDocker](https://github.com/marty90/multidocker) - Create a secure multi-user Docker machine, where each user is segregated into an indepentent container.
* [ns-enter](https://github.com/jpetazzo/nsenter) - no more ssh, enter name spaces of container by [@jpetazzo][jpetazzo]
* [Powerline-Docker](https://github.com/adrianmo/powerline-docker) - A Powerline segment for showing the status of Docker containers by [@adrianmo](https://github.com/adrianmo)
* [reg](https://github.com/jessfraz/reg) - Docker registry v2 command line client by [@jessfraz][jessfraz]
* [sen](https://github.com/TomasTomecek/sen) - Terminal user interface for docker engine, by [@TomasTomecek](https://github.com/TomasTomecek)
* [scuba](https://github.com/JonathonReinhart/scuba) - Transparently use Docker containers to encapsulate software build environments, by [@JonathonReinhart](https://github.com/JonathonReinhart)
* [supdock](https://github.com/segersniels/supdock) - :construction: Allows for slightly more visual usage of Docker with an interactive prompt by [@segersniels](https://github.com/segersniels)
* [wharfee](https://github.com/j-bennet/wharfee) - Autocompletion and syntax highlighting for Docker commands. by [@j-bennet](https://github.com/j-bennet)
* [tsaotun](https://github.com/qazbnm456/tsaotun) - Python based Assistance for Docker by [@qazbnm456](https://github.com/qazbnm456)

#### Web

* [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser) - Web Interface for the Docker Registry HTTP API v2 by [@klausmeyer](https://github.com/klausmeyer)
* [Docker Registry UI](https://github.com/atcol/docker-registry-ui) - A web UI for easy private/local Docker Registry integration by [@atc-](https://github.com/atcol)
* [docker-registry-web](https://github.com/mkuchin/docker-registry-web) - Web UI, authentication service and event recorder for private docker registry v2 by [@mkuchin](https://github.com/mkuchin)
* [DockerSurfer](https://github.com/Simone-Erba/DockerSurfer) :construction: - A web service for analyze and browse dependencies between Docker images in the Docker registry, by [@Simone-Erba](https://github.com/Simone-Erba/)
* [docker-swarm-visualizer](https://github.com/dockersamples/docker-swarm-visualizer) - Visualizes Docker services on a Docker Swarm (for running demos).
* [dockering-on-rails](https://github.com/Electrofenster/dockerding-on-rails) :skull: - Simple Web-Interface for Docker with a lot of features by [@Electrofenster](https://github.com/Electrofenster/)
* [OctoLinker](https://github.com/OctoLinker/OctoLinker) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
* [Portainer](https://github.com/portainer/portainer) - A lightweight management UI for managing your Docker hosts or Docker Swarm clusters by [@portainer](https://github.com/portainer)
* [Portus](https://github.com/SUSE/Portus) - Authorization service and frontend for Docker registry (v2) by [@SUSE](https://github.com/SUSE)
* [Rapid Dashboard](https://github.com/ozlerhakan/rapid) - A simple query dashboard to use Docker Remote API by [@ozlerhakan](https://github.com/ozlerhakan/)
* [Seagull](https://github.com/tobegit3hub/seagull) - Friendly Web UI to monitor docker daemon. by [@tobegit3hub](https://github.com/tobegit3hub)
* [Shipyard](https://github.com/shipyard/shipyard) :skull: - Shipyard enables multi-host, Docker cluster management using Docker Swarm.

## Docker Images

### Base Tools

*Tools and applications that are either installed inside containers or designed to be run as a [sidecar](https://docs.microsoft.com/en-us/azure/architecture/patterns/sidecar)*

* [amicontained](https://github.com/jessfraz/amicontained) - Container introspection tool. Find out what container runtime is being used as well as features available by [@jessfraz][jessfraz]
* [autodock](https://github.com/prologic/autodock) - Daemon for Docker Automation by [@prologic][prologic]
* [Chaperone](https://github.com/garywiz/chaperone) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package. by [@garywiz](https://github.com/garywiz)
* [CoreOS][coreos] - Linux for Massive Server Deployments
* [docker-alpine][alpine] - A super small Docker base image *(5MB)* using Alpine Linux by [@gliderlabs][gliderlabs]
* [dockerize](https://github.com/jwilder/dockerize) - Utility to simplify running applications in docker containers by [@jwilder][jwilder]
* [docker-gen](https://github.com/jwilder/docker-gen) - Generate files from docker container meta-data by [@jwilder][jwilder]
* [GoSu](https://github.com/tianon/gosu) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool) by [@tianon](https://github.com/tianon)
* [is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container by [@sindresorhus][sindresorhus]
[supercronic](https://github.com/aptible/supercronic) - crontab-compatible job runner, designed specifically to run in containers by [@aptible](https://github.com/aptible/)
* [TrivialRC](https://github.com/vorakl/TrivialRC) - A minimalistic Runtime Configuration system and process manager for containers [@vorakl](https://github.com/vorakl)

### Builder

*Applications designed to help or simplify building **new** images*

* [container-diff](https://github.com/GoogleCloudPlatform/container-diff) - An image tool for comparing and analzying container images by [@GoogleCloudPlatform][GoogleCloudPlatform]
* [container-factory](https://github.com/mutable/container-factory) - Produces Docker images from tarballs of application source code by [@mutable](https://github.com/mutable)
* [Derrick](https://github.com/alibaba/derrick) - A tool help you to automate the generation of Dockerfile and dockerize application by scanning the code. By [@alibaba](https://github.com/alibaba). 
* [dlayer](https://github.com/wercker/dlayer) - Stats collector for Docker layers by [@wercker](https://github.com/wercker)
* [docker-companion](https://github.com/mudler/docker-companion) - A command line tool written in Golang to squash and unpack docker images by [@mudler](https://github.com/mudler/)
* [docker-make](https://github.com/CtripCloud/docker-make) - Build, tag,and push a bunch of related docker images via a single command.
* [docker-replay](https://github.com/bcicen/docker-replay) - Generate `docker run`command and options from running containers. By [bcicen](https://github.com/bcicen)
* [DockerSlim](https://github.com/docker-slim/docker-slim) shrinks fat Docker images creating the smallest possible images.
* [Dockly](https://github.com/swipely/dockly) - Dockly is a gem made to ease the pain of packaging an application in Docker by [@swipely](https://github.com/swipely/)
* [dockramp](https://github.com/jlhawn/dockramp) :skull: - Proof of Concept: A Client Driven Docker Image Builder by [@jlhawn](https://github.com/jlhawn)
* [flyimg](http://flyimg.io/) - Docker image resizing, cropping, and compression on the fly.
* [MicroBadger][microbadger] - Analyze the contents of images and add metadata labels
* [packer](https://www.packer.io/docs/builders/docker.html) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible
* [portainer](https://github.com/duedil-ltd/portainer) - Apache Mesos framework for building Docker images by [@duedil-ltd](https://github.com/duedil-ltd)
* [runlike](https://github.com/lavie/runlike) 🚧 - Generate `docker run`command and options from running containers by [@lavie](https://github.com/lavie)
* [SkinnyWhale](https://github.com/djosephsen/skinnywhale) :skull: - Skinnywhale helps you make smaller (as in megabytes) Docker containers.
* [Whales](https://github.com/Gueils/whales) - A tool to automatically dockerize your applications by [@icalialabs](https://github.com/IcaliaLabs).

### Dockerfile

* [chaperone-docker](https://github.com/garywiz/chaperone-docker) - A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.
* [Dockerfile Project](http://dockerfile.github.io/) : Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
* [Dockerfile Generator](http://jrruethe.github.io/blog/2015/09/20/dockerfile-generator/)
* [Vektorcloud](https://github.com/vektorcloud) - A collection of minimal, Alpine-based Docker images

Examples by:
* [@arun-gupta](https://github.com/arun-gupta/docker-images)
* [@crosbymichael](https://github.com/crosbymichael/Dockerfiles)
* [@jessfraz](https://github.com/jessfraz/Dockerfiles)
* [@komljen](https://github.com/komljen/dockerfile-examples)
* [@kstaken](https://github.com/kstaken/dockerfile-examples)
* [@awesome-startup]( https://github.com/awesome-startup/docker-compose)
* [@ondrejmo](https://github.com/ondrejmo/Dockerfiles)
* [@pandrew](https://github.com/pandrew/dockerfiles)
* [@vimagick](https://github.com/vimagick/dockerfiles)


### Linter

* [dockfmt](https://github.com/jessfraz/dockfmt) :construction: - Dockerfile formatter and parser by [@jessfraz][jessfraz]
* [dockerfile_lint](https://github.com/projectatomic/dockerfile_lint) - A rule-based 'linter' for Dockerfiles by [@redhataccess](https://github.com/redhataccess)
* [Hadolint](https://github.com/hadolint/hadolint) - A Dockerfile linter written in Haskell by [@lukasmartinelli](https://github.com/lukasmartinelli)
* [Whale-linter](https://github.com/jeromepin/whale-linter) - A simple and small Dockerfile linter written in Python3+ without dependencies by [@jeromepin](https://github.com/jeromepin)

### Registry

Services to securely store your Docker images.

* [Amazon EC2 Container Registry](https://aws.amazon.com/ecr/) Amazon EC2 Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
* [Atomic Registry](http://www.projectatomic.io/registry/) - Red Hat Atomic Registry is an open source enterprise registry based on the Origin and Cockpit projects, enhancing the Docker registry library.
* [Azure Container Registry](https://azure.microsoft.com/de-de/services/container-registry/) Manage a Docker private registry as a first-class Azure resource
* [CargoOS](https://github.com/RedCoolBeans/cargos-buildroot) - A bare essential OS for running the Docker Engine on bare metal or Cloud. By [@RedCoolBeans](https://github.com/RedCoolBeans)
* [Cycle.io](https://cycle.io/) Bare-metal container hosting.
* [Docket](https://github.com/netvarun/docket) - Custom docker registry that allows for lightning fast deploys through bittorrent by [@netvarun](https://github.com/netvarun/)
* [Docker Hub](https://hub.docker.com/) provided by Docker Inc.
* [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
* [Europa :heavy_dollar_sign:](https://github.com/puppetlabs/europa) - Private docker registry with support for image pipelines and webhooks. By [@puppetlabs](https://github.com/puppetlabs)
* [GCE Container Registry](https://cloud.google.com/container-registry/) Fast, private Docker image storage on Google Cloud Platform
* [GitLab Container Registry](https://docs.gitlab.com/ce/user/project/container_registry.html) - Repositories focused on using it images in GitLab CI
* [Private Docker Registry :heavy_dollar_sign:](https://private-docker-registry.com) - Dedicated Conainer Registry Service with unlimited private repositories, users, teams, namespaces together with enterprise grade authentication LDAP/AD/OAuth/SAML.
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [Rescoyl](https://github.com/noteed/rescoyl) - Private Docker registry (free and open source) by [@noteed][noteed]
* [Sonatype Nexus](https://www.sonatype.com/nexus-repository-oss) - Repository with Universal Support, also for Docker images
* [TreeScale](https://github.com/treescale) - Build and Distribute container based applications. By [@tigranbs](https://github.com/tigranbs)
* [VMWare Harbor](http://vmware.github.io/harbor/) Project Harbor by VMWare is an enterprise-class registry server that stores and distributes Docker images. Harbor extends the open source Docker Distribution by adding the functionalities usually required by an enterprise, such as security, identity and management.

## Development with Docker

### API Client

* [ahab](https://github.com/instacart/ahab) - Docker event handling with Python by [@instacart](https://github.com/instacart)
* [Docker Client for JVM](https://github.com/gesellix/docker-client) - A Docker remote api client library for the JVM, written in Groovy by [@gesellix][gesellix]
* [Docker Client TypeScript](https://gitlab.com/masaeedu/docker-client) - Docker API client for JavaScript, automatically generated from Swagger API definition from moby repository. By [@masaeedu](https://github.com/masaeedu)
* [docker-it-scala](https://github.com/whisklabs/docker-it-scala) - Docker integration testing kit with Scala by [@whisklabs](https://github.com/whisklabs)
* [docker-maven-plugin (1)](https://github.com/spotify/docker-maven-plugin) - A Maven plugin for building and pushing Docker images by [@spotify][spotify]
* [docker-maven-plugin (2)](https://github.com/fabric8io/docker-maven-plugin) - A Maven plugin for running and creating Docker images by [@fabric8io](https://github.com/fabric8io)
* [Docker-PowerShell](https://github.com/Microsoft/Docker-PowerShell) - PowerShell Module for Docker
* [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet) - C#/.NET HTTP client for the Docker remote API by [@ahmetalpbalkan](ahmetalpbalkan)
* [dockerode](https://github.com/apocas/dockerode) - Docker Remote API node.js module by [@apocas](https://github.com/apocas)
* [DoMonit](https://github.com/eon01/DoMonit) - A simple Docker Monitoring wrapper For Docker API
* [go-dockerclient](https://github.com/fsouza/go-dockerclient/) - Go HTTP client for the Docker remote API by [@fsouza](https://github.com/fsouza/)
* [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) - A Docker remote api plugin for Gradle by [@gesellix][gesellix]
* [libcompose](https://github.com/docker/libcompose) - Go library for Docker Compose.
* [powerstrip](https://github.com/clusterhq/powerstrip) :skull: - A tool for prototyping Docker extensions by [@clusterhq](https://github.com/clusterhq)
* [sbt-docker](https://github.com/marcuslonnberg/sbt-docker) - Create Docker images directly from sbt by [@marcuslonnberg](https://github.com/marcuslonnberg)
* [sbt-docker-compose](https://github.com/Tapad/sbt-docker-compose) - Integrates Docker Compose functionality into sbt by [@kurtkopchik](https://github.com/kurtkopchik/)

### CI/CD

* [Buddy](https://buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development.
* [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery by [@harbur](https://github.com/harbur).
* [Cyclone](https://github.com/caicloud/cyclone) - A cloud native CI/CD platform built for container workflow by [@caicloud](https://github.com/caicloud).
* [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
* [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
* [GitLab CI](https://about.gitlab.com/gitlab-ci/) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
* [GOCD-Docker](https://github.com/gocd/gocd-docker)Go Server and Agent in docker containers to provision.
* [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment) - Continuous deployment of a microservices application.
* [mu](https://github.com/stelligent/mu) - Tool to configure CI/CD of your container applications via AWS CodePipeline, CodeBuild and ECS [@Stelligent](https://github.com/stelligent)
* [Screwdriver](http://screwdriver.cd/) - Yahoo's OpenSource buildplatform designed for Continous Delivery.
* [Skipper](https://github.com/Stratoscale/skipper) - Easily dockerize your Git repository by [@Stratoscale](https://github.com/Stratoscale)
* [SwarmCI](https://github.com/ghostsquad/swarmci) - Create a distributed, isolated task pipeline in your Docker Swarm.
* [Watchtower](https://github.com/v2tec/watchtower) - Automatically update running Docker containers by [@CenturyLinkLabs][CenturyLinkLabs]

### Development Environment

* [Binci](https://github.com/binci/binci) - Containerize your development workflow. (formerly DevLab by [@TechnologyAdvice](https://github.com/TechnologyAdvice))
* [Boot2Docker](https://github.com/boot2docker/boot2docker) - Docker for OSX and Windows
* [construi](https://github.com/lstephen/construi) - Run your builds inside a Docker defined environment by [@lstephen](https://github.com/lstephen)
* [Devstep](https://github.com/fgrehm/devstep) :skull: - Development environments powered by Docker and buildpacks by [@fgrehm][fgrehm]
* [Dinghy](https://github.com/codekitchen/dinghy) - An alternative way to use Docker on Mac OS X using Docker Machine with virtualbox, vmware, xhyve or parallels
* [DLite](https://github.com/nlf/dlite) - Simplest way to use Docker on OSX, no VM needed. By [@nlf](https://github.com/nlf)
* [Docker-Arch](https://github.com/ph3nol/docker-arch) - Generate Web/CLI projects Dockerized development environments, from 1 simple YAML file. By [@Ph3nol](https://github.com/ph3nol)
* [docker-vm](https://github.com/shyiko/docker-vm) - Simple and transparent alternative to boot2docker (backed by Vagrant) by [@shyiko](https://github.com/shyiko)
* [Docker Missing Tools](https://github.com/nandoquintana/docker-missing-tools) - A set of bash commands to shortcut typical docker dev-ops. An alternative to creating typical helper scripts like "build.sh" and "deploy.sh" inside code repositories. By [@NandoQuintana](https://github.com/nandoquintana).
* [Docker osx dev](https://github.com/brikis98/docker-osx-dev) - A productive development environment with Docker on OS X by [@brikis98](https://github.com/brikis98)
* [Docker-sync](http://docker-sync.io/) - Drastically improves performance ([50-70x](https://github.com/EugenMayer/docker-sync/wiki/4.-Performance)) when using Docker for development on Mac OS X/Windows and Linux while sharing code to the container. By [@EugenMayer](https://github.com/EugenMayer)
* [Dusty](http://dusty.gc.com/) - Managed Docker development environments on OS X
* [Eclipse Che](http://www.eclipse.org/che/) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE
* [EnvCLI](https://github.com/PhilippHeuer/EnvCLI) - Replace your local installation of Node, Go, ... with project-specific docker containers. By [@PhilippHeuer](https://github.com/PhilippHeuer)
* [forward2docker](https://github.com/bsideup/forward2docker) :skull: - Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM by [@bsideup](https://github.com/bsideup)
* [Vagga](https://github.com/tailhook/vagga) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments by [@tailhook](https://github.com/tailhook/)
* [Vessel](https://github.com/awvessel/vessel) :skull: - Automates the setup & use of dockerized development environments by [@awvessel](https://github.com/awvessel)

### Garbage Collection

* [caduc](https://github.com/tjamet/caduc) - A docker garbage collector cleaning stuff you did not use recently
* [Docker Clean](https://github.com/zzrotdesign/docker-clean) - A script that cleans Docker containers, images and volumes by [@zzrotdesign](https://github.com/zzrotdesign)
* [Docker-cleanup](https://github.com/meltwater/docker-cleanup) - Automatic Docker image, container and volume cleanup by [@meltwater](https://github.com/meltwater)
* [docker-garby](https://github.com/konstruktoid/docker-garby) - Docker garbage collection script by [@konstruktoid](https://github.com/konstruktoid).
* [docker-gc](https://github.com/spotify/docker-gc) - A cron job that will delete old stopped containers and unused images by [@spotify][spotify]
* [sherdock](https://github.com/rancher/sherdock) :skull: - Automatic GC of images based on regexp by [@rancher][rancher]

### Serverless

* [AMP](https://github.com/appcelerator/amp) - The open source unified CaaS/FaaS platform for Docker, batteries included. By [@Appcelerator](https://github.com/appcelerator/)
* [Apache OpenWhisk](https://github.com/apache/incubator-openwhisk) - a serverless, open source cloud platform that executes functions in response to events at any scale. By [@apache](https://github.com/apache)
* [Docker-Lambda](https://github.com/lambci/docker-lambda) - Docker images and test runners that replicate the live AWS Lambda environment. By [@lamb-ci](https://github.com/lambci)
* [Funker](https://github.com/bfirsh/funker-example-voting-app) - Functions as Docker containers example voting app. By [@bfirsh](https://github.com/bfirsh)
* [IronFunctions](https://github.com/iron-io/functions) - The serverless microservices platform FaaS (Funcitons as a Service) which uses Docker containers to run Any language or AWS Lambda functions
* [OpenFaaS](https://github.com/openfaas/faas) - A complete serverless functions framework for Docker and Kubernetes. By [OpenFaaS](https://github.com/openfaas)
* [SCAR](https://github.com/grycap/scar) - Serverless Container-aware Architectures (SCAR) is a serverless framework that allows easy deployment and execution of containers (e.g. Docker) in Serverless environments (e.g. Lambda) by [@grycap](https://github.com/grycap)

### Testing

* [Container Structure Test](https://github.com/GoogleCloudPlatform/container-structure-test) - A framework to validate the structure of an image by checking the outputs of commands or the contents of the filesystem. By [@GoogleCloudPlatform][GoogleCloudPlatform]
* [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss) - A fast YAML based tool for validating docker containers.
* [Dockunit](https://github.com/dockunit/platform) :skull: - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By [@dockunit](https://github.com/dockunit)
* [DockerSpec](https://github.com/zuazo/dockerspec) - A small Ruby Gem to run RSpec and Serverspec, Infrataster and Capybara tests against Dockerfiles or Docker images easily. By [@zuazo](https://github.com/zuazo)
* [Pumba](https://github.com/gaia-adm/pumba) - Chaos testing tool for Docker. Can be deployed on kubernetes and CoreOS cluster. By [@gaia](https://github.com/gaia-adm)
* [InSpec][inspec] - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements. By [@chef](https://github.com/chef)

### Wrappers

* [Ansible](http://docs.ansible.com/ansible/latest/docker_container_module.html) - Manage the life cycle of Docker containers. By RedHat
* [Azk](https://github.com/azukiapp/azk) - Orchestrate development enviornments on your local machine by [@azukiapp](https://github.com/azukiapp)
* [Beluga](https://github.com/cortexmedia/Beluga) :skull: - CLI to deploy docker containers on a single server or low amount of servers. By [@cortextmedia](https://github.com/cortexmedia)
* [dexec](https://github.com/docker-exec/dexec) - Command line interface written in Go for running code with Docker Exec images.
* [docker-do](https://github.com/benzaita/docker-do) - hassle-free docker run, like `env` but for docker by [@benzaita](https://github.com/benzaita)
* [Dray](https://github.com/CenturyLinkLabs/dray) - An engine for managing the execution of container-based workflows by [@CenturyLinkLabs][CenturyLinkLabs]
* [FuGu](https://github.com/mattes/fugu) - Docker run wrapper without orchestration by [@mattes](https://github.com/mattes)
* [SaltStack Docker module](https://docs.saltstack.com/en/latest/ref/modules/all/salt.modules.dockerng.html#module-salt.modules.dockerng) - SaltStack Docker module
* [Shutit](http://ianmiell.github.io/shutit/) - Tool for building and maintaining complex Docker deployments by [@ianmiell][ianmiell]
* [subuser](http://subuser.org) - Makes it easy to securely and portably run graphical desktop applications in Docker
* [Turbo](https://ramitsurana.github.io/turbo/) - Simple and Powerful utility for docker. By [@ramitsurana][ramitsurana]
* [Vagrant - Docker provider](https://www.vagrantup.com/docs/docker/basics.html) - Good starting point is [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example) by [@bubenkoff](https://github.com/bubenkoff)
* [udocker](https://github.com/indigo-dc/udocker) - A tool to execute simple docker containers in batch or interactive systems without root privileges by [@inidigo-dc](https://github.com/indigo-dc)


## Services based on Docker (:heavy_dollar_sign:)

### CI Services

* [CircleCI](https://circleci.com/) - Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
* [CodeFresh](https://codefresh.io) - Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
* [CodeShip](https://codeship.com/features/pro) - Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
* [ConcourseCI](https://concourse.ci) - A CI SaaS platform for developers and DevOps teams pipeline oriented.
* [Semaphore CI](https://semaphoreci.com/product/docker) — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
* [Shippable](https://app.shippable.com/) - A SaaS platform for developers and DevOps teams that significantly reduces the time taken for code to be built, tested and deployed to production.
* [IBM Bluemix Continous Delivery](https://console.bluemix.net/devops/getting-started?auth=false) - Continuous delivery using a pipeline deployment onto IBM containers on Bluemix.
* [TravisCI](https://travis-ci.org/) - A Free github projects continuous integration Saas platform for developers and Devops.
* [Wercker](http://www.wercker.com/) - A Docker-Native continous integration & deployment Automation platform for Kubernetes & Microservice Deployments.

### CaaS

* [Amazon ECS](https://aws.amazon.com/ecs/) - A management service on EC2 that supports Docker containers.
* [Arukas](https://arukas.io/) - Heroku-inspired CaaS
* [Azure ACS](https://azure.microsoft.com/en-us/services/container-service/) - A management service on Azure Virtual Machines that supports Docker containers.
* [Cloud 66](https://www.cloud66.com) - Full-stack hosted container management as a service
* [Codenvy](https://codenvy.com) - One-click Docker environments and cloud workspace for development teams
* [ContainerShip Cloud][containership] - Multi-Cloud Container Hosting Automation Platform.
* [Docker Cloud](https://cloud.docker.com/) - Former Tutum
* [Dockhero](https://dockhero.io/) - Dockhero is a Heroku add-on which turns a Docker image into a microservice attached to the Heroku app. Currently in beta.
* [Giant Swarm](https://giantswarm.io/) - Simple microservice infrastructure. Deploy your containers in seconds.
* [Google Container Engine](https://cloud.google.com/kubernetes-engine/docs/) - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
* [Hyper_](https://hyper.sh/) - Secure container hosting service with "nano-containers" and per-second billing.
* [IBM Bluemix Container Service](https://console.bluemix.net/docs/containers/container_index.html) - Run Docker containers in a hosted cloud environment on IBM Bluemix.
* [Jelastic Cloud](https://jelastic.cloud/) - "Easy-to-use" container hosting platfrom with automatic vertical and horizontal scaling. Available over 50+ hosting providers worldwide.
* [OpenShift Dedicated](https://www.openshift.com/dedicated/index.html) - A hosted [OpenShift][openshift] cluster for running your Docker containers managed by Red Hat.
* [Sloppy.io](https://sloppy.io/) - all-in-one solution for container deployment and hosting – made and hosted in Germany
* [Triton](https://www.joyent.com/) - Elastic container-native infrastructure by Joyent.

### Monitoring Services

* [AppDynamics](https://www.appdynamics.com/community/exchange/extension/docker-monitoring-extension/) - AppDynamics gives enterprises real-time insights into application performance, user performance, and business performance so they can move faster in an increasingly sophisticated, software-driven world.
* [Axibase Time-Series Database](http://axibase.com/products/axibase-time-series-database/writing-data/docker-cadvisor/) - Long-term retention of container statistics and built-in dashboards for Docker. Collected with native Google cAdvisor storage driver.
* [CA Technologies Docker Monitoring](https://www.ca.com/us/products/docker-monitoring.html) - Agile Operations solutions from CA deliver the modern Docker monitoring businesses need to accelerate and optimize the performance of microservices and the dynamic Docker environments running them. Monitor both the Docker environment and apps that run inside them.
* [Collecting docker logs and stats with Splunk](https://www.splunk.com/blog/2015/08/24/collecting-docker-logs-and-stats-with-splunk.html)
* [CoScale](https://www.coscale.com/docker-monitoring) - Full stack monitoring for containerized applications and microservices. Powered by anomaly detection to find performance problems faster.
* [Datadog](https://www.datadoghq.com/) - Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
* [Meros](https://meros.io) - Analyzes containers resources, captures logs, remote web SSH terminal and powerful DevOps alerts.
* [Prometheus](https://prometheus.io/) - Open-source service monitoring system and time series database
* [Sysdig](https://www.sysdig.org/) - An open source troubleshooting tool that provides a rich set of real-time, system-level information. It has container-specific features and is very useful in Docker environments.
* [Site24x7](https://www.site24x7.com/docker-monitoring.html) - Docker MOnitoring for DevOps and IT is a SaaS Pay per Host model
* [SPM for Docker][spm] - Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [@sematext][sematext]

# Useful Resources

* __[Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html)__ High quality articles about docker! __MUST SEE__ 
* [Cloud Native Landscape](https://github.com/cncf/landscape)
* [Docker Weekly](https://blog.docker.com/docker-weekly-archives/) Huge resource

Blogs by 
* [@codeship](https://blog.codeship.com/)  
* [@crosbymichael](http://crosbymichael.com/) 
* [@jessfraz](https://blog.jessfraz.com/) 
* [@gliderlabs](https://gliderlabs.com/devlog/) 
* [@jpetazzo](http://jpetazzo.github.io/) 
* [@jwilder](http://jasonwilder.com/) 
* [@nickjanetakis](https://nickjanetakis.com/blog/tag/docker) 
* [@progrium](http://progrium.com/blog/) 
* [@sebgoa](http://sebgoa.blogspot.be/) 
* [Container42](https://container42.com/)
* [Container solutions](http://container-solutions.com/blog/)


## Awesome Lists

* [Awesome-ciandcd](https://github.com/ciandcd/awesome-ciandcd) - Not specific to docker but relevant.
* [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes) by [@ramitsurana][ramitsurana]
* [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers) more general about container than this repo, by [@Friz-zy](https://github.com/Friz-zy).
* [Awesome Selfhosted](https://github.com/Kickball/awesome-selfhosted) list of Free Software network services and web applications which can be hosted locally by running in a classical way (setup local web server and run applications from there) or in a Docker container. By [@Kickball](https://github.com/Kickball)
* [Awesome Sysadmin](https://github.com/n1trux/awesome-sysadmin) by [@n1trux](https://github.com/n1trux)
* [ToolsOfTheTrade](https://github.com/cjbarber/ToolsOfTheTrade) a list of SaaS and On premise applications by [@cjbarber](https://github.com/cjbarber)

## Good Tips

* [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016) __MUST SEE__
* [Docker Containers on the Desktop][jessblog] - The **funniest way** to learn about docker by [@jessfraz][jessfraz] who also gave a [presentation][jessvid] about it @ DockerCon 2015
* [Container Best Practices](http://docs.projectatomic.io/container-best-practices/) - Red Hat's Project Atomic created a Container Best Practices guide which applies to everything and is updated regurlary.
* [Dealing with linked containers dependency in docker-compose](http://brunorocha.org/python/dealing-with-linked-containers-dependency-in-docker-compose.html) by [@rochacbruno](https://github.com/rochacbruno)
* [Don't Repeat Yourself with Anchors, Aliases and Extensions in Docker Compose Files](https://medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd) by [@King Chung Huang](https://github.com/kinghuang)
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm][fgrehm]
* [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](https://www.rightscale.com/blog/cloud-management-best-practices/docker-vs-vms-combining-both-cloud-portability-nirvana)


## Raspberry Pi & ARM

* [Docker Pirates ARMed with explosive stuff](http://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
* [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps)
* [git push docker containers to linux devices](https://resin.io/) Modern DevOps for IoT, leveraging git and Docker.
* [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-(ARMv7)-devices)

## Security

* [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
* [CVE Scanning Alpine images with Multi-stage builds in Docker 17.05](https://github.com/tomwillfixit/alpine-cvecheck) by [@tomwillfixit](https://twitter.com/tomwillfixit)
* [Docker and SELinux](http://www.projectatomic.io/docs/docker-and-selinux/)
* [Docker Secure Deployment Guidelines](https://github.com/GDSSecurity/Docker-Secure-Deployment-Guidelines)
* [Docker Security Cheat Sheet](https://github.com/konstruktoid/Docker/blob/master/Security/CheatSheet.adoc)
* [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](https://fr.slideshare.net/MichaelBoelen/docker-security-are-your-containers-tightly-secured-to-the-ship)
* [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448)
* [Improving Docker Security with Authenticated Volumes](https://www.blockbridge.com/improving-docker-security-with-authenticated-volumes/)
* [Lynis is an open source security auditing tool including Docker auditing](https://cisofy.com/lynis/)
* [Security Best Practices for Building Docker Images](https://linux-audit.com/tag/docker/)
* [Tuning Docker with the newest security enhancements](https://opensource.com/business/15/3/docker-security-tuning)
* [Understanding Docker security and best practices](https://blog.docker.com/2015/05/understanding-docker-security-and-best-practices/) (written 5 MAY 2015)

## Videos

* [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](https://www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
* [Deploying and scaling applications with Docker, Swarm, and a tiny bit of Python magic](https://www.youtube.com/watch?v=GpHMTR7P2Ms) (3:11:06) by [@jpetazzo][jpetazzo]
* [Docker and SELinux by Daniel Walsh from Red Hat ](https://www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
* [Docker for Developers][docker4dev] (54:26) by [@jpetazzo][jpetazzo]  <== Good introduction, context, demo
* [Docker in Production](https://www.youtube.com/watch?v=Glk5d5WP6MI) by [@jpetazzo][jpetazzo] (36:05)
* [Docker: How to Use Your Own Private Registry](https://www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
* [Extending Docker with Plugins](https://vimeo.com/110835013) (15:21)
* [From Local Docker Development to Production Deployments](https://www.youtube.com/watch?v=7CZFpHUPqXw) by [@jpetazzo][jpetazzo] @ AWS re:Invent 2015
* [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](https://www.youtube.com/watch?v=GaHzdqFithc) (42:04)
* [Introduction to Docker and containers](https://www.youtube.com/watch?v=ZVaRK10HBjo) (3:09:00) by [@jpetazzo][jpetazzo]
* [Logging on Docker: What You Need to Know][loggingDocker] (51:27)
* [Performance Analysis of Docker - Jeremy Eder](https://www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
* [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
* [State of containers: a debate with CoreOS, VMware and Google](https://www.youtube.com/watch?v=IiITP3yIRd8) (27:38)
* [SysAdminCasts: Introduction to Docker](https://sysadmincasts.com/episodes/31-introduction-to-docker) (15:49)

# Communities and Meetups

## Brazilian
* [Docker BR on Telegram](https://telegram.me/dockerbr)
* [Docker BR on Slack](http://docker-br.herokuapp.com) - Auto invite url

## Chinese
* [DockerOne](http://dockone.io/) Docker Community (in Chinese) by [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)

## English
* [Docker Community](https://www.docker.com/docker-community)
* [Docker On Line Meetup](https://www.meetup.com/Docker-Online-Meetup/)
* [Docker Reddit Community](https://www.reddit.com/r/docker/)
* [Docker Events](https://events.docker.com)

## Russian
+ [Docker Russian-speaking Community](https://t.me/docker_ru)

[ahmetalpbalkan]: https://github.com/ahmetalpbalkan
[alpine]: https://github.com/gliderlabs/docker-alpine
[anchore]: https://github.com/anchore
[arun-gupta]: https://github.com/arun-gupta
[blockbridge]: https://github.com/blockbridge
[brooklyn]: http://brooklyn.apache.org/
[calico]: https://github.com/projectcalico/calicoctl
[CenturyLinkLabs]: https://github.com/CenturyLinkLabs
[containership]: https://containership.io
[coreos]: https://github.com/coreos
[cpuguy83]: https://github.com/cpuguy83
[crosbymichael]: https://github.com/crosbymichael
[dimonomid]: https://github.com/dimonomid
[distribution]: https://github.com/docker/distribution
[docker-cheat-sheet]: https://github.com/wsargent/docker-cheat-sheet
[docker-compose]: https://docs.docker.com/compose/
[docker-quick-ref]: https://github.com/dimonomid/docker-quick-ref
[docker]: https://github.com/docker
[docker-for-mac]: https://docs.docker.com/docker-for-mac/
[docker-for-windows]: https://docs.docker.com/docker-for-windows/
[docker4dev]: https://www.youtube.com/watch?v=FdkNAjjO5yQ
[dokku]: https://github.com/dokku/dokku
[editREADME]: https://github.com/veggiemonk/awesome-docker/edit/master/README.md
[fgrehm]: https://github.com/fgrehm
[fluentd]: https://github.com/kiyoto/docker-fluentd
[GoogleCloudPlatform]: https://github.com/GoogleCloudPlatform
[gesellix]: https://github.com/gesellix
[gliderlabs]: https://github.com/gliderlabs
[gondor]: https://github.com/gondor
[grammarly]: https://github.com/grammarly
[hashicorp]: https://github.com/hashicorp
[ianmiell]: https://github.com/ianmiell
[inspec]: https://github.com/chef/inspec
[JensPiegsa]: https://github.com/JensPiegsa
[jessblog]: https://blog.jessfraz.com/post/docker-containers-on-the-desktop/
[jessvid]: https://www.youtube.com/watch?v=1qlLUf7KtAw
[jessfraz]: https://github.com/jessfraz
[jessfrazdockerfiles]: https://github.com/jessfraz/dockerfiles
[jessfrazdotfiles]: https://github.com/jessfraz/dotfiles
[jpetazzo]: https://github.com/jpetazzo
[jwilder]: https://github.com/jwilder
[kartar]: https://twitter.com/kartar
[kiyoto]: https://github.com/kiyoto
[kubernetes]: https://kubernetes.io
[labelschema]: http://label-schema.org
[loggingDocker]: https://vimeo.com/123341629
[microbadger]: https://microbadger.com
[nginxproxy]: https://github.com/jwilder/nginx-proxy
[nickjanetakis]: https://twitter.com/nickjanetakis
[noteed]: https://github.com/noteed
[ondrejmo]: https://github.com/ondrejmo
[openshift]: https://www.openshift.org/
[panamax.io]: http://panamax.io/
[pandrew]: https://github.com/pandrew
[percheron]: https://github.com/ashmckenzie/percheron
[progrium]: https://github.com/progrium
[prologic]: https://github.com/prologic
[ramitsurana]: https://github.com/ramitsurana
[rancher]: https://github.com/rancher
[sebgoa]: https://twitter.com/sebgoa
[sematext]: https://twitter.com/sematext
[sindresorhus]: https://github.com/sindresorhus/awesome
[spotify]: https://github.com/spotify
[spm]: https://github.com/sematext/sematext-agent-docker
[vegasbrianc]: https://github.com/vegasbrianc
[vfarcic]: https://github.com/vfarcic
[vimagick]: https://github.com/vimagick
[weave]: https://github.com/weaveworks/weave
[wsargent]: https://github.com/wsargent
