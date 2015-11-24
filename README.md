# Awesome Docker [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome) [![Join the chat at https://gitter.im/veggiemonk/awesome-docker](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/veggiemonk/awesome-docker?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

> A curated list of Docker resources and projects
in
Inspired by [@sindresorhus](https://github.com/sindresorhus)' [awesome](https://github.com/sindresorhus/awesome) and improved by these **[amazing contributors](https://github.com/veggiemonk/awesome-docker/graphs/contributors)**.

It's now a github project because it's considerably easier for other people to edit, fix and expand on Docker using Github. Just click [README.md](https://github.com/veggiemonk/awesome-docker/edit/master/README.md).
If this list is not complete, you can [contribute](https://github.com/veggiemonk/awesome-docker/edit/master/README.md) to make it so.

> **Please**, help organize these resources so that they are _easy to find_ and _understand_ for new comers. See how to **[Contribute](https://github.com/veggiemonk/awesome-docker/blob/master/CONTRIBUTING.md)**

#### *If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!*

# What is Docker ?

> Docker is an open platform for developers and sysadmins to build, ship, and run distributed applications. Consisting of Docker Engine, a portable, lightweight runtime and packaging tool, and Docker Hub, a cloud service for sharing applications and automating workflows, Docker enables apps to be quickly assembled from components and eliminates the friction between development, QA, and production environments. As a result, IT can ship faster and run the same app, unchanged, on laptops, data center VMs, and any cloud.

_Source:_ [What is Docker](https://www.docker.com/what-docker)

# Where to start ?

* [10-minute Interactive Tutorial](http://docs.docker.com/mac/started/)
* [Docker Training](http://training.docker.com/)
* Read this complete article: [Basics – Docker, Containers, Hypervisors, CoreOS](http://etherealmind.com/basics-docker-containers-hypervisors-coreos/)
* Watch the video: [Docker for Developers](https://www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26) by [@jpetazzo](https://github.com/jpetazzo)
* [Docker Jumpstart](https://github.com/odewahn/docker-jumpstart/): a quick introduction
* [Install Docker on your machine](https://github.com/wsargent/docker-cheat-sheet#installation) and play with a few [Useful Images](#useful-images)
* Try [Panamax: Docker Management for Humans](http://panamax.io/) It will install a CoreOS VM with VirtualBox and has nice front end
* [Install Docker Toolbox](https://www.docker.com/docker-toolbox) Docker Toolbox is an installer to quickly and easily install and setup a Docker environment on your computer. Available for both Windows and Mac, the Toolbox installs Docker Client, Machine, Compose (Mac only), Kitematic and VirtualBox.
* Check out: [Docker Cheat Sheet](https://github.com/wsargent/docker-cheat-sheet) by [@wsargent](https://github.com/wsargent) __MUST SEE__
* [Project Web Dev](http://project-webdev.blogspot.de) : (Article series) How to create your own website based on Docker
* [Docker Containers on the desktop](https://blog.jessfraz.com/post/docker-containers-on-the-desktop/) by [@jfrazelle](https://github.com/jfrazelle) The **funniest way** to learn about docker! (Tips: checkout her [dotfiles](https://github.com/jfrazelle/dotfiles) and her [dockerfiles](https://github.com/jfrazelle/dockerfiles))
* [Container Hacks and Fun Images](https://www.youtube.com/watch?v=1qlLUf7KtAw) by [@jfrazelle](https://github.com/jfrazelle) @ DockerCon 2015 **MUST WATCH VIDEO** (38:50)
* [Learn Docker](https://github.com/dwyl/learn-docker) Full environment set up, screenshots, step-by-step tutorial and more resources (video, articles, cheat sheets) by [@dwyl](https://github.com/dwyl)

# MENU

* [What is Docker ?](#what-is-docker-)
* [Where to start ?](#where-to-start-)
* [Useful Articles](#useful-articles)
 * [Main Resources](#main-resources)
 * [General Articles](#general-articles)
 * [Deep Dive](#deep-dive)
 * [Networking](#networking)
 * [Metal](#metal)
 * [Multi-Server](#multi-server)
 * [Cloud Infrastructure](#cloud-infrastructure)
 * [Good Tips](#good-tips)
 * [Newsletter](#newsletter)
 * [Continuous Integration](#continuous-integration)
 * [Optimizing Images](#optimizing-images)
 * [Service Discovery](#service-discovery)
 * [Security](#security)
 * [Performances](#performances)
 * [Raspberry Pi & ARM](#raspberry-pi--arm)
 * [Other](#other)
* [Books](#books)
* [Tools](#tools)
  * [Dev Tools](#dev-tools)
  * [Continuous Integration / Continuous Delivery](#continuous-integration--continuous-delivery)
  * [Deployment](#deployment)
  * [Hosting for repositories (registries)](#hosting-for-repositories-registries)
  * [Hosting for containers](#hosting-for-containers)
  * [Reverse Proxy](#reverse-proxy)
  * [Web Interface](#web-interface)
  * [Local Container Manager](#local-container-manager)
  * [Useful Images](#useful-images)
  * [Dockerfile](#dockerfile)
  * [Storing Images](#storing-images)
  * [Monitoring](#monitoring)
  * [Networking](#networking)
  * [Logging](#logging)
  * [Deployment and Infrastructure](#deployment-and-infrastructure)
  * [PaaS](#paas)
  * [Remote Container Manager / Orchestration](#remote-container-manager--orchestration)
  * [Security](#security-1)
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
* [Docker Cheat Sheet](https://github.com/wsargent/docker-cheat-sheet) by [@wsargent](https://github.com/wsargent) __MUST SEE__
* [CenturyLink Labs](https://labs.ctl.io/category/docker/)
* [Valuable Docker Links](http://www.nkode.io/2014/08/24/valuable-docker-links.html) Very complete
* [Docker Ecosystem](https://www.mindmeister.com/389671722/docker-ecosystem) (Mind Map) __MUST SEE__
* [Docker Ecosystem](http://comp.photo777.org/wp-content/uploads/2015/09/Docker-ecosystem-8.5.1.pdf) (PDF) __MUST SEE__    find it on [blog](http://comp.photo777.org/docker-ecosystem/) by Bryzgalov Peter.
* [Blog](https://blog.jessfraz.com/) of [@frazelledazzell](https://github.com/jfrazelle)
* [Blog](http://jpetazzo.github.io/) of [@jpetazzo](https://github.com/jpetazzo)
* [Blog](http://progrium.com/blog/) of [@progrium](https://github.com/progrium)
* [Blog](http://jasonwilder.com/) of [@jwilder](https://github.com/jwilder)
* [Blog](http://crosbymichael.com/) of [@crosbymichael](https://github.com/crosbymichael)
* [Blog](http://gliderlabs.com/blog/) of [@gliderlabs](https://github.com/gliderlabs)
* [Blog](http://sebgoa.blogspot.com/) of [@sebgoa](https://twitter.com/sebgoa)
* [Digital Ocean Community](https://www.digitalocean.com/community/search?primary_filter=tutorials&query=docker)
* [Container42](http://container42.com/)
* [Contrainer solutions](http://container-solutions.com/blog/)
* [DockerOne](http://dockone.io/) Docker Community (in Chinese) by [@LiYingJie](http://dockone.io/people/%E6%9D%8E%E9%A2%96%E6%9D%B0)
* [Project Web Dev](http://project-webdev.blogspot.de) : (Article series) How to create your own website based on Docker
* [Docker vs. VMs? Combining Both for Cloud Portability Nirvana](http://www.rightscale.com/blog/cloud-management-best-practices/docker-vs-vms-combining-both-cloud-portability-nirvana)
* [Docker Containers on the desktop](https://blog.jessfraz.com/post/docker-containers-on-the-desktop/) by [@jfrazelle](https://github.com/jfrazelle) The **funniest way** to learn about docker! (Tips: checkout her [dotfiles](https://github.com/jfrazelle/dotfiles) and her [dockerfiles](https://github.com/jfrazelle/dockerfiles))
* [Awesome Linux Container](https://github.com/Friz-zy/awesome-linux-containers) more general about container than this repo, by [@Friz-zy](https://github.com/Friz-zy).

## General Articles
* [Getting Started with Docker](https://serversforhackers.com/getting-started-with-docker/) by [@fideloper](https://github.com/fideloper) -- [Servers For Hackers](https://serversforhackers.com/editions/) is valuable resource. At some point, every programmer finds themselves needing to know their way around a server.
* [What is Docker and how do you monitor it?](http://axibase.com/docker-monitoring/)
* [How to Use Docker on OS X: The Missing Guide](https://viget.com/extend/how-to-use-docker-on-os-x-the-missing-guide)
* [Docker for (Java) Developers](http://ro14nd.de/Docker-for-Developers/)
* [Deploying NGINX with Docker](https://www.nginx.com/blog/deploying-nginx-nginx-plus-docker/)
* [Eight Docker Development Patterns](http://www.hokstad.com/docker/patterns)
* [Rails Development Environment for OS X using Docker](https://allenan.com/docker-rails-dev-environment-for-osx/)
* [Logging on Docker: What You Need to Know](https://dzone.com/articles/logging-docker-what-you-need) + see the [video](https://vimeo.com/123341629) (~50min)
* [Comparing Five Monitoring Options for Docker](http://rancher.com/comparing-monitoring-options-for-docker-deployments/)
* [Minimalistic data-only container for Docker Compose](http://dockermeetupsinbordeaux.github.io/docker-compose/data-container/2015/03/01/minimalistic-docker-data-container.html) (Written Mar 1, 2015)
* [Running Docker Containers with Systemd](http://container-solutions.com/running-docker-containers-with-systemd/)
* [Dockerizing Flask With Compose and Machine - From Localhost to the Cloud](https://realpython.com/blog/python/dockerizing-flask-with-compose-and-machine-from-localhost-to-the-cloud/) -- [GitHub](https://github.com/realpython/orchestrating-docker) Learn how to deploy an application using Docker Compose and Docker Machine (written 17 April 2015)
* [Why and How to use Docker for Development](https://medium.com/iron-io-blog/why-and-how-to-use-docker-for-development-a156c1de3b24) (written 28 APR 2015)
* [Automating Docker Logging: ElasticSearch, Logstash, Kibana, and Logspout](http://nathanleclaire.com/blog/2015/04/27/automating-docker-logging-elasticsearch-logstash-kibana-and-logspout/) (written 27 APR 2015)
* [Docker Host Volume Synchronization](http://oliverguenther.de/2015/05/docker-host-volume-synchronization/) (written 1 JUN 2015)
* [Multi-Service Local Development Environment with Vagrant and Docker](http://devbandit.com/2015/05/29/vagrant-and-docker.html) (written 29 MAY 2015)
* [From Local Development to Remote Deployment with Docker Machine and Compose](https://developer.rackspace.com/blog/dev-to-deploy-with-docker-machine-and-compose/) (written 2 JUL 2015)
* [Docker: Build, Ship and Run Any App, Anywhere](http://delftswa.github.io/chapters/docker/index.html) by [Martijn Dwars](https://github.com/MartijnDwars), [Wiebe van Geest](https://github.com/wrvangeest), [Rik Nijessen](https://github.com/gewoonrik), and [Rick Wieman](https://github.com/RickWieman) from [Delft University of Technology](http://www.tudelft.nl/) (written 2 JUL 2015)
* [Joining the Docker Ship](http://thenewstack.io/joining-the-docker-ship-and-go/) Learn how to contribute to docker (written 9 JUL 2015)
* [Continuous Deployment with Gradle and Docker](https://github.com/gesellix/pipeline-with-gradle-and-docker/blob/master/README.md) Describes a complete pipeline from source to production deploy (includes a complete Spring Boot example project) by [@gesellix](https://github.com/gesellix)
* [Containerization and the PaaS Cloud](http://www.computer.org/cms/Computer.org/ComputingNow/issues/2015/09/mcd2015030024.pdf) -- This article discusses the requirements that arise from having to facilitate applications through distributed multicloud platforms.
* [Docker for Development: Common Problems and Solutions](https://medium.com/@rdsubhas/docker-for-development-common-problems-and-solutions-95b25cae41eb) by [@rdsubhas](https://github.com/rdsubhas)
* [Docker Adoption Data](https://www.datadoghq.com/docker-adoption/) A study by Datadog on the real world Docker usage stastics and deployment patterns.
* [How to monitor Docker](https://www.datadoghq.com/blog/the-docker-monitoring-problem/) (4-part series)


## Deep Dive
* [Creating containers - Part 1](http://crosbymichael.com/creating-containers-part-1.html) This is part one of a series of blog posts detailing how docker creates containers. By [@crosbymichael](https://github.com/crosbymichael)
* [Data-only container madness](http://container42.com/2014/11/18/data-only-container-madness/)

## Networking
* [Using Docker Machine with Weave 0.10](http://blog.weave.works/2015/04/22/using-docker-machine-with-weave-0-10/) (written 22 APR 2015)
* [How to Route Traffic through a Tor Docker container](https://blog.jessfraz.com/post/routing-traffic-through-tor-docker-container/) by [@jfrazelle](https://github.com/jfrazelle) (writtent 20 JUN 2015)

## Metal
* [How to use Docker on Full Metal](http://blog.bigstep.com/big-data-performance/use-docker-full-metal-cloud/)

## Multi-Server
* [Using Fig and Flocker to build, test, deploy and migrate multi-server Dockerized apps](https://clusterhq.com/blog/fig-flocker-multi-server-docker-apps/)
* [A Docker based mini-PaaS](http://shortcircuit.net.au/~prologic/blog/article/2015/03/24/a-docker-based-mini-paas/) by [@prologic](https://github.com/prologic)

## Cloud Infrastructure
* [Cloud Infrastructure Automation for Docker Nodes](http://blog.tutum.co/2015/04/29/cloud-infrastructure-automation-for-docker-nodes/)

## Good Tips
* [24 random docker tips](http://csaba.palfi.me/random-docker-tips/) by [@csabapalfi](https://github.com/csabapalfi)
* [GUI Apps with Docker](http://fabiorehm.com/blog/2014/09/11/running-gui-apps-with-docker/) by [@fgrehm](https://github.com/fgrehm)
* [Automated Nginx Reverse Proxy for Docker](http://jasonwilder.com/blog/2014/03/25/automated-nginx-reverse-proxy-for-docker/) by [@jwilder](https://github.com/jwilder)
* [Using NSEnter with Boot2Docker](http://ro14nd.de/NSEnter-with-Boot2Docker/)
* [A Simple Way to Dockerize Applications](http://jasonwilder.com/blog/2014/10/13/a-simple-way-to-dockerize-applications/) by [@jwilder](https://github.com/jwilder)
* [Building good docker images](http://jonathan.bergknoff.com/journal/building-good-docker-images) by [@jbergknoff](https://github.com/jbergknoff)
* [10 Things Not To Forget Before Deploying Docker In Production](http://www.slideshare.net/rightscale/docker-meetup-40826948)
* [Make Your Docker Workflow Awesome With Fig.sh](http://www.syncano.com/docker-workflow-fig-sh/) Fig is a python application that helps you run groups of docker containers.
* [Docker CIFS – How to Mount CIFS as a Docker Volume](http://backdrift.org/docker-cifs-howto-mount-cifs-volume-docker-container)
* [Nginx Proxy for Docker](http://blog.danivovich.com/2015/07/09/nginx-proxy-for-docker-containers/) (written 9 JUL 2015)
* [Dealing with linked containers dependency in docker-compose](http://brunorocha.org/python/dealing-with-linked-containers-dependency-in-docker-compose.html) by [@rochacbruno](https://github.com/rochacbruno)
* [Docker Tips](http://www.mervine.net/notes/docker-tips) by [@jmervine](https://github.com/jmervine)
* [Docker on Windows behind a firewall](http://toedter.com/2015/05/11/docker-on-windows-behind-a-firewall/) by [@kaitoedter](https://twitter.com/kaitoedter)
* [Pulling Git into a Docker image without leaving SSH keys behind](http://blog.cloud66.com/pulling-git-into-a-docker-image-without-leaving-ssh-keys-behind/) by [@khash](https://github.com/khash)
* [6 Million Ways To Log In Docker](http://www.slideshare.net/raychaser/6-million-ways-to-log-in-docker-nyc-docker-meetup-12172014) by [@aychaser](https://twitter.com/aychaser)
* [Dockerfile Generator](http://jrruethe.github.io/blog/2015/09/20/dockerfile-generator/) (ruby script)
* [Running Production Hadoop Clusters in Docker Containers](http://conferences.oreilly.com/strata/big-data-conference-ca-2015/public/schedule/detail/38521)

## Newsletter
* [Docker Team](https://www.docker.com/)
* [CenturyLink Labs](https://labs.ctl.io/)
* [Tutum](https://dashboard.tutum.co/accounts/login/)
* [DevOps Weekly](http://www.devopsweekly.com)
* [Shippable](http://blog.shippable.com/)
* [WebOps weekly](http://webopsweekly.com/)

## Continuous Integration
* [Docker and Phoenix: How to Make Your Continuous Integration More Awesome](http://ariya.ofilabs.com/2014/12/docker-and-phoenix-how-to-make-your-continuous-integration-more-awesome.html)

## Optimizing Images
* [Create the smallest possible Docker container](http://blog.xebia.com/2014/07/04/create-the-smallest-possible-docker-container/)
* [Creating a Docker image from your code](http://blog.tutum.co/2014/04/10/creating-a-docker-image-from-your-code/)
* [Optimizing Docker Images](https://labs.ctl.io/optimizing-docker-images/?hvid=1OW0br)
* [How to Optimize Your Dockerfile](http://blog.tutum.co/2014/10/22/how-to-optimize-your-dockerfile/) by [@tutumcloud](https://github.com/tutumcloud)
* [Building Docker Images for Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07) by [@kelseyhightower](https://github.com/kelseyhightower)
* [Squashing Docker Images](http://jasonwilder.com/blog/2014/08/19/squashing-docker-images/) by [@jwilder](https://github.com/jwilder)
* [Dockerfile Golf (or optimizing the Docker build process)](http://www.davidmkerr.com/2014/08/dockerfile-golf-or-optimizing-docker.html)
* [ImageLayers](https://imagelayers.io/) Visualize Docker images and the layers that compose them.
* [DockerSlim](https://github.com/cloudimmunity/docker-slim) shrinks fat Docker images creating the smallest possible images.
* [SkinnyWhale](https://github.com/djosephsen/skinnywhale) Skinnywhale helps you make smaller (as in megabytes) Docker containers.

## Service Discovery
* [@progrium](https://github.com/progrium) Service Discovery articles series:
 * [Consul Service Discovery with Docker](http://progrium.com/blog/2014/08/20/consul-service-discovery-with-docker/)
 * [Understanding Modern Service Discovery with Docker](http://progrium.com/blog/2014/07/29/understanding-modern-service-discovery-with-docker/)
 * [Automatic Docker Service Announcement with Registrator](http://progrium.com/blog/2014/09/10/automatic-docker-service-announcement-with-registrator/)

## Security
* [Docker and SELinux](http://www.projectatomic.io/docs/docker-and-selinux/)
* [Bringing new security features to Docker](https://opensource.com/business/14/9/security-for-docker)
* [Docker Secure Deployment Guidelines](https://github.com/GDSSecurity/Docker-Secure-Deployment-Guidelines)
* [Security Best Practices for Building Docker Images](http://linux-audit.com/tag/docker/)
* [Docker Security: Are Your Containers Tightly Secured to the Ship? SlideShare](http://fr.slideshare.net/MichaelBoelen/docker-security-are-your-containers-tightly-secured-to-the-ship)
* [Tuning Docker with the newest security enhancements](https://opensource.com/business/15/3/docker-security-tuning)
* [Lynis is an open source security auditing tool including Docker auditing](https://cisofy.com/lynis/)
* [Understanding Docker security and best practices](https://blog.docker.com/2015/05/understanding-docker-security-and-best-practices/) (written 5 MAY 2015)

## Performances
* [Performance Analysis of Docker on Red Hat Enterprise Linux 7](http://developerblog.redhat.com/2014/08/19/performance-analysis-docker-red-hat-enterprise-linux-7/)
* [Distrubuted JMeter testing using Docker](http://srivaths.blogspot.fr/2014/08/distrubuted-jmeter-testing-using-docker.html?m=1)
* [nsinit: per-container resource monitoring of Docker containers on RHEL/Fedora](http://www.breakage.org/2014/09/03/nsinit-per-container-resource-monitoring-of-docker-containers-on-rhelfedora/)

## Raspberry Pi & ARM
* [git push docker containers to linux devices](https://resin.io/) Modern DevOps for IoT, leveraging git and Docker.
* [Docker Pirates ARMed with explosive stuff](http://blog.hypriot.com/) Huge resource on clustering, swarm, docker, pre-installed image for SD card on Raspberry Pi
* [Docker on Raspberry Pi](http://blog.xebia.com/2014/08/25/docker-on-a-raspberry-pi/)
* [Fool-Proof Recipe: Docker on the Raspberry Pi](https://www.voxxed.com/blog/2015/04/fool-proof-recipe-docker-on-the-raspberry-pi/) Same article as above but more opinionated.
* [Raspberry Pi with Docker 1.5.0](http://blog.hypriot.com/post/heavily-armed-after-major-upgrade-raspberry-pi-with-docker-1-dot-5-0/)
* [Swarming Raspberry Pi – Part 1](http://matthewkwilliams.com/index.php/2015/03/21/swarming-raspberry-pi-part-1/)
* [Swarming Raspberry Pi, Part 2: Registry & Mirror](http://matthewkwilliams.com/index.php/2015/03/29/swarming-raspberry-pi-part-2-registry-mirror/)
* [Swarming Raspberry Pi: Docker Swarm Discovery Options](http://matthewkwilliams.com/index.php/2015/04/03/swarming-raspberry-pi-docker-swarm-discovery-options/)
* [Uniform Development by Docker & QEMU](http://www.instructables.com/id/Uniform-Development-by-Docker-QEMU/)
* [Get Docker up and running on the RaspberryPi in four steps](https://github.com/umiddelb/armhf/wiki/Get-Docker-up-and-running-on-the-RaspberryPi-%28ARMv6%29-in-four-steps)
* [Installing, running, using Docker on armhf (ARMv7) devices](https://github.com/umiddelb/armhf/wiki/Installing,-running,-using-docker-on-armhf-%28ARMv7%29-devices)
* [How to run 2500 webservers on a Raspebery Pi](http://blog.loof.fr/2015/10/how-to-run-2500-webservers-on-raspberry.html)


## Other
* Presentation: Docker and JBoss - the perfect combination
 * [Vidéo](https://www.youtube.com/watch?v=4uQ6gR_xZhE)
 * [Slides](https://goldmann.pl/presentations/2014-vjbug-docker/)
 * [Code source](https://github.com/goldmann/goldmann.pl/tree/master/.presentations/2014-vjbug-docker/demos)

# Books

* [Docker Book](http://dockerbook.com/) by James Turnbul ([@kartar](https://twitter.com/kartar))
* [Docker Cookbook](http://shop.oreilly.com/product/0636920036791.do) by Sébastien Goasguen ([@sebgoa](https://twitter.com/sebgoa)) (Publisher: O'Reilly)
* [Docker Cookbook](http://dockercookbook.github.io/) by Neependra Khare ([@neependra](https://twitter.com/neependra)) (Publisher: Packt)
* [Docker in Action](https://www.manning.com/books/docker-in-action) by Jeff Nickoloff ([@allingeek](https://twitter.com/allingeek))
* [Docker in Practice](https://www.manning.com/books/docker-in-practice) by Ian Miell ([@ianmiell](https://github.com/ianmiell)) and Aidan Hobson Sayers ([@aidanhs](https://github.com/aidanhs)). ==> [Website](http://docker-in-practice.github.io/)
* [Docker Up & Running](http://newrelic.com/docker-book) by [Karl Matthias](https://twitter.com/relistan) and [Sean P. Kane](https://twitter.com/spkane)
* [Using Docker](http://shop.oreilly.com/product/0636920035671.do) by Adrian Mouat ([@adrianmouat](https://twitter.com/adrianmouat))
* [Docker: Container-Infrastruktur für Microservices](http://www.bee42.com/dockerbook/) (German) by Peter Roßbach ([@PRossbach](https://twitter.com/PRossbach))

# Tools

* [Docker](https://github.com/docker/docker)
* [Docker Images](https://hub.docker.com)
* [Docker Compose](https://github.com/docker/compose/) (Define and run multi-container applications with Docker)
* [Docker Machine](https://github.com/docker/machine) (Machine management for a container-centric world)
* [Docker Registry](https://github.com/docker/distribution) (The Docker toolset to pack, ship, store, and deliver content)
* [Docker Swarm](https://github.com/docker/swarm) (Swarm: a Docker-native clustering system)

## Dev Tools

* [GoSu](https://github.com/tianon/gosu) ("run this specific application as this specific user and get out of the pipeline" -- entrypoint script tool) by [@tianon](https://github.com/tianon)
* [Chaperone](https://github.com/garywiz/chaperone) ("A single PID1 process designed for docker containers.  Does user management, log management, startup, zombie reaping, all in one small package.) by [@garywiz](https://github.com/garywiz)
* [ns-enter](https://github.com/jpetazzo/nsenter) (no more ssh, enter name spaces of container) by [@jpetazzo](https://github.com/jpetazzo)
* [Squid-in-a-can](https://github.com/jpetazzo/squid-in-a-can) (in case of proxy problem) by [@jpetazzo](https://github.com/jpetazzo/)
* [docker-gen](https://github.com/jwilder/docker-gen) (Generate files from docker container meta-data) by [@jwilder](https://github.com/jwilder)
* [dockerize](https://github.com/jwilder/dockerize) (Utility to simplify running applications in docker containers) by [@jwilder](https://github.com/jwilder)
* [registrator](https://github.com/progrium/registrator) (Service registry bridge for Docker) by [@progrium](https://github.com/progrium)
* [Dockly](https://github.com/swipely/dockly) (Dockly is a gem made to ease the pain of packaging an application in Docker.) by [@swipely](https://github.com/swipely/)
* [docker-volumes](https://github.com/cpuguy83/docker-volumes) (Docker Volume Manager) by [@cpuguy83](https://github.com/cpuguy83)
* [dockerfile_lint](https://github.com/projectatomic/dockerfile_lint) (A rule-based 'linter' for Dockerfiles) by [@redhataccess](https://github.com/redhataccess)
* [powerstrip](https://github.com/clusterhq/powerstrip) (A tool for prototyping Docker extensions) by [@clusterhq](https://github.com/clusterhq)
* [Vagga](https://github.com/tailhook/vagga) (Vagga is a containerisation tool without daemons. It is a fully-userspace container engine inspired by Vagrant and Docker, specialized for development environments.) by [@tailhook](https://github.com/tailhook/)
* [dockerode](https://github.com/apocas/dockerode) (Not just another Docker Remote API node.js module) by [@apocas](https://github.com/apocas)
* [go-dockerclient](https://github.com/fsouza/go-dockerclient/) (Go HTTP client for the Docker remote API.) by [@fsouza](https://github.com/fsouza/)
* [container-factory](https://github.com/lsqio/container-factory) (produces Docker images from tarballs of application source code http://www.containerfactory.io) by [@lsqio](https://github.com/lsqio)
* [percheron](https://github.com/ashmckenzie/percheron) (Organise your Docker containers with muscle and intelligence) by [@ashmckenzie](https://github.com/ashmckenzie)
* [crane](https://github.com/michaelsauter/crane) (Lift containers with ease. Easy orchestration for images and containers) by [@michaelsauter](https://github.com/michaelsauter)
* [sherdock](https://github.com/rancher/sherdock) (Automatic GC of images based on regexp) by [@rancherio](https://github.com/rancherio)
* [bocker](https://github.com/p8952/bocker) (1) (Docker implemented in 100 lines of bash) by [p8952](https://github.com/p8952)
* [bocker](https://github.com/icy/bocker) (2) (Write Dockerfile completely in Bash. Extensible and simple. --> Reusable) by [@icy](https://github.com/icy)
* [docker-cli](https://github.com/j-bennet/wharfee) (A CLI with autocompletion and syntax highlighting for Docker commands.) by [@j-bennet](https://github.com/j-bennet)
* [docker-gc](https://github.com/spotify/docker-gc) (A cron job that will delete old stopped containers and unused images) by [@spotify](https://github.com/spotify)
* [dlayer](https://github.com/wercker/dlayer) (Stats collector for Docker layers) by [@wercker](https://github.com/wercker)
* [forward2docker](https://github.com/bsideup/forward2docker) (Utility to auto forward a port from localhost into ports on Docker containers running in a boot2docker VM) by [@bsideup](https://github.com/bsideup)
* [dockramp](https://github.com/jlhawn/dockramp) (Proof of Concept: A Client Driven Docker Image Builder) by [@jlhawn](https://github.com/jlhawn)
* [portainer](https://github.com/duedil-ltd/portainer) (Apache Mesos framework for building Docker images) by [@tarnfeld](https://github.com/tarnfeld)
* [Gradle Docker plugin](https://github.com/gesellix/gradle-docker-plugin) a Docker remote api plugin for Gradle by [@gesellix](https://github.com/gesellix)
* [Docker client](https://github.com/gesellix/docker-client) a Docker remote api client library for the JVM, written in Groovy by [@gesellix](https://github.com/gesellix)
* [Dropdock](http://dropdock.io/) a framework designed for Drupal to build fast, isolated development environments using Docker.
* [Devstep](https://github.com/fgrehm/devstep) (Development environments powered by Docker and buildpacks) by [@fgrehm](https://github.com/fgrehm)
* [Lorry](https://lorry.io/) (Lorry is a docker-compose.yml validator and composer) by [@CenturyLinkLabs](https://github.com/CenturyLinkLab)
* [Dray](http://dray.it/) (Dray is an engine for managing the execution of container-based workflows. Docker Workflow Engine - UNIX pipes for Docker) by [@CenturyLinkLabs](https://github.com/CenturyLinkLab)
* [docker-do](https://github.com/benzaita/docker-do) hassle-free docker run, like `env` but for docker by [@benzaita](https://github.com/benzaita)
* [Docker osx dev](https://github.com/brikis98/docker-osx-dev) A productive development environment with Docker on OS X by [@brikis98](https://github.com/brikis98)
* [rocker](https://github.com/grammarly/rocker) extended Dockerfile builder. Supports multiple FROMs, MOUNTS, templates, etc.
* [dexec](https://github.com/docker-exec/dexec) Command line interface for running code with Docker Exec images. https://docker-exec.github.io/ written in Go.
* [crowdr](https://github.com/polonskiy/crowdr) Tool for managing multiple Docker containers (docker-compose alternative)

## Continuous Integration / Continuous Delivery

* [Drone](https://github.com/drone/drone) - Continuous integration server built on Docker and configured using YAML files.
* [Docker plugin for Jenkins](https://github.com/jenkinsci/docker-plugin/) - The aim of the docker plugin is to be able to use a docker host to dynamically provision a slave, run a single build, then tear-down that slave.
* [Captain](https://github.com/harbur/captain) - Convert your Git workflow to Docker containers ready for Continuous Delivery by [@harbur](https://github.com/harbur)
* [IBM DevOps Services](https://hub.jazz.net) - Continuous delivery using a pipeline deployment onto IBM Containers on Bluemix.
* [Watchtower](https://github.com/CenturyLinkLabs/watchtower) - Automatically update running Docker containers by [@CenturyLinkLabs](https://github.com/CenturyLinkLabs)
* [Dockunit](https://github.com/dockunit/platform) - Docker based integration tests. A simple Node based utility for running Docker based unit tests. By [@dockunit](https://github.com/dockunit)

## Deployment

* [Last.Backend](https://lastbackend.com/) - Last.Backend platform is designed for automatization of all routine work with the server and deployment of applications in one click using the visual interface.
* [Conduit](https://github.com/ehazlett/conduit) - Experimental deployment system for Docker by [@ehazlett](https://github.com/ehazlett)
* [Zodiac](https://github.com/CenturyLinkLabs/zodiac) - A lightweight tool for easy deployment and rollback of dockerized applications. By [@CenturyLinkLabs](https://github.com/CenturyLinkLabs)
* [rocker-compose](https://github.com/grammarly/rocker-compose) - Docker composition tool with idempotency features for deploying apps composed of multiple containers.
* [depcon](https://github.com/gondor/depcon) - Depcon is written in Go and allows you to easily deploy Docker containers to Apache Mesos/Marathon, Amazon ECS and Kubernetes.  By [@gonodr](https://github.com/gondor)

## Hosting for repositories (registries)

Securely store your Docker images.
* [Docker Hub](https://hub.docker.com/) (provided by Docker Inc.)
* [Quay.io](https://quay.io/) (part of CoreOS) - Secure hosting for private Docker repositories
* [Reesd](https://reesd.com/) -  Private Docker repositories and redundant storage service by [@noteed](https://github.com/noteed)

## Hosting for containers

* [Orchard](https://www.orchardup.com/) (part of Docker Inc) - Get a Docker host in the cloud, instantly.
* [StackDock](https://stackdock.com/) - Docker hosting on blazing fast dedicated infrastructure
* [Tutum](https://www.tutum.co/) - Simple hosting for your Docker containers.
* [Giant Swarm](https://giantswarm.io/) - Simple microservice infrastructure. Deploy your containers in seconds.
* [Triton](https://www.joyent.com/) - Elastic container-native infrastructure by Joyent.
* [Amazon ECS](http://aws.amazon.com/ecs/) - A management service on EC2 that supports Docker containers.
* [Google Container Engine](https://cloud.google.com/container-engine/docs/) - Docker containers on Google Cloud Computing powered by [Kubernetes][kubernetes].
* [IBM Bluemix](https://console.ng.bluemix.net/) - Run Docker containers in a hosted cloud environment on IBM Bluemix.

## Reverse Proxy

* [nginx-proxy](https://github.com/jwilder/nginx-proxy) - Automated nginx proxy for Docker containers using docker-gen by [@jwilder](https://github.com/jwilder)
* [h2o-proxy](https://github.com/zchee/h2o-proxy) - Automated H2O reverse proxy for Docker containers. An alternative [jwilder/nginx-proxy](https://github.com/jwilder/nginx-proxy) by [@zchee](https://github.com/zchee)
* [docker-proxy](https://github.com/silarsis/docker-proxy) - Transparent proxy for docker containers, run in a docker container. By [@silarsis](https://github.com/silarsis)
* [muguet](https://github.com/mattallty/muguet) - DNS Server & Reverse proxy for Docker environments. By [@mattallty](https://github.com/mattallty)
* [Træfɪk](http://traefik.github.io/) - Automated reverse proxy and load-balancer for Docker, Mesos, Consul, Etcd... By [@EmileVauge](https://github.com/emilevauge)

## Web Interface

* [Docker Registry Web](https://github.com/atc-/docker-registry-ui) (A web UI for easy private/local Docker Registry integration) by [@atc-](https://github.com/atc-)
* [DockerUI](https://github.com/crosbymichael/dockerui) (DockerUI is a web interface to interact with the Remote API.) by [@crosbymichael](https://github.com/crosbymichael)
* [DockerBoard](https://github.com/dockerboard/dockerboard) (Simple dashboards, visualizations, managements for your dockers.) by [@dockerboard](https://github.com/dockerboard)
* [Portus](https://github.com/SUSE/Portus) (Authorization service and frontend for Docker registry (v2)) by [@SUSE](https://github.com/SUSE)

## Local Container Manager

* ~~[Fig](https://github.com/docker/compose) (Fast, isolated development environments using Docker) -- http://www.fig.sh~~ --> Fig has been replaced by [Docker Compose](http://docs.docker.com/compose/), and is now deprecated. The new documentation is on the [Docker website](http://docs.docker.com/compose/).
* [Shutit](http://ianmiell.github.io/shutit/) (a tool for building and maintaining complex Docker deployments) by [@ianmiell](https://github.com/ianmiell)
* [FuGu](https://github.com/mattes/fugu) (a docker run wrapper without orchestration) by [@mattes](https://github.com/mattes)
* [Boot2Docker](https://github.com/boot2docker/boot2docker) (docker for OSX and Windows) -- http://boot2docker.io/
* [docker-vm](https://github.com/shyiko/docker-vm) (A simple and transparent alternative to boot2docker (backed by Vagrant)) by [@shyiko](https://github.com/shyiko)
* [Vessel](https://github.com/awvessel/vessel) (Vessel automates the setup & use of dockerized development environments) by [@awvessel](https://github.com/awvessel)
* [subuser](http://subuser.org) (subuser makes it easy to securely and portably run graphical desktop applications in Docker)
* [OctoHost](http://www.octohost.io/) (Simple web focused Docker based mini-PaaS server. git push to deploy your websites as needed) by [@octohost](https://github.com/octohost)
* [Dokku](https://github.com/progrium/dokku) (Docker powered mini-Heroku in around 100 lines of Bash) by [@progrium](https://github.com/progrium)
* [Ansible - manage docker containers](http://docs.ansible.com/ansible/docker_module.html)
* [Vagrant - Docker provider](http://docs.vagrantup.com/v2/docker/basics.html) a good starting point is [vagrant-docker-example](https://github.com/bubenkoff/vagrant-docker-example) by [@bubenkoff](https://github.com/bubenkoff)
* [Dray](https://github.com/CenturyLinkLabs/dray)  An engine for managing the execution of container-based workflows. http://Dray.it by [@CenturyLinkLabs](https://github.com/CenturyLinkLabs/)
* [percheron](https://github.com/ashmckenzie/percheron) (Organise your Docker containers with muscle and intelligence) by [@ashmckenzie](https://github.com/ashmckenzie)
* [Dusty](http://dusty.gc.com/) Managed Docker development environments on OS X
* [Beluga](https://github.com/cortexmedia/Beluga) Beluga a CLI to deploy docker containers on a single server or low amount of servers. By [@cortextmedia](https://github.com/cortexmedia)

## Volume management and plugins

* [Convoy](https://github.com/rancher/convoy) - an open-source Docker volume driver that can snapshot, backup and restore Docker volumes anywhere. By [@rancher](https://github.com/rancher)
* [Azure Files Volume Driver](https://github.com/ahmetalpbalkan/azurefile-dockervolumedriver) - A Docker volume driver that allows you to mount persistent volumes backed by Microsoft Azure File Service. By [@ahmetalpbalkan](https://github.com/ahmetalpbalkan)
* [Docker Unison](https://github.com/leighmcculloch/docker-unison) A docker volume container using Unison for fast two-way folder sync. Created as an alternative to slow boot2docker volumes on OS X. By [@leighmcculloch](https://github.com/leighmcculloch)
* [Netshare](https://github.com/gondor/docker-volume-netshare) A Docker volume plugin written in Go that supports mounting NFS, AWS EFS & CIFS volumes within a container. By [@gondor](https://github.com/gondor)


## Useful Images

* [Official Images from Docker Hub](https://github.com/docker-library/official-images)
* [Base Image](https://github.com/phusion/baseimage-docker) by [@phusion](https://github.com/phusion/)
* [Busybox](https://github.com/jpetazzo/docker-busybox) (with either `buildroot` or Ubuntu's `busybox-static`) by [@jpetazzo](https://github.com/jpetazzo)
* ~~[Busybox](https://github.com/progrium/busybox) (with `opkg`) by [@progrium](https://github.com/progrium)~~ [@deprecated use [docker-alpine](https://github.com/gliderlabs/docker-alpine) instead]
* [OpenWRT](http://www.zoobab.com/docker-openwrt-image) by [@zoobab](https://github.com/zoobab)
* [Phusion Docker Hub Account](https://hub.docker.com/u/phusion/)
* [passenger-docker](https://github.com/phusion/passenger-docker) (Docker base images for Ruby, Python, Node.js and Meteor web apps) by [@phusion](https://github.com/phusion)
* [docker-alpine](https://github.com/gliderlabs/docker-alpine) (A super small Docker base image *(5MB)* using Alpine Linux) by [@gliderlabs](https://github.com/gliderlabs)
* [docker-fluentd](https://github.com/kiyoto/docker-fluentd) (the Container to Log Other Containers' Logs) by [@kiyoto](https://github.com/kiyoto)
* [chaperone-docker](https://github.com/garywiz/chaperone-docker) (A set of images using the Chaperone process manager, including a lean Alpine image, LAMP, LEMP, and bare-bones base kits.)


## Dockerfile

* [Dockerfile Project](http://dockerfile.github.io/) : Trusted Automated Docker Builds. Dockerfile Project maintains a central repository of Dockerfile for various popular open source software services runnable on a Docker container.
* [Collection of Dockerfiles](https://github.com/crosbymichael/Dockerfiles) by [@crosbymichael](https://github.com/crosbymichael)
* [Dockerfile Example](https://github.com/komljen/dockerfile-examples) by [@komljen](https://github.com/komljen)
* [Dockerfile Example 2](https://github.com/kstaken/dockerfile-examples) by [@kstaken](https://github.com/kstaken)
* [Dockerfile @jfrazelle](https://github.com/jfrazelle/dockerfiles) by [@jfrazelle](https://github.com/jfrazelle) **MUST SEE** for a fully containerized desktop!

## Storing Images

* ~~[Docker Registry](https://github.com/docker/docker-registry) (Registry server for Docker (hosting/delivering of repositories and images))~~ deprecated
* [Docker Registry v2](https://github.com/docker/distribution) (The Docker toolset to pack, ship, store, and deliver content)
* [Rescoyl](https://github.com/noteed/rescoyl) (Private Docker registry) by [@noteed](https://github.com/noteed)

## Monitoring

* [Axibase Time-Series Database](http://axibase.com/products/axibase-time-series-database/writing-data/docker-cadvisor/) (Long-term retention of container statistics and built-in dashboards for Docker. Collected with native Google cAdvisor storage driver.)
* [cAdvisor](https://github.com/google/cadvisor) (Analyzes resource usage and performance characteristics of running containers. created by [@Google](https://github.com/google)
* [Datadog](https://www.datadoghq.com/) Datadog is a full-stack monitoring service for large-scale cloud environments that aggregates metrics/events from servers, databases, and applications. It includes support for Docker, Kubernetes, and Mesos.
* [Dockerana](https://github.com/dockerana/dockerana) (packaged version of Graphite and Grafana, specifically targeted at metrics from Docker.)
* [Docker-mon](https://github.com/icecrime/docker-mon) (Console-based Docker monitoring) by [@icecrime](https://github.com/icecrime)
* [InfluxDB, cAdvisor, Grafana](https://github.com/vegasbrianc/docker-monitoring) (InfluxDB Time series DB in combination with Grafana and cAdvisor) by [@vegasbrianc](https://github.com/vegasbrianc)
* [New Relic](http://newrelic.com/docker) New Relics Docker Monitoring tool
* [Prometheus](http://prometheus.io/) (Open-source service monitoring system and time series database)
* [Ruxit](https://ruxit.com/docker-monitoring/#docker_understand-performance) Monitor containerized applications without installing agents or modifiying your Run commands
* [Seagull](https://github.com/tobegit3hub/seagull) (Friendly Web UI to monitor docker daemon.) by [@tobegit3hub](https://github.com/tobegit3hub)
* [Site24x7](https://www.site24x7.com/docker-monitoring.html) Docker MOnitoring for DevOps and IT is a SaaS Pay per Host model
* [Sysdig](http://www.sysdig.org/): An open source troubleshooting tool that provides a rich set of real-time, system-level information. It has container-specific features and is very useful in Docker environments.
* [Zabbix Docker module](https://github.com/monitoringartist/Zabbix-Docker-Monitoring): Zabbix module that provides discovery of running containers, CPU/memory/blk IO/net container metrics. Systemd Docker and LXC execution driver is also supported. It's a dynamically linked shared object library, so its performance is (~10x) better, than any script solution.
* [SPM for Docker](https://github.com/sematext/spm-agent-docker) Monitoring of host and container metrics, Docker events and logs. Automatic log parser. Anomaly Detection and alerting for metrics and logs. [@sematext](https://twitter.com/sematext)
* [Zabbix Docker](https://github.com/gomex/docker-zabbix) - Monitor containers automatically using zabbix LLD feature.
* [Collecting docker logs and stats with Splunk](http://blogs.splunk.com/2015/08/24/collecting-docker-logs-and-stats-with-splunk/)

## Networking

* [Weave][weave] (The Docker network) -- Weave creates a virtual network that connects Docker containers deployed across multiple hosts.
* [Calico-Docker](http://www.projectcalico.org/getting-started/docker/) - Calico is a pure layer 3 virtual network that allows containers over multiple docker-hosts to talk to each other.

## Logging

* [LogJam](https://github.com/gocardless/logjam) (Logjam is a log forwarder designed to listen on a local port, receive log entries over UDP, and forward these messages on to a log collection server (such as logstash).) by [@gocardless](https://github.com/gocardless)
* [Docker-Fluentd](https://github.com/kiyoto/docker-fluentd): (Docker container to Log Other Containers' Logs. One can aggregate the logs of Docker containers running on the same host using Fluentd.) by [@kiyoto](https://github.com/kiyoto)
* [Logspout](https://github.com/gliderlabs/logspout) (Log routing for Docker container logs) by [@gliderlabs](https://github.com/gliderlabs)
* [SPM for Docker](https://github.com/sematext/spm-agent-docker) Monitoring of Metrics, Events and Logs implemented in Node.js. Integrated [logagent-js](https://github.com/sematext/logagent-js) to detect and parse various log formats. [@sematext](https://twitter.com/sematext)

## Deployment and Infrastructure

* [Centurion](https://github.com/newrelic/centurion): Centurion is a mass deployment tool for Docker fleets. It takes containers from a Docker registry and runs them on a fleet of hosts with the correct environment variables, host volume mappings, and port mappings. By [@newrelic](https://github.com/newrelic)
* [Clocker](https://github.com/brooklyncentral/clocker): Clocker creates and manages a Docker cloud infrastructure. Clocker supports single-click deployments and runtime management of multi-node applications that run as containers distributed across multiple hosts. It leverages [Weave][weave] for networking and [Brooklyn][brooklyn] for application blueprints. By [@brooklyncentral](https://github.com/brooklyncentral)
* [Cloud 66](http://www.cloud66.com) - Full-stack hosted container management as a service
* [deploy](https://github.com/Perennials/deploy) - Git and Docker deployment tool. A middle ground between simple Docker composition tools and full blown cluster orchestration. Declarative configuration and short commands for managing (syncing, building, running) of infrastructures of more than a few services. Able to deploy whole preconfigured server or system of services with a single line (without having to scroll the line).
* [Docket](https://github.com/netvarun/docket): Custom docker registry that allows for lightning fast deploys through bittorrent by [@netvarun](https://github.com/netvarun/)
* [Longshoreman](https://github.com/longshoreman/longshoreman): Longshoreman automates application deployment using Docker. Just create a Docker repository (or use a service), configure the cluster using AWS or Digital Ocean (or whatever you like) and deploy applications using a Heroku-like CLI tool. By [longshoreman](https://github.com/longshoreman)
* [StackEngine](http://stackengine.com/free-docker-management/): The StackEngine Free Startup Edition enables you to build, deploy and operate your Docker applications across unlimited hosts.  Free [Community](http://community.stackengine.com) support via Slack channel gets you rapid access to a community of users and developers ready to help.

## PaaS

* [Dokku](https://github.com/progrium/dokku) (Docker powered mini-Heroku in around 100 lines of Bash) by [@progrium](https://github.com/progrium)
* [Tsuru](https://github.com/tsuru/tsuru) (Tsuru is an extensible and open source Platform as a Service software.) -- https://tsuru.io/
* [Flynn](https://github.com/flynn/flynn) (A next generation open source platform as a service) -- https://flynn.io/
* [Deis](https://github.com/deis/deis) (Your PaaS, your rules) -- http://deis.io/

## Remote Container Manager / Orchestration

* [Kontena](https://github.com/kontena/kontena) (Application Containers for Masses) -- http://www.kontena.io/
* [Kubernetes][kubernetes] (Open source orchestration system for Docker containers by Google)  -- http://kubernetes.io
* [Shipyard](https://github.com/shipyard/shipyard) (Composable Docker Management) -- http://shipyard-project.com/
* [Panamax](https://github.com/CenturyLinkLabs/panamax-ui/wiki) (Docker Management for Humans) -- http://panamax.io/
* ~~[Gaudi](https://github.com/marmelab/gaudi) (Gaudi allows to share multi-component applications, based on Docker, Go, and YAM) -- http://gaudi.io/~~ project discontinued.
* [CoreOS](https://github.com/coreos) (Linux for Massive Server Deployments) -- https://coreos.com/
* [Rancher](https://github.com/rancher/rancher) (Portable AWS-style infrastructure service for Docker) -- http://rancher.com/
* [dokku-alt](https://github.com/dokku-alt/dokku-alt) (Dokku fork with Dockerfile support, database plugins, ACL and more)
* [Docker container on Mesos](https://mesosphere.com/learn/launch-docker-container-on-mesosphere/) (Docker plus Mesosphere provides an easy way to automate and scale deployment of containers in a production environment)
* [Marathon](https://mesosphere.github.io/marathon/docs/) (Marathon is a private PaaS built on Mesos. It automatically handles hardware or software failures and ensures that an app is "always on")
* [Serf](https://github.com/hashicorp/serf) (Service orchestration and management tool) by [@hashicorp](https://github.com/hashicorp)
* [Flocker](https://github.com/ClusterHQ/flocker) (Flocker is a data volume manager and multi-host Docker cluster management tool) by [@ClusterHQ](https://github.com/ClusterHQ)
* [Decking](http://decking.io/): (Decking aims to simplify the creation, organsation and running of clusters of Docker containers in a way which is familiar to developers)
* [Maestro](https://github.com/toscanini/maestro) (Maestro provides the ability to easily launch, orchestrate and manage mulitiple Docker containers as single unit) by [@tascanini](https://github.com/toscanini)
* [Citadel](http://citadeltoolkit.org/) (Citadel is a toolkit for scheduling containers on a Docker cluster)
* [CloudSlang](http://www.cloudslang.io/) (CloudSlang is a workflow engine to create Docker process automation)
* [autodock](https://github.com/prologic/autodock) (Daemon for Docker Automation) by [@prologic](https://github.com/prologic)
* [blimp](https://github.com/tubesandlube/blimp) Uses Docker Machine to easily move a container from one Docker host to another, show containers running against all of your hosts, replicate a container across multiple hosts and more. By [@defermat](https://github.com/defermat) and [@schvin](https://github.com/schvin)
* [Nomad Project] (https://www.nomadproject.io/) Easily deploy applications at any scale. A Distributed, Highly Available, Datacenter-Aware Scheduler.

## Security

* [docker-bench-security](https://github.com/docker/docker-bench-security) script that checks for dozens of common best-practices around deploying Docker containers in production. By [@docker](https://github.com/docker)
* [notary](https://github.com/docker/notary) a server and a client for running and interacting with trusted collections. By [@docker](https://github.com/docker)
* [Twistlock](https://www.twistlock.com/)  Twistlock Security Suite detects vulnerabilities, hardens container images, and enforces security policies across the lifecycle of applications.

## Service Discovery

* [docker-consul](https://github.com/gliderlabs/docker-consul) by [@progrium](https://github.com/progrium)
* [etcd](https://github.com/coreos/etcd): A highly-available key value store for shared configuration and service discovery by [@coreOS](https://github.com/coreos)
* [Docker Grand Ambassador](https://github.com/cpuguy83/docker-grand-ambassador) This is a fully dynamic docker link ambassador. + [Article](http://docs.docker.com/engine/articles/ambassador_pattern_linking/) by [@cpuguy83](https://github.com/cpuguy83)
* [confd](http://www.confd.io/): Manage local application configuration files using templates and data from etcd or consul.
* [proxy](https://github.com/factorish/proxy): lightweight nginx based load balancer self using service discovery provided by registrator. by [@factorish](https://github.com/factorish)


# Slides

* [Docker Slideshare Account](http://www.slideshare.net/dotCloud)
* [Docker Security](http://www.slideshare.net/jpetazzo) with [@jpetazzo](https://github.com/jpetazzo)

# Videos

## Main Account

* [Docker Youtube Account](https://www.youtube.com/user/dockerrun)
* [CenturyLink Labs Docker Interviews](https://www.youtube.com/playlist?list=PL_q4Fk7SVBCIjyuCBFBItXnzGI3qBa2L1)
* [YLD Event](https://www.youtube.com/channel/UCvksXSnLqIVM_uFB7xyrsSg/videos) Conference about *containers*!!! [@YLDio](https://twitter.com/YLDio)

## Useful videos

* [Docker for Developers](https://www.youtube.com/watch?v=FdkNAjjO5yQ) (54:26) by [@jpetazzo](https://github.com/jpetazzo)  <== Good introduction, context, demo
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
* [Extending Docker with Plugins](https://vimeo.com/110835013) (15:21)
* [Logging on Docker: What You Need to Know](https://vimeo.com/123341629) (51:27)
* [Container Hacks and Fun Images](https://www.youtube.com/watch?v=1qlLUf7KtAw) by [@jfrazelle](https://github.com/jfrazelle) @ DockerCon 2015 (**MUST WATCH VIDEO**: 38:50)
* [From Local Docker Development to Production Deployments](https://www.youtube.com/watch?v=7CZFpHUPqXw) by [@jpetazzo](https://github.com/jpetazzo) @ AWS re:Invent 2015


# Interesting Twitter Accounts

* [Docker](https://twitter.com/docker)
* [CenturyLink Labs](https://twitter.com/CenturyLinkLabs)
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
* [Jessie Frazelle](https://twitter.com/frazelledazzell) Work @docker and uses full containerized desktop, lots of fun.

[weave]: https://github.com/weaveworks/weave
[brooklyn]: https://brooklyn.incubator.apache.org/
[kubernetes]: http://kubernetes.io
