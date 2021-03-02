#!/usr/bin/env bash
set -e

apt-get update
apt-get install -y --force-yes psmisc inotify-tools bc
echo 'export TERM=xterm' >> /root/.bashrc
