#!/usr/bin/env python3

import os
import sys
import builder

def main():
    envs_list = ['GIT_CLONE_URL', 'GIT_REF', '_WORKFLOW_GIT_CLONE_URL', '_WORKFLOW_GIT_REF', 'FILES']

    envs = {}

    for env_name in envs_list:
        envs[env_name] = os.environ.get(env_name)
    
    try:
        if builder.Builder(envs).run():
            print("BUILD SUCCEED.", file=sys.stdout)
        else:
            print("BUILD FAILED.", file=sys.stdout)
            exit(1)
    except Exception as e:
        print("BUILD FAILED: %s" % str(e), file=sys.stderr)
        exit(1)

main()