#!/usr/bin/env python3

import os
import builder
import sys

def main():
    env_list = [
        'GIT_CLONE_URL', 'GIT_REF',
        "_WORKFLOW_GIT_CLONE_URL", "_WORKFLOW_GIT_REF", "_WORKFLOW_FLAG_CACHE"
        'BEARS', 'FILES']
    envs = {}
    for env_name in env_list:
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
