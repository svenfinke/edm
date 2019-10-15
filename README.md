[![Travis (.org)](https://img.shields.io/travis/svenfinke/edm?style=flat-square)](https://travis-ci.org/svenfinke/edm) 
[![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate issues](https://img.shields.io/codeclimate/issues/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate technical debt](https://img.shields.io/codeclimate/tech-debt/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate coverage](https://img.shields.io/codeclimate/coverage/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm)

# External Dependency Manager

**edm** is supposed to help you to manage external dependencies by providing a file which contains all deps. These can easily be downloaded into the project directory. A dependency can be a binary, a script or a config file. Fetching an zip-archive that will be automatically unzipped is supported as well.

To ensure a basic level of security, you can provide hash to check the file integrity after it has been downloaded. If the file has been altered or replaced, the download will fail and you will receive a corresponding error message.

# Usage

 `edm` is a CLI tool that will provide nested commands with their own arguments and flags. Type `edm help` for additional information. To find help for a specific command, type `edm [command] help`. This will provide you with a list and descriptions for flags, arguments and/or subcommands that are available.
 
## Init
 
`edm init` will generate an empty config file for you that you can extend to define your dependencies. The default file is `.edm.yaml`.

## Fetch

`edm fetch` will download all the dependencies and put them into the target directory.

# Configuration

## Types

The types identify the kind of file that your are fetching. Depending on this type, some additional actions may be done. A zip file might be unarchived and the contents will be moved into the target, or a binary is made executable. These are the types that are available right now:

| type | description | additional config |
|---|---|---|
| default | Downloads file via http | - |
| binary | Download file via http and make it executable | - |
