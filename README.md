# travis-test
Some tests of Travis CI + GitHub collaboration

## Deploy notes
See [docs](https://docs.travis-ci.com/user/deployment/releases/) for some description on deployment

Needs _travis_ CLI tool. Link is in the docs. For installation needed _ruby-dev_

Simple _travis setup releases_ didn't work for me - by default this tool works with travis-ci.org while my repos are registered at travis-ci.com. Thus the following actions were done:
* travis login --com (logging into your account)
* travis setup releases --com --force (--force as I had basic .travis.yml in place)
