# install

Easily just go to releases page and download the one that fits your os

# What is proto fetch?

<hr/>
Proto fetch is individually built for managing shared resources between project parts. For example shared .proto files
between backend and android application. With proto fetch you can:
<li>fetch shared resources from private or public repositories</li>
<li>complete support of github and other git based services</li>
<li>support of after fetch and before fetch commands</li>
<li>support of environmental variables in output directory or after and before fetch commands</li>

Fore more information checkout this link: [https://medium.com/namely-labs/how-we-build-grpc-services-at-namely-52a3ae9e7c35#a944](https://medium.com/namely-labs/how-we-build-grpc-services-at-namely-52a3ae9e7c35#a944)

# Why proto fetch?

<hr/>
proto fetch is a great application for managing shared resources between application parts.
Lets just define a part of application as parts that don't have direct access to each others resources.
For example android application don't have access to backend resources. Managing
these types of shared resources is a very boring procedure that must be done 
everytime that an update occurs. But with proto fetch, this procedure is easily removed and you can
just focus on your coding.

# How to use proto fetch?

<hr />
Everything you need to know is just about proto-fetch.yaml file. we'll describe options one by one:

## base-url & repository-owner & repository options

If we assume a repository like this: [github.com/zytell3301/proto-fetch](https://github.com/zytell3301/proto-fetch),
base url would be github.com and repository-owner is zytell3301 and repository option is proto-fetch.

## auth-token

If you are using resources placed in a private repository, you must set this option, so the application will authorize
itself to server to get protected resources.

## output-dir

This is the path that all the shared resources will go

## files option

In this option you determine that which files must get fetched from repository. Consider the repository a directory. The
path is starts from the root of the repository. Putting a ./ at the beginning of the path is an optional argument. Also,
you can override the output-dir per every file like this: path/to/file->path/to/override/path. If you don't determine
the second path (->path/to/override/path), the file will go to output-dir. Please pay attention that directory structure
will be similar to the structure in the repository itself. For example if your resource is in example folder, an example
folder will be created in the determined path.


## before-fetch-commands & after-fetch-commands options
These options are commands that will be executed before and after fetching ALL resources.
For example, you can cleanup the proto directory before fetching files and after fetching proto
files, you can again compile the updated proto files.

## env-variables

This is where you define your environmental variables. These variables can be
used in output-dir or override paths or even after fetch and before fetch commands.
The format is like: VARIABLE_NAME=variable value. <br />
It is highly recommended defining a variable that contains absolute path of you project
and then using it for commands and paths. This will prevent unwanted removals because of
giving paths wrongly.