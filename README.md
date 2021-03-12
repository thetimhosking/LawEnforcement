# LawEnforcement
This is a repository for several microservices that support law enforcement. 

The microservices belong in two groups:
1. Enterprise domains - used by the application domains and are more general in application
1. Applications domains - implement features and use cases that are specific to law enforcement

## Enterprise domains
Enterprise domains offer common features that may be used by several application domains.

The enterprise domains are:
1. Party: the identity of individual persons and organisations
2. Object: both tangible and intangible "things" that can be involved in incidents or may be used for evidence
3. Location: the geographical place, or address, or electronic space where an event occurs
4. Event: the occurrence at a given time that is of interest, involving people and/or organisations, objects, and locations
5. Activity: the things the organisation does in response to events, or to provide a service
6. Resource: facilities, assets, information, etc. that may be used in many business processes
7. Service: a business service that the organisation provides; not to be confused with a software service that might support or enable business services. For example, a protection service may be provided for VIPs, protected witnesses in a matter, informants, etc. A missing person service may be provided to find missing persons and to support their relatives and friends

## Application Domains
The application domains are:
1. Incident Response
2. Custody Management
3. Investigation Management
4. Intelligence Gathering
5. Legal Process

## Delivery
The intention is to deliver these microservices in Golang because:
1. It is fast
1. It runs on many platforms
1. Provides many libraries
1. Is simple to use

