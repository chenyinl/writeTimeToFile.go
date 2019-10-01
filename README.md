# Write File in Append Mode
![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)
![version](https://img.shields.io/badge/version-0.0.1-green.svg)
> Example of Golang on how to open a file and write the file at the end.

This code opens a file, move the pointer to the end and write time at the end..

## Table of Contents
- [Background](#background)
- [Install](#install)
- [Deployment](#deployment)
- [Usage](#usage)
- [Output](#output)
- [Contributing](#contributing)
- [License](#license)
- [Reference](#reference)

## Background
When we are running a CRON job, we sometimes needs to know when the last time the job was executed. Even though the job was scheduled, but sometimes the job run longer then expected, or there was a problem from third party so the job was not successfully. So I wrote this code to log down when the last job was done.

## Install
Install like regular Golang file.

## Deployment
Deploy like regular Golang work.

## Usage
If the file was not compiled, we can just run like:
```
> go run writeFile.go
```
## Output
Below is an example of how the file looks like after being called for 16 times.
```
2019-07-18T15:27:20-07:00
2019-07-18T15:51:15-07:00
2019-07-18T16:02:02-07:00
2019-07-18T16:06:50-07:00
2019-07-18T16:08:14-07:00
2019-07-18T16:12:09-07:00
2019-07-18T16:12:13-07:00
2019-07-18T16:12:18-07:00
2019-07-18T16:19:24-07:00
2019-07-18T16:19:39-07:00
2019-07-18T16:32:12-07:00
2019-07-18T16:33:21-07:00
2019-07-18T16:33:42-07:00
2019-07-18T16:33:44-07:00
2019-07-18T16:36:24-07:00
2019-10-01T15:42:09-07:00

```
## Files
writeFile.go is the only file.

## Contributing
PRs accepted.

## License

Private © 

## Reference
```
https://gobyexample.com/writing-files
```
