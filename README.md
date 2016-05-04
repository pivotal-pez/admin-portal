# admin-portal

### this is an admin portal to show some metrics about your foundation.

## Deploying a release binary to cloud foundry

```
# go to:
# https://github.com/pivotal-pez/admin-portal/releases/latest
# and download adminportal.tgz

$ tar -xvzf adminportal.tgz

$ cd adminportal

$ cf login -a [cf.api.domain] -u [cfadminuser] -p [adminuserpass] -o
[mytargetorg] -s [mytargetspace] --skip-ssl-validation

# this will setup a user provided service containing foundation api url and user
information.
# the user will need to have uaa.admin and cloudcontroller.admin privileges.
$ cat cups.txt | sh

$ cf push adminportal

```


## Running locally for development

```

# install the wercker cli
$ curl -L https://install.wercker.com | sh

#copy the sample env config file
$ cp myenv.example myenv # fill in your details in the newly created myenv file

#copy the sample wercker local deployment manifest
$ cp wercker_local_deploy.example wercker_local_deploy.yml # fill in your details in the newly created file


#copy the sample vcap_services definition
$ cp vcap_services_template.json.example vcap_services_template.json # fill in your details in the newly created file



# make sure a docker host is running
$ boot2docker up && $(boot2docker shellinit)

# run the app locally using wercker magic
$ ./runlocaldeploy myenv

$ echo "open ${DOCKER_HOST} in your browser to view this app locally"

```

## SSL verification

If you're running in with a self signed cert (pcfdev, bosh lite etc.) run the
following to disable certificate validation.

```
cf set-env adminportal SKIP_SSL_VERIFY true
```
