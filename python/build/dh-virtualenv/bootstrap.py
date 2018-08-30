#!/usr/bin/env python3

import os
import sys
import subprocess
import glob

REPO_PATH = 'git-repo'

def git_clone(url):
    r = subprocess.run(['git', 'clone', url, REPO_PATH])

    if r.returncode != 0:
        print("[COUT] Git clone error: Invalid argument to exit", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return False
    
    print()
    return True

def git_upload(upload):
    file_name = glob.glob('*.deb')[0]
    r1 = subprocess.run(['curl', '-XPUT', '-d', '@' + file_name, upload])
    if r1.returncode != 0:
        print("[COUT] upload error", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return False
    print()

    return True

def build():
    r = subprocess.run('cd {}; yes | mk-build-deps -ri; dpkg-buildpackage -us -uc -b'.format(REPO_PATH), shell=True)

    if r.returncode != 0:
        print("[COUT] build error", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return False
    return True

def parse_argument():
    ret = {}
    gitCloneUrl = os.environ.get('GIT_CLONE_URL', None)
    if not gitCloneUrl:
        return {}
    ret['git_clone_url'] = gitCloneUrl


    gitUploadUrl = os.environ.get('GIT_UPLOAD_URL', None)
    if not gitUploadUrl:
       return {}
    ret['git_upload_url'] = gitUploadUrl

    return ret

def main():
    argv = parse_argument()

    git_clone_url = argv.get('git_clone_utl')
    if not git_clone_url:
        print("[COUT] The git-clone-url value is null", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return

    git_upload_url = argv.get('git_upload_utl')
    if not git_upload_url:
        print("[COUT] The git-upload-url value is null", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return

    if not git_clone(git_clone_url):
        return
    if not build():
        return

main()