# Release Process

This document outlines the steps to create a new release for gokid using GoReleaser and GitHub Actions.

## Prerequisites

1. Ensure you have push access to the main repository.
2. Make sure all changes you want to include in the release are merged into the main branch.

## Steps to Create a Release

1. Update the version number in your project files (if applicable).

2. Commit any changes:
   ```
   git add .
   git commit -m "Prepare for release vX.Y.Z"
   ```

3. Create and push a new tag with the version number:
   ```
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```

4. The GitHub Actions workflow will automatically trigger when you push the new tag.

5. Go to the GitHub repository's "Actions" tab to monitor the release process.

6. Once the workflow completes successfully, go to the "Releases" page on GitHub to find your new release with attached artifacts.

7. The release will automatically be added to the Homebrew tap (martinbernstorff/homebrew-tap).

8. Update the release notes on GitHub with any relevant information about the new version.

## Troubleshooting

If the release fails, check the GitHub Actions logs for any error messages. Common issues include:

- Incorrect version formatting (should be vX.Y.Z)
- Missing permissions for the GitHub token
- Issues with the .goreleaser.yaml configuration

For any persistent problems, please open an issue in the repository.

## Local Testing

To test the release process locally before pushing:

1. Install GoReleaser (https://goreleaser.com/install/)

2. Run:
   ```
   goreleaser release --snapshot --clean
   ```

This will simulate the release process without publishing anything.