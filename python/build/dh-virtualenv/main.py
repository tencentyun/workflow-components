#!/usr/bin/env python3

import os
import sys
import builder

def main():
    envs_list = ['GIT_CLONE_URL', 'GIT_REF', "_WORKFLOW_GIT_CLONE_URL", "_WORKFLOW_GIT_REF"]
    envs = {}

    for env_name in envs_list:
        envs[env_name] = os.environ.get(env_name)
    
    try:
        if builder.Builder(envs).run():
            print("build succeed", file=sys.stdout)
        else:
            print("build failed", file=sys.stdout)
            exit(1)
    except Exception as e:
        print("build failed: %s" % str(e), file=sys.stderr)
        exit(1)
main()