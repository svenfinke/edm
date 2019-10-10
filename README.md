# External Dependency Manager

**edm** is supposed to help you to manage external dependencies by providing a file which contains all deps. These can easily be downloaded into the project directory. A dependency can be a binary, a script or a config file. Fetching an zip-archive that will be automatically unzipped is supported as well.

To ensure a basic level of security, you can provide hash to check the file integrity after it has been downloaded. If the file has been altered or replaced, the download will fail and you will receive a corresponding error message.

# Usage

 