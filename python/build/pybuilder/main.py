#!/usr/bin/env python

import os
import builder
import sys

def main():
    envs_list = ['GIT_CLONE_URL', 'GIT_REF', '_WORKFLOW_GIT_CLONE_URL', '_WORKFLOW_GIT_REF', 'ENTRY_FILE']

    envs = {}
    for env_name in envs_list:
        envs[env_name] = os.environ.get(env_name)
    try:
        if builder.Builder(envs).run():
            print >> sys.stdout, "BUILD SUCCEED"
        else:
            print >> sys.stdout, "BUILD FAILED"
            exit(1)
    except Exception as e:
        print >> sys.stderr, "BUILD FAILED: %s" % str(e)
        exit(1)
main()