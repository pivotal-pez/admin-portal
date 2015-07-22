# admin-portal

### this is an admin portal to show some metrics about your foundation.

## Running locally

```

# install the wercker cli
$ curl -L https://install.wercker.com | sh

#copy the sample env config file
$ cp myenv.example myenv # fill in your details in the newly created myenv file

# make sure a docker host is running
$ boot2docker up && $(boot2docker shellinit)

# run the app locally using wercker magic
$ ./runlocaldeploy myenv

$ echo "open ${DOCKER_HOST} in your browser to view this app locally"

```
