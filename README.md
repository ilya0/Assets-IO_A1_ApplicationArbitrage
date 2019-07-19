# Assets-IO_A1_ApplicationArbitrage

Application Arbitrage is a template created specifically to help with quickly demonstrating the capabilities of blockchain and its ability to interface with mulitple applications and users, while keeping a central authority to various data that could be saved in the blockchain.



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

## Deployment and Using the App



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


## Running tests

No test need to be ran at this time




## Built With

* Angular 7
* Node.js
* oracle blockchain

##Unsolved problems, etc.

* CORs issue is currently underdevelopment, and in order to use the front end cors plugin needs to be used in chrome
* could use a beutified layout
* styling

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

1.0.0

## Authors

* **Ilya Osovets** -- [ilya0](https://github.com/ilya0)


See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* App Dev Oracle Hub
