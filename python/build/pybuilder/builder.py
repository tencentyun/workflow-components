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
            print >> sys.stderr, "environment variable GIT_CLONE_URL is required"
            
            exit(1)

        self.project_name = os.path.basename(self.git_clone_URL.rstrip('.git'))

    def run(self):
        os.chdir(BASE_SPACE)

        return self.git_pull() and self.git_reset() and self.build() 

    def git_pull(self):
        cmd = ['git', 'clone', '--recurse-submodules', self.git_clone_URL, self.project_name]
        print >> sys.stdout, "Run CMD %s" % ' '.join(cmd) 
        r = subprocess.call(cmd)

        if r != 0:
            print >> sys.stderr, "Git clone failed"
            
            return False

        return True

    def git_reset(self):
        cmd = ['git', 'checkout', self.git_ref, '--']
        print >> sys.stdout, "Run CMD %s" % ' '.join(cmd)   
        r = subprocess.call(cmd, cwd=os.path.join(BASE_SPACE, self.project_name), stdout=subprocess.PIPE)
        if r != 0:
            print >> sys.stderr, "Git checkout failed"
            
            return False

        return True

    def build(self):
        r = subprocess.call('pyb {}/{}'.format(BASE_SPACE, self.project_name), shell=True)

        if r != 0:

            # print("Nuitka error", file=sys.stderr)
            return False

        return True
