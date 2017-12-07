# Welcome to Revel-Api-Seed

*Warning : this project is currently under active development, use it with caution. Contributions and PR are welcome!* 

Revel-Api-Seed aims to be a complete backend but light starting point for new full-stack projects, using these technologies

* [Revel](https://revel.github.io) : A high-productivity web framework for the [Go language](http://www.golang.org/) as API backend.
* [GORM](https://github.com/jinzhu/gorm) : The fantastic ORM library for Golang, aims to be developer friendly.
* [Ionic](https://ionicframework.com) ( [Angular](https://angular.io) ) : The beautiful, free and open source mobile SDK for developing native and progressive web apps with ease.

... and providing out-of-the-box

* Authentication endpoints (`/login`, `/register`) and logic
* Usage of JWT Tokens
* CRUD API example
* GORM Transactions
* Fully bootstraped and backend-compatible Ionic application

... and soon more !
* WebSockets


### Start the web server:

    # Get the seed
    go get github.com/obitux/revel-api-seed
    
    # Install Ionic dependencies and build the project so it will be available in `prod` mode at `/`
    cd $GOPATH/src/github.com/obitux/revel-api-seed/examples/ionic/app
    npm install
    ionic cordova build browser --prod

    # Run the server
    revel run github.com/obitux/revel-api-seed

    # Run Ionic in dev mode (at `http://localhost:4200`)
    cd $GOPATH/src/github.com/obitux/revel-api-seed/examples/ionic/app
    ionic serve

Go to http://localhost:9000/ and you'll see the Ionic demo app in production mode.


## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    examples/         Frontend examples ready-to-work

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

