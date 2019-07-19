# Assets-IO_A1_ApplicationArbitrage

Application Arbitrage is a template created specifically to help with quickly demonstrating the capabilities of blockchain and its ability to interface with mulitple applications and users, while keeping a central authority to various data that could be saved in the blockchain.

### Overview

Since the begining of time people have wondered, what music are people listening to right now?

![](https://github.com/ilya0/Project-4/blob/master/ERD/enstein.jpg)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

![](https://github.com/ilya0/Project1/blob/master/readme%20files/Main%20setup%20front%20page.png)

Link to [Item](http://ilya0.github.io/Project1)

Link to [Trello](https://trello.com/b/0pPnXkD1/project1-pvp)



### Prerequisities
------

What things you need to install the software and how to install them

```
Give examples
```

### Installing
------
The application is a multipart application which has three parts. 

Front end - Interfaces with the middleware inorder to show all the data 

Middleware - interfaces with the blockchain and holds all the api interfaces

Blockchain - records all the data for immutiblity 

Steps should be as follows 
#### Creating blockchain instance
This guide is made with the assumtion that a blockchain instance has already been created. In order to create a blockchain, if one is not available please follow the creating a blockchain instance guide.

#### Checking the blockchain for runtime

1. Access block chain from www.cloud.oracle.com
2. Click on signin on the top bar
3. Sign into the tenancy 
4. Sign in with your SSO in order to access the environement
5. Double check your instance is running in the list
6. In order to check any details of the instance click on the hamburger menu to the left of the instance



#### Starting the Middleware

```
cd UI/Middleware
npm install
npm start

```


#### Starting the Frontend

```
cd UI/Frontend
npm install
npm start
open http://localhost:4200

```
Frontend test page will open on localhost:4200



End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment and Using the App

Add additional notes about how to deploy this on a live system

## Built With

* Dropwizard - Bla bla bla
* Maven - Maybe
* Atom - ergaerga

##Unsolved problems, etc.

* needs description in profiles
* could use a beutified layout
* search functionality
* styling

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Ilya Osovets** -- [ilya0](https://github.com/ilya0)


See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone who's code was used
* Inspiration
* etc
