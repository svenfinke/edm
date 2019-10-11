[![Travis (.org)](https://img.shields.io/travis/svenfinke/edm?style=flat-square)](https://travis-ci.org/svenfinke/edm) 
[![Code Climate maintainability](https://img.shields.io/codeclimate/maintainability/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate issues](https://img.shields.io/codeclimate/issues/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate technical debt](https://img.shields.io/codeclimate/tech-debt/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm) 
[![Code Climate coverage](https://img.shields.io/codeclimate/coverage/svenfinke/edm?style=flat-square)](https://codeclimate.com/github/svenfinke/edm)
[![Codacy branch grade](https://img.shields.io/codacy/grade/01cfe390e56c47f2aa61273b3a2061a1/master?style=flat-square)](https://app.codacy.com/manual/svenfinke/edm/dashboard)

# External Dependency Manager

**edm** is supposed to help you to manage external dependencies by providing a file which contains all deps. These can easily be downloaded into the project directory. A dependency can be a binary, a script or a config file. Fetching an zip-archive that will be automatically unzipped is supported as well.

To ensure a basic level of security, you can provide hash to check the file integrity after it has been downloaded. If the file has been altered or replaced, the download will fail and you will receive a corresponding error message.

# Usage

 