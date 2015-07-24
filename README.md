# admin-portal

### this is an admin portal to show some metrics about your foundation.

## Running locally

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
