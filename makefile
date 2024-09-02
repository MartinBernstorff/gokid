patch:
	git tag $(git-semver -target minor)
	git push origin --tags

minor:
	git tag $(git-semver -target minor)
	git push origin --tags