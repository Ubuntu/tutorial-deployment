# tutorial-deployment
Serve and help to generate tutorial deployment for https://tutorials.ubuntu.com

[![Test Status](https://travis-ci.org/ubuntu/tutorial-deployment.svg?branch=master)](https://travis-ci.org/ubuntu/tutorial-deployment)

Those couple of tools are used in conjunction with https://github.com/canonical-websites/tutorials.ubuntu.com to generate a tutorial website. Those can be written in Markdown or google doc, [using the claat google's library](https://github.com/googlecodelabs/tools).

## Important note
We are using currently [a fork](https://github.com/didrocks/codelab-ubuntu-tools) of the claat tools as our fixes for the markdown parser are getting reviewed and merged by the Google team.

As this is a more robust deployment procedure, some structural changed were needed and are in progress [there](https://github.com/didrocks/tutorials.ubuntu.com/tree/reformat-tooling). This tool works exclusively with that branch.

## Generate
The `generate` command will generate tutorials in **HTML**, using [Polymerjs](https://www.polymer-project.org/), to be compatible with the aforementioned tutorial source code.

It fetches in well-known places the code lab list and sources (both in **google doc** or **markdown** format), the general events and categories metadata, to generate the desired output and API files.

Every default directories will be detected by the tool if present in the tutorial directories. Arguments and options can tweak this behavior.

## Serve
The `serve` command will generate the same code lab content generated on the fly in a temporary directory, but also install watchers on local source files (code lab markdown file or any referenced local images).

Any save on any of those files will retrigger the corresponding code lab build and API generation, serve by this local HTTP web server (default port is **8080**)

Changes are all done in temporary files and not destructive on the tutorial repository. Note that source and web server paths can be overridden as for the generate command.
