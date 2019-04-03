import os
import sys
import subprocess
import glob

BASE_SPACE = "/root/src"

class Builder:
    def __init__(self, envs):
        if envs.get('GIT_CLONE_URL'):
            self.git_clone_URL = envs.get('GIT_CLONE_URL').rstrip('/')
            self.git_ref = envs.get('GIT_REF') or 'master'
        elif envs.get('_WORKFLOW_GIT_CLONE_URL'):
            self.git_clone_URL = envs.get('_WORKFLOW_GIT_CLONE_URL').rstrip('/')
            self.git_ref = envs.get('_WORKFLOW_GIT_REF') or 'master'
        else:
            print("environment variable GIT_CLONE_URL is required", flle=sys.stderr)
            exit(1)

        if envs.get('ENTRY_MOD'):
            self.entry_mod = envs.get('ENTRY_MOD')
        else:
            print("environment variable ENTRY_MOD is required", file=sys.stderr)
            exit(1)

        self.project_name = os.path.basename(self.git_clone_URL.rstrip('.git'))
        self.setup = []
        self.pipup = []

    def run(self):
        os.chdir(BASE_SPACE)

        return self.git_pull() and self.git_reset() and self.python_install() and self.pip_install() and self.build() 

    def git_pull(self):
        cmd = ['git', 'clone', '--recurse-submodules', self.git_clone_URL, self.project_name]
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd)

        if r.returncode != 0:
            print("Git clone failed", file=sys.stderr)
            return False

        for file_name in glob.glob('{}/setup.py'.format(BASE_SPACE)):
            self.setup.append(file_name)
        
        for file_name in glob.glob('{}/*/setup.py'.format(BASE_SPACE)):
            self.setup.append(file_name)

        for file_name in glob.glob('{}/requirements.txt'.format(BASE_SPACE)):
            self.pipup.append(file_name)

        for file_name in glob.glob('{}/*/requirements.txt'.format(BASE_SPACE)):
            self.pipup.append(file_name)

        return True

    def git_reset(self):
        cmd = ['git', 'checkout', self.git_ref, '--']
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd, cwd=os.path.join(BASE_SPACE, self.project_name), stdout=subprocess.PIPE)
        if r.returncode != 0:
            print("Git checkout failed", file=sys.stderr)
            return False

        return True

    def python_install(self):
        for path in self.setup:
            file_name = os.path.basename(path)
            dir_name = os.path.dirname(path)
            r = subprocess.run('cd {}; python3 {} install'.format(dir_name, file_name), shell=True)
            if r.returncode != 0:
                print("Install dependences failed", file=sys.stderr)
                return False
        return True

    def pip_install(self):
        for file_name in self.pipup:
            r = subprocess.run(['pip3', 'install', '-r', file_name])
            if r.returncode != 0:
                print("Install dependences failed", file=sys.stderr)
                return False
        return True

    def build(self):
        cmd = ['pdoc', '--html-dir', '/tmp/output', '--html', self.entry_mod, '--all-submodules']
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd)

        if r.returncode != 0:
            print("pdoc error", file=sys.stderr)
            return False

        return True

