# basic directory structure for a microservice

easily choose what kind of api's you want to build 
- rest,grpc .....
or even implement both

cmd-
    contains main intialize everthing there
    spwan go routine for all kind of apis seperately

internal-
    api-
        contains implementation of different kind of apis
    
    config-
        contains configuration which is required by the service can be 
        database hosts,ports,passwords or keys, maybe more more configuraion 
        factors

    db-
        contains methods to connect to the databases and have connection refrences
        to them example postgres,redis...

    error-
        custom errors related to you organization or service

    event-
        contains event processors(consumers ) or event producers

    logger-
        all of logging lies here wether it's an external logger or just
        simple logging

    models-
        all the models lie here

    respository-
        repository is the abstraction over databases 

    service-
        various services which can be required or not
        for example starting a grpc service

    util-
        all the funciton generally which are required all over the 
        project or maybe which can't be categories

    
