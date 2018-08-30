#!/usr/bin/env python3

import os
import builder
import sys

def main():
    envs_list = ['GIT_CLONE_URL', 'GIT_REF', '_WORKFLOW_GIT_CLONE_URL', '_WORKFLOW_GIT_REF']

    envs = {}
    for env_name in envs_list:
        envs[env_name] = os.environ.get(env_name)