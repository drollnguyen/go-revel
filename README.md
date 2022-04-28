# Getting Started

## INSTALL GO
    Before you can use Revel, first you need to install Go.

    See the official Go installation guide.
    
    <ul>
        <li><a href="https://github.com/golang/go/wiki/Ubuntu">Ubuntu</a></li>
        <li><a href="https://golang.org/doc/install#windows">Windows</a></li>
    </ul>

## SET UP YOUR GOPATH
    If you have not created a GOPATH as part of the installation, do so now. The GOPATH is a directory where all of your Go code will live. Here is one way of setting it up:

    <ol>
        <li>Make a directory: <code class="language-plaintext highlighter-rouge">mkdir ~/gocode</code></li>
        <li>Tell Go to use that as your GOPATH: <code class="language-plaintext highlighter-rouge">export GOPATH=~/gocode</code></li>
        <li>Save your GOPATH so that it will apply to all future shell sessions: <code class="language-plaintext highlighter-rouge">echo export GOPATH=$GOPATH &gt;&gt; ~/.bash_profile</code></li>
    </ol>

    Now your Go installation is complete.

## INSTALL GIT AND HG
    Git and Mercurial are required to allow go get to clone various dependencies.

    <ul>
        <li><a href="http://git-scm.com/book/en/Getting-Started-Installing-Git">Installing Git</a></li>
        <li><a href="https://www.mercurial-scm.org/downloads">Installing Mercurial</a></li>
    </ul>

## GET THE REVEL FRAMEWORK
    To get the Revel framework, run

    go get github.com/revel/revel


# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

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

