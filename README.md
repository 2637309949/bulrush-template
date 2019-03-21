# Bulrush Framework

![Bulrush flash](./assets/flash.jpg)

## Usage

### Project structure

    bulrush-template
    │
    ├── addition        # global reference module
    ├── assets          # assets
    │   └── public
    │       ├── apidoc  # apidoc assets
    │       └── upload  # upload assets
    ├── bin             # bin execScript 
    ├── binds           # bind models for jsonMap
    ├── build           # files assets after build
    │   ├── assets
    │   │   └── public
    │   │       ├── apidoc
    │   │       └── upload
    │   ├── conf
    │   └── logs
    ├── conf            # project conf
    ├── logs            # logs files
    ├── models          # project models defined
    ├── routes          # http routes
    ├── services        # services layers
    ├── tmp             # tmp file for fresh
    ├── utils           # utils tools
    └── vendor          # dependence listing

#### For Dev
```shell
$ make -f Makefile.dep
$ make -f Makefile.dev
```

#### For Apidoc
```shell
$ make -f Makefile.api
```

#### For Prod
```shell
$ make
```

## MIT License
Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.