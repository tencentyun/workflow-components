import os
import sys
import subprocess

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

        self.project_name = os.path.basename(self.git_clone_URL.rstrip('.git'))

    def run(self):
        os.chdir(BASE_SPACE)

        return self.git_pull() and self.git_reset() and self.build()

    def git_pull(self):
        cmd = ['git', 'clone', '--recurse-submodules', self.git_clone_URL, self.project_name]
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd)

        if r.returncode != 0:
            print("Git clone failed", file=sys.stderr)
            return False

        return True

    def git_reset(self):
        cmd = ['git', 'checkout', self.git_ref, '--']
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd, cwd=os.path.join(BASE_SPACE, self.project_name), stdout=subprocess.PIPE)
        if r.returncode != 0:
            print("Git checkout failed", file=sys.stderr)
            return False

        return True

    def build(self):
        r = subprocess.run('cd {}/{}; yes | mk-build-deps -ri; dpkg-buildpackage -us -uc -b'.format(BASE_SPACE, self.project_name), shell=True)

        if r.returncode != 0:
            print("Nuitka error", file=sys.stderr)
            return False

        return True