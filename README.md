# Awesome Docker

> A curated list of Docker resources

Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome](https://github.com/sindresorhus/awesome) and improved by these **[amazing contributors](https://github.com/veggiemonk/awesome-docker/graphs/contributors)**.

This list is not exhaustive, contribute to make it so: add your docker resource by forking / sending pull requests.

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)**

# What is Docker ?

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.


# MENU

* [Useful Articles](#useful-articles)
 * [Main Resources](#main-resources)
 * [General Articles](#general-articles)
 * [Deep Dive](#deep-dive)
 * [Metal](#metal)
 * [Multi-Server](#multi-server)
 * [Good Tips](#good-tips)
 * [Continuous Integration](#continuous-integration)
 * [Optimizing Images](#optimizing-images)
 * [Service Discovery](#service-discovery)
 * [Security](#security)
 * [Performances](#performances)
 * [Other](#other)
* [Books](#books)
* [Tools](#tools)
  * [Continous Integration / Continous Delivery](#continous-integration--continous-delivery) 
  * [Hosting for repositories (registries)](#hosting-for-repositories-registries)
  * [Hosting for containers](#hosting-for-containers-paying-services-)
  * [Dev Tools](#dev-tools)
  * [Useful Images](#useful-images)
  * [Docker Files](#docker-files)
  * [Storing Images](#storing-images)
  * [Web Interface](#web-interface)
  * [Monitoring](#monitoring)
  * [Running Containers](#running-containers--orchestration)
  * [Networking](#networking)
  * [Logging](#logging)
  * [Service Discovery](#service-discovery-1)
* [Slides](#slides)
* [Videos](#videos)
  * [Main Account](#main-account)
  * [Useful videos](#useful-videos)
* [Interesting Twitter Accounts](#interesting-twitter-accounts)
  
----

# Useful Articles

## Main Resources

* [Docker Weekly](http://blog.docker.com/docker-weekly-archives/) Huge resource
* [Century Links Labs](http://www.centurylinklabs.com/category/docker/)
* [Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html) Very complete
* [Docker Ecosystem](http://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) __MUST SEE__
* [Blog](http://jpetazzo.github.io/) of [@jpetazzo](https://github.com/jpetazzo)
* [Blog](http://progrium.com/blog/) of [@progrium](https://github.com/progrium)
* [Blog](http://jasonwilder.com/) of [@jwilder](https://github.com/jwilder)
* [Blog](http://crosbymichael.com/) of [@crosbymichael](https://github.com/crosbymichael)
* [Digital Ocean Community](https://www.digitalocean.com/community/search?primary_filter=tutorials&query=docker)
* [Container42](http://container42.com/)
* [Contrainer solutions](http://container-solutions.com/blog/)

## General Articles
* [Getting Started with Docker](https://serversforhackers.com/articles/2014/03/20/getting-started-with-docker/) by [@fideloper](https://github.com/fideloper) -- [Servers For Hackers](https://serversforhackers.com/editions/) is valuable resource. At some point, every programmer finds themselves needing to know their way around a server.
* [How to Use Docker on OS X: The Missing Guide](http://viget.com/extend/how-to-use-docker-on-os-x-the-missing-guide)
* [Docker for (Java) Developers](http://ro14nd.de/Docker-for-Developers/)
* [Deploying NGINX with Docker](http://nginx.com/blog/deploying-nginx-nginx-plus-docker/)
* [Eight Docker Development Patterns](http://www.hokstad.com/docker/patterns)

## Deep Dive
* [Creating containers - Part 1](http://crosbymichael.com/creating-containers-part-1.html) This is part one of a series of blog posts detailing how docker creates containers. By [@crosbymichael](https://github.com/crosbymichael)
* [Data-only container madness](http://container42.com/2014/11/18/data-only-container-madness/)

## Metal
* [How to use Docker on Full Metal](http://blog.bigstep.com/big-data-performance/use-docker-full-metal-cloud/)

## Multi-Server
* [Using Fig and Flocker to build, test, deploy and migrate multi-server Dockerized apps](https://clusterhq.com/blog/fig-flocker-multi-server-docker-apps/)
* [blimp](https://github.com/tubesandlube/blimp) Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more. By [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin)

## Good Tips
* [24 random docker tips](http://csaba.palfi.me/random-docker-tips/) by [@csabapalfi](https://github.com/csabapalfi)
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm](https://github.com/fgrehm)
* [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/) by [@jwilder](https://github.com/jwilder)
* [Using NSEnter with Boot2Docker](http://ro14nd.de/NSEnter-with-Boot2Docker/)
* [A Simple Way to Dockerize Applications](http://jasonwilder.com/blog/2014/10/13/a-simple-way-to-dockerize-applications) by [@jwilder](https://github.com/jwilder)
* [Building good docker images](http://jonathan.bergknoff.com/journal/building-good-docker-images) by [@jbergknoff](https://github.com/jbergknoff)
* [10 Things Not To Forget Before Deploying Docker In Production](http://www.slideshare.net/rightscale/docker-meetup-40826948)
* [Make Your Docker Workflow Awesome With Fig.sh](http://www.syncano.com/docker-workflow-fig-sh/) Fig is a python application that helps you run groups of docker containers.

## Continuous Integration
* [Docker and Phoenix: How to Make Your Continuous Integration More Awesome](http://ariya.ofilabs.com/2014/12/docker-and-phoenix-how-to-make-your-continuous-integration-more-awesome.html)

## Optimizing Images
* [Create the smallest possible Docker container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
* [Creating a Docker image from your code](http://blog.tutum.co/2014/04/10/creating-a-docker-image-from-your-code/)
* [Optimizing Docker Images](http://www.centurylinklabs.com/optimizing-docker-images/?hvid=1OW0br)
* [How to Optimize Your Dockerfile](http://blog.tutum.co/2014/10/22/how-to-optimize-your-dockerfile/) by [@tutumcloud](https://github.com/tutumcloud)
* [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07) by [@kelseyhightower](https://github.com/kelseyhightower)
* [Squashing Docker Images](http://jasonwilder.com/blog/2014/08/19/squashing-docker-images/) by [@jwilder](https://github.com/jwilder)
* [Dockerfile Golf (or optimizing the Docker build process)](http://www.davidmkerr.com/2014/08/dockerfile-golf-or-optimizing-docker.html)

## Service Discovery
* [@progrium](https://github.com/progrium) Service Discovery articles series:
 * [Consul Service Discovery with Docker](http://progrium.com/blog/2014/08/20/consul-service-discovery-with-docker/)
 * [Understanding Modern Service Discovery with Docker](http://progrium.com/blog/2014/07/29/understanding-modern-service-discovery-with-docker/)
 * [Automatic Docker Service Announcement with Registrator](http://progrium.com/blog/2014/09/10/automatic-docker-service-announcement-with-registrator/) 

## Security
* [Docker and SELinux](http://www.projectatomic.io/docs/docker-and-selinux/)
* [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
* [Docker Secure Deployment Guidelines](https://github.com/GDSSecurity/Docker-Secure-Deployment-Guidelines)

## Performances
* [Performance Analysis of Docker on Red Hat Enterprise Linux 7](http://developerblog.redhat.com/2014/08/19/performance-analysis-docker-red-hat-enterprise-linux-7/)
* [Distrubuted JMeter testing using Docker](http://srivaths.blogspot.fr/2014/08/distrubuted-jmeter-testing-using-docker.html?m=1)
* [nsinit: per-container resource monitoring of Docker containers on RHEL/Fedora](http://www.breakage.org/2014/09/03/nsinit-per-container-resource-monitoring-of-docker-containers-on-rhelfedora/)

## Other
* Presentation: Docker and JBoss - the perfect combination
 * [Vidéo](http://www.youtube.com/watch?v=4uQ6gR_xZhE)
 * [Slides](https://goldmann.pl/presentations/2014-vjbug-docker)
 * [Code source](https://github.com/goldmann/goldmann.pl/tree/master/.presentations/2014-vjbug-docker/demos)


# Books

* [Docker Book](http://dockerbook.com/) by James Turnbul ([@kartar](https://twitter.com/kartar))
* [Docker Cookbook](http://shop.oreilly.com/product/0636920036791.do) Solutions and examples for building distributed applications By Sébastien Goasguen 

# Tools 

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Cheat Sheet](https://github.com/wsargent/docker-cheat-sheet)  by [@wsargent](https://github.com/wsargent)

## Continous Integration / Continous Delivery
* [Drone](https://github.com/drone/drone) - https://drone.io/
* [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.

## Hosting for repositories (registries)
* [Docker Hub](https://hub.docker.com/) (provided by Docker Inc.)
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [Reesd](https://reesd.com/) -  Private Docker repositories and redundant storage service by [@noteed](https://github.com/noteed)

## Hosting for containers
* [Orchard](https://www.orchardup.com/) (part of Docker Inc) - Get a Docker host in the cloud, instantly.
* [StackDock](https://stackdock.com/) - Docker hosting on blazing fast dedicated infrastructure 
* [Tutum](https://www.tutum.co/) - Simple hosting for your Docker containers.

## Dev Tools

* [Fig](https://github.com/docker/fig) (Fast, isolated development environments using Docker) -- http://www.fig.sh
* [Shutit](http://ianmiell.github.io/shutit/) (a tool for building and maintaining complex Docker deployments) by [@ianmiell](https://github.com/ianmiell)
* [FuGu](https://github.com/mattes/fugu) (a docker run wrapper without orchestration) by [@mattes](https://github.com/mattes)
* [GoSu](https://github.com/tianon/gosu) ("run this specific application as this specific user and get out of the pipeline" -- entrypoint script tool) by [@tianon](https://github.com/tianon)
* [ns-enter](https://github.com/jpetazzo/nsenter) (no more ssh, enter name spaces of container) by [@jpetazzo](https://github.com/jpetazzo)
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) (in case of proxy problem) by [@jpetazzo](https://github.com/jpetazzo/)
* [docker-gen](https://github.com/jwilder/docker-gen) (Generate files from docker container meta-data) by [@jwilder](https://github.com/jwilder)
* [dockerize](https://github.com/jwilder/dockerize) (Utility to simplify running applications in docker containers) by [@jwilder](https://github.com/jwilder)
* [registrator](https://github.com/progrium/registrator) (Service registry bridge for Docker) by [@progrium](https://github.com/progrium)
* [Boot2Docker](https://github.com/boot2docker/boot2docker) (docker for OSX and Windows) -- http://boot2docker.io/
* [OctoHost](http://www.octohost.io/) (Simple web focused Docker based mini-PaaS server. git push to deploy your websites as needed) by [@octohost](https://github.com/octohost)
* [Dockly](https://github.com/swipely/dockly): Dockly is a gem made to ease the pain of packaging an application in Docker.
* [docker-volumes](https://github.com/cpuguy83/docker-volumes) (Docker Volume Manager) by [@cpuguy83](https://github.com/cpuguy83)
* [dockerfile_lint](https://github.com/redhataccess/dockerfile_lint) (A rule-based 'linter' for Dockerfiles) by [@redhataccess](https://github.com/redhataccess)
 
## Useful Images
* [Base Image](https://github.com/phusion/baseimage-docker) by [@phusion](https://github.com/phusion/)
* [Busybox](https://github.com/jpetazzo/docker-busybox) (with either `buildroot` or Ubuntu's `busybox-static`) by [@jpetazzo](https://github.com/jpetazzo)
* [Busybox](https://github.com/progrium/busybox) (with `opkg`) by [@progrium](https://github.com/progrium)
* [DockerFile Example](https://github.com/komljen/dockerfile-examples) by [@komljen](https://github.com/komljen)
* [OpenWRT](http://www.zoobab.com/docker-openwrt-image) by [@zoobab](https://github.com/zoobab)
* [Phusion Docker Hub Account](https://hub.docker.com/u/phusion/)

## Docker Files
* [Dockerfile Project](http://dockerfile.github.io/) : Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
* [Collection of Dockerfiles](https://github.com/crosbymichael/Dockerfiles) by [@crosbymichael](https://github.com/crosbymichael)

## Storing Images
* [Docker Registry](https://github.com/docker/docker-registry) (Registry server for Docker (hosting/delivering of repositories and images))

## Web Interface
* [Docker Registry Web](https://github.com/atc-/docker-registry-web) (A web UI for easy private/local Docker Registry integration) by [@atc-](https://github.com/atc-)
* [DockerUI](https://github.com/crosbymichael/dockerui) (DockerUI is a web interface to interact with the Remote API.) by [@crosbymichael](https://github.com/crosbymichael)
* [DockerBoard](https://github.com/dockerboard/dockerboard) (Simple dashboards, visualizations, managements for your dockers.) by [@dockerboard](https://github.com/dockerboard)
 
## Monitoring
* [Seagull](https://github.com/tobegit3hub/seagull) (Friendly Web UI to monitor docker daemon.)
* [Dockerana](https://github.com/dockerana/dockerana) (packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.)

## Running Containers / Orchestration
* [Kubernetes](http://kubernetes.io) Open source orchestration system for Docker containers by Google  -- http://kubernetes.io
* [Shipyard](https://github.com/shipyard/shipyard) (Composable Docker Management) -- http://shipyard-project.com/
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui/wiki) (Docker Management for Humans) -- http://panamax.io/
* [Flynn](https://github.com/flynn/flynn) (A next generation open source platform as a service (PaaS)) -- https://flynn.io/
* [Deis](https://github.com/deis/deis) (Your PaaS, your rules) -- http://deis.io/
* [Gaudi](https://github.com/marmelab/gaudi) (Gaudi allows to share multi-component applications, based on Docker, Go, and YAM) -- http://gaudi.io/
* [CoreOS](https://github.com/coreos) (Linux for Massive Server Deployments) -- https://coreos.com/
* [Rancher](https://github.com/rancherio/rancher) (Portable AWS-style infrastructure service for Docker) -- http://www.rancher.io/
* [Dokku](https://github.com/progrium/dokku) (Docker powered mini-Heroku in around 100 lines of Bash) by [@progrium](https://github.com/progrium)
  * [dokku-alt](https://github.com/dokku-alt/dokku-alt) (Dokku fork with Dockerfile support, database plugins, ACL and more)
* [cAdvisor](https://github.com/google/cadvisor) (Analyzes resource usage and performance characteristics of running containers)
* [Ansible - manage docker containers](http://docs.ansible.com/docker_module.html)
* [Launch Docker container on Mesos](https://mesosphere.com/learn/launch-docker-container-on-mesosphere/) (Docker plus Mesosphere provides an easy way to automate and scale deployment of containers in a production environment.)
* [Marathon](https://mesosphere.github.io/marathon/docs/) (Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on".)
* [Serf](https://github.com/hashicorp/serf)Service orchestration and management tool. by [@hashicorp](https://github.com/hashicorp)
* [Flocker](https://github.com/ClusterHQ/flocker) (Flocker is a data volume manager and multi-host Docker cluster management tool) by [@ClusterHQ](https://github.com/ClusterHQ)
* [Decking](http://decking.io/): Decking aims to simplify the creation, organsation and running of clusters of Docker containers in a way which is familiar to developers.
* [Maestro](https://github.com/toscanini/maestro): Maestro provides the ability to easily launch, orchestrate and manage mulitple Docker containers as single unit.

## Networking
* [Weave][weave] (The Docker network) -- Weave creates a virtual network that connects Docker containers deployed across multiple hosts.
 
## Logging
* [LogJam](https://github.com/gocardless/logjam) (Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collecton server (such as logstash).)

## Service Discovery
* [docker-consul](https://github.com/progrium/docker-consul) by [@progrium](https://github.com/progrium)
* [etcd](https://github.com/coreos/etcd): A highly-available key value store for shared configuration and service discovery
* [Docker Grand Ambassador](https://github.com/cpuguy83/docker-grand-ambassador) This is a fully dynamic docker link ambassador. + [Article](http://docs.docker.com/articles/ambassador_pattern_linking/)
* [confd](http://www.confd.io/): Manage local application configuration files using templates and data from etcd or consul.

## Deployment and Infrastructure
* [Centurion](https://github.com/newrelic/centurion): Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings.
* [Longshoreman](https://github.com/longshoreman/longshoreman): Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool.
* [Clocker](https://github.com/brooklyncentral/clocker): Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts. It leverages [Weave][weave] for networking and [Brooklyn][boorklyn] for application blueprints.

# Slides
* [Docker Slideshare Account](http://www.slideshare.net/dotCloud)
* [Docker Security](http://www.slideshare.net/jpetazzo) with [@jpetazzo](https://github.com/jpetazzo)

# Videos

## Main Account
* [Docker Youtube Account](http://www.youtube.com/user/dockerrun)
* [CenturyLink Labs Docker Interviews](http://www.youtube.com/playlist?list=PL_q4Fk7SVBCIjyuCBFBItXnzGI3qBa2L1)
* [YLD Event](https://www.youtube.com/channel/UCvksXSnLqIVM_uFB7xyrsSg/videos) Conference about *containers*!!! [@YLDio](https://twitter.com/YLDio)

## Useful videos
* [SysAdminCasts: Introduction to Docker](https://sysadmincasts.com/episodes/31-introduction-to-docker) (15:49)
* [Orchestrating Docker containers in production using Fig](https://www.youtube.com/watch?v=SEtRg8siQWw) (7:11)
* [Development Environments with Fig](http://youtu.be/QpSFOHvFyMc) by Aanand Prasad (17:58)
* [Docker in Production](http://youtu.be/Glk5d5WP6MI) by [@jpetazzo](https://github.com/jpetazzo) (36:05)
* [Docker: How to Use Your Own Private Registry](https://www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
* [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](https://www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
* [Performance Analysis of Docker - Jeremy Eder](https://www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
* [Docker and SELinux by Daniel Walsh from Red Hat ](https://www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
* [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](https://www.youtube.com/watch?v=GaHzdqFithc) (42:04)
* [Ansible and Docker HP](https://www.youtube.com/watch?v=oZ45v8AeE7k) (32:38)
* [Run Any App on Mesos on Any Infrastructure Using Docker](https://www.youtube.com/watch?v=u5jd9YT9EsY) (17:44)
* [Extending Docker with Plugins](http://vimeo.com/110835013) (15:21)


# Interesting Twitter Accounts
* [Docker](https://twitter.com/docker)
* [Century Link Labs](https://twitter.com/CenturyLinkLabs)
* [Flux7Labs](https://twitter.com/Flux7Labs)
* [TutumCloud](https://twitter.com/tutumcloud)
* [Project Atomic](https://twitter.com/ProjectAtomic)
* [Openshift By Red Hat](https://twitter.com/openshift)
* [YLD](https://twitter.com/YLDio)

## People
* [Solomon Hykes](https://twitter.com/solomonstre) Founder of Docker
* [Gabriel Monroy](https://twitter.com/gabrtv) Creator of Deis
* [Jérôme Petazzoni](https://twitter.com/jpetazzo) Docker Developer
* [Michael Crosby](https://twitter.com/crosbymichael) Docker Developer
* [James Turnbull](https://twitter.com/kartar) Author of Docker Book
* [Jeff Lindsay](https://twitter.com/progrium) Design-minded software architect

[weave]: https://github.com/zettio/weave
[brooklyn]: https://brooklyn.incubator.apache.org/
