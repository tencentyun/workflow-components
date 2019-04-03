import os
import sys
import json
import subprocess
import pycco.main as pyccoLib
import anymarkup

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

    def git_reset(self):
        cmd = ['git', 'checkout', self.git_ref, '--']
        print("Run CMD %s" % ' '.join(cmd), file=sys.stdout)
        r = subprocess.run(cmd, cwd=os.path.join(BASE_SPACE, self.project_name), stdout=subprocess.PIPE)
        if r.returncode != 0:
            print("Git checkout failed", file=sys.stderr)
            return False

        return True

    def build(self):
        for root, dirs, files in os.walk(os.path.join(BASE_SPACE, self.project_name)):
            for file_name in files:
                if file_name.endswith('.py'):
                    o = self.pycco(os.path.join(root, file_name))
                    all_true = all_true and o

    def trim_repo_path(self, n):
        return n[len(BASE_SPACE) + 1:]

    def pycco(self, file_name):
        code = open(file_name, 'rb').read().decode('utf-8')
        language = pyccoLib.get_language(file_name, code)
        sections = pyccoLib.parse(code, language)
        pyccoLib.highlight(sections, language, outdir="docs")

        data = { 'file_path': self.trim_repo_path(file_name), 'sections': sections }
        
        print('[COUT] CO_JSON_CONTENT {}'.format(json.dumps(data)))

        return True
