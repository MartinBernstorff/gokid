patch:
	git tag $(shell git-semver -target patch)
	git push origin --tags

minor:
	git tag $(shell git-semver -target minor)
	git push origin --tags
