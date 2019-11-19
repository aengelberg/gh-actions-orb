# gh-actions orb [![CircleCI Build Status](https://circleci.com/gh/aengelberg/gh-actions-orb.svg?style=shield "CircleCI Build Status")](https://circleci.com/gh/aengelberg/gh-actions-orb) [![CircleCI Orb Version](https://img.shields.io/badge/endpoint.svg?url=https://badges.circleci.io/orb/aengelberg/gh-actions)](https://circleci.com/orbs/registry/orb/aengelberg/gh-actions-orb) [![GitHub license](https://img.shields.io/badge/license-Apache-blue.svg)](https://raw.githubusercontent.com/aengelberg/gh-actions-orb/master/LICENSE)

This orb allows CircleCI users to run the workflow automation scripts on the [GitHub Marketplace](https://github.com/marketplace?type=actions), otherwise known as "Actions". It's particularly useful for developers that want to combine the speed and flexiblity of CircleCI with third-party extensibility of GitHub Actions.

This is still a work-in-progress and not feature-complete.

Behind the scenes, this orb makes use of the [lights-camera-action](https://github.com/aengelberg/lights-camera-action) project to download and execute GitHub Actions.

## Usage

To invoke a GitHub Action in your CircleCI config, specify its repository, source revision, and inputs:

```yaml
- gh-action:
    uses: actions/hello-world-docker-action@v1
    with: |
      who-to-greet: 'Mona The Octocat'
```

This is intentionally very similar to the syntax if you were to write it in a GitHub Workflow:

```yaml
- uses: actions/hello-world-docker-action@v1
  with:
    who-to-greet: 'Mona The Octocat'
```

## TODO

* Implement nodejs Actions.
* Implement non-Dockerfile-based Docker Actions.
* Capture [logging commands](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/development-tools-for-github-actions#logging-commands) that appear in stdout and properly simulate the corresponding behavior.