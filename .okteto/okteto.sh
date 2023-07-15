#!/bin/sh -eu

helm upgrade --install --wait morning-night-dream-oidc ./charts/oidc --namespace=oidc-morning-night-dreamer
