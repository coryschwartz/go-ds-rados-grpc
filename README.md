# rados-ds grpc proxy

# WARNING

This is proof-of-concept quality code.
Don't use it in production.

# What is it?

Proxy datastore with a backing Ceph/Rados cluster

# Why?

Rados supports VERY large storage on the order of multiple Petabytes. It is replicated, redundant,
authenticated, performant, and has many features that are desirable in a storage layer.

# Demonstration

The following demonstrates using a remote datastore on Lotus.
In this example, lotus is setup to connect to an instance of this datastore for the cold splitstore.

![gif](https://gist.github.com/coryschwartz/43e845517bb9b228b005e79ee6490868/raw/6950a27789d17bbd47c19501cae3fc6331ffa282/small.gif)

[mkv](https://gist.github.com/coryschwartz/43e845517bb9b228b005e79ee6490868/raw/773609722c26d90eed24fe9d412426f11e90e593/video.mkv)

# Building

RADOS apis vary somewhat between ceph versions. (https://docs.ceph.com/en/latest/releases/)
To build for the appropriate version, you must have the correct library and header files installed.
Additionally, go tags are used to control the ceph version built. See 
(go-ceph)[https://github.com/ceph/go-ceph] for details. At a minimum, you must install librados-dev 
for the rados version you want to support.

Provided that you have dependencies installed, build as usual with `go build`

# configuration

Configuration is done by environment variables.

A lot of the configuration comes from rados, which is provided as a configuration file. Ceph is an 
authenticated storage system. You'll need to create a rados pool and a user with permission to write 
to the pool. 

By default, rados is configured in /etc/ceph, which is the typical location for ceph configuration 
files.
