<?php
class Builder {
  const workdir = '/root/src';

  function __construct( $envs ) {
    if ($envs["GIT_CLONE_URL"]) {
      $this->gitCloneURL = rtrim($envs["GIT_CLONE_URL"], '/');
      $this->gitRef = $envs["GIT_REF"];
    } else if ($envs["_WORKFLOW_GIT_CLONE_URL"]) {
      $this->gitCloneURL = rtrim($envs["_WORKFLOW_GIT_CLONE_URL"], '/');
      $this->gitRef = $envs["_WORKFLOW_GIT_REF"];
    }

    if (!$this->gitRef) {
      $this->gitRef = 'master';
    }

    if (!$this->gitCloneURL) {
      fwrite(STDERR, "envionment variables GIT_CLONE_URL is required\n");
      exit(1);
    }

    $this->projectName = basename($this->gitCloneURL, '.git');
    $this->projectPath = self::workdir . '/' . $this->projectName;

    $this->hubRepo = rtrim($envs["HUB_REPO"], '/');
    if (!$this->hubRepo) {
      return; // no need upload
    }

    $this->hubUser = $envs["HUB_USER"];
    $this->hubToken = $envs["HUB_TOKEN"];

    if (!$this->hubUser || !$this->hubToken) {
      $this->hubUser = $envs["_WORKFLOW_HUB_USER"];
      $this->hubToken = $envs["_WORKFLOW_HUB_TOKEN"];
    }

    if (!$this->hubUser || !$this->hubToken) {
      fwrite(STDERR, "envionment variable HUB_USER, HUB_TOKEN are required\n");
      exit(1);
    }

    $this->artifactPath = rtrim($envs["ARTIFACT_PATH"], '/');
    $this->artifactTag = $envs["ARTIFACT_TAG"];
    if (!$this->artifactTag) {
      $this->artifactTag = "latest";
    }
  }

  function run() {
    if (!$this->gitPull()) {
      fwrite(STDERR, "gitPull failed\n");
      return false;
    }

    if (!$this->gitReset()) {
      fwrite(STDERR, "gitReset failed\n");
      return false;
    }

    if (!$this->installDependence()) {
      fwrite(STDERR, "installDependence failed\n");
      return false;
    }

    if ($this->hubRepo && !$this->uploadDependence()) {
      fwrite(STDERR, "uploadDependence failed\n");
      return false;
    }

    return true;
  }

  static function runCMD($cmd, &$exitCode, $workdir=self::workdir) {
    fwrite(STDOUT, "Run CMD: $cmd in $workdir\n");
    chdir($workdir);
    exec($cmd, $output, $exitCode);

    $len=count($output);
    for($i=0;$i<$len;$i++) {
      fwrite(STDOUT, $output[$i]);
    }
    return $output;
  }

  function gitPull() {
    $cmd = "git clone --recurse-submodules $this->gitCloneURL $this->projectName";
    self::runCMD($cmd, $exitCode);
    return $exitCode == 0;
  }

  function gitReset() {
    $cmd = "git reset --hard $this->gitRef";
    self::runCMD($cmd, $exitCode, $this->projectPath);
    return $exitCode == 0;
  }

  function installDependence() {
    $cmd = "composer install";
    self::runCMD($cmd, $exitCode, $this->projectPath);
    return $exitCode == 0;
  }

  function uploadDependence(){
    $cmd = "tar -cjf vendor.tar.bz vendor";
    self::runCMD($cmd, $exitCode, $this->projectPath);

    if ($exitCode != 0) {
      return false;
    }

    $path = $this->artifactPath . '/vendor.tar.bz';
    $cmd = "/.workflow/bin/thub push --username=$this->hubUser --password=$this->hubToken --repo=$this->hubRepo --localpath=vendor.tar.bz --path=$path --tag=$this->artifactTag";
    self::runCMD($cmd, $exitCode, $this->projectPath);

    if ($exitCode != 0) {
      return false;
    }

    $artifactURL = $this->hubRepo . '/' . ltrim($path, '/');
    fwrite(STDOUT, "[JOB_OUT] ARTIFACT_URL = $artifactURL\n");
    return true;
  }

}

?>

