# gh-actions-orb

A CircleCI orb that runs GitHub Actions.

Still a work-in-progress.

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

* Publish the orb.
* Implement nodejs Actions.
* Implement non-Dockerfile-based Docker Actions.
* Capture [logging commands](https://help.github.com/en/actions/automating-your-workflow-with-github-actions/development-tools-for-github-actions#logging-commands) that appear in stdout and properly simulate the corresponding behavior.