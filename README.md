# Awesome Docker

> A curated list of Docker resources

Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome](https://github.com/sindresorhus/awesome)

Share some love: add your docker resource by forking / sending pull requests.

This list is not exhaustive (nor is it meant to be) but can probably help you get up and running faster. 

* [Useful Articles](#useful-articles)
* [Books](#books)
* [Dev](#dev)
  * [Continous Integration / Continous Delivery](#continous-integration-continous-delivery) 
  * [Dev Tools](#dev-tools)
  * [Images](#images)
  * [Managing Images](#managing-images)
* [Slides](#slides)
* [Videos](#videos)
  * [Main Account](#main-account)
  * [Useful videos](#useful-videos)
* [(Interesting) Twitter Accounts](#interesting-twitter-accounts)
  
----

## Useful Articles

* [Docker Weekly](http://blog.docker.com/docker-weekly-archives/)
* [Century Links Labs](http://www.centurylinklabs.com/category/docker-posts/)
* [Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html)
* [Docker Ecosystem](http://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) <--- MUST SEE
* [Blog](http://jpetazzo.github.io/) of [@jpetazzo](https://github.com/jpetazzo)
* [Getting Started with Docker](https://serversforhackers.com/articles/2014/03/20/getting-started-with-docker/) by [@fideloper](https://github.com/fideloper)
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm](https://github.com/fgrehm)
* [Automatic Docker Service Announcement with Registrator](http://progrium.com/blog/2014/09/10/automatic-docker-service-announcement-with-registrator/) by [@progrium](https://github.com/progrium)

## Books

* [Docker Book](http://dockerbook.com/) by James Turnbul ([@kartar](https://twitter.com/kartar))
 
## Dev 

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Cheat Sheet](https://github.com/wsargent/docker-cheat-sheet)  by [@wsargent](https://github.com/wsargent)

### Continous Integration / Continous Delivery

* [Drone](https://github.com/drone/drone) - https://drone.io/

### Dev Tools

* [Fig](https://github.com/docker/fig) (Fast, isolated development environments using Docker) -- http://www.fig.sh
* [Sandstorm](https://github.com/sandstorm-io/sandstorm) (Personal Cloud Sandbox /!\ uses LXC, not docker, but it's so nice) -- https://sandstorm.io/
* [ns-enter](https://github.com/jpetazzo/nsenter) (no more ssh, enter name spaces of container) by @jpetazzo
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) (in case of proxy problem) by [@jpetazzo](https://github.com/jpetazzo/)
* [docker-gen](https://github.com/jwilder/docker-gen) (Generate files from docker container meta-data) by [@jwilder](https://github.com/jwilder)
* [registrator](https://github.com/progrium/registrator) (Service registry bridge for Docker) by [@progrium](https://github.com/progrium)

### Images

* [Base Image by @phusion](https://github.com/phusion/baseimage-docker)
* [Busybox](https://github.com/jpetazzo/docker-busybox) (with either `buildroot` or Ubuntu's `busybox-static`) by [@jpetazzo](https://github.com/jpetazzo)
* [Busybox](https://github.com/progrium/busybox) (with `opkg`) by [@progrium](https://github.com/progrium)
* [DockerFile Example](https://github.com/komljen/dockerfile-examples) by [@komljen](https://github.com/komljen)
* [OpenWRT](http://www.zoobab.com/docker-openwrt-image) by [@zoobab](https://github.com/zoobab)

### Managing Images

* [Shipyard](https://github.com/shipyard/shipyard) (Composable Docker Management) -- http://shipyard-project.com/
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui/wiki) (Docker Management for Humans) -- http://panamax.io/
* [Docker Registry](https://github.com/docker/docker-registry)
* [Docker Registry Web](https://github.com/atc-/docker-registry-web) (A web UI for easy private/local Docker Registry integration) by [@atc-](https://github.com/atc-)
* [Flynn](https://github.com/flynn/flynn) (A next generation open source platform as a service (PaaS)) -- https://flynn.io/
* [Deis](https://github.com/deis/deis) (Your PaaS, your rules) -- http://deis.io/
* [Gaudi](https://github.com/marmelab/gaudi) (Gaudi allows to share multi-component applications, based on Docker, Go, and YAM) -- http://gaudi.io/
* [CoreOS](https://github.com/coreos) (Linux for Massive Server Deployments) -- https://coreos.com/
* [Boot2Docker](https://github.com/boot2docker/boot2docker) (docker for OSX and Windows) -- http://boot2docker.io/
* [Dokku](https://github.com/progrium/dokku) (Docker powered mini-Heroku in around 100 lines of Bash)
  * [dokku-alt](https://github.com/dokku-alt/dokku-alt) (Dokku fork with Dockerfile support, database plugins, ACL and more)

## Slides

* [Docker Slideshare Account](http://www.slideshare.net/dotCloud)
* [Docker Security](http://www.slideshare.net/jpetazzo) with [@jpetazzo](https://github.com/jpetazzo)

## Videos

### Main Account

* [Docker Youtube Account](http://www.youtube.com/user/dockerrun)
* [CenturyLink Labs Docker Interviews](http://www.youtube.com/playlist?list=PL_q4Fk7SVBCIjyuCBFBItXnzGI3qBa2L1)
* [YLD Event](https://www.youtube.com/channel/UCvksXSnLqIVM_uFB7xyrsSg/videos) Conference about *containers*!!! [@YLDio](https://twitter.com/YLDio)

### Useful videos

* [Orchestrating Docker containers in production using Fig](https://www.youtube.com/watch?v=SEtRg8siQWw) (7:11)
* [Development Environments with Fig](http://youtu.be/QpSFOHvFyMc) by Aanand Prasad (17:58)
* [Docker in Production](http://youtu.be/Glk5d5WP6MI) by [@jpetazzo](https://github.com/jpetazzo) (36:05)
* [Docker: How to Use Your Own Private Registry](https://www.youtube.com/watch?v=CAewZCBT4PI) (15:01)
* [Contributing to Docker by Andrew "Tianon" Page (InfoSiftr)](https://www.youtube.com/watch?v=1jwo8-1HYYg) (34:31)
* [Performance Analysis of Docker - Jeremy Eder](https://www.youtube.com/watch?v=6f2E6PKYb0w) (1:36:58)
* [Docker and SELinux by Daniel Walsh from Red Hat ](https://www.youtube.com/watch?v=zWGFqMuEHdw) (40:23)
* [Immutable Infrastructure with Docker and EC2 by Michael Bryzek (Gilt)](https://www.youtube.com/watch?v=GaHzdqFithc) (42:04)
* [Ansible and Docker HP](https://www.youtube.com/watch?v=oZ45v8AeE7k) (32:38)


## (Interesting) Twitter Accounts

* [Docker](https://twitter.com/docker)
* [Century Link Labs](https://twitter.com/CenturyLinkLabs)
* [Flux7Labs](https://twitter.com/Flux7Labs)
* [TutumCloud](https://twitter.com/tutumcloud)
* [Project Atomic](https://twitter.com/ProjectAtomic)
* [Openshift By Red Hat](https://twitter.com/openshift)
* [YLD](https://twitter.com/YLDio)

### People
* [Solomon Hykes](https://twitter.com/solomonstre) Founder of Docker
* [Gabriel Monroy](https://twitter.com/gabrtv) Creator of Deis
* [Jérôme Petazzoni](https://twitter.com/jpetazzo) Docker Developer
* [Michael Crosby](https://twitter.com/crosbymichael) Docker Developer
* [James Turnbull](https://twitter.com/kartar) Author of Docker Book
* [Jeff Lindsay](https://twitter.com/progrium) Design-minded software architect
