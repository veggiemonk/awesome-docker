# Awesome Docker

> A curated list of Docker resources

Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome](https://github.com/sindresorhus/awesome)

Share some love: add your docker resource by forking / sending pull requests.

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)**

This list is not exhaustive (nor is it meant to be) but can probably help you get up and running faster. 

---

* [Useful Articles](#useful-articles)
 * [Main Resources](#main-resources)
 * [General Articles](#general-articles)
 * [Good Tips](#good-tips)
 * [Optimizing Images](#optimizing-images)
 * [Service Discovery](#service-discovery)
 * [Security](#security)
 * [Performances](#performances)
 * [Other](#other)
* [Books](#books)
* [Dev](#dev)
  * [Continous Integration / Continous Delivery](#continous-integration-continous-delivery) 
  * [Hosting for repositories (registries)](#hosting-for-repositories-registries)
  * [Hosting for containers](#hosting-for-containers)
  * [Dev Tools](#dev-tools)
  * [Useful Images](#userful-images)
  * [Storing Images](#storing-images)
  * [Running Containers](#running-containers)
  * [Service Discovery](#service-discovery)
* [Slides](#slides)
* [Videos](#videos)
  * [Main Account](#main-account)
  * [Useful videos](#useful-videos)
* [Interesting Twitter Accounts](#interesting-twitter-accounts)
  
----

## Useful Articles

### Main Resources

* [Docker Weekly](http://blog.docker.com/docker-weekly-archives/) Huge resource
* [Century Links Labs](http://www.centurylinklabs.com/category/docker-posts/)
* [Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html) Very complete
* [Docker Ecosystem](http://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) <--- MUST SEE
* [Blog](http://jpetazzo.github.io/) of [@jpetazzo](https://github.com/jpetazzo)
* [Blog](http://progrium.com/blog/) of [@progrium](https://github.com/progrium)

### General Articles
* [Getting Started with Docker](https://serversforhackers.com/articles/2014/03/20/getting-started-with-docker/) by [@fideloper](https://github.com/fideloper) -- [Servers For Hackers](https://serversforhackers.com/editions/) is valuable resource. At some point, every programmer finds themselves needing to know their way around a server.
* [How to Use Docker on OS X: The Missing Guide](http://viget.com/extend/how-to-use-docker-on-os-x-the-missing-guide)
* [Docker for (Java) Developers](http://ro14nd.de/Docker-for-Developers/)

### Good Tips
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm](https://github.com/fgrehm)
* [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/) 
* [Using NSEnter with Boot2Docker](http://ro14nd.de/NSEnter-with-Boot2Docker/)

### Optimizing Images
* [Create the smallest possible Docker container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
* [Creating a Docker image from your code](http://blog.tutum.co/2014/04/10/creating-a-docker-image-from-your-code/)
* [Optimizing Docker Images](http://www.centurylinklabs.com/optimizing-docker-images/?hvid=1OW0br)
* [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07) by @kelseyhightower
* [Squashing Docker Images](http://jasonwilder.com/blog/2014/08/19/squashing-docker-images/) by [@jwilder](https://github.com/jwilder)
* [Dockerfile Golf (or optimizing the Docker build process)](http://www.davidmkerr.com/2014/08/dockerfile-golf-or-optimizing-docker.html)

### Service Discovery
* [@progrium](https://github.com/progrium) Service Discovery articles series:
 * [Consul Service Discovery with Docker](http://progrium.com/blog/2014/08/20/consul-service-discovery-with-docker/) 
 * [Understanding Modern Service Discovery with Docker](http://progrium.com/blog/2014/07/29/understanding-modern-service-discovery-with-docker/)
 * [Automatic Docker Service Announcement with Registrator](http://progrium.com/blog/2014/09/10/automatic-docker-service-announcement-with-registrator/) 

### Security
* [Docker and SELinux](http://www.projectatomic.io/docs/docker-and-selinux/)
* [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
* 
### Perfomances
* [Performance Analysis of Docker on Red Hat Enterprise Linux 7](http://developerblog.redhat.com/2014/08/19/performance-analysis-docker-red-hat-enterprise-linux-7/)
* [Distrubuted JMeter testing using Docker](http://srivaths.blogspot.fr/2014/08/distrubuted-jmeter-testing-using-docker.html?m=1)
* [nsinit: per-container resource monitoring of Docker containers on RHEL/Fedora](http://www.breakage.org/2014/09/03/nsinit-per-container-resource-monitoring-of-docker-containers-on-rhelfedora/)

### Other
* Presentation: Docker and JBoss - the perfect combination
 * [Vidéo](http://www.youtube.com/watch?v=4uQ6gR_xZhE)
 * [Slides](https://goldmann.pl/presentations/2014-vjbug-docker)
 * [Code source](https://github.com/goldmann/goldmann.pl/tree/master/.presentations/2014-vjbug-docker/demos)


## Books

* [Docker Book](http://dockerbook.com/) by James Turnbul ([@kartar](https://twitter.com/kartar))
 
## Dev 

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Cheat Sheet](https://github.com/wsargent/docker-cheat-sheet)  by [@wsargent](https://github.com/wsargent)

### Continous Integration / Continous Delivery

* [Drone](https://github.com/drone/drone) - https://drone.io/

### Hosting for repositories (registries)

* [Docker Hub](https://hub.docker.com/) (provided by Docker Inc.)
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [Reesd](https://reesd.com/) -  Private Docker repositories and redundant storage service by [@noteed](https://github.com/noteed)

### Hosting for containers (Paying Services $$$)

* [Orchard](https://www.orchardup.com/) (part of Docker Inc) - Get a Docker host in the cloud, instantly.
* [StackDock](https://stackdock.com/) - Docker hosting on blazing fast dedicated infrastructure 
* [Tutum](https://www.tutum.co/) - Simple hosting for your Docker containers.

### Dev Tools

* [Fig](https://github.com/docker/fig) (Fast, isolated development environments using Docker) -- http://www.fig.sh
* [Sandstorm](https://github.com/sandstorm-io/sandstorm) (Personal Cloud Sandbox /!\ uses LXC, not docker, but it's so nice) -- https://sandstorm.io/
* [ns-enter](https://github.com/jpetazzo/nsenter) (no more ssh, enter name spaces of container) by [@jpetazzo](https://github.com/jpetazzo)
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) (in case of proxy problem) by [@jpetazzo](https://github.com/jpetazzo/)
* [docker-gen](https://github.com/jwilder/docker-gen) (Generate files from docker container meta-data) by [@jwilder](https://github.com/jwilder)
* [registrator](https://github.com/progrium/registrator) (Service registry bridge for Docker) by [@progrium](https://github.com/progrium)
* [Boot2Docker](https://github.com/boot2docker/boot2docker) (docker for OSX and Windows) -- http://boot2docker.io/
* [OctoHost](http://www.octohost.io/) (Simple web focused Docker based mini-PaaS server. git push to deploy your websites as needed) by [@octohost](https://github.com/octohost)

### Useful Images

* [Base Image](https://github.com/phusion/baseimage-docker) by [@phusion](https://github.com/phusion/)
* [Busybox](https://github.com/jpetazzo/docker-busybox) (with either `buildroot` or Ubuntu's `busybox-static`) by [@jpetazzo](https://github.com/jpetazzo)
* [Busybox](https://github.com/progrium/busybox) (with `opkg`) by [@progrium](https://github.com/progrium)
* [DockerFile Example](https://github.com/komljen/dockerfile-examples) by [@komljen](https://github.com/komljen)
* [OpenWRT](http://www.zoobab.com/docker-openwrt-image) by [@zoobab](https://github.com/zoobab)
* [Phusion Docker Hub Account](https://hub.docker.com/u/phusion/)

### Storing Images

* [Docker Registry](https://github.com/docker/docker-registry) (Registry server for Docker (hosting/delivering of repositories and images))
* [Docker Registry Web](https://github.com/atc-/docker-registry-web) (A web UI for easy private/local Docker Registry integration) by [@atc-](https://github.com/atc-)

### Running Containers

* [Shipyard](https://github.com/shipyard/shipyard) (Composable Docker Management) -- http://shipyard-project.com/
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui/wiki) (Docker Management for Humans) -- http://panamax.io/
* [Flynn](https://github.com/flynn/flynn) (A next generation open source platform as a service (PaaS)) -- https://flynn.io/
* [Deis](https://github.com/deis/deis) (Your PaaS, your rules) -- http://deis.io/
* [Gaudi](https://github.com/marmelab/gaudi) (Gaudi allows to share multi-component applications, based on Docker, Go, and YAM) -- http://gaudi.io/
* [CoreOS](https://github.com/coreos) (Linux for Massive Server Deployments) -- https://coreos.com/
* [Dokku](https://github.com/progrium/dokku) (Docker powered mini-Heroku in around 100 lines of Bash) by [@progrium](https://github.com/progrium)
  * [dokku-alt](https://github.com/dokku-alt/dokku-alt) (Dokku fork with Dockerfile support, database plugins, ACL and more)
* [Ansible - manage docker containers](http://docs.ansible.com/docker_module.html)

### Service Discovery

* [docker-consul](https://github.com/progrium/docker-consul) by [@progrium](https://github.com/progrium)
* [etcd](https://github.com/coreos/etcd): A highly-available key value store for shared configuration and service discovery

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


## Interesting Twitter Accounts

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
