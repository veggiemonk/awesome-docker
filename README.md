
# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/veggiemonk/awesome-docker](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/veggiemonk/awesome-docker?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Build Status](https://travis-ci.org/veggiemonk/awesome-docker.svg?branch=master)](https://travis-ci.org/veggiemonk/awesome-docker) [![Say Thanks](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg)](https://saythanks.io/to/veggiemonk)

> A curated list of Docker resources and projects
Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome][sindresorhus] and improved by these **[amazing contributors](https://github.com/veggiemonk/awesome-docker/graphs/contributors)**.

It's now a GitHub project because it's considerably easier for other people to edit, fix and expand on Docker using GitHub. Just click [README.md][editREADME] to submit a [pull request][editREADME].
If this list is not complete, you can [contribute][editREADME] to make it so.

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)** for tips!

#### *If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editREADME] to improve this file. Thank you!*

The creators and maintainers of this list do not receive and should not receive any form of payment to accept a change made by any contributor. 
The goal of this repo is to index articles, learning materials and projects, not to advertise for profit. 
**All pull requests are merged by default** and removed if inappropriate or unavailable, or fixed when necessary.

All the links are monitored and tested with [awesome_bot](https://github.com/dkhamsing/awesome_bot) made by  [@dkhamsing](https://github.com/dkhamsing)

# MENU

- [What is Docker ?](#what-is-docker-)
- [Where to start ?](#where-to-start-)
- [MENU](#menu)
- [Tools](#tools)
	- [Continuous Integration / Continuous Delivery](#continuous-integration--continuous-delivery)
	- [Deployment and Infrastructure](#deployment-and-infrastructure)
	- [Deployment](#deployment)
	- [Dev Tools](#dev-tools)
	- [Docker Compose file](#docker-compose-file)
	- [Dockerfile](#dockerfile)
	- [Hosting for containers](#hosting-for-containers)
	- [Hosting for repositories (registries)](#hosting-for-repositories-registries)
	- [Local Container Manager](#local-container-manager)
	- [Logging](#logging)
	- [Metadata](#metadata)
	- [Monitoring](#monitoring)
	- [Networking](#networking)
	- [PaaS](#paas)
	- [Remote Container Manager / Orchestration](#remote-container-manager--orchestration)
	- [Reverse Proxy](#reverse-proxy)
	- [Security](#security)
	- [Serverless](#serverless)
	- [Service Discovery](#service-discovery)
	- [Storing Images and Registries](#storing-images-and-registries)
	- [Terminal User Interface](#terminal-user-interface)
	- [Useful Images](#useful-images)
	- [Volume management and plugins](#volume-management-and-plugins)
	- [Web Interface](#web-interface)
- [Slides](#slides)
- [Useful Articles](#useful-articles)
	- [Cloud Infrastructure](#cloud-infrastructure)
	- [Continuous Integration](#continuous-integration)
	- [Deep Dive](#deep-dive)
	- [General Articles](#general-articles)
	- [Good Tips](#good-tips)
	- [Main Resources](#main-resources)
	- [Metal](#metal)
	- [Multi-Server](#multi-server)
	- [Networking](#networking)
	- [Newsletter](#newsletter)
	- [Optimizing Images](#optimizing-images)
	- [Other](#other)
	- [Performances](#performances)
	- [Raspberry Pi & ARM](#raspberry-pi--arm)
	- [Security](#security)
	- [Service Discovery](#service-discovery)
- [Books](#books)
	- [English](#in-english)
	- [Chinese](#chinese)
	- [German](#german)
	- [Portuguese](#portuguese)
- [Videos](#videos)
	- [Main Account](#main-account)
	- [Useful videos](#useful-videos)
- [Interactive Learning Environments](#interactive-learning-environments)
- [Interesting Twitter Accounts](#interesting-twitter-accounts)
	- [People](#people)
- [Communities and Meetups](#communities-and-meetups)



# What is Docker ?

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ [What is Docker](https://www.docker.com/what-docker)

# Where to start ?
* [10-minute Interactive Tutorial](https://docs.docker.com/docker-for-mac/)
* [Docker Training](http://training.docker.com/)
* Read this complete article: [Basics – Docker, Containers, Hypervisors, CoreOS](http://etherealmind.com/basics-docker-containers-hypervisors-coreos/)
* Watch the video: [Docker for Developers][docker4dev] (54:26) by [@jpetazzo][jpetazzo]
* [Docker Jumpstart](https://github.com/odewahn/docker-jumpstart/): a quick introduction
* [Docker Curriculum](http://prakhar.me/docker-curriculum/): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
* [Install Docker on your machine](docker-cheat-sheet#installation) and play with a few [Useful Images](#useful-images)
* Try [Panamax: Docker Management for Humans][panamax.io] It will install a CoreOS VM with VirtualBox and has nice front end
* [Install Docker Toolbox](https://www.docker.com/products/docker-toolbox) Docker Toolbox is an installer to quickly and easily install and setup a Docker environment on your computer. Available for both Windows and Mac, the Toolbox installs Docker Client, Machine, Compose (Mac only), Kitematic and VirtualBox.
* Check out [Docker Cheat Sheet][docker-cheat-sheet] by [@wsargent][wsargent] __MUST SEE__
* [Project Web Dev][projwebdev] : (Article series) How to create your own website based on Docker
* [Docker Containers on the desktop][jessblog] by [@jfrazelle][jfrazelle]) The **funniest way** to learn about docker! (Tips: checkout her [dotfiles][jfrazelledotfiles] and her [dockerfiles][jfrazelledockerfiles])
* [Container Hacks and Fun Images][jessvid] by [@jfrazelle][jfrazelle] @ DockerCon 2015 **MUST WATCH VIDEO** (38:50)
* [Learn Docker](https://github.com/dwyl/learn-docker) Full environment set up, screenshots, step-by-step tutorial and more resources (video, articles, cheat sheets) by [@dwyl](https://github.com/dwyl)
* [Docker Caveats](http://docker-saigon.github.io/post/Docker-Caveats/) What You Should Know About Running Docker In Production (written 11 APRIL 2016) __MUST SEE__
* [How to Whale](https://howtowhale.com/) Learn Docker in your web browser, no setup or installation required.
* [Docker for all - Developers, Testers, DevOps, Product Owners + Videos](https://github.com/machzqcq/docker-for-all) Docker Training Videos for all
* [Play With Docker](http://play-with-docker.com/) - A great way to get started with Docker. Docker runs directly in your browser.

# Where to start - on Windows ?
* [Windows Containers Quick Start](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick_start/quick_start) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016
* [Build And Run Your First Docker Windows Server Container](https://blog.docker.com/2016/09/build-your-first-docker-windows-server-container/) Walkthrough installing Docker on Windows 10, building a Docker image and running a Windows container
* Video: [Windows Containers and Docker: The 101](https://www.youtube.com/watch?v=N7SG2wEyQtM) A 20-minute overview, using Docker to run  PowerShell, ASP.NET Core and ASP.NET apps
* [A Comparative Study of Docker Engine on Windows Server vs Linux](http://collabnix.com/archives/1965) Comparing the feature sets and implementations of Docker on Windows and Linux
* [Docker with Microsoft SQL 2016 + ASP.NET](http://blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
* [Running a Legacy ASP.NET App in a Windows Container](https://blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and runnning as a Windows container
* [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](http://www.hanselman.com/blog/ExploringASPNETCoreWithDockerInBothLinuxAndWindowsContainers.aspx) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows](https://docs.docker.com/docker-for-windows/)


# Tools

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
* [Docker Machine](https://github.com/docker/machine) (Machine management for a container-centric world)
* [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)
* [Docker Swarm](https://github.com/docker/swarm) (Swarm: a Docker-native clustering system)

## Terminal User Interface
* [ctop](https://github.com/yadutaf/ctop) - A command line / text based Linux Containers monitoring tool that works just like you expect by [@yadutaf](https://github.com/yadutaf)
* [dockercraft](https://github.com/docker/dockercraft) - Docker + Minecraft = Dockercraft by [@docker][docker]
* [dockersql](https://github.com/crosbymichael/dockersql) - A command line interface to query Docker using SQL by [@crosbymichael][crosbymichael]
* [dockly](https://github.com/lirantal/dockly) - An interactive shell UI for managing Docker containers by [@lirantal](https://github.com/lirantal)
* [dry](https://github.com/moncho/dry) - An interactive CLI for Docker containers by [@moncho](https://github.com/moncho)
* [sen](https://github.com/TomasTomecek/sen) - Terminal user interface for docker engine, by [@TomasTomecek](https://github.com/TomasTomecek)
* [wharfee](https://github.com/j-bennet/wharfee) - Autocompletion and syntax highlighting for Docker commands. by [@j-bennet](https://github.com/j-bennet)
* [tsaotun](https://github.com/qazbnm456/tsaotun) - Python based Assistance for Docker by [@qazbnm456](https://github.com/qazbnm456)

## Dev Tools
* [ahab](https://github.com/instacart/ahab) - Docker event handling with Python by [@instacart](https://github.com/instacart)
* [athena](https://github.com/athena-oss/athena) - An automation platform with a plugin architecture that allows you to easily create and share services. Some useful plugins are : [php plugin](https://github.com/athena-oss/plugin-php), [selenium plugin](https://github.com/athena-oss/plugin-selenium) and [proxy plugin](https://github.com/athena-oss/plugin-proxy).
* [bocker](https://github.com/p8952/bocker) (1) - Docker implemented in 100 lines of bash by [p8952](https://github.com/p8952)
* [bocker](https://github.com/icy/bocker) (2) - Write Dockerfile completely in Bash. Extensible and simple. --> Reusable by [@icy](https://github.com/icy)
* [caduc](https://github.com/tjamet/caduc) - A docker garbage collector cleaning stuff you did not use recently
* [Chaperone](https://github.com/garywiz/chaperone) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package. by [@garywiz](https://github.com/garywiz)
* [codelift](https://codelift.io/) - CodeLift is an automated Docker image build utility for 'dockerizing' services by [@BoozAllen](https://twitter.com/BoozAllen)
* [Codenvy](https://codenvy.io) - One-click Docker environments and cloud workspace for development teams
* [Compose Registry](https://www.composeregistry.com) - A very handy search engine for Compose Files
* [Composerize](https://github.com/magicmark/composerize) - Convert docker run commands into docker-compose files
* [container-factory](https://github.com/lsqio/container-factory) - Produces Docker images from tarballs of application source code by [@lsqio](https://github.com/lsqio)
* [crowdr](https://github.com/polonskiy/crowdr) - Tool for managing multiple Docker containers (`docker-compose` alternative) by [@polonskiy](https://github.com/polonskiy/)
* [DevLab](https://github.com/TechnologyAdvice/DevLab) - Utility for running containerized development environments
* [Devstep](https://github.com/fgrehm/devstep) - Development environments powered by Docker and buildpacks by [@fgrehm][fgrehm]
* [dexec](https://github.com/docker-exec/dexec) - Command line interface for running code with Docker Exec images. https://docker-exec.github.io/ written in Go.
* [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin) - Search the Docker Registry with the Dext smart launcher.
* [dlayer](https://github.com/wercker/dlayer) - Stats collector for Docker layers by [@wercker](https://github.com/wercker)
* [Docker Clean](https://github.com/zzrotdesign/docker-clean) - A script that cleans Docker containers, images and volumes by [@zzrotdesign](https://github.com/zzrotdesign)
* [Docker Client for JVM](https://github.com/gesellix/docker-client) - A Docker remote api client library for the JVM, written in Groovy by [@gesellix][gesellix]
* [Docker meets the IDE](http://domeide.github.io/) - Integrating your favorite containers in the editor of your choice by [domeide](https://github.com/domeide)
* [Docker osx dev](https://github.com/brikis98/docker-osx-dev) - A productive development environment with Docker on OS X by [@brikis98](https://github.com/brikis98)
* [Docker Volume Clone Utility](https://github.com/gdiepen/docker-convenience-scripts) - A Docker Utility to Clone Volumes [@gdiepen](https://twitter.com/gdiepen)
* [docker-companion](https://github.com/mudler/docker-companion) - A command line tool written in Golang to squash and unpack docker images by [@mudler](https://github.com/mudler/)
* [docker-compose-search](https://github.com/francescou/docker-compose-search) - A search engine for Docker Compose application stacks by [@francescou](https://github.com/francescou/)
* [docker-do](https://github.com/benzaita/docker-do) - hassle-free docker run, like `env` but for docker by [@benzaita](https://github.com/benzaita)
* [docker-garby](https://github.com/konstruktoid/docker-garby) - Docker garbage collection script by [@konstruktoid](https://github.com/konstruktoid).
* [docker-gc](https://github.com/spotify/docker-gc) - A cron job that will delete old stopped containers and unused images by [@spotify](https://github.com/spotify)
* [docker-gen](https://github.com/jwilder/docker-gen) - Generate files from docker container meta-data by [@jwilder][jwilder]
* [docker-ls](https://github.com/mayflower/docker-ls) - CLI tools for browsing and manipulating docker registries by [@mayflower](https://github.com/mayflower)
* [docker-make](https://github.com/CtripCloud/docker-make) - Build, tag,and push a bunch of related docker images via a single command.
* [Docker-PowerShell](https://github.com/Microsoft/Docker-PowerShell) - PowerShell Module for Docker
* [docker-replay](https://github.com/bcicen/docker-replay) - Generate `docker run`command and options from running containers
* [Docker-sync](http://docker-sync.io/) - Tool from [@EugenMayer](https://github.com/EugenMayer) to improve performance when using Docker on Mac OS X and your volume has several folders and files like gulp projects for instance.
* [docker-volumes](https://github.com/cpuguy83/docker-volumes) - Docker Volume Manager by [@cpuguy83][cpuguy83]
* [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet) - C#/.NET HTTP client for the Docker remote API by [@ahmetalpbalkan](https://github.com/ahmetalpbalkan/)
* [dockerfile_lint](https://github.com/projectatomic/dockerfile_lint) - A rule-based 'linter' for Dockerfiles by [@redhataccess](https://github.com/redhataccess)
* [dockerize](https://github.com/jwilder/dockerize) - Utility to simplify running applications in docker containers by [@jwilder][jwilder]
* [dockerode](https://github.com/apocas/dockerode) - Not just another Docker Remote API node.js module by [@apocas](https://github.com/apocas)
* [Dockly](https://github.com/swipely/dockly) - Dockly is a gem made to ease the pain of packaging an application in Docker by [@swipely](https://github.com/swipely/)
* [dockramp](https://github.com/jlhawn/dockramp) - Proof of Concept: A Client Driven Docker Image Builder by [@jlhawn](https://github.com/jlhawn)
* [draw-compose](https://github.com/Alexis-benoist/draw-compose) - Utility to draw a schema of a docker compose by [@Alexis-benoist](https://github.com/Alexis-benoist)
* [Dray](http://dray.it/) - Dray is an engine for managing the execution of container-based workflows. Docker Workflow Engine - UNIX pipes for Docker by [@CenturyLinkLabs][CenturyLinkLabs]
* [Dropdock](http://dropdock.io/) - A framework designed for Drupal to build fast, isolated development environments using Docker.
* [Dupper](https://github.com/athakwani/dupper) - Use git repository as Object Oriented Container. Setup Cloud Development environments with few commands for any git repository by [@athakwani](https://github.com/athakwani/)
* [DVM](https://github.com/getcarina/dvm) - Docker version manager by [@getcarina](https://github.com/getcarina)
* [Eclipse Che](http://www.eclipse.org/che) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE
* [elsy](https://github.com/cisco/elsy) - An opinionated, multi-language, build tool based on Docker and Docker Compose
* [forward2docker](https://github.com/bsideup/forward2docker) - Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM by [@bsideup](https://github.com/bsideup)
* [go-dockerclient](https://github.com/fsouza/go-dockerclient/) - Go HTTP client for the Docker remote API by [@fsouza](https://github.com/fsouza/)
* [GoSu](https://github.com/tianon/gosu) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool) by [@tianon](https://github.com/tianon)
* [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) - A Docker remote api plugin for Gradle by [@gesellix][gesellix]
* [habitus](https://github.com/cloud66/habitus) - A Build Flow Tool for Docker http://www.habitus.io by [@cloud66](https://github.com/cloud66)
* [is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container by [@sindresorhus][sindresorhus]
* [Lorry](https://lorry.io/) - Lorry is a docker-compose.yml validator and composer by [@CenturyLinkLabs][CenturyLinkLabs]
* [ns-enter](https://github.com/jpetazzo/nsenter) - no more ssh, enter name spaces of container by [@jpetazzo][jpetazzo]
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
* [percheron][percheron] - Organise your Docker containers with muscle and intelligence by [@ashmckenzie](https://github.com/ashmckenzie)
* [portainer](https://github.com/duedil-ltd/portainer) - Apache Mesos framework for building Docker images by [@duedil-ltd](https://github.com/duedil-ltd)
* [Powerline-Docker](https://github.com/adrianmo/powerline-docker) - A Powerline segment for showing the status of Docker containers by [@adrianmo](https://github.com/adrianmo)
* [powerstrip](https://github.com/clusterhq/powerstrip) - A tool for prototyping Docker extensions by [@clusterhq](https://github.com/clusterhq)
* [registrator](https://github.com/progrium/registrator) - Service registry bridge for Docker by [@progrium][progrium]
* [rocker](https://github.com/grammarly/rocker) - Extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc. by [grammarly](https://github.com/grammarly).
* [sbt-docker-compose](https://github.com/Tapad/sbt-docker-compose) - Integrates Docker Compose functionality into sbt by [@kurtkopchik](https://github.com/kurtkopchik/)
* [sherdock](https://github.com/rancher/sherdock) - Automatic GC of images based on regexp by [@rancher][rancher]
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) - in case of proxy problem by [@jpetazzo][jpetazzo]
* [Vagga](https://github.com/tailhook/vagga) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments by [@tailhook](https://github.com/tailhook/)
* [Whale-linter](https://github.com/jeromepin/whale-linter) - A simple and small Dockerfile linter written in Python3+ without dependencies.
* [flyimg](http://flyimg.io/) - Docker image resizing, cropping, and compression on the fly. 

## Continuous Integration / Continuous Delivery
* [Awesome-ciandcd](https://github.com/ciandcd/awesome-ciandcd) - Not specific to docker but relevant.
* [Buddy](https://buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development
* [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery by [@harbur](https://github.com/harbur)
* [CircleCI](https://circleci.com/) - Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
* [CodeFresh](https://codefresh.io) - Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
* [CodeShip](https://codeship.com/features/pro) - Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
* [Cyclone](https://github.com/caicloud/cyclone) - A cloud native CI/CD platform built for container workflow by [@caicloud](https://github.com/caicloud).
* [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
* [Dockunit](https://github.com/dockunit/platform) - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By [@dockunit](https://github.com/dockunit)
* [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
* [GitLab CI](https://about.gitlab.com/gitlab-ci/) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
* [GOCD-Docker](https://github.com/gocd/gocd-docker)Go Server and Agent in docker containers to provision.
* [IBM DevOps Services](https://hub.jazz.net) - Continuous delivery using a pipeline deployment onto IBM Containers on Bluemix.
* [InSpec](https://github.com/chef/inspec) - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements.
* [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment) - Continuous deployment of a microservices application
* [Pumba](https://github.com/gaia-adm/pumba) - Chaos testing tool for Docker. Can be deployed on Kubernets and CoreOS clusters.
* [Semaphore CI](https://semaphoreci.com/product/docker) — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
* [Shippable](https://app.shippable.com/) - A SaaS platform for developers and DevOps teams that significantly reduces the time taken for code to be built, tested and deployed to production.
* [Skipper](https://github.com/Stratoscale/skipper) - Easily dockerize your Git repository by [@Stratoscale](https://github.com/Stratoscale)
* [SwarmCI](https://github.com/ghostsquad/swarmci) - Create a distributed, isolated task pipeline in your Docker Swarm.
* [Watchtower](https://github.com/v2tec/watchtower) - Automatically update running Docker containers by [@CenturyLinkLabs][CenturyLinkLabs]

## Deployment
* [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker by [@ehazlett](https://github.com/ehazlett)
* [depcon](https://github.com/gondor/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes.  By [@gonodr][gondor]
* [dockit](https://github.com/humblec/dockit) - Do docker actions and Deploy gluster containers!
* [rocker-compose](https://github.com/grammarly/rocker-compose) - Docker composition tool with idempotency features for deploying apps composed of multiple containers.
* [Zodiac](https://github.com/CenturyLinkLabs/zodiac) - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs][CenturyLinkLabs]

## Hosting for repositories (registries)
Securely store your Docker images.
* [Amazon EC2 Container Registry](https://aws.amazon.com/ecr/) Amazon EC2 Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images. 
* [Azure Container Registry](https://azure.microsoft.com/de-de/services/container-registry/) Manage a Docker private registry as a first-class Azure resource
* [Carina](https://getcarina.com) Bare-metal container hosting. Developed by Rackspace. Currently in Beta.
* [Compose Registry](https://www.composeregistry.com/) - Project from [@francescou]Francesco Uliana that stores docker-compose.yml files with stack examples like LAMP/LEMP stacks, Django stacks, Flask stacks and ELK stacks for instance.
* [Container Compliance](https://github.com/OpenSCAP/container-compliance) Open Source tool for assesing running containers and cold images for vulnerabilites and audits.
* [Cycle.io](https://cycle.io/) Bare-metal container hosting.
* [Docker Hub](https://hub.docker.com/) provided by Docker Inc.
* [GCE Container Registry](https://cloud.google.com/container-registry/) Fast, private Docker image storage on Google Cloud Platform
* [GitLab Container Registry](https://docs.gitlab.com/ce/container_registry/README.html) - Repositories focused on using it images in GitLab CI
* [Private Docker Registry](https://private-docker-registry.com) Hosted registry service for companies with unlimited private repositories, users, teams and namespaces. Federeated identity with LDAP/AD/OAuth/SAML and VPN Tunneling.
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [TreeScale](https://treescale.com/) - Build and Distriubute container based applications.

## Hosting for containers
* [Amazon ECS](http://aws.amazon.com/ecs/) - A management service on EC2 that supports Docker containers.
* [ContainerShip Cloud][containership] - Multi-Cloud Container Hosting Automation Platform.
* [Docker Cloud](https://cloud.docker.com/) - Former Tutum
* [Giant Swarm](https://giantswarm.io/) - Simple microservice infrastructure. Deploy your containers in seconds.
* [Google Container Engine](https://cloud.google.com/container-engine/docs/) - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
* [Hyper_](https://hyper.sh/) - Secure container hosting service with "nano-containers" and per-second billing.
* [IBM Bluemix](https://console.ng.bluemix.net/) - Run Docker containers in a hosted cloud environment on IBM Bluemix.
* [OpenShift Dedicated](https://www.openshift.com/dedicated/index.html) - A hosted [OpenShift][openshift] cluster for running your Docker containers managed by Red Hat.
* [Orchard](https://www.orchardup.com/) (part of Docker Inc) - Get a Docker host in the cloud, instantly.
* [Triton](https://www.joyent.com/) - Elastic container-native infrastructure by Joyent.

## Reverse Proxy
* [docker-proxy](https://github.com/silarsis/docker-proxy) - Transparent proxy for docker containers, run in a docker container. By [@silarsis](https://github.com/silarsis)
* [fabio](https://github.com/eBay/fabio) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By [@eBay](https://github.com/eBay)
* [h2o-proxy](https://github.com/zchee/h2o-proxy) - Automated H2O reverse proxy for Docker containers. An alternative to [jwilder/nginx-proxy][nginxproxy] by [@zchee](https://github.com/zchee)
* [Let's Encrypt Nginx-proxy Companion](https://github.com/JrCs/docker-letsencrypt-nginx-proxy-companion) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By [@JrCs](https://github.com/JrCs)
* [muguet](https://github.com/mattallty/muguet) - DNS Server & Reverse proxy for Docker environments. By [@mattallty](https://github.com/mattallty)
* [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [@jwilder][jwilder]
* [Swarm Ingress Router](https://github.com/tpbowden/swarm-ingress-router) - Route DNS names to Swarm services based on labels.
* [Træfɪk](https://traefik.io/) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [@EmileVauge](https://github.com/emilevauge)

## Web Interface
* [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser) - Web Interface for the Docker Registry HTTP API v2 by [@klausmeyer](https://github.com/klausmeyer)
* [Docker Registry UI](https://github.com/atc-/docker-registry-ui) - A web UI for easy private/local Docker Registry integration by [@atc-](https://github.com/atc-)
* [docker-registry-web](https://github.com/mkuchin/docker-registry-web) - Web UI, authentication service and event recorder for private docker registry v2 by [@mkuchin](https://github.com/mkuchin)
* [docker-swarm-visualizer](https://github.com/manomarks/docker-swarm-visualizer) - Visualizes Docker services on a Docker Swarm (for running demos).
* [dockering-on-rails](https://github.com/Electrofenster/dockerding-on-rails) - Simple Web-Interface for Docker with a lot of features by [@Electrofenster](https://github.com/Electrofenster/)
* [DockerUI](https://github.com/kevana/ui-for-docker) - DockerUI is a web interface to interact with the Remote API by [@crosbymichael][crosbymichael]
* [Portus](https://github.com/SUSE/Portus) - Authorization service and frontend for Docker registry (v2) by [@SUSE](https://github.com/SUSE)
* [Rapid Dashboard](https://github.com/ozlerhakan/rapid) - A simple query dashboard to use Docker Remote API by [@ozlerhakan](https://github.com/ozlerhakan/)

## Local Container Manager
* [Ansible - manage docker containers](http://docs.ansible.com/ansible/docker_module.html)
* [Azk](http://www.azk.io/) - Orchestrate development enviornments on your local machine by [@azukiapp](https://github.com/azukiapp)
* [Beluga](https://github.com/cortexmedia/Beluga) - CLI to deploy docker containers on a single server or low amount of servers. By [@cortextmedia](https://github.com/cortexmedia)
* [Boot2Docker](https://github.com/boot2docker/boot2docker) - Docker for OSX and Windows -- http://boot2docker.io/
* [Dinghy](https://github.com/codekitchen/dinghy) - An alternative way to use Docker on Mac OS X using Docker Machine with virtualbox, vmware, xhyve or parallels
* [DLite](https://github.com/nlf/dlite) - Simplest way to use Docker on OSX, no VM needed. By [@nlf](https://github.com/nlf)
* [docker-vm](https://github.com/shyiko/docker-vm) - Simple and transparent alternative to boot2docker (backed by Vagrant) by [@shyiko](https://github.com/shyiko)
* [Dokku][dokku] - Docker powered mini-Heroku in around 100 lines of Bash by [@progrium][progrium]
* [Dray](https://github.com/CenturyLinkLabs/dray) - An engine for managing the execution of container-based workflows. http://Dray.it by [@CenturyLinkLabs][CenturyLinkLabs]
* [Dusty](http://dusty.gc.com/) - Managed Docker development environments on OS X
* [FuGu](https://github.com/mattes/fugu) - Docker run wrapper without orchestration by [@mattes](https://github.com/mattes)
* [libcompose](https://github.com/docker/libcompose) - Go library for Docker Compose.
* [OctoHost](http://octohost.io/) - Simple web focused Docker based mini-PaaS server. git push to deploy your websites as needed) by [@octohost](https://github.com/octohost)
* [percheron][percheron] - Organise your Docker containers with muscle and intelligence by [@ashmckenzie](https://github.com/ashmckenzie)
* [Shutit](http://ianmiell.github.io/shutit/) - Tool for building and maintaining complex Docker deployments by [@ianmiell][ianmiell]
* [subuser](http://subuser.org) - Makes it easy to securely and portably run graphical desktop applications in Docker
* [Turbo](https://ramitsurana.github.io/turbo/) - Simple and Powerful utility for docker. By [@ramitsurana][ramitsurana]
* [Vagrant - Docker provider](https://www.vagrantup.com/docs/docker/basics.html) - Good starting point is [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example) by [@bubenkoff](https://github.com/bubenkoff)
* [Vessel](https://github.com/awvessel/vessel) - Automates the setup & use of dockerized development environments by [@awvessel](https://github.com/awvessel)

## Volume management and plugins
* [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume) - The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By [@blockbridge][blockbridge]
* [Convoy](https://github.com/rancher/convoy) - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher][rancher]
* [Azure Files Volume Driver](https://github.com/ahmetalpbalkan/azurefile-dockervolumedriver) - A Docker volume driver that allows you to mount persistent volumes backed by Microsoft Azure File Service. By [@ahmetalpbalkan][ahmetalpbalkan]
* [Docker Unison](https://github.com/leighmcculloch/docker-unison) A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By [@leighmcculloch](https://github.com/leighmcculloch)
* [Netshare](https://github.com/gondor/docker-volume-netshare) A Docker volume plugin written in Go that supports mounting NFS, AWS EFS & CIFS volumes within a container. By [@gondor][gondor]
* [Docker Machine NFS](https://github.com/adlogix/docker-machine-nfs) Activates NFS for an existing boot2docker box created through Docker Machine on OS X.
* [REX-Ray](https://github.com/codedellemc/rexray) Vendor agnostic storage orchestration engine to provide persistent storage for Docker containers as well as Mesos frameworks and tasks.
* [Local Persist](https://github.com/CWSpear/local-persist) Specify a mountpoint for your local volumes (created via `docker volume create`) so that files will always persist and so you can mount to different directories in different containers.
* [Minio](https://github.com/jelastic-jps/minio) - S3 compatible object storage server in Docker containers


## Useful Images
* [Base Image](https://github.com/phusion/baseimage-docker) by [@phusion](https://github.com/phusion/)
* [Busybox](https://github.com/jpetazzo/docker-busybox) - with either `buildroot` or Ubuntu's `busybox-static` by [@jpetazzo][jpetazzo]
* [chaperone-docker](https://github.com/garywiz/chaperone-docker) - A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.
* [docker-alpine][alpine] - A super small Docker base image *(5MB)* using Alpine Linux by [@gliderlabs][gliderlabs]
* [docker-fluentd][fluentd] - the Container to Log Other Containers' Logs by [@kiyoto][kiyoto]
* [nvidia-docker](https://github.com/NVIDIA/nvidia-docker) - Build and run Docker containers leveraging NVIDIA GPUs.
* [Official Images from Docker Hub](https://github.com/docker-library/official-images)
* [OpenWRT](http://www.zoobab.com/docker-openwrt-image) by [@zoobab](https://github.com/zoobab)
* [passenger-docker](https://github.com/phusion/passenger-docker) - Docker base images for Ruby, Python, Node.js and Meteor web apps by [@phusion](https://github.com/phusion)
* [Phusion Docker Hub Account](https://hub.docker.com/u/phusion/)


## Dockerfile
* [Collection of Dockerfiles](https://github.com/crosbymichael/Dockerfiles) by [@crosbymichael][crosbymichael]
* [Collection of Dockerfiles 2](https://github.com/pandrew/dockerfiles) by [@pandrew][pandrew]
* [Collection of Dockerfiles 3](https://github.com/vimagick/dockerfiles) by [@vimagick][vimagick]
* [Collection of Dockerfiles 4](https://github.com/ondrejmo/Dockerfiles) by [@ondrejmo][ondrejmo]
* [Collection of Dockerfiles 5](https://github.com/arun-gupta/docker-images) by [@arun-gupta][arun-gupta]
* [Dockerfile Project](http://dockerfile.github.io/) : Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
* [Dockerfile Example](https://github.com/komljen/dockerfile-examples) by [@komljen](https://github.com/komljen)
* [Dockerfile Example 2](https://github.com/kstaken/dockerfile-examples) by [@kstaken](https://github.com/kstaken)
* [Dockerfile @jfrazelle][jfrazelledockerfiles] by [@jfrazelle][jfrazelle] **MUST SEE** for a fully containerized desktop!
* [Vektorcloud](https://github.com/vektorcloud) - A collection of minimal, Alpine-based Docker images

## Docker Compose file
* [Docker Compose Example](https://github.com/llitfkitfk/docker-compose-demo) by [@llitfkitfk](https://github.com/llitfkitfk)

## Storing Images and Registries
* [Atomic Registry](http://www.projectatomic.io/registry/) - Red Hat Atomic Registry is an open source enterprise registry based on the Origin and Cockpit projects, enhancing the Docker registry library.
* [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
* [Rescoyl](https://github.com/noteed/rescoyl) - Private Docker registry by [@noteed][noteed]
* [VMWare Harbor](http://vmware.github.io/harbor/) Project Harbor by VMWare is an enterprise-class registry server that stores and distributes Docker images. Harbor extends the open source Docker Distribution by adding the functionalities usually required by an enterprise, such as security, identity and management.

## Monitoring
* [AppDynamics](https://www.appdynamics.com/community/exchange/extension/docker-monitoring-extension/) - AppDynamics gives enterprises real-time insights into application performance, user performance, and business performance so they can move faster in an increasingly sophisticated, software-driven world.
* [Axibase Time-Series Database](http://axibase.com/products/axibase-time-series-database/writing-data/docker-cadvisor/) - Long-term retention of container statistics and built-in dashboards for Docker. Collected with native Google cAdvisor storage driver.
* [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers. Created by [@Google](https://github.com/google)
* [Collecting docker logs and stats with Splunk](http://blogs.splunk.com/2015/08/24/collecting-docker-logs-and-stats-with-splunk/)
* [CoScale](http://www.coscale.com/docker-monitoring) - Full stack monitoring for containerized applications and microservices. Powered by anomaly detection to find performance problems faster.
* [Datadog](https://www.datadoghq.com/) - Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
* [Docker-mon](https://github.com/icecrime/docker-mon) - Console-based Docker monitoring by [@icecrime](https://github.com/icecrime)
* [Dockerana](https://github.com/dockerana/dockerana) - packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.
* [DoMonit](https://github.com/eon01/DoMonit) - A simple Docker Monitoring wrapper For Docker API
* [Dynatrace](https://www.dynatrace.com/technologies/cloud-and-microservices/docker-monitoring/) - Monitor containerized applications without installing agents or modifying your Run commands
* [Glances](https://nicolargo.github.io/glances/) - A cross-platform curses-based system monitoring tool written in Python by [@nicolargo](https://github.com/nicolargo)
* [Grafana Docker Dashboard Template](https://grafana.net/dashboards/179) - A template for your Docker, Grafana and Prometheus stack [@vegasbrianc][vegasbrianc]
* [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) - InfluxDB Time series DB in combination with Grafana and cAdvisor by [@vegasbrianc][vegasbrianc]
* [Meros](https://meros.io) - Analyzes containers resources, captures logs, remote web SSH terminal and powerful DevOps alerts.
* [New Relic](https://newrelic.com/partner/docker) - New Relics Docker Monitoring tool
* [Prometheus](https://prometheus.io/) - Open-source service monitoring system and time series database
* [Seagull](https://github.com/tobegit3hub/seagull) - Friendly Web UI to monitor docker daemon. by [@tobegit3hub](https://github.com/tobegit3hub)
* [Site24x7](https://www.site24x7.com/docker-monitoring.html) - Docker MOnitoring for DevOps and IT is a SaaS Pay per Host model
* [SPM for Docker][spm] - Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [@sematext][sematext]
* [Sysdig](http://www.sysdig.org/) - An open source troubleshooting tool that provides a rich set of real-time, system-level information. It has container-specific features and is very useful in Docker environments.
* [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
* [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.

## Networking
* [Calico-Docker](https://www.projectcalico.org/getting-started/docker/) - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
* [Flannel](https://github.com/coreos/flannel/) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes.
* [Wagl](https://github.com/ahmetalpbalkan/wagl) - DNS Service Discovery for Docker Swarm (by [@ahmetalpbalkan][ahmetalpbalkan] ) http://ahmetalpbalkan.github.io/wagl/
* [Weave][weave] (The Docker network) - Weave creates a virtual network that connects Docker containers deployed across multiple hosts.


## Logging
* [Docker-Fluentd][fluentd] - Docker container to Log Other Containers' Logs. One can aggregate the logs of Docker containers running on the same host using Fluentd by [@kiyoto][kiyoto]
* [LogJam](https://github.com/gocardless/logjam) - Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash) by [@gocardless](https://github.com/gocardless)
* [Logspout](https://github.com/gliderlabs/logspout) - Log routing for Docker container logs by [@gliderlabs][gliderlabs]
* [Logsene for Docker][spm] Monitoring of Metrics, Events and Logs implemented in Node.js. Integrated [logagent-js](https://github.com/sematext/logagent-js) to detect and parse various log formats. [@sematext][sematext]
* [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting) Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].

## Deployment and Infrastructure

* [Centurion](https://github.com/newrelic/centurion) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [@newrelic](https://github.com/newrelic)
* [Clocker](https://github.com/brooklyncentral/clocker) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and [Brooklyn][brooklyn] for application blueprints. By [@brooklyncentral](https://github.com/brooklyncentral)
* [Cloud 66](http://www.cloud66.com) - Full-stack hosted container management as a service
* [deploy](https://github.com/ttiny/deploy) - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration. Declarative configuration and short commands for managing (syncing, building, running) of infrastructures of more than a few services. Able to deploy whole preconfigured server or system of services with a single line (without having to scroll the line).
* [Docket](https://github.com/netvarun/docket) - Custom docker registry that allows for lightning fast deploys through bittorrent by [@netvarun](https://github.com/netvarun/)
* [Longshoreman](https://github.com/longshoreman/longshoreman) - Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By [longshoreman](https://github.com/longshoreman)


## PaaS
* [Atlantis](https://github.com/ooyala/atlantis) - Atlantis is an Open Source PaaS for HTTP applications built on Docker and written in Go
* [Convox Rack](https://github.com/convox/rack) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
* [Dcw](https://github.com/pbertera/dcw) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels. 
* [Deis](https://github.com/deis/deis) - Your PaaS, your rules http://deis.io/
* [Dockhero](https://dockhero.io/) - Dockhero is a Heroku add-on which turns a Docker image into a microservice attached to the Heroku app. Currently in beta.
* [Dokku][dokku] - Docker powered mini-Heroku in around 100 lines of Bash by [@progrium][progrium]
* [Empire](https://github.com/remind101/empire) - A PaaS built on top of Amazon EC2 Container Service (ECS)
* [Flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service https://flynn.io/
* [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by [Red Hat](https://www.redhat.com/)
* [Rancher][rancher] - Rancher is an open source project that provides a complete platform for operating Docker in production
* [Tsuru](https://github.com/tsuru/tsuru) - Tsuru is an extensible and open source Platform as a Service software https://tsuru.io/
* [DataMC](http://datamc.io/) - DataMc is a PaaS for Production ready and fully managed Data Platform

## Remote Container Manager / Orchestration
* [autodock](https://github.com/prologic/autodock) - Daemon for Docker Automation by [@prologic][prologic]
* [blimp](https://github.com/tubesandlube/blimp) - Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more. By [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin)
* [Capitan](https://github.com/byrnedo/capitan) - Composable docker orchestration with added scripting support by [@byrnedo](https://github.com/byrnedo).
* [CloudSlang](http://www.cloudslang.io/) - CloudSlang is a workflow engine to create Docker process automation
* [clusterdock](https://github.com/clusterdock/framework) - Docker container orchestration to enable the testing of long-running cluster deployments.
* [ContainerShip](https://github.com/containership/containership) A simple container management platform [containership]
* [CoreOS][coreos] - Linux for Massive Server Deployments https://coreos.com/
* [Crane](https://github.com/Dataman-Cloud/crane) - Control plane based on docker built-in swarm [@Dataman-Cloud](https://github.com/Dataman-Cloud)
* [Deploying a Containerized App on a Public Node with Mesos](https://docs.mesosphere.com/usage/tutorials/containerized-app/) - Docker plus Mesosphere provides an easy way to automate and scale deployment of containers in a production environment
* [ElasticKube](https://github.com/ElasticBox/elastickube) - Open source management platform for Kubernetes.
* [Fleet](https://github.com/coreos/fleet) - A Distributed init System providing low-level orchestration [coreos.com]
* [Flocker](https://github.com/ClusterHQ/flocker) - Flocker is a data volume manager and multi-host Docker cluster management tool by [@ClusterHQ](https://github.com/ClusterHQ)
* [Kontena](https://github.com/kontena/kontena) - Application Containers for Masses https://www.kontena.io/
* [Kubernetes][kubernetes] - Open source orchestration system for Docker containers by Google [kubernetes] See Also [awesome-kubernetes](https://github.com/ramitsurana/awesome-kubernetes) by [@ramitsurana][ramitsurana]
* [Maestro](https://github.com/toscanini/maestro) - Maestro provides the ability to easily launch, orchestrate and manage mulitiple Docker containers as single unit by [@tascanini](https://github.com/toscanini)
* [Mantl](https://github.com/mantl/mantl) - Mantl is a modern platform for rapidly deploying globally distributed services [@mantl](https://github.com/mantl/)
* [Marathon](https://mesosphere.github.io/marathon/docs/) - Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on")
* [MCollective Docker Agent](https://github.com/m4ce/mcollective-docker-agent) - Uses MCollective to orchestrate your Docker containers and images [@m4ce](https://github.com/m4ce)
* [Nomad Project](https://www.nomadproject.io/) - Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler.
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui/wiki) - Docker Management for Humans [panamax.io]
* [Portainer](http://portainer.io/) - A lightweight management UI for managing your Docker host or Docker Swarm cluster
* [Rancher](https://github.com/rancher/rancher) - Portable AWS-style infrastructure service for Docker http://rancher.com/
* [Serf](https://github.com/hashicorp/serf) - Service orchestration and management tool by [@hashicorp](https://github.com/hashicorp)
* [Shipyard](https://github.com/shipyard/shipyard) - Composable Docker Management http://shipyard-project.com/

## Security
* [Clair](https://github.com/coreos/clair) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [@coreos][CoreOS]
* [docker-bench-security](https://github.com/docker/docker-bench-security) - script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker][docker]
* [notary](https://github.com/docker/notary) - a server and a client for running and interacting with trusted collections. By [@docker][docker]
* [Twistlock](https://twistlock.com/)  - Twistlock Security Suite detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.

## Serverless
* [AMP](https://github.com/appcelerator/amp) - The open source Container-as-a-Service & microservice platform for Docker
* [Docker-Lambda](https://github.com/lambci/docker-lambda) - Docker images and test runners that replicate the live AWS Lambda environment
* [Funker](https://github.com/bfirsh/funker-example-voting-app) - Functions as Docker containers
* [IronFunctions](https://github.com/iron-io/functions) - The serverless microservices platform FaaS (Funcitons as a Service) which uses Docker containers to run Any language or AWS Lambda functions

## Service Discovery
* [Docker Grand Ambassador](https://github.com/cpuguy83/docker-grand-ambassador) - This is a fully dynamic docker link ambassador. + [Article](https://docs.docker.com/engine/admin/ambassador_pattern_linking/) by [@cpuguy83][cpuguy83]
* [docker-consul](https://github.com/gliderlabs/docker-consul) by [@progrium][progrium]
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery by [@coreOS][coreos]
* [proxy](https://github.com/factorish/proxy) - lightweight nginx based load balancer self using service discovery provided by registrator. by [@factorish](https://github.com/factorish)
* [wagl](https://github.com/ahmetalpbalkan/wagl/) - Service discovery for docker swarm using DNS

## Metadata
* [label-schema.org][labelschema] - a convention and shared namespace for Docker labels.
* [MicroBadger][microbadger] - add metadata to Docker images using labels.

# Slides
* [Docker Slideshare Account](http://www.slideshare.net/Docker)
* [Docker Security](http://www.slideshare.net/jpetazzo) with [@jpetazzo][jpetazzo]
* [Hide your DEV ENV in a container](http://www.slideshare.net/JohanJanssen4/hide-your-development-environment-and-application-in-a-container) by [@johanjanssen42](https://twitter.com/johanjanssen42)
* [Docker for the new era](https://www.slideshare.net/ramitsurana/docker-for-the-new-era)  by [@ramitsurana][ramitsurana]

# Useful Articles

## Main Resources
* [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers) more general about container than this repo, by [@Friz-zy](https://github.com/Friz-zy).
* [Blog](http://crosbymichael.com/) of [@crosbymichael][crosbymichael]
* [Blog](http://gliderlabs.com/devlog/) of [@gliderlabs][gliderlabs]
* [Blog](http://jasonwilder.com/) of [@jwilder][jwilder]
* [Blog](http://jpetazzo.github.io/) of [@jpetazzo][jpetazzo]
* [Blog](http://progrium.com/blog/) of [@progrium][progrium]
* [Blog](http://sebgoa.blogspot.be/) of [@sebgoa][sebgoa]
* [Blog](https://blog.codeship.com/) of [@codeship](https://github.com/codeship)
* [Blog](https://blog.jessfraz.com/) of [@frazelledazzell][jfrazelle]
* [Container solutions](http://container-solutions.com/blog/)
* [Container42](http://container42.com/)
* [Digital Ocean Community](https://www.digitalocean.com/community/search?q=docker&type=tutorials)
* [Docker Cheat Sheet](http://docker.jens-piegsa.com) by [@JensPiegsa][JensPiegsa] *(updated for Docker 1.13)*
* [Docker Cheat Sheet][docker-cheat-sheet] by [@wsargent][wsargent] __MUST SEE__
* [Docker Containers on the desktop][jessblog] by [@jfrazelle][jfrazelle] The **funniest way** to learn
* [Docker Ecosystem](http://comp.photo777.org/wp-content/uploads/2015/09/Docker-ecosystem-8.5.1.pdf) (PDF) __MUST SEE__    find it on [blog](http://comp.photo777.org/docker-ecosystem/) by Bryzgalov Peter.
* [Docker Ecosystem](https://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) __MUST SEE__
* [Docker Kubernetes Lab Handbook](https://github.com/xiaopeng163/docker-k8s-lab)
* [Docker Printable Refcard][docker-quick-ref] by [@dimonomid][dimonomid]
* [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](http://www.rightscale.com/blog/cloud-management-best-practices/docker-vs-vms-combining-both-cloud-portability-nirvana)
* [Docker Weekly](https://blog.docker.com/docker-weekly-archives/) Huge resource
* [DockerOne](http://dockone.io/) Docker Community (in Chinese) by [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)
* [Project Web Dev][projwebdev] : (Article series) How to create your own website based on Docker
* [Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html) Very complete about docker! (Tips: checkout her [dotfiles][jfrazelledotfiles] and her [dockerfiles][jfrazelledockerfiles]))

## General Articles
* [Automating Docker Logging: ElasticSearch, Logstash, Kibana, and Logspout](https://nathanleclaire.com/blog/2015/04/27/automating-docker-logging-elasticsearch-logstash-kibana-and-logspout/) (written 27 APR 2015)
* [Comparing Five Monitoring Options for Docker](http://rancher.com/comparing-monitoring-options-for-docker-deployments/)
* [Containerization and the PaaS Cloud](https://www.computer.org/cms/Computer.org/ComputingNow/issues/2015/09/mcd2015030024.pdf) -- This article discusses the requirements that arise from having to facilitate applications through distributed multicloud platforms.
* [Continuous Deployment with Gradle and Docker](https://github.com/gesellix/pipeline-with-gradle-and-docker/blob/master/README.md) Describes a complete pipeline from source to production deploy (includes a complete Spring Boot example project) by [@gesellix][gesellix]
* [Deploying NGINX with Docker](https://www.nginx.com/blog/deploying-nginx-nginx-plus-docker/)
* [Docker Adoption Data](https://www.datadoghq.com/docker-adoption/) A study by Datadog on the real world Docker usage statistics and deployment patterns.
* [Docker for (Java) Developers](https://ro14nd.de/Docker-for-Developers)
* [Docker for Development: Common Problems and Solutions](https://medium.com/@rdsubhas/docker-for-development-common-problems-and-solutions-95b25cae41eb) by [@rdsubhas](https://github.com/rdsubhas)
* [Docker Host Volume Synchronization](http://oliverguenther.de/2015/05/docker-host-volume-synchronization/) (written 1 JUN 2015)
* [Docker: Build, Ship and Run Any App, Anywhere](http://delftswa.github.io/chapters/docker/index.html) by [Martijn Dwars](https://github.com/MartijnDwars), [Wiebe van Geest](https://github.com/wrvangeest), [Rik Nijessen](https://github.com/gewoonrik), and [Rick Wieman](https://github.com/RickWieman) from [Delft University of Technology](http://www.tudelft.nl/) (written 2 JUL 2015)
* [Dockerizing Flask With Compose and Machine - From Localhost to the Cloud](https://realpython.com/blog/python/dockerizing-flask-with-compose-and-machine-from-localhost-to-the-cloud/) -- [GitHub](https://github.com/realpython/orchestrating-docker) Learn how to deploy an application using Docker Compose and Docker Machine (written 17 April 2015)
* [Eight Docker Development Patterns](http://hokstad.com/docker/patterns)
* [From Local Development to Remote Deployment with Docker Machine and Compose](https://developer.rackspace.com/blog/dev-to-deploy-with-docker-machine-and-compose/) (written 2 JUL 2015)
* [Getting Started with Docker](https://serversforhackers.com/getting-started-with-docker) by [@fideloper](https://github.com/fideloper) -- [Servers For Hackers](https://serversforhackers.com/editions) is valuable resource. At some point, every programmer finds themselves needing to know their way around a server.
* [How to monitor Docker](https://www.datadoghq.com/blog/the-docker-monitoring-problem/) (4-part series)
* [How to Use Docker on OS X: The Missing Guide](https://www.viget.com/articles/how-to-use-docker-on-os-x-the-missing-guide)
* [Joining the Docker Ship](http://thenewstack.io/joining-the-docker-ship-and-go/) Learn how to contribute to docker (written 9 JUL 2015)
* [Logging on Docker: What You Need to Know](https://dzone.com/articles/logging-docker-what-you-need) + see the [video][loggingDocker] (~50min)
* [Make better software with Docker](https://www.shivering-isles.com/how-docker-helps-to-create-better-software/) A short article about how docker changes the way of writing software. And provide an easy way to convert from a monolite application to a microserviced application
* [Minimalistic data-only container for Docker Compose](http://dockermeetupsinbordeaux.github.io/docker-compose/data-container/2015/03/01/minimalistic-docker-data-container.html) (Written Mar 1, 2015)
* [Node.js in Containers Using Docker](https://medium.com/software-engineering/node-js-in-containers-using-docker-48649e4a9d67) Good introduction to `docker` and `docker-compose` (Written Jan, 2017)
* [Running Docker Containers with Systemd](http://container-solutions.com/running-docker-containers-with-systemd/)
* [Swarm v. Fleet v. Kubernetes v. Mesos](https://www.oreilly.com/ideas/swarm-v-fleet-v-kubernetes-v-mesos) Comparing different orchestration tools. (written OCT 2015)
* [The Shortlist of Docker Hosting](https://blog.codeship.com/the-shortlist-of-docker-hosting) There are so many specialized and optimized Docker hosting services available, it’s high time for a review to see what’s on offer (by Chris Ward).
* [Using Ansible with Docker Machine to Bootstrap Host Nodes](https://nathanleclaire.com/blog/2015/11/10/using-ansible-with-docker-machine-to-bootstrap-host-nodes/) by [@nathanleclaire](https://github.com/nathanleclaire)
* [What is Docker and how do you monitor it?](http://axibase.com/docker-monitoring/)
* [Why and How to use Docker for Development](https://medium.com/iron-io-blog/why-and-how-to-use-docker-for-development-a156c1de3b24) (written 28 APR 2015)

## Portuguese Articles
* [Uma rápida introdução ao Docker e instalação no Ubuntu](https://woliveiras.com.br/posts/uma-rapida-introducao-ao-docker-e-instalacao-no-ubuntu/)
* [O que é uma imagem e o que é um container Docker?](https://woliveiras.com.br/posts/imagem-docker-ou-um-container-docker/)
* [Criando uma imagem Docker personalizada](https://woliveiras.com.br/posts/Criando-uma-imagem-Docker-personalizada/)
* [Comandos mais utilizados no Docker](https://woliveiras.com.br/posts/comandos-mais-utilizados-no-docker/)

## Deep Dive
* [Creating containers - Part 1](http://crosbymichael.com/creating-containers-part-1.html) This is part one of a series of blog posts detailing how docker creates containers. By [@crosbymichael][crosbymichael]
* [Data-only container madness](http://container42.com/2014/11/18/data-only-container-madness/)

## Networking
* [Using Docker Machine with Weave 0.10](https://www.weave.works/using-docker-machine-with-weave-0-10/) (written 22 APR 2015)
* [How to Route Traffic through a Tor Docker container](https://blog.jessfraz.com/post/routing-traffic-through-tor-docker-container/) by [@jfrazelle][jfrazelle] (writtent 20 JUN 2015)
* [Demystifing Docker overlay networking](http://blog.nigelpoulton.com/demystifying-docker-overlay-networking/). By [@nigelpoulton] (written 12 OCT 2016)

## Metal
* [How to use Docker on Full Metal](http://blog.bigstep.com/use-docker-full-metal-cloud/)
* [CargoOS](https://cargos.io/) A bare essential OS for running the Docker Engine on bare metal or Cloud.

## Multi-Server
* [A Docker based mini-PaaS](http://shortcircuit.net.au/~prologic/blog/article/2015/03/24/a-docker-based-mini-paas/)
by [@prologic][prologic]
* [A multi-host scalable web services demo using Docker swarm, Docker compose, NGINX, and Blockbridge](https://www.blockbridge.com/a-scalable-web-services-demo-using-docker-swarm-compose-and-blockbridge/)

## Cloud Infrastructure
* [Cloud Infrastructure Automation for Docker Nodes](https://blog.tutum.co/2015/04/29/cloud-infrastructure-automation-for-docker-nodes/)
* [3 clear trends in Amazon ECS adoption and usage](https://www.datadoghq.com/blog/3-clear-trends-in-ecs-adoption/)

## Good Tips
* [10 practical docker tips](http://www.smartjava.org/content/10-practical-docker-tips-day-day-docker-usage) (Dec 2015) by [@josdirksen](https://github.com/josdirksen)
* [10 Things Not To Forget Before Deploying Docker In Production](http://www.slideshare.net/rightscale/docker-meetup-40826948)
* [24 random docker tips](https://csabapalfi.github.io/random-docker-tips/) by [@csabapalfi](https://github.com/csabapalfi)
* [6 Million Ways To Log In Docker](http://www.slideshare.net/raychaser/6-million-ways-to-log-in-docker-nyc-docker-meetup-12172014) by [@raychaser](https://twitter.com/raychaser)
* [A Simple Way to Dockerize Applications](http://jasonwilder.com/blog/2014/10/13/a-simple-way-to-dockerize-applications/) by [@jwilder][jwilder]
* [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/) by [@jwilder][jwilder]
* [Building good docker images](http://jonathan.bergknoff.com/journal/building-good-docker-images) by [@jbergknoff](https://github.com/jbergknoff)
* [Container Best Practices](http://docs.projectatomic.io/container-best-practices/) - Red Hat's Project Atomic created a Container Best Practices guide which applies to everything and is updated regurlary.
* [Dealing with linked containers dependency in docker-compose](http://brunorocha.org/python/dealing-with-linked-containers-dependency-in-docker-compose.html) by [@rochacbruno](https://github.com/rochacbruno)
* [Docker CIFS – How to Mount CIFS as a Docker Volume](http://backdrift.org/docker-cifs-howto-mount-cifs-volume-docker-container)
* [Docker on Windows behind a firewall](http://toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by [@kaitoedter](https://twitter.com/kaitoedter)
* [Docker Tips](http://www.mervine.net/notes/docker-tips) by [@jmervine](https://github.com/jmervine)
* [Dockerfile Generator](http://jrruethe.github.io/blog/2015/09/20/dockerfile-generator/) (ruby script)
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm][fgrehm]
* [Kubernetes Cheatsheet](http://k8s.info/cs.html) - A great resource for managing your Kubernetes installation
* [Nginx Proxy for Docker](https://blog.danivovich.com/2015/07/09/nginx-proxy-for-docker-containers/) (written 9 JUL 2015)
* [Production Meteor and Node Using Docker, Part I](https://projectricochet.com/blog/production-meteor-and-node-using-docker-part-i) by [@projectricochet](https://github.com/projectricochet)
* [Pulling Git into a Docker image without leaving SSH keys behind](http://blog.cloud66.com/pulling-git-into-a-docker-image-without-leaving-ssh-keys-behind/) by [@khash](https://github.com/khash)
* [Resource Management in Docker](https://goldmann.pl/blog/2014/09/11/resource-management-in-docker/) by [@marekgoldmann](https://twitter.com/marekgoldmann)
* [Running Production Hadoop Clusters in Docker Containers](https://conferences.oreilly.com/strata/big-data-conference-ca-2015/public/schedule/detail/38521)
* [Using NSEnter with Boot2Docker](https://ro14nd.de/NSEnter-with-Boot2Docker)

## Newsletter
* [Docker Team](https://www.docker.com/)
* [Tutum](https://dashboard.tutum.co/)
* [Shippable](http://blog.shippable.com/)
* [WebOps weekly](http://webopsweekly.com/)

## Continuous Integration
* [Docker and Phoenix: How to Make Your Continuous Integration More Awesome](https://ariya.io/2014/12/docker-and-phoenix-how-to-make-your-continuous-integration-more-awesome)
* [Jenkins 2.0 - Screencast Series](http://theremotelab.com/blog/jenkins2.0-screencast-series/) by [Virendra Bhalothia](https://twitter.com/bhalothiaa)
* [Pushing to ECR Using Jenkins Pipeline Plugin](https://blog.mikesir87.io/2016/04/pushing-to-ecr-using-jenkins-pipeline-plugin/) by [@mikesir87](https://github.com/mikesir87)

## Optimizing Images
* [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07) by [@kelseyhightower](https://github.com/kelseyhightower)
* [Create the smallest possible Docker container](http://blog.xebia.com/create-the-smallest-possible-docker-container/)
* [Creating a Docker image from your code](https://blog.tutum.co/2014/04/10/creating-a-docker-image-from-your-code/)
* [Docker - Minimize your containers with alpine linux](https://www.shivering-isles.com/docker-minimize-your-containers-with-alpine-linux/) A article about how refactor a Dockerfile to switch from a heavyweight container to a lightweight container based on alpine linux.
* [Dockerfile Golf (or optimizing the Docker build process)](http://www.davidmkerr.com/2014/08/dockerfile-golf-or-optimizing-docker.html)
* [DockerSlim](https://github.com/docker-slim/docker-slim) shrinks fat Docker images creating the smallest possible images.
* [How to Optimize Your Dockerfile](https://blog.tutum.co/2014/10/22/how-to-optimize-your-dockerfile/) by [@tutumcloud](https://github.com/tutumcloud)
* [MicroBadger][microbadger] - Analyze the contents of images and add metadata labels
* [Optimizing Docker Images](https://www.ctl.io/developers/blog/post/optimizing-docker-images/)
* [SkinnyWhale](https://github.com/djosephsen/skinnywhale) Skinnywhale helps you make smaller (as in megabytes) Docker containers.
* [Squashing Docker Images](http://jasonwilder.com/blog/2014/08/19/squashing-docker-images/) by [@jwilder][jwilder]

## Service Discovery
* [@progrium][progrium] Service Discovery articles series:
* [Consul Service Discovery with Docker](http://progrium.com/blog/2014/08/20/consul-service-discovery-with-docker/)
* [Understanding Modern Service Discovery with Docker](http://progrium.com/blog/2014/07/29/understanding-modern-service-discovery-with-docker/)
* [Automatic Docker Service Announcement with Registrator](http://progrium.com/blog/2014/09/10/automatic-docker-service-announcement-with-registrator/)

## Security
* [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
* [Docker and SELinux](http://www.projectatomic.io/docs/docker-and-selinux/)
* [Docker Secure Deployment Guidelines](https://github.com/GDSSecurity/Docker-Secure-Deployment-Guidelines)
* [Docker Security Cheat Sheet](https://github.com/konstruktoid/Docker/blob/master/Security/CheatSheet.adoc)
* [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](http://fr.slideshare.net/MichaelBoelen/docker-security-are-your-containers-tightly-secured-to-the-ship)
* [How CVE's are handled on Offical Docker Images](https://github.com/docker-library/official-images/issues/1448)
* [Improving Docker Security with Authenticated Volumes](https://www.blockbridge.com/improving-docker-security-with-authenticated-volumes/)
* [Lynis is an open source security auditing tool including Docker auditing](https://cisofy.com/lynis/)
* [Security Best Practices for Building Docker Images](https://linux-audit.com/tag/docker/)
* [Tuning Docker with the newest security enhancements](https://opensource.com/business/15/3/docker-security-tuning)
* [Understanding Docker security and best practices](https://blog.docker.com/2015/05/understanding-docker-security-and-best-practices/) (written 5 MAY 2015)

## Performances
* [Distrubuted JMeter testing using Docker](http://srivaths.blogspot.fr/2014/08/distrubuted-jmeter-testing-using-docker.html?m=1)
* [nsinit: per-container resource monitoring of Docker containers on RHEL/Fedora](http://www.breakage.org/2014/09/03/nsinit-per-container-resource-monitoring-of-docker-containers-on-rhelfedora/)
* [Performance Analysis of Docker on Red Hat Enterprise Linux 7](http://developerblog.redhat.com/2014/08/19/performance-analysis-docker-red-hat-enterprise-linux-7/)

## Raspberry Pi & ARM
* [Docker on Raspberry Pi](http://blog.xebia.com/docker-on-a-raspberry-pi/)
* [Docker Pirates ARMed with explosive stuff](http://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
* [Fool-Proof Recipe: Docker on the Raspberry Pi](https://www.voxxed.com/blog/2015/04/fool-proof-recipe-docker-on-the-raspberry-pi/) Same article as above but more opinionated.
* [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps)
* [git push docker containers to linux devices](https://resin.io/) Modern DevOps for IoT, leveraging git and Docker.
* [How to run 2500 webservers on a Raspberry Pi](http://blog.loof.fr/2015/10/how-to-run-2500-webservers-on-raspberry.html)
* [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-(ARMv7)-devices)
* [Raspberry Pi with Docker 1.5.0](http://blog.hypriot.com/post/heavily-armed-after-major-upgrade-raspberry-pi-with-docker-1-dot-5-0/)
* [Swarming Raspberry Pi – Part 1](http://matthewkwilliams.com/index.php/2015/03/21/swarming-raspberry-pi-part-1/)
* [Swarming Raspberry Pi, Part 2: Registry & Mirror](http://matthewkwilliams.com/index.php/2015/03/29/swarming-raspberry-pi-part-2-registry-mirror/)
* [Swarming Raspberry Pi: Docker Swarm Discovery Options](http://matthewkwilliams.com/index.php/2015/04/03/swarming-raspberry-pi-docker-swarm-discovery-options/)
* [Uniform Development by Docker & QEMU](http://www.instructables.com/id/Uniform-Development-by-Docker-QEMU/)


## Other
* Presentation: Docker and JBoss - the perfect combination
 * [Vidéo](https://www.youtube.com/watch?v=4uQ6gR_xZhE)
 * [Code source](https://github.com/goldmann/goldmann.pl/tree/master/.presentations/2014-vjbug-docker/demos)
 * [JBoss and Docker Presentation](https://goldmann.pl/presentations/2014-vjbug-docker/ )

# Books

## English
* [Docker Book](https://dockerbook.com/) by James Turnbul ([@kartar][kartar])
* [Docker Cookbook](http://dockercookbook.github.io/) by Neependra Khare ([@neependra](https://twitter.com/neependra)) (Publisher: Packt)
* [Docker Cookbook](http://shop.oreilly.com/product/0636920036791.do) by Sébastien Goasguen ([@sebgoa][sebgoa]) (Publisher: O'Reilly)
* [Docker High Performance](https://www.packtpub.com/networking-and-servers/docker-high-performance) by Allan Espinosa (Publisher: Packt)
* [Docker in Action](https://www.manning.com/books/docker-in-action) by Jeff Nickoloff ([@allingeek](https://twitter.com/allingeek))
* [Docker in Practice](https://www.manning.com/books/docker-in-practice) by Ian Miell ([@ianmiell][ianmiell]) and Aidan Hobson Sayers ([@aidanhs](https://github.com/aidanhs)). ==> [Website](http://docker-in-practice.github.io/)
* [Docker in Production: Lessons from the Trenches](http://www.amazon.com/Docker-Production-Trenches-Joe-Johnston-ebook/dp/B0141W6KYC) by Joe Johnston (Author), John Fiedler (Author), Milos Gajdos (Author), Antoni Batchelli (Author), Justin Cormack (Author)
* [Docker Security](https://www.openshift.com/promotions/docker-security.html) by Adrian Mouat ([@adrianmouat](https://twitter.com/adrianmouat)) (Publisher: O'Reilly)
* [Docker Up & Running](https://newrelic.com/docker-book) by [Karl Matthias](https://twitter.com/relistan) and [Sean P. Kane](https://twitter.com/spkane)
* [Extending Docker](https://www.packtpub.com/networking-and-servers/extending-docker) by Russ McKendrick (Publisher: Packt)
* [Kubernetes Up and Running: Dive into the Future of Infrastructure](http://shop.oreilly.com/product/0636920043874.do) by Kelsey Hightower ([@kelseyhightower](https://twitter.com/kelseyhightower)) (Publisher: O'Reilly)
* [Kubernetes](https://www.openshift.com/promotions/kubernetes.html) by [David Rensin](https://research.google.com/pubs/DavidRensin.html) (Publisher: O'Reilly)
* [Learning Docker Networking](https://www.packtpub.com/networking-and-servers/learning-docker-networking) by Rajdeep Dua, Vaibhav Kohli and Santosh Kumar Konduri (Publisher: Packt)
* [Learning Docker](https://www.packtpub.com/virtualization-and-cloud/learning-docker) by Pethuru Raj, Jeeva S. Chelladhurai and Vinod Singh (Publisher: Packt)
* [Mastering Docker](https://www.packtpub.com/virtualization-and-cloud/mastering-docker) by Scott Gallagher (Publisher: Packt)
* [Orchestrating Docker](https://www.packtpub.com/virtualization-and-cloud/orchestrating-docker) by Shrikrishna Holla (Publisher: Packt)
* [Securing Docker](https://www.packtpub.com/virtualization-and-cloud/securing-docker) by Scott Gallagher (Publisher: Packt)
* [Troubleshooting Docker](https://www.packtpub.com/virtualization-and-cloud/troubleshooting-docker) by John Wooten, Navid Shaikh (Publisher: Packt)
* [Using Docker](http://shop.oreilly.com/product/0636920035671.do) by Adrian Mouat ([@adrianmouat](https://twitter.com/adrianmouat)) (Publisher: O'Reilly)

## Chinese
* [The Source Code Analysis of Docker](https://www.amazon.cn/图书/dp/B012ROMRUM) (Chinese) by [Allen Sun](https://github.com/allencloud)
* [Docker Container and Container Cloud](https://www.amazon.cn/图书/dp/B014ETH1IG) (Chinese) by [Harry Zhang](https://twitter.com/resouer) & Jianbo Sun & Zhejiang University SEL Laboratory

## Portuguese
* [Containers com Docker do desenvolvimento à produção](https://www.casadocodigo.com.br/products/livro-docker) by Daniel Romero ([@infoslack](https://twitter.com/infoslack))
* [Aprendendo Docker: Do básico à orquestração de contêineres](http://aprendendodocker.com.br/) by Wellington F. Silva ([@_wsilva](https://twitter.com/_wsilva)) (Publisher: Editora Novatec)
* [Docker para Desenvolvedores](https://leanpub.com/dockerparadesenvolvedores) by Rafael Gomes ([@gomex](https://twitter.com/gomex)) (Publisher: Leanpub) - 80% finished

# Videos

## Main Account
* [Docker Youtube Account](https://www.youtube.com/user/dockerrun)
* [CenturyLink Labs Docker Interviews](https://www.youtube.com/playlist?list=PL_q4Fk7SVBCIjyuCBFBItXnzGI3qBa2L1)
* [Container Camp](https://www.youtube.com/channel/UCvksXSnLqIVM_uFB7xyrsSg/videos) Conference about *containers*!!! [@containercamp](https://twitter.com/containercamp)
* [Quoi d'neuf Docker](https://www.youtube.com/channel/UCOAhkxpryr_BKybt9wIw-NQ/videos) **FRENCH** chronique vidéo sur Youtube proposant de courtes vidéos (maximum 15 minutes) sur la thématique "Docker et son écosystème" [Site Web](http://www.quoidneufdocker.xyz/)

## Useful videos
* [Ansible and Docker HP](https://www.youtube.com/watch?v=oZ45v8AeE7k) (32:38)
* [Container Hacks and Fun Images][jessvid] by [@jfrazelle][jfrazelle] @ DockerCon 2015 (**MUST WATCH VIDEO**: 38:50)
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
* [Run Any App on Mesos on Any Infrastructure Using Docker](https://www.youtube.com/watch?v=u5jd9YT9EsY) (17:44)
* [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course 
* [State of containers: a debate with CoreOS, VMware and Google](https://www.youtube.com/watch?v=IiITP3yIRd8) (27:38)
* [SysAdminCasts: Introduction to Docker](https://sysadmincasts.com/episodes/31-introduction-to-docker) (15:49)

# Interactive Learning Environments
* [Katacoda](https://www.katacoda.com/): Learn Docker using Interactive Browser-Based Labs

# Interesting Twitter Accounts
* [CenturyLink Labs](https://twitter.com/CenturyLinkLabs)
* [Docker Captains Twitter List](https://twitter.com/EltonStoneman/lists/docker-captains)
* [Docker News](https://twitter.com/dockernews)
* [Docker](https://twitter.com/docker)
* [Flux7Labs](https://twitter.com/Flux7Labs)
* [OpenShift by Red Hat](https://twitter.com/openshift)
* [Project Atomic](https://twitter.com/ProjectAtomic)
* [The New Stack](https://twitter.com/thenewstack)
* [YLD](https://twitter.com/YLDio)

## People
* [Docker Captains](https://www.docker.com/community/docker-captains) - Docker experts and community leaders
* [Gabriel Monroy](https://twitter.com/gabrtv) Creator of Deis
* [James Turnbull][kartar] Author of Docker Book
* [Jeff Lindsay](https://twitter.com/progrium) Design-minded software architect
* [Jérôme Petazzoni](https://twitter.com/jpetazzo) Docker Developer
* [Jessie Frazelle](https://twitter.com/jessfraz) Ex-@docker maintainer and uses full containerized desktop, lots of fun.
* [Michael Crosby](https://twitter.com/crosbymichael) Docker Developer
* [Solomon Hykes](https://twitter.com/solomonstre) Founder of Docker

# Communities and Meetups
* [Docker Community](https://www.docker.com/docker-community)
* [Docker On Line Meetup](https://www.meetup.com/Docker-Online-Meetup/)
* [Docker Regional Meetup List](https://www.docker.com/community/meetup-groups)

## Other Brazilian Channels
* [Docker BR on Telegram](https://telegram.me/dockerbr)
* [Docker BR on Slack](http://docker-br.herokuapp.com) - Auto invite url


[ahmetalpbalkan]: https://github.com/ahmetalpbalkan
[alpine]: https://github.com/gliderlabs/docker-alpine
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
[docker4dev]: https://www.youtube.com/watch?v=FdkNAjjO5yQ
[dokku]: https://github.com/dokku/dokku
[editREADME]: https://github.com/veggiemonk/awesome-docker/edit/master/README.md
[fgrehm]: https://github.com/fgrehm
[fluentd]: https://github.com/kiyoto/docker-fluentd
[gesellix]: https://github.com/gesellix
[gliderlabs]: https://github.com/gliderlabs
[gondor]: https://github.com/gondor
[ianmiell]: https://github.com/ianmiell
[JensPiegsa]: https://github.com/JensPiegsa
[jessblog]: https://blog.jessfraz.com/post/docker-containers-on-the-desktop/
[jessvid]: https://www.youtube.com/watch?v=1qlLUf7KtAw
[jfrazelle]: https://github.com/jfrazelle
[jfrazelledockerfiles]: https://github.com/jessfraz/dockerfiles
[jfrazelledotfiles]: https://github.com/jessfraz/dotfiles
[jpetazzo]: https://github.com/jpetazzo
[jwilder]: https://github.com/jwilder
[kartar]: https://twitter.com/kartar
[kiyoto]: https://github.com/kiyoto
[kubernetes]: https://kubernetes.io
[labelschema]: http://label-schema.org
[loggingDocker]: https://vimeo.com/123341629
[microbadger]: https://microbadger.com
[nginxproxy]: https://github.com/jwilder/nginx-proxy
[noteed]: https://github.com/noteed
[ondrejmo]: https://github.com/ondrejmo
[openshift]: https://www.openshift.org/
[panamax.io]: http://panamax.io/
[pandrew]: https://github.com/pandrew
[percheron]: https://github.com/ashmckenzie/percheron
[progrium]: https://github.com/progrium
[projwebdev]: http://project-webdev.blogspot.de
[prologic]: https://github.com/prologic
[ramitsurana]: https://github.com/ramitsurana
[rancher]: https://github.com/rancher
[sebgoa]: https://twitter.com/sebgoa
[sematext]: https://twitter.com/sematext
[sindresorhus]: https://github.com/sindresorhus/awesome
[spm]: https://github.com/sematext/sematext-agent-docker
[vegasbrianc]: https://github.com/vegasbrianc
[vimagick]: https://github.com/vimagick
[weave]: https://github.com/weaveworks/weave
[wsargent]: https://github.com/wsargent
