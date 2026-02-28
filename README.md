# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)][sindresorhus] [![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/veggiemonk/awesome-docker/)[![Last Commit](https://img.shields.io/github/last-commit/veggiemonk/awesome-docker)](https://github.com/veggiemonk/awesome-docker/commits/main)<!-- omit in toc -->

> A curated list of Docker resources and projects

If you would like to contribute, please read [CONTRIBUTING.md][contributing] first.
It contains a lot of tips and guidelines to help keep things organized.
Just click [README.md][editreadme] to submit a [pull request][editreadme].
If this list is not complete, you can [contribute][editreadme] to make it so. Here is a great video tutorial to learn how to [contribute on Github](https://egghead.io/lessons/javascript-identifying-how-to-contribute-to-an-open-source-project-on-github).

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for newcomers. See how to **[Contribute][contributing]** for tips!

**_If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editreadme] to improve this file. Thank you!_**

The creators and maintainers of this list do not receive any form of payment to accept a change made by any contributor. This page is not an official Docker product in any way. It is a list of links to projects and is maintained by volunteers. Everybody is welcome to contribute. The goal of this repo is to index open-source projects, not to advertise for profit.

# Contents <!-- omit in toc -->

<!-- TOC -->

- [Legend](#legend)
- [What is Docker](#what-is-docker)
- [Where to start](#where-to-start)
- [Where to start (Windows)](#where-to-start-windows)
- [Projects](#projects)
  - [Container Operations](#container-operations)
    - [Container Composition](#container-composition)
    - [Deployment and Infrastructure](#deployment-and-infrastructure)
    - [Monitoring](#monitoring)
    - [Networking](#networking)
    - [Orchestration](#orchestration)
    - [PaaS](#paas)
    - [Reverse Proxy](#reverse-proxy)
    - [Runtime](#runtime)
    - [Security](#security)
    - [Service Discovery](#service-discovery)
    - [Volume Management / Data](#volume-management--data)
    - [User Interface](#user-interface)
      - [IDE integrations](#ide-integrations)
      - [Desktop](#desktop)
      - [Terminal](#terminal)
        - [Terminal UI](#terminal-ui)
        - [CLI tools](#cli-tools)
        - [Other](#other)
      - [Web](#web)
  - [Docker Images](#docker-images)
    - [Base Tools](#base-tools)
    - [Builder](#builder)
    - [Dockerfile](#dockerfile)
    - [Linter](#linter)
    - [Metadata](#metadata)
    - [Registry](#registry)
  - [Development with Docker](#development-with-docker)
    - [API Client](#api-client)
    - [CI/CD](#cicd)
    - [Development Environment](#development-environment)
    - [Garbage Collection](#garbage-collection)
    - [Serverless](#serverless)
    - [Testing](#testing)
    - [Wrappers](#wrappers)
  - [Services based on Docker (mostly :heavy\_dollar\_sign:)](#services-based-on-docker-mostly-heavy_dollar_sign)
    - [CI Services](#ci-services)
    - [CaaS](#caas)
    - [Monitoring Services](#monitoring-services)
- [Useful Resources](#useful-resources)
  - [Awesome Lists](#awesome-lists)
  - [Demos and Examples](#demos-and-examples)
  - [Good Tips](#good-tips)
  - [Raspberry Pi \& ARM](#raspberry-pi--arm)
  - [Security](#security-1)
  - [Videos](#videos)
- [Communities and Meetups](#communities-and-meetups)
  - [Brazilian](#brazilian)
  - [English](#english)
  - [Russian](#russian)
  - [Spanish](#spanish)
  - [Stargazers over time](#stargazers-over-time)

<!-- /TOC -->

# Legend

-   Monetized :heavy_dollar_sign:

# What is Docker

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ [What is Docker](https://www.docker.com/why-docker/)

# Where to start

-   [Benefits of using Docker](https://semaphore.io/blog/docker-benefits) for development and delivery, with a practical roadmap for adoption.
- [Bootstrapping Microservices](https://www.manning.com/books/bootstrapping-microservices-with-docker-kubernetes-and-terraform) - A practical and project-based guide to building applications with microservices, starts by building a Docker image for a single microservice and publishing it to a private container registry, finishes by deploying a complete microservices application to a production Kubernetes cluster.
-   [Docker Curriculum](https://github.com/prakhar1989/docker-curriculum): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
-   [Docker Documentation](https://docs.docker.com/): the official documentation.
-   [Docker for beginners](https://github.com/groda/big_data/blob/master/docker_for_beginners.md): A tutorial for beginners who need to learn the basics of Docker—from "Hello world!" to basic interactions with containers, with simple explanations of the underlying concepts.
-   [Docker for novices](https://www.youtube.com/watch?v=xsjSadjKXns) An introduction to Docker for developers and testers who have never used it. (Video 1h40, recorded linux.conf.au 2019 — Christchurch, New Zealand) by Alex Clews.

-   [Docker katas](https://github.com/eficode-academy/docker-katas) A series of labs that will take you from "Hello Docker" to deploying a containerized web application to a server.
-   [Docker simplified in 55 seconds](https://www.youtube.com/watch?v=vP_4DlOH1G4): An animated high-level introduction to Docker. Think of it as a visual tl;dr that makes it easier to dive into more complex learning materials.
-   [Docker Training](https://training.mirantis.com) :heavy_dollar_sign:
-   [Dockerlings](https://github.com/furkan/dockerlings): Learn docker from inside your terminal, with a modern TUI and bite sized exercises (by [furkan](https://github.com/furkan))

-   [Introduction à Docker](https://blog.stephane-robert.info/docs/conteneurs/moteurs-conteneurs/docker/) A dedicated section to master Docker on a French site about DevSecOps: From the basics to best practices, including optimizing, securing your containers...
-   [Learn Docker](https://github.com/dwyl/learn-docker): step-by-step tutorial and more resources (video, articles, cheat sheets) by [dwyl](https://github.com/dwyl)
- [Learn Docker (Visually)](https://pagertree.com/learn/docker/overview) - A beginner-focused high-level overview of all the major components of Docker and how they fit together. Lots of high-quality images, examples, and resources.
-   [Play With Docker](https://training.play-with-docker.com/): PWD is a great way to get started with Docker from beginner to advanced users. Docker runs directly in your browser.
-   [Practical Guide about Docker Commands in Spanish](https://github.com/brunocascio/docker-espanol) This Spanish guide contains the use of basic docker commands with real life examples.
-   [Setting Python Development Environment with VScode and Docker](https://github.com/RamiKrispin/vscode-python): A step-by-step tutorial for setting up a dockerized Python development environment with VScode, Docker, and the Dev Container extension.
-   [The Docker Handbook](https://docker-handbook.farhan.dev/) An open-source book that teaches you the fundamentals, best practices and some intermediate Docker functionalities. The book is hosted on [fhsinchy/the-docker-handbook](https://github.com/fhsinchy/the-docker-handbook) and the projects are hosted on [fhsinchy/docker-handbook-projects](https://github.com/fhsinchy/docker-handbook-projects) repository.

**Cheatsheets** by

-   [eon01](https://github.com/eon01/DockerCheatSheet)
-   [dimonomid](https://github.com/dimonomid/docker-quick-ref) (PDF)
-   [JensPiegsa](https://github.com/JensPiegsa/docker-cheat-sheet)
-   [wsargent](https://github.com/wsargent/docker-cheat-sheet) (Most popular)

# Where to start (Windows)

-   [Docker on Windows behind a firewall](https://toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by [kaitoedter](https://twitter.com/kaitoedter)
- [Docker Reference Architecture: Modernizing Traditional .NET Framework Applications](https://docs.mirantis.com/containers/v3.0/dockeree-ref-arch/app-dev/modernize-dotnet-apps.html) - You will learn to identify the types of .NET Framework applications that are good candidates for containerization, the "lift-and-shift" approach to containerization.
-   [Docker with Microsoft SQL 2016 + ASP.NET](https://blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
-   [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](https://www.hanselman.com/blog/exploring-aspnet-core-with-docker-in-both-linux-and-windows-containers) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows][docker-for-windows]
-   [Running a Legacy ASP.NET App in a Windows Container](https://blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and running as a Windows container
- [Windows Containers and Docker: The 101](https://www.youtube.com/watch?v=N7SG2wEyQtM) - A 20-minute overview, using Docker to run PowerShell, ASP.NET Core and ASP.NET apps.
-   [Windows Containers Quick Start](https://learn.microsoft.com/en-us/virtualization/windowscontainers/about/) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016

---

# Projects

-   Moby = open source development
-   Docker CE = free product release based on Moby
-   Docker EE = commercial product release based on Docker CE.

> Docker EE is on the same code base as Docker CE, so also built from Moby, with commercial components added, such as "docker data center / universal control plane"

-   [Moby](https://github.com/moby/moby)
-   [Docker Images](https://hub.docker.com)
-   [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
-   [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)

## Container Operations

### Container Composition

- [Capitan](https://github.com/byrnedo/capitan) - Composable docker orchestration with added scripting support by [byrnedo].
- [Composerize](https://github.com/magicmark/composerize) - Convert docker run commands into docker-compose files.
- [crowdr](https://github.com/polonskiy/crowdr) - Tool for managing multiple Docker containers (`docker-compose` alternative).
- [ctk](https://github.com/ctk-hq/ctk) - Visual composer for container based workloads. By [corpulent](https://github.com/corpulent).
- [docker-config-update](https://github.com/sudo-bmitch/docker-config-update) - Utility to update docker configs and secrets for deploying in a compose file.
- [elsy](https://github.com/cisco/elsy) - An opinionated, multi-language, build tool based on Docker and Docker Compose.
- [habitus](https://github.com/cloud66-oss/habitus) - A Build Flow Tool for Docker.
- [kompose](https://github.com/kubernetes/kompose) - Go from Docker Compose to Kubernetes.
- [LLM Harbor](https://github.com/av/harbor) - A CLI and companion app to effortlessly run LLM backends, APIs, frontends, and services with one command. By [av](https://github.com/av).
- [plash](https://github.com/ihucos/plash) - A container run and build engine - runs inside docker.
- [podman-compose](https://github.com/containers/podman-compose) - A script to run docker-compose.yml using podman.
-   [Smalte](https://github.com/roquie/smalte) – Dynamically configure applications that require static configuration in docker container. By [roquie](https://github.com/roquie)
- [Stitchocker](https://github.com/alexaandrov/stitchocker) - A lightweight and fast command line utility for conveniently grouping your docker-compose multiple container services. By [alexaandrov](https://github.com/alexaandrov).

### Deployment and Infrastructure

- [awesome-stacks](https://github.com/ethibox/awesome-stacks) - Deploy 150+ open-source web apps with one Docker command.
- [blackfish](https://gitlab.com/blackfish/blackfish) - A CoreOS VM to build swarm clusters for Dev & Production.
- [BosnD](https://gitlab.com/n0r1sk/bosnd) - BosnD, the boatswain daemon - A dynamic configuration file writer & service reloader for dynamically changing container environments.
- [Centurion](https://github.com/newrelic/centurion) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [newrelic](https://github.com/newrelic).
- [Clocker](https://github.com/brooklyncentral/clocker) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and [Brooklyn](https://brooklyn.apache.org/) for application blueprints. By [brooklyncentral](https://github.com/brooklyncentral).
- [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker.
- [depcon](https://github.com/ContainX/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes. By [ContainX][containx].
- [docker-to-iac](https://github.com/deploystackio/docker-to-iac) - Translate docker run and commit into Infrastructure as Code templates for AWS, Render.com and DigitalOcean.
- [gitkube](https://github.com/hasura/gitkube) - Gitkube is a tool for building and deploying docker images on Kubernetes using `git push`. By [Hasura](https://github.com/hasura/).
- [Grafeas](https://github.com/grafeas/grafeas) - A common API for metadata about containers, from image and build details to security vulnerabilities. By [grafeas](https://github.com/grafeas).
- [swarm-ansible](https://github.com/LombardiDaniel/swarm-ansible?tab=readme-ov-file) - Swarm-Ansible bootstraps a production-ready swarm cluster using ansible. Comes with tools to automate CI, help monitoring and traefik pre-configured for SSL certificates and simple-auth. Comes with a private registry and more!.
- [SwarmManagement](https://github.com/hansehe/SwarmManagement) - Swarm Management is a python application, installed with pip. The application makes it easy to manage a Docker Swarm by configuring a single yaml file describing which stacks to deploy, and which networks, configs or secrets to create.
- [werf](https://github.com/werf/werf) - Werf is a CI/CD tool for building Docker images efficiently and deploying them to Kubernetes using GitOps.

### Monitoring

- [Autoheal](https://github.com/willfarrell/docker-autoheal) - Monitor and restart unhealthy docker containers automatically.
- [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers.
- [Checkmate](https://github.com/bluewave-labs/checkmate) - Checkmate is an open-source, self-hosted tool designed to track and monitor server hardware, uptime, response times, and incidents in real-time with beautiful visualizations.
- [DLIA](https://github.com/zorak1103/dlia) - DLIA is an AI-powered Docker log monitoring agent that uses Large Language Models (LLMs) to intelligently analyze container logs, detect anomalies, and provide contextual insights over time. By [zorak1103](https://github.com/zorak1103).
- [Docker-Alertd](https://github.com/deltaskelta/docker-alertd) - Monitor and send alerts based on docker container resource usage/statistics.
- [Docker-Flow-Monitor](https://github.com/docker-flow/docker-flow-monitor) - Reconfigures Prometheus when a new service is updated or deployed automatically.
- [DockProc](https://gitlab.com/n0r1sk/dockproc) - I/O monitoring for containers on processlevel.
- [dockprom](https://github.com/stefanprodan/dockprom) - Docker hosts and containers monitoring with Prometheus, Grafana, cAdvisor, NodeExporter and AlertManager.
- [Doku](https://github.com/amerkurev/doku) - Doku is a simple web-based application that allows you to monitor Docker disk usage. [amerkurev](https://github.com/amerkurev).
- [Dozzle](dozzle) - Monitor container logs in real-time with a browser or mobile device. [amir20](https://github.com/amir20).
- [Dynatrace](https://docs.dynatrace.com/docs/observe/infrastructure-observability/container-platform-monitoring) - :heavy_dollar_sign: Monitor containerized applications without installing agents or modifying your Run commands.
- [Glances](https://github.com/nicolargo/glances) - A cross-platform curses-based system monitoring tool written in Python.
- [Grafana Docker Dashboard Template](https://grafana.com/grafana/dashboards/179-docker-prometheus-monitoring/) - A template for your Docker, Grafana and Prometheus stack [vegasbrianc][vegasbrianc].
- [HertzBeat](https://github.com/dromara/hertzbeat) - An open-source real-time monitoring system with custom-monitor and agentless.
- [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) - InfluxDB Time series DB in combination with Grafana and cAdvisor.
- [Logspout](https://github.com/gliderlabs/logspout) - Log routing for Docker container logs.
- [monit-docker](https://github.com/decryptus/monit-docker) - Monitor docker containers resources usage or status and execute docker commands or inside containers. [decryptus][decryptus].
- [NexClipper](https://github.com/NexClipper/NexClipper) - NexClipper is the container monitoring and performance management solution specialized in Docker, Apache Mesos, Marathon, DC/OS, Mesosphere, Kubernetes.
- [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting) - Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].
- [Sidekick](https://github.com/runsidekick/sidekick) - Open source live application debugger like Chrome DevTools for your backend. Collect traces and generate logs on-demand without stopping & redeploying your applications.
- [SwarmAlert](https://github.com/gpulido/SwarmAlert) - Monitors a Docker Swarm and sends Pushover alerts when it finds a container with no healthy service task running.
- [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.
- [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.

### Networking

-   [Calico][calico] - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
- [Flannel](https://github.com/coreos/flannel/) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes. By [coreos][coreos].
- [Freeflow](https://github.com/Microsoft/Freeflow) - High performance container overlay networks on Linux. Enabling RDMA (on both InfiniBand and RoCE) and accelerating TCP to bare metal performance. By [Microsoft](https://github.com/Microsoft).
- [MyIP](https://github.com/jason5ng32/MyIP) - All in one IP Toolbox. Easy to check all your IPs, IP geolocation, check for DNS leaks, examine WebRTC connections, speed test, ping test, MTR test, check website availability, whois search and more. By [jason5ng32](https://github.com/jason5ng32).
- [netshoot](https://github.com/nicolaka/netshoot) - The netshoot container has a powerful set of networking tools to help troubleshoot Docker networking issues.
- [Pipework](https://github.com/jpetazzo/pipework) - Software-Defined Networking for Linux Containers, Pipework works with "plain" LXC containers, and with the awesome Docker. By [jpetazzo][jpetazzo].

### Orchestration

- [Ansible Linux Docker](https://github.com/Peco602/ansible-linux-docker) - Run Ansible from a Linux container. By [Peco602][peco602].
- [athena](https://github.com/athena-oss/athena) - An automation platform with a plugin architecture that allows you to easily create and share services.
- [CloudSlang](https://github.com/CloudSlang/cloud-slang) - CloudSlang is a workflow engine to create Docker process automation.
- [clusterdock](https://github.com/clusterdock/clusterdock) - Docker container orchestration to enable the testing of long-running cluster deployments.
- [Crane](https://github.com/Dataman-Cloud/crane) - Control plane based on docker built-in swarm [Dataman-Cloud](https://github.com/Dataman-Cloud).
- [Docker Flow Swarm Listener](https://github.com/docker-flow/docker-flow-swarm-listener) - Docker Flow Swarm Listener project is to listen to Docker Swarm events and send requests when a change occurs. By [docker-flow][docker-flow].
- [docker rollout](https://github.com/Wowu/docker-rollout) - Zero downtime deployment for Docker Compose services.
- [Haven](https://github.com/codeabovelab/haven-platform) - Haven is a simplified container management platform that integrates container, application, cluster, image, and registry managements. By [codeabovelab](https://github.com/codeabovelab).
- [Kubernetes](https://github.com/kubernetes/kubernetes) - Open source orchestration system for Docker containers by Google.
- [ManageIQ](https://github.com/ManageIQ/manageiq) - Discover, optimize and control your hybrid IT. By [ManageIQ](https://github.com/ManageIQ).
- [Mesos](https://github.com/apache/mesos) - Resource/Job scheduler for containers, VM's and physical hosts [apache](https://mesos.apache.org/).
- [Nebula](https://github.com/nebula-orchestrator) - A Docker orchestration tool designed to manage massive scale distributed clusters.
- [Nomad](https://github.com/hashicorp/nomad) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler.
- [Rancher](https://github.com/rancher/rancher) - An open source project that provides a complete platform for operating Docker in production.
- [RedHerd Framework](https://github.com/redherd-project/redherd-framework) - RedHerd is a collaborative and serverless framework for orchestrating a geographically distributed group of assets capable of simulating complex offensive cyberspace operations. By [RedHerdProject](https://github.com/redherd-project).
- [Swarm-cronjob](https://github.com/crazy-max/swarm-cronjob) - Create jobs on a time-based schedule on Swarm by [crazy-max].

### PaaS

- [caprover](https://github.com/caprover/caprover) - [Previously known as CaptainDuckDuck] Automated Scalable Webserver Package (automated Docker+nginx) - Heroku on Steroids.
- [Convox Rack](https://github.com/convox/rack) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
- [Dcw](https://github.com/pbertera/dcw) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels.
- [Dokku](https://github.com/dokku/dokku) - Docker powered mini-Heroku that helps you build and manage the lifecycle of applications (originally by [progrium][progrium]).
- [Empire](https://github.com/remind101/empire) - A PaaS built on top of Amazon EC2 Container Service (ECS).
- [Exoframe](https://github.com/exoframejs/exoframe) - A self-hosted tool that allows simple one-command deployments using Docker.
- [Hephy Workflow](https://github.com/teamhephy/workflow) - Open source PaaS for Kubernetes that adds a developer-friendly layer to any Kubernetes cluster, making it easy to deploy and manage applications. Fork of [Deis Workflow](https://github.com/deis/workflow).
- [Krane](https://github.com/krane/krane) - Toolset for managing container workloads on remote servers.
- [Nanobox](https://github.com/nanobox-io/nanobox) - :heavy_dollar_sign: An application development platform that creates local environments that can then be deployed and scaled in the cloud.
-   [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by [Red Hat](https://www.redhat.com/en)
- [Tsuru](https://github.com/tsuru/tsuru) - Tsuru is an extensible and open source Platform as a Service software.

### Reverse Proxy

- [BunkerWeb](https://github.com/bunkerity/bunkerweb) - Open-source and next-gen Web Application Firewall (WAF). By [Bunkerity](https://www.bunkerity.com).
- [caddy-docker-proxy](https://github.com/lucaslorentz/caddy-docker-proxy) - Caddy-based reverse proxy, configured with service or container labels. By [lucaslorentz](https://github.com/lucaslorentz).
- [caddy-docker-upstreams](https://github.com/invzhi/caddy-docker-upstreams) - Docker upstreams module for Caddy, configured with container labels. By [invzhi](https://github.com/invzhi).
- [Docker Dnsmasq Updater](https://github.com/moonbuggy/docker-dnsmasq-updater) - Update a remote dnsmasq server with Docker container hostnames.
- [docker-flow-proxy](https://github.com/docker-flow/docker-flow-proxy) - Reconfigures proxy every time a new service is deployed, or when a service is scaled. By [docker-flow][docker-flow].
- [fabio](https://github.com/fabiolb/fabio) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By [magiconair](https://github.com/magiconair) (Frank Schroeder).
- [Let's Encrypt Nginx-proxy Companion](https://github.com/nginx-proxy/docker-letsencrypt-nginx-proxy-companion) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By [JrCs](https://github.com/JrCs).
- [mesh-router](https://github.com/Yundera/mesh-router) - Free domain(nsl.sh) provider for Docker containers with automatic HTTPS routing. Uses Wireguard VPN to securely route subdomain requests across networks. Ideal for self-hosted NAS and cloud deployments. By [Yundera](https://github.com/Yundera).
- [Nginx Proxy Manager](https://github.com/jc21/nginx-proxy-manager) - A beautiful web interface for proxying web based services with SSL. By [jc21](https://github.com/jc21).
-   [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [jwilder][jwilder]
- [OpenResty Manager](https://github.com/Safe3/openresty-manager) - The easiest using, powerful and beautiful OpenResty Manager(Nginx Enhanced Version), open source alternative to OpenResty Edge. By [Safe3](https://github.com/Safe3/).
- [Swarm Router](https://github.com/flavioaiello/swarm-router) - A «zero config» service name based router for docker swarm mode with a fresh and more secure approach. By [flavioaiello](https://github.com/flavioaiello).
- [Træfɪk](https://github.com/containous/traefik) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [EmileVauge](https://github.com/emilevauge).

### Runtime

- [cri-o](https://github.com/cri-o/cri-o) - Open Container Initiative-based implementation of Kubernetes Container Runtime Interface by [cri-o](https://github.com/cri-o).
- [lxc](https://github.com/lxc/lxc) - LXC - Linux Containers.
- [podman](https://github.com/containers/libpod) - Libpod is a library used to create container pods. Home of Podman.
- [rlxc](https://github.com/brauner/rlxc) - LXC binary written in Rust.
- [runtime-tools](https://github.com/opencontainers/runtime-tools) - Oci-runtime-tool is a collection of tools for working with the OCI runtime specification.

### Security

- [Anchor](https://github.com/SongStitch/anchor/) - A tool to ensure reproducible builds by pinning dependencies inside your Dockerfiles [SongStitch](https://github.com/songStitch/).
- [Anchor Enterprise](https://anchore.com/) - :heavy_dollar_sign: Analyze images for CVE vulnerabilities and against custom security policies.
- [Aqua Security](https://www.aquasec.com) - :heavy_dollar_sign: Securing container-based applications from Dev to Production on any platform.
- [bane](https://github.com/genuinetools/bane) - AppArmor profile generator for Docker containers.
- [buildcage](https://github.com/dash14/buildcage) - Restricts outbound network access during Docker builds to prevent supply chain attacks, working as a drop-in BuildKit remote driver for Docker Buildx, with ready-to-use GitHub Actions.
- [CetusGuard](https://github.com/hectorm/cetusguard) - CetusGuard is a tool that protects the Docker daemon socket by filtering calls to its API endpoints.
- [Checkov](https://github.com/bridgecrewio/checkov) - Static analysis for infrastructure as code manifests (Terraform, Kubernetes, Cloudformation, Helm, Dockerfile, Kustomize) find security misconfiguration and fix them. By [bridgecrew](https://github.com/bridgecrewio).
- [CIS Docker Benchmark](https://github.com/dev-sec/cis-docker-benchmark) - This [InSpec][inspec] compliance profile implement the CIS Docker 1.12.0 Benchmark in an automated way to provide security best-practice tests around Docker daemon and containers in a production environment. By [dev-sec](https://github.com/dev-sec).
- [Clair](https://github.com/quay/clair) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [coreos][coreos].
- [crowdsec-blocklist-import](https://github.com/wolffcatskyy/crowdsec-blocklist-import) - Aggregates 36 free threat intelligence feeds into 120k+ malicious IPs for CrowdSec bouncers, providing 10-20x more blocks than default lists. By [wolffcatskyy](https://github.com/wolffcatskyy).
- [Dagda](https://github.com/eliasgranderubio/dagda) - Dagda is a tool to perform static analysis of known vulnerabilities, trojans, viruses, malware & other malicious threats in docker images/containers and to monitor the docker daemon and running docker containers for detecting anomalous activities. By [eliasgranderubio](https://github.com/eliasgranderubio).
- [Deepfence Enterprise](https://deepfence.io) - :heavy_dollar_sign: Full life cycle Cloud Native Workload Protection platform for kubernetes, virtual machines and serverless. By [deepfence][deepfence].
- [Deepfence Threat Mapper](https://github.com/deepfence/ThreatMapper) - Powerful runtime vulnerability scanner for kubernetes, virtual machines and serverless. By [deepfence][deepfence].
- [docker-bench-security](https://github.com/docker/docker-bench-security) - Script that checks for dozens of common best-practices around deploying Docker containers in production. By [docker][docker].
- [docker-explorer](https://github.com/google/docker-explorer) - A tool to help forensicate offline docker acquisitions.
- [docker-lock](https://github.com/safe-waters/docker-lock) - A cli-plugin for docker to automatically manage image digests by tracking them in a separate Lockfile. By [safe-waters][safe-waters].
- [dvwassl](https://github.com/Peco602/dvwassl) - SSL-enabled Damn Vulnerable Web App to test Web Application Firewalls. By [Peco602][peco602].
- [KICS](https://github.com/checkmarx/kics) - An infrastructure-as-code scanning tool, find security vulnerabilities, compliance issues, and infrastructure misconfigurations early in the development cycle. Can be extended for additional policies. By [Checkmarx](https://github.com/Checkmarx).
- [notary](https://github.com/theupdateframework/notary) - A server and a client for running and interacting with trusted collections. By [TUF](https://github.com/theupdateframework).
- [oscap-docker](https://github.com/OpenSCAP/openscap) - OpenSCAP provides oscap-docker tool which is used to scan Docker containers and images. By [OpenSCAP](https://github.com/OpenSCAP).
- [Prisma Cloud](https://www.paloaltonetworks.com/prisma/cloud) - :heavy_dollar_sign: (Previously Twistlock Security Suite) detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.
- [segspec](https://github.com/dormstern/segspec) - Extracts network dependencies from Docker Compose, Kubernetes manifests, Helm charts, and other config files to generate Kubernetes NetworkPolicies with evidence tracing. By [dormstern](https://github.com/dormstern).
- [Syft](https://github.com/anchore/syft) - CLI tool and library for generating a Software Bill of Materials (SBOM) from container images and filesystems.
- [Sysdig Falco](https://github.com/falcosecurity/falco) - Sysdig Falco is an open source container security monitor. It can monitor application, container, host, and network activity and alert on unauthorized activity.
- [Sysdig Secure](https://www.sysdig.com/solutions/cloud-detection-and-response-cdr) - :heavy_dollar_sign: Sysdig Secure addresses run-time security through behavioral monitoring and defense, and provides deep forensics based on open source Sysdig for incident response.
- [Trend Micro DeepSecurity](https://www.trendmicro.com/en_us/business/products/hybrid-cloud/deep-security.html) - :heavy_dollar_sign: Trend Micro DeepSecurity offers runtime protection for container workloads and hosts as well as preruntime scanning of images to identify vulnerabilities, malware and content such as hardcoded secrets.
- [Trivy](https://github.com/aquasecurity/trivy) - Aqua Security's open source simple and comprehensive vulnerability scanner for containers (suitable for CI).

### Service Discovery

-   [docker-consul](https://github.com/gliderlabs/docker-consul) by [progrium][progrium]
- [docker-dns](https://github.com/bytesharky/docker-dns) - Lightweight DNS forwarder for Docker containers, resolves container names with custom suffixes (e.g. `.docker`) on the host to simplify service discovery.
- [etcd](https://github.com/etcd-io/etcd) - Distributed reliable key-value store for the most critical data of a distributed system by [etcd-io](https://github.com/etcd-io) (former part of CoreOS).
- [istio](https://github.com/istio/istio) - An open platform to connect, manage, and secure microservices.
- [registrator](https://github.com/gliderlabs/registrator) - Service registry bridge for Docker by [gliderlabs][gliderlabs] and [progrium][progrium].

### Volume Management / Data

-   [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume) :heavy_dollar_sign:- The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By [blockbridge](https://github.com/blockbridge)
-   - [Label Backup](https://github.com/resulgg/label-backup) - A lightweight, Docker-aware backup agent that automatically discovers and backs up containerized databases (PostgreSQL, MySQL, MongoDB, Redis) based on Docker labels. Supports local storage and S3-compatible destinations with flexible scheduling via cron expressions.
-   [Docker Volume Backup](https://github.com/offen/docker-volume-backup) Backup Docker volumes locally or to any S3 compatible storage. By [offen](https://github.com/offen)
- [Minio](https://github.com/minio/minio) - S3 compatible object storage server in Docker containers.
-   [Netshare](https://github.com/ContainX/docker-volume-netshare) Docker NFS, AWS EFS, Ceph & Samba/CIFS Volume Plugin. By [ContainX][containx]
- [portworx](https://portworx.com) - :heavy_dollar_sign: Decentralized storage solution for persistent, shared and replicated volumes.
- [quobyte](https://www.quobyte.com/) - :heavy_dollar_sign: Fully fault-tolerant distributed file system with a docker volume driver.
-   [REX-Ray](https://github.com/rexray/rexray) provides a vendor agnostic storage orchestration engine. The primary design goal is to provide persistent storage for Docker, Kubernetes, and Mesos. By[thecodeteam](https://github.com/thecodeteam) (DELL Technologies)
- [Storidge](https://github.com/Storidge/quick-start) - :heavy_dollar_sign: Software-defined Persistent Storage for Kubernetes and Docker Swarm.

### User Interface

#### IDE integrations

-   JetBrains IDEs (IntelliJ IDEA, GoLand, WebStorm, CLion etc.) has [built-in Docker plugin](https://www.jetbrains.com/help/idea/docker.html#managing-images)
-   Eclipse [Docker Tooling plugin](https://www.eclipse.org/community/eclipse_newsletter/2016/july/article2.php)
-   [docker.el](https://github.com/Silex/docker.el) Manage docker from Emacs by [Silex](https://github.com/Silex)

#### Desktop

Native desktop applications for managing and monitoring docker hosts and clusters

- [Docker DB Manager](https://github.com/AbianS/docker-db-manager) - Desktop app for managing Docker database containers with visual interface and one-click operations.
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) - Official native app. Only for Windows and MacOS.
- [Simple Docker UI](https://github.com/felixgborrego/simple-docker-ui) - Built on Electron. By [felixgborrego](https://github.com/felixgborrego/).
- [Stevedore](https://github.com/slonopotamus/stevedore) - Good Docker Desktop replacement for Windows. Both Linux and Windows Containers are supported. [slonopotamus](https://github.com/slonopotamus).

#### Terminal

##### Terminal UI
- [d4s](https://github.com/jr-k/d4s) - A fast, keyboard-driven terminal UI to manage Docker containers, Compose stacks, and Swarm services with the ergonomics of K9s.
- [dive](https://github.com/wagoodman/dive) - A tool for exploring each layer in a docker image. By [wagoodman](https://github.com/wagoodman).
-   [dockdash](https://github.com/byrnedo/dockdash) detailed stats. By [byrnedo]
- [dockly](https://github.com/lirantal/dockly) - An interactive shell UI for managing Docker containers.
- [DockMate](https://github.com/shubh-io/dockmate) - Lightweight terminal-based Docker and Podman manager with a text-based user interface,.
- [DockSTARTer](https://github.com/GhostWriters/DockSTARTer) - DockSTARTer helps you get started with home server apps running in Docker by [GhostWriters](https://github.com/GhostWriters).
- [dprs](https://github.com/durableprogramming/dprs) - A developer-focused TUI for managing Docker containers with real-time log streaming and container management. Built with Rust. By [durableprogramming](https://github.com/durableprogramming).
- [dry](https://github.com/moncho/dry) - An interactive CLI for Docker containers.
- [goManageDocker](https://github.com/ajayd-san/gomanagedocker) - TUI tool to view and manage your docker objects blazingly fast with sensible keybindings, also supports VIM navigation out of the box.
- [lazydocker](https://github.com/jesseduffield/lazydocker) - The lazier way to manage everything docker. A simple terminal UI for both docker and docker-compose, written in Go with the gocui library. By [jesseduffield](https://github.com/jesseduffield).
- [lazyjournal](https://github.com/Lifailon/lazyjournal) - A interface for reading and filtering the logs output of Docker and Podman containers like [Dozzle](dozzle) but for the terminal with support for fuzzy find, regex and output coloring.
- [oxker](https://github.com/mrjackwills/oxker) - A simple tui to view & control docker containers. Written in [Rust](https://rust-lang.org/), making heavy use of [ratatui](https://github.com/tui-rs-revival/ratatui) & [Bollard](https://github.com/fussybeaver/bollard),.

##### CLI tools

- [captain](https://github.com/jenssegers/captain) - Easily start and stop docker compose projects from any directory. By [jenssegers](https://github.com/jenssegers).
- [dcinja](https://github.com/Falldog/dcinja) - The powerful and smallest binary size of template engine for docker command line environment. By [Falldog](https://github.com/Falldog).
- [dcp](https://github.com/exdx/dcp) - A simple tool for copying files from container filesystems. By [exdx](https://github.com/exdx).
- [dctl](https://github.com/FabienD/docker-stack) - Dctl is a Cli tool that helps developers by allowing them to execute all docker compose commands anywhere in the terminal, and more. By [FabienD](https://github.com/FabienD).
- [decompose](https://github.com/s0rg/decompose) - Reverse-engineering tool for docker environments. By [s0rg](https://github.com/s0rg).
- [docker pushrm](https://github.com/christian-korneck/docker-pushrm) - A Docker CLI plugin that lets you push the README.md file from the current directory to Docker Hub. Also supports Quay and Harbor. By [christian-korneck](https://github.com/christian-korneck).
- [docker-captain](https://github.com/lucabello/docker-captain) - A friendly CLI to manage multiple Docker Compose deployments with style — powered by Typer, Rich, questionary, and sh.
- [docker-ls](https://github.com/mayflower/docker-ls) - CLI tools for browsing and manipulating docker registries.
- [DVM](https://github.com/howtowhale/dvm) - Docker version manager.
- [goinside](https://github.com/iamsoorena/goinside) - Get inside a running docker container easily.
- [Pdocker](https://github.com/g31s/Pdocker) - A simple tool to manage and maintain Docker for personal projects.
- [proco](https://github.com/shiwaforce/poco) - Proco will help you to organise and manage Docker, Docker-Compose, Kubernetes projects of any complexity using simple YAML config files to shorten the route from finding your project to initialising it in your local environment.
- [scuba](https://github.com/JonathonReinhart/scuba) - Transparently use Docker containers to encapsulate software build environments,.
- [skopeo](https://github.com/containers/skopeo) - Work with remote images registries - retrieving information, images, signing content.
- [supdock](https://github.com/segersniels/supdock) - Allows for slightly more visual usage of Docker with an interactive prompt. By [segersniels](https://github.com/segersniels).
- [tsaotun](https://github.com/qazbnm456/tsaotun) - Python based Assistance for Docker.

##### Other

- [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin) - Search the Docker Registry with the Dext smart launcher. By [vutran](https://github.com/vutran).
- [docker-ssh](https://github.com/jeroenpeeters/docker-ssh) - SSH Server for Docker containers ~ Because every container should be accessible. By [jeroenpeeters](https://github.com/jeroenpeeters).
-   [dockerfile-mode](https://github.com/spotify/dockerfile-mode) An emacs mode for handling Dockerfiles by [spotify][spotify]
- [MultiDocker](https://github.com/marty90/multidocker) - Create a secure multi-user Docker machine, where each user is segregated into an indepentent container.
- [Powerline-Docker](https://github.com/adrianmo/powerline-docker) - A Powerline segment for showing the status of Docker containers.

#### Web

- [CASA](https://github.com/knrdl/casa) - Outsource the administration of a handful of containers to your co-workers,.
- [Container Web TTY](https://github.com/wrfly/container-web-tty) - Connect your containers via a web-tty [wrfly](https://github.com/wrfly).
- [dockemon](https://github.com/ProductiveOps/dokemon) - Docker Container Management GUI.
- [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser) - Web Interface for the Docker Registry HTTP API v2.
- [docker-registry-web](https://github.com/mkuchin/docker-registry-web) - Web UI, authentication service and event recorder for private docker registry v2.
- [docker-swarm-visualizer](https://github.com/dockersamples/docker-swarm-visualizer) - Visualizes Docker services on a Docker Swarm (for running demos).
- [dockge](https://github.com/louislam/dockge) - Easy-to-use and reactive self-hosted docker compose.yaml stack-oriented manager.
- [Komodo](https://github.com/mbecker20/komodo) - A tool to build and deploy software on many servers.
- [Kubevious](https://github.com/kubevious/kubevious) - A highly visual web UI for Kubernetes which renders configuration and state in an application centric way.
- [Mafl](https://github.com/hywax/mafl) - Minimalistic flexible homepage.
- [netdata](https://github.com/netdata/netdata) - Real-time performance monitoring.
- [OctoLinker](https://github.com/OctoLinker/OctoLinker) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
- [Portainer](https://github.com/portainer/portainer) - A lightweight management UI for managing your Docker hosts or Docker Swarm clusters.
- [Rapid Dashboard](https://github.com/ozlerhakan/rapid) - A simple query dashboard to use Docker Remote API.
- [Seagull](https://github.com/tobegit3hub/seagull) - Friendly Web UI to monitor docker daemon.
- [Swarmpit](https://github.com/swarmpit/swarmpit) - Swarmpit provides simple and easy to use interface for your Docker Swarm cluster. You can manage your stacks, services, secrets, volumes, networks etc.
- [Swirl](https://github.com/cuigh/swirl) - Swirl is a web management tool for Docker, focused on swarm cluster By [cuigh](https://github.com/cuigh/).
- [Theia](https://github.com/eclipse-theia/theia) - Extensible platform to develop full-fledged multi-language Cloud & Desktop IDE-like products with state-of-the-art web technologies.
- [usulnet](https://github.com/fr4nsys/usulnet) - A complete and modern Docker management platform designed for sysadmin, devops with enterprise grade tools, cve scanner, ssh, rdp on web and much more. By [fr4nsys](https://github.com/fr4nsys).

## Docker Images

### Base Tools

Tools and applications that are either installed inside containers or designed to be run as a [sidecar](https://learn.microsoft.com/en-us/azure/architecture/patterns/sidecar)

- [amicontained](https://github.com/genuinetools/amicontained) - Container introspection tool. Find out what container runtime is being used as well as features available.
- [Chaperone](https://github.com/garywiz/chaperone) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package.
- [ckron](https://github.com/nicomt/ckron) - A cron-style job scheduler for docker,.
-   [CoreOS][coreos] - Linux for Massive Server Deployments
- [distroless](https://github.com/GoogleContainerTools/distroless) - Language focused docker images, minus the operating system,.
- [docker-alpine](https://github.com/gliderlabs/docker-alpine) - A super small Docker base image _(5MB)_ using Alpine Linux.
- [docker-gen](https://github.com/jwilder/docker-gen) - Generate files from docker container meta-data.
- [dockerize](https://github.com/powerman/dockerize) - Utility to simplify running applications in docker containers by [jwilder][jwilder], [powerman][powerman].
- [GoSu](https://github.com/tianon/gosu) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool).
- [is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container.
- [lstags](https://github.com/ivanilves/lstags) - Sync Docker images across registries.
- [microcheck](https://github.com/tarampampam/microcheck) - Lightweight health check utilities for Docker containers (75 KB instead of 9.3 MB for httpcheck versus cURL) in pure C - http(s), port checks, and parallel execution are included.
- [Ofelia](https://github.com/mcuadros/ofelia/) - Ofelia is a modern and low footprint job scheduler for docker environments, built on Go. Ofelia aims to be a replacement for the old fashioned cron. Supports configuration from container labels and/or configuration files.
- [SparkView](https://github.com/beyondssl/sparkview-container) - Access VMs, desktops, servers or applications anytime and from anywhere, without complex and costly client roll-outs or user management.
- [su-exec](https://github.com/ncopa/su-exec) - This is a simple tool that will simply execute a program with different privileges. The program will be executed directly and not run as a child, like su and sudo does, which avoids TTY and signal issues. Why reinvent gosu? This does more or less exactly the same thing as gosu but it is only 10kb instead of 1.8MB. By [ncopa](https://github.com/ncopa).
- [sue](https://github.com/theAkito/sue) - Executes a program as a user different from the user running sue. This is a maintainable alternative to ncopa/su-exec, which is the better tianon/gosu. This one is far better (higher performance, smaller size), than the original gosu, however it is far easier to maintain, than su-exec, which is written in plain C. Made by [Akito][akito].
- [supercronic](https://github.com/aptible/supercronic) - Crontab-compatible job runner, designed specifically to run in containers.
- [TrivialRC](https://github.com/vorakl/TrivialRC) - A minimalistic Runtime Configuration system and process manager for containers [vorakl](https://github.com/vorakl).

### Builder

Applications designed to help or simplify building **new** images

- [ansible-bender](https://github.com/ansible-community/ansible-bender) - A tool utilising `ansible` and `buildah`.
- [buildah](https://github.com/containers/buildah) - A tool that facilitates building OCI images.
- [BuildKit](https://github.com/moby/buildkit) - Concurrent, cache-efficient, and Dockerfile-agnostic builder toolkit.
- [cekit](https://github.com/cekit/cekit) - A tool used by openshift to build base images using different build engines.
- [container-factory](https://github.com/mutable/container-factory) - Produces Docker images from tarballs of application source code.
- [copy-docker-image](https://github.com/mdlavin/copy-docker-image) - Copy a Docker image between registries without a full Docker installation.
- [Derrick](https://github.com/alibaba/derrick) - A tool help you to automate the generation of Dockerfile and dockerize application by scanning the code. By [alibaba](https://github.com/alibaba).
- [dlayer](https://github.com/orisano/dlayer) - Docker layer analyzer.
- [docker-companion](https://github.com/mudler/docker-companion) - A command line tool written in Golang to squash and unpack docker images.
- [docker-make](https://github.com/CtripCloud/docker-make) - Build, tag,and push a bunch of related docker images via a single command.
- [docker-repack](https://github.com/orf/docker-repack) - Repacks a Docker image into a smaller, more efficient version that makes it significantly faster to pull. By [orf](https://github.com/orf).
- [docker-replay](https://github.com/bcicen/docker-replay) - Generate `docker run`command and options from running containers. By [bcicen](https://github.com/bcicen).
-   [DockerSlim](https://github.com/docker-slim/docker-slim) shrinks fat Docker images creating the smallest possible images.
- [Dockly](https://github.com/swipely/dockly) - Dockly is a gem made to ease the pain of packaging an application in Docker.
- [essex](https://github.com/utensils/essex) - Boilerplate for Docker Based Projects: Essex is a CLI utility written in bash to quickly setup clean and consistent Docker projects with Makefile driven workflows. [jamesbrink](https://github.com/jamesbrink).
- [HPC Container Maker](https://github.com/NVIDIA/hpc-container-maker) - Generates Dockerfiles from a high level Python recipe, including building blocks for High-Performance Computing components.
- [img](https://github.com/genuinetools/img) - Standalone, daemon-less, unprivileged Dockerfile and OCI compatible container image builder.
- [packer](https://developer.hashicorp.com/packer/integrations/hashicorp/docker/latest/components/builder/docker) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible.
- [portainer](https://github.com/duedil-ltd/portainer) - Apache Mesos framework for building Docker images.
- [Production-Ready Python Containers :heavy_dollar_sign:](https://pythonspeed.com/products/pythoncontainer/) - A template for creating production-ready Docker images for Python applications.
- [RAUDI](https://github.com/cybersecsi/RAUDI) - A tool to automatically update (and optionally push to Docker Hub) Docker Images for 3rd party software whenever theres is a new release/update/commit. By [SecSI](https://github.com/cybersecsi).
- [runlike](https://github.com/lavie/runlike) - Generate `docker run`command and options from running containers.
- [userdef](https://github.com/theAkito/userdef) - An advanced `adduser` for your Alpine based Docker images. Made by [Akito][akito].
- [Whaler](https://github.com/P3GLEG/Whaler) - Program to reverse Docker images into Dockerfiles.
- [Whales](https://github.com/Gueils/whales) - A tool to automatically dockerize your applications.

### Dockerfile

- [chaperone-docker](https://github.com/garywiz/chaperone-docker) - A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.
-   [Dockerfile Generator](https://github.com/ozankasikci/dockerfile-generator) `dfg` is both a Go library and an executable that produces valid Dockerfiles using various input channels.
- [Dockerfile Project](https://dockerfile.github.io/) - Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
- [dockerfilegraph](https://github.com/patrickhoefler/dockerfilegraph) - Visualize your multi-stage Dockerfiles. By [PatrickHoefler](https://github.com/patrickhoefler).
- [Dockershelf](https://github.com/Dockershelf/dockershelf) - A repository that serves as a collector for docker recipes that are universal, efficient and slim. Images are updated, tested and published daily via a Travis cron job.
- [Vektorcloud](https://github.com/vektorcloud) - A collection of minimal, Alpine-based Docker images.

Examples by:

-   [0xy](https://gitlab.com/0xy/dockerfiles)
-   [arun-gupta](https://github.com/arun-gupta/docker-images)
-   [awesome-startup](https://github.com/awesome-startup/docker-compose)
-   [crosbymichael](https://github.com/crosbymichael/Dockerfiles)
-   [jessfraz](https://github.com/jessfraz/dockerfiles)
-   [komljen](https://github.com/komljen/dockerfile-examples)
-   [kstaken](https://github.com/kstaken/dockerfile-examples)
-   [ondrejmo](https://github.com/ondrejmo/Dockerfiles)
-   [vimagick](https://github.com/vimagick/dockerfiles)

### Linter

- [Dockadvisor](https://github.com/deckrun/dockadvisor) - Lightweight Dockerfile linter with 60+ rules, quality scoring, and security checks.
- [docker-image-size-limit](https://github.com/wemake-services/docker-image-size-limit) - A tool to keep an eye on your docker images size.
- [Dockerfile Linter action](https://github.com/buddy-works/dockerfile-linter) - The linter lets you verify Dockerfile syntax to make sure it follows the best practices for building efficient Docker images.
- [FROM:latest](https://github.com/replicatedhq/dockerfilelint) - An opinionated Dockerfile linter.
- [Hadolint](https://github.com/hadolint/hadolint) - A Dockerfile linter that checks for best practices, common mistakes, and is also able to lint any bash written in `RUN` instructions;.

### Metadata

- [opencontainer](https://github.com/opencontainers/image-spec/blob/main/annotations.md) - A convention and shared namespace for Docker labels defined by OCI Image Spec.

### Registry

Services to securely store your Docker images.

- [Amazon Elastic Container Registry :heavy_dollar_sign:](https://aws.amazon.com/ecr/) - Amazon Elastic Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
- [Azure Container Registry :heavy_dollar_sign:](https://azure.microsoft.com/en-us/products/container-registry/#overview) - Manage a Docker private registry as a first-class Azure resource.
- [CargoOS](https://github.com/RedCoolBeans/cargos-buildroot) - A bare essential OS for running the Docker Engine on bare metal or Cloud. By [RedCoolBeans](https://github.com/RedCoolBeans).
- [cleanreg](https://github.com/hcguersoy/cleanreg) - A small tool to delete image manifests from a Docker Registry implementing the API v2, dereferencing them for the GC.
- [Cloudsmith :heavy_dollar_sign:](https://cloudsmith.com/product/formats/docker-registry) - A fully managed package management SaaS, with first-class support for public and private Docker registries (and many others, incl. Helm charts for the Kubernetes ecosystem). Has a generous free-tier and is also completely free for open-source.
- [Container Registry Service :heavy_dollar_sign:](https://container-registry.com/) - Harbor based Container Management Solution as a Service for teams and organizations. Free tier offers 1 GB storage for private repositories.
- [Cycle.io :heavy_dollar_sign:](https://cycle.io/) - Bare-metal container hosting.
- [DigitalOcean :heavy_dollar_sign:](https://www.digitalocean.com/products/container-registry) - DigitalOcean Container Registry.
-   [Docker Hub](https://hub.docker.com/) provided by Docker Inc.
-   [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
- [Docket](https://github.com/netvarun/docket) - Custom docker registry that allows for lightning fast deploys through bittorrent.
- [Dragonfly](https://github.com/dragonflyoss/Dragonfly2) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology.
-   [GCP Artifact Registry :heavy_dollar_sign:](https://cloud.google.com/artifact-registry/docs) Fast, private Docker image storage on Google Cloud Platform.
- [Gitea Container Registry](https://docs.gitea.com/usage/packages/container) - Integrated Docker registry in Gitea, ideal for private, small-scale image hosting.
- [GitHub Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) - GitHub's solution for storing and managing Docker images, with tight integration into GitHub Actions.
- [GitLab Container Registry](https://docs.gitlab.com/user/packages/container_registry/) - Registry focused on using its images in GitLab CI.
-   [Harbor](https://github.com/goharbor/harbor) An open source trusted cloud native registry project that stores, signs, and scans content. Supports replication, user management, access control and activity auditing. By [CNCF](https://www.cncf.io) formerly [VMWare][vmware]
- [JFrog Artifactory :heavy_dollar_sign:](https://jfrog.com/artifactory/) - Artifact Repository Manager, can be used as private Docker Registry as well.
- [Kraken](https://github.com/uber/kraken) - Uber's Highly scalable P2P docker registry, capable of distributing TBs of data in seconds.
- [nscr](https://github.com/jhstatewide/nscr) - A light-weight, self-contained container registry that's easy to run and maintain.
- [Quay.io :heavy_dollar_sign:](https://quay.io/) - Secure hosting for private Docker repositories.
- [Registryo](https://github.com/inmagik/registryo) - UI and token based authentication server for onpremise docker registry.
- [RepoFlow](https://www.repoflow.io) - A simple and easy-to-use package management platform with Docker support alongside other formats like PyPI, Maven, npm, and Helm. Includes smart search, built-in Docker image scanning, and a great free option for both self-hosted and cloud use.
- [Rescoyl](https://github.com/noteed/rescoyl) - Private Docker registry (free and open source).
- [Sonatype Nexus Repository](https://www.sonatype.com/products/sonatype-nexus-repository) - Manage binaries and build artifacts across your software supply chain.

## Development with Docker

### API Client

- [ahab](https://github.com/instacart/ahab) - Docker event handling with Python.
- [contajners](https://github.com/lispyclouds/contajners) - An idiomatic, data-driven, REPL friendly Clojure client for OCI container engines. By [lispyclouds][lispyclouds].
- [Docker Client for JVM](https://github.com/gesellix/docker-client) - A Docker remote api client library for the JVM, written in Groovy.
- [Docker Client TypeScript](https://gitlab.com/masaeedu/docker-client) - Docker API client for JavaScript, automatically generated from Swagger API definition from moby repository. By [masaeedu](https://github.com/masaeedu).
- [docker-controller-bot](https://github.com/dgongut/docker-controller-bot) - Telegram bot to control docker containers. By [dgongut](https://github.com/dgongut/).
- [docker-it-scala](https://github.com/whisklabs/docker-it-scala) - Docker integration testing kit with Scala.
- [docker-java-api](https://github.com/amihaiemil/docker-java-api) - Lightweight, truly object-oriented, Java client for Docker's API. By [amihaiemil](https://github.com/amihaiemil).
- [docker-maven-plugin](https://github.com/fabric8io/docker-maven-plugin) - A Maven plugin for running and creating Docker images.
- [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet) - C#/.NET HTTP client for the Docker remote API.
- [Docker.Registry.DotNet](https://github.com/ChangemakerStudios/Docker.Registry.DotNet) - .NET (C#) Client Library for interacting with a Docker Registry API (v2) [rquackenbush](https://github.com/rquackenbush).
- [dockerode](https://github.com/apocas/dockerode) - Docker Remote API node.js module.
- [DoMonit](https://github.com/eon01/DoMonit) - A simple Docker Monitoring wrapper For Docker API.
- [go-dockerclient](https://github.com/fsouza/go-dockerclient/) - Go HTTP client for the Docker remote API.
- [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) - A Docker remote api plugin for Gradle.
- [Portainer stack utils](https://github.com/greenled/portainer-stack-utils) - Bash script to deploy/update/undeploy Docker stacks in a Portainer instance from a docker-compose yaml file. By [greenled](https://github.com/greenled).
- [sbt-docker](https://github.com/marcuslonnberg/sbt-docker) - Create Docker images directly from sbt.

### CI/CD

- [Buddy :heavy_dollar_sign:](https://buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development.
- [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery.
- [Cyclone](https://github.com/caicloud/cyclone) - Powerful workflow engine and end-to-end pipeline solutions implemented with native Kubernetes resources.
- [Defang](https://github.com/DefangLabs/defang) - Deploy Docker Compose to your favorite cloud in minutes.
- [Depot :heavy_dollar_sign:](https://depot.dev) - Build Docker images fast, in the cloud. Blazing fast compute, automatic intelligent caching, and zero configuration. [Done in seconds](https://depot.dev/#benchmarks).
- [Diun](https://github.com/crazy-max/diun) - Receive notifications when an image or repository is updated on a Docker registry by [crazy-max].
- [dockcheck](https://github.com/mag37/dockcheck) - A script checking updates for docker images without pulling then auto-update selected/all containers. With notifications, pruning and more.
- [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
- [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
- [Drydock](https://github.com/CodesWhat/drydock) - Open-source container update monitoring with web dashboard, 15 registries, 16 notification triggers, and security scanning. Drop-in WUD replacement. By [CodesWhat](https://github.com/CodesWhat).
- [Gantry](https://github.com/shizunge/gantry) - Automatically update selected Docker swarm services.
- [GitLab Runner](https://gitlab.com/gitlab-org/gitlab-runner) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
- [Jaypore CI](https://github.com/theSage21/jaypore_ci) - Simple, very flexible, powerful CI / CD / automation system configured in Python. Offline and local first.
- [Kraken CI](https://github.com/Kraken-CI/kraken) - Modern CI/CD, open-source, on-premise system that is highly scalable and focused on testing. One of its executors is Docker. Developed.
- [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment) - Continuous deployment of a microservices application.
- [mu](https://github.com/stelligent/mu) - Tool to configure CI/CD of your container applications via AWS CodePipeline, CodeBuild and ECS [Stelligent](https://github.com/stelligent).
- [Popper](https://github.com/systemslab/popper) - Github actions workflow (HCL syntax) execution engine.
- [Screwdriver :heavy_dollar_sign:](https://screwdriver.cd/) - Yahoo's OpenSource buildplatform designed for Continous Delivery.
- [Skipper](https://github.com/Stratoscale/skipper) - Easily dockerize your Git repository.
- [SwarmCI](https://github.com/ghostsquad/swarmci) - Create a distributed, isolated task pipeline in your Docker Swarm.
- [Tekton CD](https://tekton.dev/) - A cloud-native pipeline resource.

### Development Environment

- [Binci](https://github.com/binci/binci) - Containerize your development workflow. (formerly DevLab by [TechnologyAdvice](https://github.com/TechnologyAdvice)).
- [coder](https://github.com/coder/coder) - Remote development machines powered by Terraform or Docker.
- [construi](https://github.com/lstephen/construi) - Run your builds inside a Docker defined environment.
- [dde](https://github.com/whatwedo/dde) - Local development environment toolset based on Docker. By [whatwedo](https://github.com/whatwedo).
- [DIP](https://github.com/bibendi/dip) - CLI utility for straightforward provisioning and interacting with an application configured by docker-compose. By [bibendi](https://github.com/bibendi).
- [dobi](https://github.com/dnephin/dobi) - A build automation tool for Docker applications. By [dnephin](https://github.com/dnephin).
- [Docker Missing Tools](https://github.com/nandoquintana/docker-missing-tools) - A set of bash commands to shortcut typical docker dev-ops. An alternative to creating typical helper scripts like "build.sh" and "deploy.sh" inside code repositories. By [NandoQuintana](https://github.com/nandoquintana).
- [Docker-Arch](https://github.com/Ph3nol/Docker-Arch) - Generate Web/CLI projects Dockerized development environments, from 1 simple YAML file. By [Ph3nol](https://github.com/ph3nol).
- [Docker-sync](https://github.com/EugenMayer/docker-sync) - Drastically improves performance ([50-70x](https://github.com/EugenMayer/docker-sync/wiki/4.-Performance)) when using Docker for development on Mac OS X/Windows and Linux while sharing code to the container. By [EugenMayer](https://github.com/EugenMayer).
- [docker-vm](https://github.com/shyiko/docker-vm) - Simple and transparent alternative to boot2docker (backed by Vagrant).
- [DockerDL](https://github.com/matifali/dockerdl) - Deep Learning Docker Images. Don't waste time setting up a deep learning env when you can get a deep learning environment with everything pre-installed.
- [Eclipse Che](https://github.com/eclipse/che) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE.
- [EnvCLI](https://github.com/EnvCLI/EnvCLI) - Replace your local installation of Node, Go, ... with project-specific docker containers. By [EnvCLI](https://github.com/EnvCLI).
- [ESP32 Linux - Docker builder](https://github.com/hpsaturn/esp32s3-linux) - Container solution to compile Linux and develop it for ESP32 microcontrollers - By [Hpsaturn](https://github.com/hpsaturn).
- [Gebug](https://github.com/moshebe/gebug) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
- [Kitt](https://github.com/senges/kitt) - A portable and disposable Shell environment, based on Docker and Nix. By [senges](https://github.com/senges).
- [Lando](https://github.com/lando/lando) - Lando is for developers who want to quickly specify and painlessly spin up the services and tools needed to develop their projects. By [Tandem](https://www.thinktandem.io/).
- [Rust Universal Compiler](https://github.com/Peco602/rust-universal-compiler) - Container solution to compile Rust projects for Linux, macOS and Windows. By [Peco602][peco602].
- [uniget](https://github.com/uniget-org/cli) - Uni(versal)get, the installer and updater for container tools and beyond (formerly docker-setup). By [nicholasdille](https://github.com/nicholasdille).
- [Vagga](https://github.com/tailhook/vagga) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments.
- [Zsh-in-Docker](https://github.com/deluan/zsh-in-docker) - Install Zsh, Oh-My-Zsh and plugins inside a Docker container with one line! By [Deluan](https://www.deluan.com).

### Garbage Collection

- [caduc](https://github.com/tjamet/caduc) - A docker garbage collector cleaning stuff you did not use recently.
- [Docker Clean](https://github.com/ZZROTDesign/docker-clean) - A script that cleans Docker containers, images and volumes.
- [docker-custodian](https://github.com/Yelp/docker-custodian) - Keep docker hosts tidy. By [Yelp](https://github.com/Yelp).
- [docker_gc](https://github.com/pdacity/docker_gc) - Image for automatic removing unused Docker Swarm objects. Also works just as Docker Service.
- [Docuum](https://github.com/stepchowfun/docuum) - Least recently used (LRU) eviction of Docker images.

### Serverless

- [Apache OpenWhisk](https://github.com/apache/openwhisk) - A serverless, open source cloud platform that executes functions in response to events at any scale. By [apache](https://github.com/apache).
- [Funker](https://github.com/bfirsh/funker-example-voting-app) - Functions as Docker containers example voting app. By [bfirsh](https://github.com/bfirsh).
- [IronFunctions](https://github.com/iron-io/functions) - The serverless microservices platform FaaS (Functions as a Service) which uses Docker containers to run Any language or AWS Lambda functions.
- [Koyeb](https://www.koyeb.com/) - :heavy_dollar_sign: Koyeb is a developer-friendly serverless platform to deploy apps globally. Seamlessly run Docker containers, web apps, and APIs with git-based deployment, native autoscaling, a global edge network, and built-in service mesh and discovery.
- [OpenFaaS](https://github.com/openfaas/faas) - A complete serverless functions framework for Docker and Kubernetes. By [OpenFaaS](https://github.com/openfaas).
- [SCAR](https://github.com/grycap/scar) - Serverless Container-aware Architectures (SCAR) is a serverless framework that allows easy deployment and execution of containers (e.g. Docker) in Serverless environments (e.g. Lambda).

### Testing

- [Container Structure Test](https://github.com/GoogleContainerTools/container-structure-test) - A framework to validate the structure of an image by checking the outputs of commands or the contents of the filesystem. By [GoogleContainerTools][googlecontainertools].
- [dgoss](https://github.com/goss-org/goss/tree/master/extras/dgoss) - A fast YAML based tool for validating docker containers.
- [DockerSpec](https://github.com/zuazo/dockerspec) - A small Ruby Gem to run RSpec and Serverspec, Infrataster and Capybara tests against Dockerfiles or Docker images easily. By [zuazo](https://github.com/zuazo).
- [EZDC](https://github.com/lynchborg/ezdc) - Golang test harness for easily setting up tests that rely on services in a docker-compose.yml. By [byrnedo].
-   [InSpec][inspec] - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements. By [chef](https://github.com/chef)
- [Kurtosis](https://github.com/kurtosis-tech/kurtosis) - A composable build system for multi-container test environments that provides developers with: a powerful Python-like SDK for environment configuration, a compile-time validator to verify environment behavior & setup, and a runtime for environment execution, monitoring, & debugging capabilities. By [Kurtosis](https://www.kurtosis.com/).
- [Pull Dog](https://github.com/apps/pull-dog) - A GitHub app that automatically creates Docker-based test environments for your pull requests, from your docker-compose files. Not open source.
- [Pumba](https://github.com/alexei-led/pumba) - Chaos testing tool for Docker. Can be deployed on kubernetes and CoreOS cluster. By [alexei-led](https://github.com/alexei-led).

### Wrappers

- [Ansible](https://docs.ansible.com/projects/ansible/latest/collections/community/general/docker_container_module.html) - Manage the life cycle of Docker containers. By RedHat.
- [dexec](https://github.com/docker-exec/dexec) - Command line interface written in Go for running code with Docker Exec images.
- [dockerized](https://github.com/benzaita/dockerized-cli) - Seamlessly execute commands in a container.
- [Dray](https://github.com/CenturyLinkLabs/dray) - An engine for managing the execution of container-based workflows.
- [Hokusai](https://github.com/artsy/hokusai) - A Docker + Kubernetes CLI for application developers; used to containerize an application and to manage its lifecycle throughout development, testing, and release cycles. From [artsy](https://github.com/artsy).
- [Preevy](https://github.com/livecycle/preevy) - Preview environments for Docker and Docker Compose projects. Test your changes and get feedback from devs and non-devs (Product/Design) by deploying pull requests to the your cloud provider as part of your CI pipeline.
- [Shutit](https://github.com/ianmiell/shutit) - Tool for building and maintaining complex Docker deployments.
- [subuser](https://github.com/subuser-security/subuser) - Makes it easy to securely and portably run graphical desktop applications in Docker.
- [Terraform cloud-init config](https://github.com/christippett/terraform-cloudinit-container-server) - Terraform module for deploying a single Docker image or `docker-compose.yaml` file to any Cloud™ VM.
- [Turbo](https://github.com/ramitsurana/turbo) - Simple and Powerful utility for docker. By [ramitsurana][ramitsurana].
- [udocker](https://github.com/indigo-dc/udocker) - A tool to execute simple docker containers in batch or interactive systems without root privileges.
- [Vagrant - Docker provider](https://developer.hashicorp.com/vagrant/docs/providers/docker/basics) - Good starting point is [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example).

## Services based on Docker (mostly :heavy_dollar_sign:)

### CI Services

- [CircleCI](https://circleci.com/) - :heavy_dollar_sign: Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
- [CodeFresh](https://codefresh.io) - :heavy_dollar_sign: Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
- [CodeShip](https://www.cloudbees.com/blog/how-to-run-codeship-parallel-test-pipelines-efficiently-for-optimal-ci-parallelization) - :heavy_dollar_sign: Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
- [ConcourseCI](https://concourse-ci.org) - :heavy_dollar_sign: A CI SaaS platform for developers and DevOps teams pipeline oriented.
-   [Semaphore CI](https://semaphore.io/) :heavy_dollar_sign: — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
- [TravisCI](https://www.travis-ci.com/) - :heavy_dollar_sign: A Free github projects continuous integration Saas platform for developers and Devops.

### CaaS

- [Amazon ECS](https://aws.amazon.com/ecs/) - :heavy_dollar_sign: A management service on EC2 that supports Docker containers.
- [Appfleet](https://appfleet.com/) - :heavy_dollar_sign: Edge platform to deploy and manage containerized services globally. The system will route the traffic to the closest location for lower latency.
- [Azure AKS](https://azure.microsoft.com/en-us/products/kubernetes-service/) - :heavy_dollar_sign: Simplify Kubernetes management, deployment, and operations. Use a fully managed Kubernetes container orchestration service.
- [Cloud 66](https://www.cloud66.com) - :heavy_dollar_sign: Full-stack hosted container management as a service.
- [Giant Swarm](https://www.giantswarm.io/) - :heavy_dollar_sign: Simple microservice infrastructure. Deploy your containers in seconds.
- [Google Container Engine](https://docs.cloud.google.com/kubernetes-engine/docs) - :heavy_dollar_sign: Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
- [Mesosphere DC/OS Platform](https://d2iq.com/products/dcos) - :heavy_dollar_sign: Integrated platform for data and containers built on Apache Mesos.
- [Red Hat OpenShift Dedicated](https://www.redhat.com/en/technologies/cloud-computing/openshift/dedicated) - :heavy_dollar_sign: Fully-managed Red Hat® OpenShift® service on Amazon Web Services and Google Cloud.
- [Triton](https://www.joyent.com/) - :heavy_dollar_sign: Elastic container-native infrastructure by Joyent.

### Monitoring Services

- [AppDynamics](https://github.com/Appdynamics/docker-monitoring-extension) - Docker Monitoring extension gathers metrics from the Docker Remote API, either using Unix Socket or TCP.
- [Better Stack](https://betterstack.com/community/guides/scaling-docker/) - :heavy_dollar_sign: A Docker-compatible observability stack that delivers robust log aggregation and uptime monitoring capabilities for various software application.
- [Broadcom Docker Monitoring](https://www.broadcom.com/info/aiops/docker-monitoring) - :heavy_dollar_sign: Agile Operations solutions from Broadcom deliver the modern Docker monitoring businesses need to accelerate and optimize the performance of microservices and the dynamic Docker environments running them. Monitor both the Docker environment and apps that run inside them. (former CA Technologies).
-   [Collecting docker logs and stats with Splunk](https://www.splunk.com/en_us/blog/tips-and-tricks/collecting-docker-logs-and-stats-with-splunk.html)
- [Datadog](https://www.datadoghq.com/) - :heavy_dollar_sign: Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
- [Prometheus](https://prometheus.io/) - :heavy_dollar_sign: Open-source service monitoring system and time series database.
- [Site24x7](https://www.site24x7.com/docker-monitoring.html) - :heavy_dollar_sign: Docker Monitoring for DevOps and IT is a SaaS Pay per Host model.
- [SPM for Docker](https://github.com/sematext/sematext-agent-docker) - :heavy_dollar_sign: Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [sematext](https://github.com/sematext).
- [Sysdig Monitor](https://www.sysdig.com/products/monitor) - :heavy_dollar_sign: Sysdig Monitor can be used as either software or a SaaS service to monitor, alert, and troubleshoot containers using system calls. It has container-specific features for Docker and Kubernetes.

# Useful Resources

-   **[Valuable Docker Links](http://nane.kratzke.pages.mylab.th-luebeck.de/about/blog/2014/08/24/valuable-docker-links/)** High quality articles about docker! **MUST SEE**
-   [Cloud Native Landscape](https://github.com/cncf/landscape)
- [Docker Blog](https://www.docker.com/blog/) - Regular updates about Docker, the community and tools.
-   [Docker Certification](https://intellipaat.com/docker-training-course/?US) :heavy_dollar_sign: will help you to will Learn Docker containerization, running Docker containers, Image creation, Dockerfile, Docker orchestration, security best practices, and more through hands-on projects and case studies and helps to clear Docker Certified Associate.

- [Docker dev bookmarks](https://www.codever.dev/search?q=docker) - Use the tag [docker](https://www.codever.dev/bookmarks/t/docker).
-   [Docker in Action, Second Edition](https://www.manning.com/books/docker-in-action-second-edition)
-   [Docker in Practice, Second Edition](https://www.manning.com/books/docker-in-practice-second-edition)
- [Docker packaging guide for Python](https://pythonspeed.com/docker/) - A series of detailed articles on the specifics of Docker packaging for Python.
-   [Learn Docker in a Month of Lunches](https://www.manning.com/books/learn-docker-in-a-month-of-lunches)
- [Learn Docker](https://coursesity.com/blog/best-docker-tutorials/) - Learn Docker - curated list of the top online docker tutorials and courses.
-   [Programming Community Curated Resources for learning Docker](https://hackr.io/tutorials/learn-docker)

## Awesome Lists

- [Awesome CI/CD](https://github.com/cicdops/awesome-ciandcd) - Not specific to docker but relevant.
- [Awesome Compose](https://github.com/docker/awesome-compose) - Docker Compose samples.
-   [Awesome Kubernetes](https://github.com/ramitsurana/awesome-kubernetes) by [ramitsurana][ramitsurana]
-   [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers) more general about container than this repo, by [Friz-zy](https://github.com/Friz-zy).
-   [Awesome Selfhosted](https://github.com/awesome-selfhosted/awesome-selfhosted) list of Free Software network services and web applications which can be hosted locally by running in a classical way (setup local web server and run applications from there) or in a Docker container. By [Kickball](https://github.com/Kickball)
-   [Awesome Sysadmin](https://github.com/n1trux/awesome-sysadmin) by [n1trux](https://github.com/n1trux)
-   [ToolsOfTheTrade](https://github.com/cjbarber/ToolsOfTheTrade) a list of SaaS and On premise applications by [cjbarber](https://github.com/cjbarber)

## Demos and Examples

-   [An Annotated Docker Config for Frontend Web Development](https://nystudio107.com/blog/an-annotated-docker-config-for-frontend-web-development) A local development environment with Docker allows you to shrink-wrap the devops your project needs as config, making onboarding frictionless.
-   [Local Docker DB](https://github.com/alexmacarthur/local-docker-db) a list of docker-compose samples for a lot of databases by [alexmacarthur](https://github.com/alexmacarthur)
-   [Webstack-micro](https://github.com/ferbs/webstack-micro) Demo web app showing how Docker Compose might be used to set up an API Gateway, centralized authentication, background workers, and WebSockets as containerized services.

## Good Tips

-   [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016) **MUST SEE**
- [Docker Containers on the Desktop](https://blog.jessfraz.com/post/docker-containers-on-the-desktop/) - The **funniest way** to learn about docker by [jessfraz][jessfraz] who also gave a [presentation](https://www.youtube.com/watch?v=1qlLUf7KtAw) about it @ DockerCon 2015.
-   [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](https://www.flexera.com/blog/finops/)
- [Dockerfile best practices](https://github.com/hexops/dockerfile) - This repository has best-practices for writing Dockerfiles.
-   [Don't Repeat Yourself with Anchors, Aliases and Extensions in Docker Compose Files](https://medium.com/@kinghuang/docker-compose-anchors-aliases-extensions-a1e4105d70bd) by [King Chung Huang](https://github.com/kinghuang)
-   [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [fgrehm][fgrehm]

## Raspberry Pi & ARM

-   [Docker Pirates ARMed with explosive stuff](https://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
-   [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps)
-   [git push docker containers to linux devices](https://www.balena.io) Modern DevOps for IoT, leveraging git and Docker.
-   [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-%28ARMv7%29-devices)

## Security

-   [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
-   [CVE Scanning Alpine images with Multi-stage builds in Docker 17.05](https://github.com/tomwillfixit/alpine-cvecheck) by [tomwillfixit](https://twitter.com/tomwillfixit)
-   [Docker Secure Deployment Guidelines](https://github.com/AonCyberLabs/Docker-Secure-Deployment-Guidelines)
-   [Docker Security - Quick Reference](https://binarymist.io/publication/docker-security/)
-   [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](https://www.slideshare.net/slideshow/docker-security-are-your-containers-tightly-secured-to-the-ship/43834790)
-   [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448)
-   [Lynis is an open source security auditing tool including Docker auditing](https://cisofy.com/lynis/)
-   [Security Best Practices for Building Docker Images](https://linux-audit.com/tags/docker/)
-   [Software Engineering Radio interview of Docker Security Team Lead (Diogo Mónica)](https://www.se-radio.net/2017/05/se-radio-episode-290-diogo-monica-on-docker-security/)
-   [Ten Docker Image Security Best Practices Cheat Sheet](https://snyk.io/blog/10-docker-image-security-best-practices/)
-   [Top ten most popular docker images each contain at least 30 vulnerabilities](https://snyk.io/blog/top-ten-most-popular-docker-images-each-contain-at-least-30-vulnerabilities/)
-   [Tuning Docker with the newest security enhancements](https://opensource.com/business/15/3/docker-security-tuning)
-   [10 best practices to containerize Node.js web applications with Docker](https://snyk.io/blog/10-best-practices-to-containerize-nodejs-web-applications-with-docker/)

## Videos

-   [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](https://www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
-   [Deploying and scaling applications with Docker, Swarm, and a tiny bit of Python magic](https://www.youtube.com/watch?v=GpHMTR7P2Ms) (3:11:06) by [jpetazzo][jpetazzo]
-   [Docker and SELinux by Daniel Walsh from Red Hat](https://www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
-   [Docker Course](https://www.youtube.com/watch?v=UZpyvK6UGFo) (Spanish) by [pablokbs](https://github.com/pablokbs)
-   [Docker for Developers](https://www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26) by [jpetazzo][jpetazzo] <== Good introduction, context, demo
-   [Docker from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxD-btrjrqdEfQHbkZnQrmqE) (1:22:01) on YouTube by Paris Nakita Kejser
-   [Docker: How to Use Your Own Private Registry](https://www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
-   [Docker in Production](https://www.youtube.com/watch?v=Glk5d5WP6MI) by [jpetazzo][jpetazzo] (36:05)
-   [Docker Primer to Docker Compose](https://www.youtube.com/watch?v=G-s2GXGAjTk) (1:56:45) on YouTube by LoginRadius
-   [Docker Registry from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxAz3d4Fj3edq7UcxEhdTCBm) (44:40) on YouTube by Paris Nakita Kejser
-   [Docker Swarm from scratch](https://www.youtube.com/playlist?list=PLLhEJK7fQIxAY4gZd1Wl-GsLvg-e9Ap1e) (1:41:28) on YouTube by Paris Nakita Kejser
-   [Extending Docker with Plugins](https://vimeo.com/110835013) (15:21)
-   [From Local Docker Development to Production Deployments](https://www.youtube.com/watch?v=7CZFpHUPqXw) by [jpetazzo][jpetazzo] @ AWS re:Invent 2015
-   [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](https://www.youtube.com/watch?v=GaHzdqFithc) (42:04)
-   [Introduction to Docker and containers](https://www.youtube.com/watch?v=ZVaRK10HBjo) (3:09:00) by [jpetazzo][jpetazzo]
-   [Logging on Docker: What You Need to Know](https://vimeo.com/123341629) (51:27)
-   [Performance Analysis of Docker - Jeremy Eder](https://www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
-   [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
-   [State of containers: a debate with CoreOS, VMware and Google](https://www.youtube.com/watch?v=IiITP3yIRd8) (27:38)

# Communities and Meetups

## Brazilian

-   [Docker BR on Telegram](https://telegram.me/dockerbr)

## English

-   [Docker Community](https://www.docker.com/community/)
-   [Docker Events](https://www.docker.com/events/)
-   [Docker Online Meetup](https://www.meetup.com/en-AU/Docker-Online-Meetup/)
-   [Docker Reddit Community](https://www.reddit.com/r/docker/)

## Russian

-   [Docker Russian-speaking Community](https://t.me/docker_ru)

## Spanish

-   [Docker Tips](https://dockertips.com/)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/veggiemonk/awesome-docker.svg)](https://starchart.cc/veggiemonk/awesome-docker)

[contributing]: https://github.com/veggiemonk/awesome-docker/blob/master/.github/CONTRIBUTING.md
[akito]: https://github.com/theAkito
[calico]: https://github.com/projectcalico/calico
[centurylinklabs]: https://github.com/CenturyLinkLabs
[containx]: https://github.com/ContainX
[containers]: https://github.com/containers
[coreos]: https://github.com/coreos
[deepfence]: https://github.com/deepfence
[distribution]: https://github.com/docker/distribution
[docker-flow]: https://github.com/docker-flow
[docker-for-windows]: https://docs.docker.com/desktop/setup/install/windows-install/
[docker]: https://github.com/docker
[dozzle]: https://github.com/amir20/dozzle
[editreadme]: https://github.com/veggiemonk/awesome-docker/edit/master/README.md
[fgrehm]: https://github.com/fgrehm
[gesellix]: https://github.com/gesellix
[genuinetools]: https://github.com/genuinetools
[gliderlabs]: https://github.com/gliderlabs
[google]: https://github.com/google
[googlecontainertools]: https://github.com/GoogleContainerTools
[inspec]: https://github.com/inspec/inspec
[jessfraz]: https://github.com/jessfraz
[jpetazzo]: https://github.com/jpetazzo
[jwilder]: https://github.com/jwilder
[kubernetes]: https://kubernetes.io
[lispyclouds]: https://github.com/lispyclouds
[nvidia]: https://github.com/nvidia
[nginxproxy]: https://github.com/nginx-proxy/nginx-proxy
[openshift]: https://okd.io/
[oracle]: https://github.com/oracle
[peco602]: https://github.com/Peco602
[powerman]: https://github.com/powerman
[progrium]: https://github.com/progrium
[ramitsurana]: https://github.com/ramitsurana
[rancher]: https://github.com/rancher
[safe-waters]: https://github.com/safe-waters
[sindresorhus]: https://github.com/sindresorhus/awesome
[spotify]: https://github.com/spotify
[tomastomecek]: https://github.com/TomasTomecek
[vegasbrianc]: https://github.com/vegasbrianc
[weave]: https://github.com/weaveworks/weave
[vmware]: https://github.com/vmware
[byrnedo]: https://github.com/byrnedo
[crazy-max]: https://github.com/crazy-max
[grammarly]: https://github.com/grammarly
[skanehira]: https://github.com/skanehira
