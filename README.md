# drone-dockerhub

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-dockerhub/status.svg)](http://beta.drone.io/drone-plugins/drone-dockerhub)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-dockerhub/coverage.svg)](https://aircover.co/drone-plugins/drone-dockerhub)
[![](https://badge.imagelayers.io/plugins/drone-dockerhub:latest.svg)](https://imagelayers.io/?images=plugins/drone-dockerhub:latest 'Get your own badge on imagelayers.io')

Drone plugin to trigger a DockerHub remote build

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-dockerhub <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "token": "be579c82-7c0e-11e4-81c4-0242ac110020",
        "repo": "foo/bar"
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-dockerhub <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "token": "be579c82-7c0e-11e4-81c4-0242ac110020",
        "repo": "foo/bar"
    }
}
EOF
```
