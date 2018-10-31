=============================
GoLang resources and snippets
=============================

As I walk through getting GOLang knowledge, I'll list here all the useful stuff I find on my journey.

Documentations
==============

Many documentation resources.

The official documentations
---------------------------

- https://golang.org/
- https://golang.org/doc


The non official documentations
-------------------------------

- Todd McLoad's docs : https://drive.google.com/drive/u/0/folders/0B22KXlqHz6ZNfjNXTzk1U3JHUkJ6VjJ3dnJKNzVtNjRUM3Q2WFNqWGI2Q3RadERqUlVrOEU


You need to achieve following steps in order to be fine ready for developping with Go.

Install GOLang
--------------

Depending on your OS, you'll need to follow the steps describe by the `official documentation`_.

.. _official documentation: https://golang.org/doc/install


Build you workspace
-------------------

The proper way to build you environment is explained in the official documentation

https://golang.org/doc/code.html

Some tooling can help you building a workspace for a project like `virtualgo`_.

Even so, you'll have to set a couple of environement variables.

Recognize your current Go env
.............................

The ``go env`` command helps you to know your current GoLang environement, listing all the variables recognnized by the go interpreter/compiler.

.. code-block:: bash

    $ go env

    GOARCH="amd64"
    GOBIN=""
    GOCACHE="/Users/vdagoury/Library/Caches/go-build"
    GOEXE=""
    GOFLAGS=""
    GOHOSTARCH="amd64"
    GOHOSTOS="darwin"
    GOOS="darwin"
    GOPATH="/Users/vdagoury/go"
    GOPROXY=""
    GORACE=""
    GOROOT="/usr/local/Cellar/go/1.11.1/libexec"
    GOTMPDIR=""
    GOTOOLDIR="/usr/local/Cellar/go/1.11.1/libexec/pkg/tool/darwin_amd64"
    GCCGO="gccgo"
    CC="clang"
    CXX="clang++"
    CGO_ENABLED="1"
    GOMOD=""
    CGO_CFLAGS="-g -O2"
    CGO_CPPFLAGS=""
    CGO_CXXFLAGS="-g -O2"
    CGO_FFLAGS="-g -O2"
    CGO_LDFLAGS="-g -O2"
    PKG_CONFIG="pkg-config"
    GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/5w/jwq8mykj5k3b5t1mg21dp0h80000gn/T/go-build832488294=/tmp/go-build -gno-record-gcc-switches -fno-common"

The important variables to notice are :
- GOROOT : the GOLang installation dir.
- GOPATH : you workspace dir (with all your projects)


Tools for GO
============

dep
---

The dependency manager tool.
Even if 'go get' is here and achieves its job quite well, this tool is part of the mandatory tool set you need.
So it's strongly recommended to install this one day 1.

You'll need to have set your environment first.

- Project : https://github.com/golang/dep
- Documentation https://golang.github.io/dep/

Insallation :

.. code:: bash

    $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh


virtualgo
---------

The 'virtualenv' tool for Go. It brings isolation for your workspace.

- https://github.com/GetStream/vg

.. code:: bash

    $ brew install bindfs

    $ go get -u github.com/GetStream/vg

    $ vg init  # initial creation of workspace 

    # Now all commands will be executed from within the example workspace

    (example) $ go get github.com/pkg/errors # package only present in workspace
    (example) $ vg ensure  # installs the dependencies of the example project using dep
    (example) $ vg deactivate


GVM
---

The (not mandatory at all) GO installer.

- https://github.com/moovweb/gvm


.. code:: bash

    # Install
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    
    # List Go versions
    gvm list

    # List installable Go versions
    gvm listall

    # Remove all
    gvm implode

