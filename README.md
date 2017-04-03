
# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/veggiemonk/awesome-docker](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/veggiemonk/awesome-docker?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge) [![Say Thanks](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg)](https://saythanks.io/to/veggiemonk)

> A curated list of Docker resources and projects
Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome][sindresorhus] and improved by these **[amazing contributors](https://github.com/veggiemonk/awesome-docker/graphs/contributors)**.

It's now a GitHub project because it's considerably easier for other people to edit, fix and expand on Docker using GitHub.
Just click [README.md][editREADME] to submit a [pull request][editREADME].
If this list is not complete, you can [contribute][editREADME] to make it so.

***You can see the updates from [TWITTER](https://twitter.com/awesome_docker)***

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)** for tips!

#### *If you see a link here that is not (any longer) a good fit, you can fix it by submitting a [pull request][editREADME] to improve this file. Thank you!*

The creators and maintainers of this list do not receive and should not receive any form of payment to accept a change made by any contributor.
The goal of this repo is to index open-source projects, not to advertise for profit.

All the links are monitored and tested with [awesome_bot](https://github.com/dkhamsing/awesome_bot) made by [@dkhamsing](https://github.com/dkhamsing)

# Contents

- [What is Docker ?](#what-is-docker-)
- [Where to start ?](#where-to-start-)
- [Tools](#tools)
	- [Continuous Integration / Continuous Delivery](#continuous-integration--continuous-delivery)
	- [Deployment and Infrastructure](#deployment-and-infrastructure)
	- [Developer Tools](#developer-tools)
	- [Development Environments](#development-environments)
	- [Docker Compose file](#docker-compose-file)
	- [Dockerfile](#dockerfile)
	- [Garbage Collection](#garbage-collection)
	- [Hosting Images (registries)](#hosting-images-registries)
	- [Image Builder](#image-builder)
	- [Images](#images)
	- [Linter / Validator](#linter--validator)
	- [Local Container Manager](#local-container-manager)
	- [Monitoring & Logging](#monitoring--logging)
	- [Networking](#networking)
	- [PaaS](#paas)
	- [Remote Container Manager / Orchestration](#remote-container-manager--orchestration)
	- [Reverse Proxy](#reverse-proxy)
	- [Security](#security)
	- [Serverless](#serverless)
	- [Service Discovery](#service-discovery)
	- [CaaS - Services for running containers](#services-for-running-containers)
	- [Terminal User Interface](#terminal-user-interface)
	- [Testing](#testing)
	- [Utilities](#utilities)
	- [Volume management and plugins](#volume-management-and-plugins)
	- [Web Interface](#web-interface)
- [Slides](#slides)
- [Useful Resources](#useful-resources)
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
* [Docker Curriculum](https://prakhar.me/docker-curriculum/): A comprehensive tutorial for getting started with Docker. Teaches how to use Docker and deploy dockerized apps on AWS with Elastic Beanstalk and Elastic Container Service.
* [Install Docker on your machine](https://github.com/wsargent/docker-cheat-sheet#installation) and play with a few [Useful Images](#useful-images)
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
* [Play With Docker](http://labs.play-with-docker.com/) - A great way to get started with Docker. Docker runs directly in your browser.
* [Katacoda](https://www.katacoda.com/): Learn Docker using Interactive Browser-Based Labs

# Where to start - on Windows ?
* [Windows Containers Quick Start](https://docs.microsoft.com/en-us/virtualization/windowscontainers/quick_start/quick_start) Overview of Windows containers, drilling down to Quick Starts for Windows 10 and Windows Server 2016
* [Build And Run Your First Docker Windows Server Container](https://blog.docker.com/2016/09/build-your-first-docker-windows-server-container/) Walkthrough installing Docker on Windows 10, building a Docker image and running a Windows container
* Video: [Windows Containers and Docker: The 101](https://www.youtube.com/watch?v=N7SG2wEyQtM) A 20-minute overview, using Docker to run  PowerShell, ASP.NET Core and ASP.NET apps
* [A Comparative Study of Docker Engine on Windows Server vs Linux](http://collabnix.com/archives/1965) Comparing the feature sets and implementations of Docker on Windows and Linux
* [Docker with Microsoft SQL 2016 + ASP.NET](http://blog.alexellis.io/docker-does-sql2016-aspnet/) Demonstration running ASP.NET and SQL Server workloads in Docker
* [Running a Legacy ASP.NET App in a Windows Container](https://blog.sixeyed.com/dockerizing-nerd-dinner-part-1-running-a-legacy-asp-net-app-in-a-windows-container/) Steps for Dockerizing a legacy ASP.NET app and runnning as a Windows container
* [Exploring ASP.NET Core with Docker in both Linux and Windows Containers](http://www.hanselman.com/blog/ExploringASPNETCoreWithDockerInBothLinuxAndWindowsContainers.aspx) Running ASP.NET Core apps in Linux and Windows containers, using [Docker for Windows](https://docs.docker.com/docker-for-windows/)

----

# Tools

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
* [Docker Machine](https://github.com/docker/machine) (Machine management for a container-centric world)
* [Docker Registry][distribution] (The Docker toolset to pack, ship, store, and deliver content)
* [Docker Swarm](https://github.com/docker/swarm) (Swarm: a Docker-native clustering system)

## Continuous Integration / Continuous Delivery
* [Awesome-ciandcd](https://github.com/ciandcd/awesome-ciandcd) - Not specific to docker but relevant.
* [Buddy](https://buddy.works) - The best of Git, build & deployment tools combined into one powerful tool that supercharged our development
* [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery by [@harbur](https://github.com/harbur)
* [Cyclone](https://github.com/caicloud/cyclone) - A cloud native CI/CD platform built for container workflow by [@caicloud](https://github.com/caicloud).
* [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
* [Dockunit](https://github.com/dockunit/platform) - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By [@dockunit](https://github.com/dockunit)
* [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
* [GitLab CI](https://about.gitlab.com/gitlab-ci/) - GitLab has integrated CI to test, build and deploy your code with the use of GitLab runners.
* [GOCD-Docker](https://github.com/gocd/gocd-docker)Go Server and Agent in docker containers to provision.
* [InSpec](https://github.com/chef/inspec) - InSpec is an open-source testing framework for infrastructure with a human- and machine-readable language for specifying compliance, security and policy requirements.
* [Microservices Continuous Deployment](https://github.com/francescou/docker-continuous-deployment) - Continuous deployment of a microservices application
* [Screwdriver](http://screwdriver.cd/) - Yahoo's OpenSource buildplatform designed for Continous Delivery
* [Skipper](https://github.com/Stratoscale/skipper) - Easily dockerize your Git repository by [@Stratoscale](https://github.com/Stratoscale)
* [SwarmCI](https://github.com/ghostsquad/swarmci) - Create a distributed, isolated task pipeline in your Docker Swarm.
* [Watchtower](https://github.com/v2tec/watchtower) - Automatically update running Docker containers by [@CenturyLinkLabs][CenturyLinkLabs]

### CI Services
* [CircleCI](https://circleci.com/) - Push or pull Docker images from your build environment, or build and run containers right on CircleCI.
* [CodeFresh](https://codefresh.io) - Everything you need to build, test, and share your Docker applications. Provides automated end to end testing.
* [CodeShip](https://codeship.com/features/pro) - Work with your established Docker workflows while automating your testing and deployment tasks with our hosted platform dedicated to speed and security.
* [Semaphore CI](https://semaphoreci.com/product/docker) — A high-performance cloud solution that makes it easy to build, test and ship your containers to production.
* [Shippable](https://app.shippable.com/) - A SaaS platform for developers and DevOps teams that significantly reduces the time taken for code to be built, tested and deployed to production.
* [IBM DevOps Services](https://hub.jazz.net) - Continuous delivery using a pipeline deployment onto IBM Containers on Bluemix.

## Deployment and Infrastructure
* [Centurion](https://github.com/newrelic/centurion) - Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [@newrelic](https://github.com/newrelic)
* [Clocker](https://github.com/brooklyncentral/clocker) - Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts, on both Docker and Marathon. It leverages [Calico][calico] and [Weave][weave] for networking and [Brooklyn][brooklyn] for application blueprints. By [@brooklyncentral](https://github.com/brooklyncentral)
* [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker by [@ehazlett](https://github.com/ehazlett)
* [depcon](https://github.com/gondor/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes.  By [@gonodr][gondor]
* [deploy](https://github.com/ttiny/deploy) - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration. Declarative configuration and short commands for managing (syncing, building, running) of infrastructures of more than a few services. Able to deploy whole preconfigured server or system of services with a single line (without having to scroll the line).
* [Docket](https://github.com/netvarun/docket) - Custom docker registry that allows for lightning fast deploys through bittorrent by [@netvarun](https://github.com/netvarun/)
* [dockit](https://github.com/humblec/dockit) - Do docker actions and Deploy gluster containers! By [@humblec](https://github.com/humblec)
* [Longshoreman](https://github.com/longshoreman/longshoreman) - Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By [longshoreman](https://github.com/longshoreman)
* [rocker-compose](https://github.com/grammarly/rocker-compose) - Docker composition tool with idempotency features for deploying apps composed of multiple containers. By [@grammarly](grammarly)
* [Zodiac](https://github.com/CenturyLinkLabs/zodiac) - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs][CenturyLinkLabs]

## Developer Tools
* [ahab](https://github.com/instacart/ahab) - Docker event handling with Python by [@instacart](https://github.com/instacart)
* [bocker](https://github.com/p8952/bocker) (1) - Docker implemented in 100 lines of bash by [p8952](https://github.com/p8952)
* [construi](https://github.com/lstephen/construi) - Run your builds inside a Docker defined environment by [@lstephen](https://github.com/lstephen)
* [Docker Client for JVM](https://github.com/gesellix/docker-client) - A Docker remote api client library for the JVM, written in Groovy by [@gesellix][gesellix]
* [docker-it-scala](https://github.com/whisklabs/docker-it-scala) - Docker integration testing kit with Scala by [@whisklabs](https://github.com/whisklabs)
* [docker-maven-plugin (1)](https://github.com/spotify/docker-maven-plugin) - A Maven plugin for building and pushing Docker images by [@spotify](https://github.com/spotify/)
* [docker-maven-plugin (2)](https://github.com/fabric8io/docker-maven-plugin) - A Maven plugin for running and creating Docker images by [@fabric8io](https://github.com/fabric8io)
* [Docker-PowerShell](https://github.com/Microsoft/Docker-PowerShell) - PowerShell Module for Docker
* [Docker.DotNet](https://github.com/Microsoft/Docker.DotNet) - C#/.NET HTTP client for the Docker remote API by [@ahmetalpbalkan](ahmetalpbalkan)
* [dockerode](https://github.com/apocas/dockerode) - Docker Remote API node.js module by [@apocas](https://github.com/apocas)
* [go-dockerclient](https://github.com/fsouza/go-dockerclient/) - Go HTTP client for the Docker remote API by [@fsouza](https://github.com/fsouza/)
* [is-docker](https://github.com/sindresorhus/is-docker) - Check if the process is running inside a Docker container by [@sindresorhus][sindresorhus]
* [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) - A Docker remote api plugin for Gradle by [@gesellix][gesellix]
* [powerstrip](https://github.com/clusterhq/powerstrip) - A tool for prototyping Docker extensions by [@clusterhq](https://github.com/clusterhq)
* [sbt-docker](https://github.com/marcuslonnberg/sbt-docker) - Create Docker images directly from sbt by [@marcuslonnberg](https://github.com/marcuslonnberg)
* [sbt-docker-compose](https://github.com/Tapad/sbt-docker-compose) - Integrates Docker Compose functionality into sbt by [@kurtkopchik](https://github.com/kurtkopchik/)

## Development Environments
* [DevLab](https://github.com/TechnologyAdvice/DevLab) - Utility for running containerized development environments
* [Devstep](https://github.com/fgrehm/devstep) - Development environments powered by Docker and buildpacks by [@fgrehm][fgrehm]
* [Docker osx dev](https://github.com/brikis98/docker-osx-dev) - A productive development environment with Docker on OS X by [@brikis98](https://github.com/brikis98)
* [Docker-sync](http://docker-sync.io/) - Tool to improve performance when using Docker on Mac OS X and your volume has several folders and files like gulp projects for instance. By [@EugenMayer](https://github.com/EugenMayer)
* [Dray](http://dray.it/) - Dray is an engine for managing the execution of container-based workflows. Docker Workflow Engine - UNIX pipes for Docker by [@CenturyLinkLabs][CenturyLinkLabs]
* [Vagga](https://github.com/tailhook/vagga) - Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments by [@tailhook](https://github.com/tailhook/)

## Docker Compose file
* [Docker Compose Example](https://github.com/llitfkitfk/docker-compose-demo) by [@llitfkitfk](https://github.com/llitfkitfk)

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

## Garbage Collection
* [caduc](https://github.com/tjamet/caduc) - A docker garbage collector cleaning stuff you did not use recently
* [Docker Clean](https://github.com/zzrotdesign/docker-clean) - A script that cleans Docker containers, images and volumes by [@zzrotdesign](https://github.com/zzrotdesign)
* [docker-garby](https://github.com/konstruktoid/docker-garby) - Docker garbage collection script by [@konstruktoid](https://github.com/konstruktoid).
* [docker-gc](https://github.com/spotify/docker-gc) - A cron job that will delete old stopped containers and unused images by [@spotify](https://github.com/spotify)
* [sherdock](https://github.com/rancher/sherdock) - Automatic GC of images based on regexp by [@rancher][rancher]

## Hosting Images (registries)
Services to securely store your Docker images.
* [Amazon EC2 Container Registry](https://aws.amazon.com/ecr/) Amazon EC2 Container Registry (ECR) is a fully-managed Docker container registry that makes it easy for developers to store, manage, and deploy Docker container images.
* [Atomic Registry](http://www.projectatomic.io/registry/) - Red Hat Atomic Registry is an open source enterprise registry based on the Origin and Cockpit projects, enhancing the Docker registry library.
* [Azure Container Registry](https://azure.microsoft.com/de-de/services/container-registry/) Manage a Docker private registry as a first-class Azure resource
* [CargoOS](https://cargos.io/) A bare essential OS for running the Docker Engine on bare metal or Cloud.
* [Carina](https://getcarina.com) Bare-metal container hosting. Developed by Rackspace. Currently in Beta.
* [Compose Registry](https://www.composeregistry.com/) - Project from [@francescou]Francesco Uliana that stores docker-compose.yml files with stack examples like LAMP/LEMP stacks, Django stacks, Flask stacks and ELK stacks for instance.
* [Container Compliance](https://github.com/OpenSCAP/container-compliance) Open Source tool for assesing running containers and cold images for vulnerabilites and audits.
* [Cycle.io](https://cycle.io/) Bare-metal container hosting.
* [Docker Hub](https://hub.docker.com/) provided by Docker Inc.
* [Docker Registry v2][distribution] - The Docker toolset to pack, ship, store, and deliver content
* [GCE Container Registry](https://cloud.google.com/container-registry/) Fast, private Docker image storage on Google Cloud Platform
* [GitLab Container Registry](https://docs.gitlab.com/ce/container_registry/README.html) - Repositories focused on using it images in GitLab CI
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [Rescoyl](https://github.com/noteed/rescoyl) - Private Docker registry (free and open source) by [@noteed][noteed]
* [Sonatype Nexus](https://www.sonatype.com/nexus-repository-oss) - Repository with Universal Support, also for Docker images
* [TreeScale](https://treescale.com/) - Build and Distribute container based applications.
* [VMWare Harbor](http://vmware.github.io/harbor/) Project Harbor by VMWare is an enterprise-class registry server that stores and distributes Docker images. Harbor extends the open source Docker Distribution by adding the functionalities usually required by an enterprise, such as security, identity and management.

## Images
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

## Image Builder
* [bocker](https://github.com/icy/bocker) (2) - Write Dockerfile completely in Bash. Extensible and simple. --> Reusable by [@icy](https://github.com/icy)
* [container-factory](https://github.com/lsqio/container-factory) - Produces Docker images from tarballs of application source code by [@lsqio](https://github.com/lsqio)
* [dlayer](https://github.com/wercker/dlayer) - Stats collector for Docker layers by [@wercker](https://github.com/wercker)
* [docker-companion](https://github.com/mudler/docker-companion) - A command line tool written in Golang to squash and unpack docker images by [@mudler](https://github.com/mudler/)
* [docker-make](https://github.com/CtripCloud/docker-make) - Build, tag,and push a bunch of related docker images via a single command.
* [DockerSlim](https://github.com/docker-slim/docker-slim) shrinks fat Docker images creating the smallest possible images.
* [elsy](https://github.com/cisco/elsy) - An opinionated, multi-language, build tool based on Docker and Docker Compose
* [flyimg](http://flyimg.io/) - Docker image resizing, cropping, and compression on the fly.
* [habitus](https://github.com/cloud66/habitus) - A Build Flow Tool for Docker http://www.habitus.io by [@cloud66](https://github.com/cloud66)
* [MicroBadger][microbadger] - Analyze the contents of images and add metadata labels
* [packer](https://www.packer.io/docs/builders/docker.html) - Hashicorp tool to build machine images including docker image integrated with configuration management tools like chef, puppet, ansible
* [portainer](https://github.com/duedil-ltd/portainer) - Apache Mesos framework for building Docker images by [@duedil-ltd](https://github.com/duedil-ltd)
* [rocker](https://github.com/grammarly/rocker) - Extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc. by [grammarly](grammarly).
* [SkinnyWhale](https://github.com/djosephsen/skinnywhale) Skinnywhale helps you make smaller (as in megabytes) Docker containers.

## Linter / Validator
* [dockerfile_lint](https://github.com/projectatomic/dockerfile_lint) - A rule-based 'linter' for Dockerfiles by [@redhataccess](https://github.com/redhataccess)
* [Lorry](https://lorry.io/) - Lorry is a docker-compose.yml validator and composer by [@CenturyLinkLabs][CenturyLinkLabs]
* [Whale-linter](https://github.com/jeromepin/whale-linter) - A simple and small Dockerfile linter written in Python3+ without dependencies.

## Local Container Manager
* [Ansible - manage docker containers](http://docs.ansible.com/ansible/docker_module.html)
* [Azk](http://www.azk.io/) - Orchestrate development enviornments on your local machine by [@azukiapp](https://github.com/azukiapp)
* [Beluga](https://github.com/cortexmedia/Beluga) - CLI to deploy docker containers on a single server or low amount of servers. By [@cortextmedia](https://github.com/cortexmedia)
* [Boot2Docker](https://github.com/boot2docker/boot2docker) - Docker for OSX and Windows -- http://boot2docker.io/
* [crowdr](https://github.com/polonskiy/crowdr) - Tool for managing multiple Docker containers (`docker-compose` alternative) by [@polonskiy](https://github.com/polonskiy/)
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

## Monitoring & Logging
* [cAdvisor](https://github.com/google/cadvisor) - Analyzes resource usage and performance characteristics of running containers. Created by [@Google](https://github.com/google)
* [Docker-Fluentd][fluentd] - Docker container to Log Other Containers' Logs. One can aggregate the logs of Docker containers running on the same host using Fluentd by [@kiyoto][kiyoto]
* [Docker-mon](https://github.com/icecrime/docker-mon) - Console-based Docker monitoring by [@icecrime](https://github.com/icecrime)
* [Dockerana](https://github.com/dockerana/dockerana) - packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.
* [DoMonit](https://github.com/eon01/DoMonit) - A simple Docker Monitoring wrapper For Docker API
* [Dynatrace](https://www.dynatrace.com/technologies/cloud-and-microservices/docker-monitoring/) - Monitor containerized applications without installing agents or modifying your Run commands
* [Glances](https://nicolargo.github.io/glances/) - A cross-platform curses-based system monitoring tool written in Python by [@nicolargo](https://github.com/nicolargo)
* [Grafana Docker Dashboard Template](https://grafana.com/dashboards/179) - A template for your Docker, Grafana and Prometheus stack [@vegasbrianc][vegasbrianc]
* [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) - InfluxDB Time series DB in combination with Grafana and cAdvisor by [@vegasbrianc][vegasbrianc]
* [LogJam](https://github.com/gocardless/logjam) - Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash) by [@gocardless](https://github.com/gocardless)
* [Logsene for Docker][spm] Monitoring of Metrics, Events and Logs implemented in Node.js. Integrated [logagent-js](https://github.com/sematext/logagent-js) to detect and parse various log formats. [@sematext][sematext]
* [Logspout](https://github.com/gliderlabs/logspout) - Log routing for Docker container logs by [@gliderlabs][gliderlabs]
* [Out-of-the-box Host/Container Monitoring/Logging/Alerting Stack](https://github.com/uschtwill/docker_monitoring_logging_alerting) - Docker host and container monitoring, logging and alerting out of the box using cAdvisor, Prometheus, Grafana for monitoring, Elasticsearch, Kibana and Logstash for logging and elastalert and Alertmanager for alerting. Set up in 5 Minutes. Secure mode for production use with built-in [Automated Nginx Reverse Proxy (jwilder's)][nginxproxy].
* [Seagull](https://github.com/tobegit3hub/seagull) - Friendly Web UI to monitor docker daemon. by [@tobegit3hub](https://github.com/tobegit3hub)
* [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring) - Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
* [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.
* [Docker-Alertd](https://github.com/deltaskelta/docker-alertd) - Monitor and send alerts based on docker container resource usage/statistics

### Monitoring & Logging Services
* [AppDynamics](https://www.appdynamics.com/community/exchange/extension/docker-monitoring-extension/) - AppDynamics gives enterprises real-time insights into application performance, user performance, and business performance so they can move faster in an increasingly sophisticated, software-driven world.
* [Axibase Time-Series Database](http://axibase.com/products/axibase-time-series-database/writing-data/docker-cadvisor/) - Long-term retention of container statistics and built-in dashboards for Docker. Collected with native Google cAdvisor storage driver.
* [Collecting docker logs and stats with Splunk](http://blogs.splunk.com/2015/08/24/collecting-docker-logs-and-stats-with-splunk/)
* [CoScale](http://www.coscale.com/docker-monitoring) - Full stack monitoring for containerized applications and microservices. Powered by anomaly detection to find performance problems faster.
* [Datadog](https://www.datadoghq.com/) - Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
* [Meros](https://meros.io) - Analyzes containers resources, captures logs, remote web SSH terminal and powerful DevOps alerts.
* [New Relic](https://newrelic.com/partner/docker) - New Relics Docker Monitoring tool
* [Prometheus](https://prometheus.io/) - Open-source service monitoring system and time series database
* [Sysdig](http://www.sysdig.org/) - An open source troubleshooting tool that provides a rich set of real-time, system-level information. It has container-specific features and is very useful in Docker environments.
* [Site24x7](https://www.site24x7.com/docker-monitoring.html) - Docker MOnitoring for DevOps and IT is a SaaS Pay per Host model
* [SPM for Docker][spm] - Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [@sematext][sematext]

## Networking
* [Calico-Docker](https://www.projectcalico.org/getting-started/docker/) - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.
* [Flannel](https://github.com/coreos/flannel/) - Flannel is a virtual network that gives a subnet to each host for use with container runtimes.
* [Weave][weave] (The Docker network) - Weave creates a virtual network that connects Docker containers deployed across multiple hosts.

## PaaS
* [Atlantis](https://github.com/ooyala/atlantis) - Atlantis is an Open Source PaaS for HTTP applications built on Docker and written in Go
* [Convox Rack](https://github.com/convox/rack) - Convox Rack is open source PaaS built on top of expert infrastructure automation and devops best practices.
* [Dcw](https://github.com/pbertera/dcw) - Docker-compose SSH wrapper: a very poor man PaaS, exposing the docker-compose and custom-container commands defined in container labels.
* [Deis](https://github.com/deis/deis) - Your PaaS, your rules http://deis.io/
* [Dokku][dokku] - Docker powered mini-Heroku in around 100 lines of Bash by [@progrium][progrium]
* [Empire](https://github.com/remind101/empire) - A PaaS built on top of Amazon EC2 Container Service (ECS)
* [Flynn](https://github.com/flynn/flynn) - A next generation open source platform as a service https://flynn.io/
* [Nanobox](https://github.com/nanobox-io/nanobox) - A micro-PaaS (μPaaS) for creating consistent, isolated, development environments deployable anywhere https://nanobox.io.
* [OpenShift][openshift] - An open source PaaS built on [Kubernetes][kubernetes] and optimized for Dockerized app development and deployment by [Red Hat](https://www.redhat.com/)
* [Rancher][rancher] - Rancher is an open source project that provides a complete platform for operating Docker in production
* [Tsuru](https://github.com/tsuru/tsuru) - Tsuru is an extensible and open source Platform as a Service software https://tsuru.io/

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

## Reverse Proxy
* [docker-proxy](https://github.com/silarsis/docker-proxy) - Transparent proxy for docker containers, run in a docker container. By [@silarsis](https://github.com/silarsis)
* [fabio](https://github.com/eBay/fabio) - A fast, modern, zero-conf load balancing HTTP(S) router for deploying microservices managed by consul. By [@eBay](https://github.com/eBay)
* [h2o-proxy](https://github.com/zchee/h2o-proxy) - Automated H2O reverse proxy for Docker containers. An alternative to [jwilder/nginx-proxy][nginxproxy] by [@zchee](https://github.com/zchee)
* [Let's Encrypt Nginx-proxy Companion](https://github.com/JrCs/docker-letsencrypt-nginx-proxy-companion) - A lightweight companion container for the nginx-proxy. It allow the creation/renewal of Let's Encrypt certificates automatically. By [@JrCs](https://github.com/JrCs)
* [muguet](https://github.com/mattallty/muguet) - DNS Server & Reverse proxy for Docker environments. By [@mattallty](https://github.com/mattallty)
* [nginx-proxy][nginxproxy] - Automated nginx proxy for Docker containers using docker-gen by [@jwilder][jwilder]
* [Swarm Ingress Router](https://github.com/tpbowden/swarm-ingress-router) - Route DNS names to Swarm services based on labels.
* [Træfɪk](https://traefik.io/) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [@EmileVauge](https://github.com/emilevauge)

## Security
* [Clair](https://github.com/coreos/clair) - Clair is an open source project for the static analysis of vulnerabilities in appc and docker containers. By [@coreos][CoreOS]
* [docker-bench-security](https://github.com/docker/docker-bench-security) - script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker][docker]
* [notary](https://github.com/docker/notary) - a server and a client for running and interacting with trusted collections. By [@docker][docker]
* [Twistlock](https://twistlock.com/)  - Twistlock Security Suite detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.

## Serverless
* [AMP](https://github.com/appcelerator/amp) - The open source Container-as-a-Service & microservice platform for Docker
* [Docker-Lambda](https://github.com/lambci/docker-lambda) - Docker images and test runners that replicate the live AWS Lambda environment
* [FaaS](https://github.com/alexellis/faas) - Docker Serverless/Functions as a Service (on Docker Swarm)
* [Funker](https://github.com/bfirsh/funker-example-voting-app) - Functions as Docker containers
* [IronFunctions](https://github.com/iron-io/functions) - The serverless microservices platform FaaS (Funcitons as a Service) which uses Docker containers to run Any language or AWS Lambda functions

## Service Discovery
* [Docker Grand Ambassador](https://github.com/cpuguy83/docker-grand-ambassador) - This is a fully dynamic docker link ambassador. + [Article](https://docs.docker.com/engine/admin/ambassador_pattern_linking/) by [@cpuguy83][cpuguy83]
* [docker-consul](https://github.com/gliderlabs/docker-consul) by [@progrium][progrium]
* [etcd](https://github.com/coreos/etcd) - A highly-available key value store for shared configuration and service discovery by [@coreOS][coreos]
* [proxy](https://github.com/factorish/proxy) - lightweight nginx based load balancer self using service discovery provided by registrator. by [@factorish](https://github.com/factorish)
* [registrator](https://github.com/progrium/registrator) - Service registry bridge for Docker by [@progrium][progrium]

## Services for running containers
* [Amazon ECS](http://aws.amazon.com/ecs/) - A management service on EC2 that supports Docker containers.
* [Arukas](https://arukas.io/) - Heroku-inspired CaaS
* [Cloud 66](http://www.cloud66.com) - Full-stack hosted container management as a service
* [ContainerShip Cloud][containership] - Multi-Cloud Container Hosting Automation Platform.
* [DataMC](http://datamc.io/) - DataMc is a PaaS for Production ready and fully managed Data Platform
* [Docker Cloud](https://cloud.docker.com/) - Former Tutum
* [Dockhero](https://dockhero.io/) - Dockhero is a Heroku add-on which turns a Docker image into a microservice attached to the Heroku app. Currently in beta.
* [Giant Swarm](https://giantswarm.io/) - Simple microservice infrastructure. Deploy your containers in seconds.
* [Google Container Engine](https://cloud.google.com/container-engine/docs/) - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
* [Hyper_](https://hyper.sh/) - Secure container hosting service with "nano-containers" and per-second billing.
* [IBM Bluemix](https://console.ng.bluemix.net/) - Run Docker containers in a hosted cloud environment on IBM Bluemix.
* [OpenShift Dedicated](https://www.openshift.com/dedicated/index.html) - A hosted [OpenShift][openshift] cluster for running your Docker containers managed by Red Hat.
* [Orchard](https://www.orchardup.com/) (part of Docker Inc) - Get a Docker host in the cloud, instantly.
* [Sloppy.io](https://sloppy.io/) - all-in-one solution for container deployment and hosting – made and hosted in Germany
* [Triton](https://www.joyent.com/) - Elastic container-native infrastructure by Joyent.

## Terminal User Interface
* [ctop (1)](https://github.com/yadutaf/ctop) - A command line / text based Linux Containers monitoring tool that works just like you expect (Python) by [@yadutaf](https://github.com/yadutaf)
* [ctop (2)](https://github.com/bcicen/ctop) - Top-like interface for container metrics (Golang) by [@bcicen](https://github.com/bcicen/)
* [dockercraft](https://github.com/docker/dockercraft) - Docker + Minecraft = Dockercraft by [@docker][docker]
* [dockersql](https://github.com/crosbymichael/dockersql) - A command line interface to query Docker using SQL by [@crosbymichael][crosbymichael]
* [dockly](https://github.com/lirantal/dockly) - An interactive shell UI for managing Docker containers by [@lirantal](https://github.com/lirantal)
* [dry](https://github.com/moncho/dry) - An interactive CLI for Docker containers by [@moncho](https://github.com/moncho)
* [sen](https://github.com/TomasTomecek/sen) - Terminal user interface for docker engine, by [@TomasTomecek](https://github.com/TomasTomecek)
* [wharfee](https://github.com/j-bennet/wharfee) - Autocompletion and syntax highlighting for Docker commands. by [@j-bennet](https://github.com/j-bennet)
* [tsaotun](https://github.com/qazbnm456/tsaotun) - Python based Assistance for Docker by [@qazbnm456](https://github.com/qazbnm456)

## Testing
* [dgoss](https://github.com/aelsabbahy/goss/tree/master/extras/dgoss) - A fast YAML based tool for validating docker containers.
* [elgalu/docker-selenium](https://github.com/elgalu/docker-selenium) - Selenium in Docker with Chrome & Firefox plus video recording support.
* [Pumba](https://github.com/gaia-adm/pumba) - Chaos testing tool for Docker. Can be deployed on Kubernets and CoreOS clusters.
* [SeleniumHQ/docker-selenium](https://github.com/SeleniumHQ/docker-selenium) - Docker images for Selenium Standalone Server, Hub, and Node configurations with Chrome and Firefox.
* [zalenium](https://github.com/zalando/zalenium) - A Selenium Grid extension to scale up and down your local grid dynamically with docker containers.

## Utilities
* [athena](https://github.com/athena-oss/athena) - An automation platform with a plugin architecture that allows you to easily create and share services.
* [Chaperone](https://github.com/garywiz/chaperone) - A single PID1 process designed for docker containers. Does user management, log management, startup, zombie reaping, all in one small package. by [@garywiz](https://github.com/garywiz)
* [codelift](https://codelift.io/) - CodeLift is an automated Docker image build utility for 'dockerizing' services by [@BoozAllen](https://twitter.com/BoozAllen)
* [Codenvy](https://codenvy.io) - One-click Docker environments and cloud workspace for development teams
* [Compose Registry](https://www.composeregistry.com) - A very handy search engine for Compose Files
* [Composerize](https://github.com/magicmark/composerize) - Convert docker run commands into docker-compose files
* [dexec](https://github.com/docker-exec/dexec) - Command line interface written in Go for running code with Docker Exec images.
* [dext-docker-registry-plugin](https://github.com/vutran/dext-docker-registry-plugin) - Search the Docker Registry with the Dext smart launcher.
* [Docker meets the IDE](http://domeide.github.io/) - Integrating your favorite containers in the editor of your choice by [domeide](https://github.com/domeide)
* [Docker Volume Clone Utility](https://github.com/gdiepen/docker-convenience-scripts) - A Docker Utility to Clone Volumes [@gdiepen](https://twitter.com/gdiepen)
* [docker-compose-search](https://github.com/francescou/docker-compose-search) - A search engine for Docker Compose application stacks by [@francescou](https://github.com/francescou/)
* [docker-do](https://github.com/benzaita/docker-do) - hassle-free docker run, like `env` but for docker by [@benzaita](https://github.com/benzaita)
* [docker-gen](https://github.com/jwilder/docker-gen) - Generate files from docker container meta-data by [@jwilder][jwilder]
* [docker-ls](https://github.com/mayflower/docker-ls) - CLI tools for browsing and manipulating docker registries by [@mayflower](https://github.com/mayflower)
* [docker-replay](https://github.com/bcicen/docker-replay) - Generate `docker run`command and options from running containers. By [bcicen](https://github.com/bcicen)
* [docker-volumes](https://github.com/cpuguy83/docker-volumes) - Docker Volume Manager by [@cpuguy83][cpuguy83]
* [dockerize](https://github.com/jwilder/dockerize) - Utility to simplify running applications in docker containers by [@jwilder][jwilder]
* [Dockly](https://github.com/swipely/dockly) - Dockly is a gem made to ease the pain of packaging an application in Docker by [@swipely](https://github.com/swipely/)
* [dockramp](https://github.com/jlhawn/dockramp) - Proof of Concept: A Client Driven Docker Image Builder by [@jlhawn](https://github.com/jlhawn)
* [draw-compose](https://github.com/Alexis-benoist/draw-compose) - Utility to draw a schema of a docker compose by [@Alexis-benoist](https://github.com/Alexis-benoist)
* [Dropdock](http://dropdock.io/) - A framework designed for Drupal to build fast, isolated development environments using Docker.
* [Dupper](https://github.com/athakwani/dupper) - Use git repository as Object Oriented Container. Setup Cloud Development environments with few commands for any git repository by [@athakwani](https://github.com/athakwani/)
* [DVM](https://github.com/howtowhale/dvm) - Docker version manager by [@howtowhale](https://github.com/howtowhale)
* [Eclipse Che](http://www.eclipse.org/che) - Developer workspace server with Docker runtimes, cloud IDE, next-generation Eclipse IDE
* [forward2docker](https://github.com/bsideup/forward2docker) - Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM by [@bsideup](https://github.com/bsideup)
* [GoSu](https://github.com/tianon/gosu) - Run this specific application as this specific user and get out of the pipeline (entrypoint script tool) by [@tianon](https://github.com/tianon)
* [ns-enter](https://github.com/jpetazzo/nsenter) - no more ssh, enter name spaces of container by [@jpetazzo][jpetazzo]
* [OctoLinker](https://github.com/OctoLinker/browser-extension) - A browser extension for GitHub that makes the image name in a `Dockerfile` clickable and redirect you to the related Docker Hub page.
* [percheron][percheron] - Organise your Docker containers with muscle and intelligence by [@ashmckenzie](https://github.com/ashmckenzie)
* [Powerline-Docker](https://github.com/adrianmo/powerline-docker) - A Powerline segment for showing the status of Docker containers by [@adrianmo](https://github.com/adrianmo)
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) - in case of proxy problem by [@jpetazzo][jpetazzo]
* [TrivialRC](https://github.com/vorakl/TrivialRC) - A minimalistic Runtime Configuration system and process manager for containers [@vorakl](https://github.com/vorakl)

## Volume management and plugins
* [Blockbridge](https://github.com/blockbridge/blockbridge-docker-volume) - The Blockbridge plugin is a volume plugin that provides access to an extensible set of container-based persistent storage options. It supports single and multi-host Docker environments with features that include tenant isolation, automated provisioning, encryption, secure deletion, snapshots and QoS. By [@blockbridge][blockbridge]
* [Convoy](https://github.com/rancher/convoy) - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher][rancher]
* [Docker Unison](https://github.com/leighmcculloch/docker-unison) A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By [@leighmcculloch](https://github.com/leighmcculloch)
* [Netshare](https://github.com/gondor/docker-volume-netshare) A Docker volume plugin written in Go that supports mounting NFS, AWS EFS & CIFS volumes within a container. By [@gondor][gondor]
* [Docker Machine NFS](https://github.com/adlogix/docker-machine-nfs) Activates NFS for an existing boot2docker box created through Docker Machine on OS X.
* [REX-Ray](https://github.com/codedellemc/rexray) Vendor agnostic storage orchestration engine to provide persistent storage for Docker containers as well as Mesos frameworks and tasks.
* [Local Persist](https://github.com/CWSpear/local-persist) Specify a mountpoint for your local volumes (created via `docker volume create`) so that files will always persist and so you can mount to different directories in different containers.
* [Minio](https://github.com/jelastic-jps/minio) - S3 compatible object storage server in Docker containers

## Web Interface
* [Docker Registry Browser](https://github.com/klausmeyer/docker-registry-browser) - Web Interface for the Docker Registry HTTP API v2 by [@klausmeyer](https://github.com/klausmeyer)
* [Docker Registry UI](https://github.com/atc-/docker-registry-ui) - A web UI for easy private/local Docker Registry integration by [@atc-](https://github.com/atc-)
* [docker-registry-web](https://github.com/mkuchin/docker-registry-web) - Web UI, authentication service and event recorder for private docker registry v2 by [@mkuchin](https://github.com/mkuchin)
* [docker-swarm-visualizer](https://github.com/manomarks/docker-swarm-visualizer) - Visualizes Docker services on a Docker Swarm (for running demos).
* [dockering-on-rails](https://github.com/Electrofenster/dockerding-on-rails) - Simple Web-Interface for Docker with a lot of features by [@Electrofenster](https://github.com/Electrofenster/)
* [DockerUI](https://github.com/kevana/ui-for-docker) - DockerUI is a web interface to interact with the Remote API by [@crosbymichael][crosbymichael]
* [Portus](https://github.com/SUSE/Portus) - Authorization service and frontend for Docker registry (v2) by [@SUSE](https://github.com/SUSE)
* [Rapid Dashboard](https://github.com/ozlerhakan/rapid) - A simple query dashboard to use Docker Remote API by [@ozlerhakan](https://github.com/ozlerhakan/)

# Useful Resources
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
* [Docker Cheat Sheet](http://docker.jens-piegsa.com) by [@JensPiegsa][JensPiegsa] *(updated for Docker 1.13)*
* [Docker Cheat Sheet][docker-cheat-sheet] by [@wsargent][wsargent] __MUST SEE__
* [Docker Containers on the desktop][jessblog] by [@jfrazelle][jfrazelle] The **funniest way** to learn
* [Docker Ecosystem](http://comp.photo777.org/wp-content/uploads/2015/09/Docker-ecosystem-8.5.1.pdf) (PDF) __MUST SEE__    find it on [blog](http://comp.photo777.org/docker-ecosystem/) by Bryzgalov Peter.
* [Docker Ecosystem](https://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) __MUST SEE__
* [Docker Kubernetes Lab Handbook](https://github.com/xiaopeng163/docker-k8s-lab)
* [Docker Printable Refcard][docker-quick-ref] by [@dimonomid][dimonomid]
* [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](http://www.rightscale.com/blog/cloud-management-best-practices/docker-vs-vms-combining-both-cloud-portability-nirvana)
* [Docker Weekly](https://blog.docker.com/docker-weekly-archives/) Huge resource
* [Project Web Dev][projwebdev] : (Article series) How to create your own website based on Docker
* __[Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html)__ High quality articles about docker! __MUST SEE__

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
* [Shippable](http://blog.shippable.com/)
* [WebOps weekly](http://webopsweekly.com/)
* [DevOpsLinks](http://devopslinks.com)

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

## Raspberry Pi & ARM
* [Docker Pirates ARMed with explosive stuff](http://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
* [Get Docker up and running on the RaspberryPi in three steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-three-steps)
* [git push docker containers to linux devices](https://resin.io/) Modern DevOps for IoT, leveraging git and Docker.
* [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-(ARMv7)-devices)

# Videos

## Main Account
* [Docker Youtube Account](https://www.youtube.com/user/dockerrun)
* [CenturyLink Labs Docker Interviews](https://www.youtube.com/playlist?list=PL_q4Fk7SVBCIjyuCBFBItXnzGI3qBa2L1)
* [Container Camp](https://www.youtube.com/channel/UCvksXSnLqIVM_uFB7xyrsSg/videos) Conference about *containers*!!! [@containercamp](https://twitter.com/containercamp)
* [Quoi d'neuf Docker](https://www.youtube.com/channel/UCOAhkxpryr_BKybt9wIw-NQ/videos) **FRENCH** chronique vidéo sur Youtube proposant de courtes vidéos (maximum 15 minutes) sur la thématique "Docker et son écosystème" [Site Web](http://www.quoidneufdocker.xyz/)

## Useful videos
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
* [Scalable Microservices with Kubernetes](https://www.udacity.com/course/scalable-microservices-with-kubernetes--ud615) Free Udacity course
* [State of containers: a debate with CoreOS, VMware and Google](https://www.youtube.com/watch?v=IiITP3yIRd8) (27:38)
* [SysAdminCasts: Introduction to Docker](https://sysadmincasts.com/episodes/31-introduction-to-docker) (15:49)

# Interesting Twitter Accounts
* [CenturyLink Labs](https://twitter.com/CenturyLinkLabs)
* [Docker Captains Twitter List](https://twitter.com/EltonStoneman/lists/docker-captains)
* [Docker](https://twitter.com/docker)
* [Flux7Labs](https://twitter.com/Flux7Labs)
* [OpenShift by Red Hat](https://twitter.com/openshift)
* [Project Atomic](https://twitter.com/ProjectAtomic)
* [The New Stack](https://twitter.com/thenewstack)

# Communities and Meetups

## Brazilian
* [Docker BR on Telegram](https://telegram.me/dockerbr)
* [Docker BR on Slack](http://docker-br.herokuapp.com) - Auto invite url

## Chinese
* [DockerOne](http://dockone.io/) Docker Community (in Chinese) by [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)

## English
* [Docker Community](https://www.docker.com/docker-community)
* [Docker On Line Meetup](https://www.meetup.com/Docker-Online-Meetup/)
* [Docker Regional Meetup List](https://www.docker.com/community/meetup-groups)

## Russian
+ [Docker Russian-speaking Community](https://t.me/docker_ru)

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
[grammarly]: https://github.com/grammarly
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
