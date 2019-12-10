#!/bin/sh -x

gcloud auth login # expect browser to pop-up for interactive login

gcloud config set project gotrainingxebia # or your own <gcloud-project>

gcloud app deploy app.yaml  --quiet --version 1

