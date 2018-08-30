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

def build(entry_file):
    r = subprocess.run(['nuitka', '--recurse-all', '{}/{}'.format(REPO_PATH, entry_file)])

    if r.returncode != 0:
        print("[COUT] nuitka error", file=sys.stderr)
        return False

    return True

def get_pip_cmd(version):
    if version == 'python3' or version == 'py3k':
        return 'pip3'

    return 'pip'

def get_python_cmd(version):
    if version == 'python3' or version == 'py3k':
        return 'python3'

    return 'python'

def init_env(version):
    subprocess.run([get_pip_cmd(version), 'install', 'nuitka'])

def validate_version(version):
    valid_version = ['python', 'python2', 'python3', 'py3k']
    if version not in valid_version:
        print("[COUT] Check version failed: the valid version is {}".format(valid_version), file=sys.stderr)
        return False

    return True

def setup(path, version='py3k'):
    file_name = os.path.basename(path)
    dir_name = os.path.dirname(path)
    r = subprocess.run('cd {}; {} {} install'.format(dir_name, get_python_cmd(version), file_name))
    if r.returncode != 0:
        print("[COUT] install dependences failed", file=sys.stderr)
        return False
    
    return True

def pip_install(file_name, version='py3k'):
    r = subprocess.run([get_pip_cmd(version), 'install', '-r', file_name])

    if r.returncode != 0:
        print("[COUT] install dependences failed", file=sys.stderr)
        return False

    return True

def parse_argument():
    ret = {}

    gitCloneUrl = os.environ.get('GIT_CLONE_URL', None)
    if not gitCloneUrl:
        return {}
    ret['git_clone_url'] = gitCloneUrl

    entryFile = os.environ.get('ENTRY_FILE', None)
    if not entryFile:
        return {}
    ret['entry_file'] = entryFile

    gitUploadUrl = os.environ.get('GIT_UPLOADURL', None)
    if not gitUploadUrl:
        return {}
    ret['git_upload_url'] = gitUploadUrl

    version = os.environ.get('VERSION')
    ret['version'] = version

    return ret

def main():
    argv = parse_argument()

    git_clone_url = argv.get('git_clone_url')
    if not git_clone_url:
        print("[COUT] The git-clone-url value is null", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return
    
    git_upload_url = argv.get('git_uplaod_url')
    if not git_upload_url:
        print("[COUT] The git-upload-url value is null", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return 

    entry_file = argv.get('entry_file')
    if not entry_file:
        print("[COUT] The entry_file value is null", file=sys.stderr)
        print("[COUT] CO_RESULT = false")
        return
    
    version = argv.get('version')

    if not validate_version(version):
        print("[COUT] CO_RESULT = false")
        return

    init_env(version)

    if not git_clone(git_clone_url):
        return

    for file_name in glob.glob('{}/setup.py'.format(REPO_PATH)):
        setup(file_name, version)

    for file_name in glob.glob('{}/*/setup.py'.format(REPO_PATH)):
        setup(file_name, version)

    for file_name in glob.glob('{}/requirements.txt'.format(REPO_PATH)):
        pip_install(file_name, version)

    for file_name in glob.glob('{}/*/requirements.txt'.format(REPO_PATH)):
        pip_install(file_name, version)

    if not build(entry_file):
        print("[COUT] CO_RESULT = false")
        return

    return

main()