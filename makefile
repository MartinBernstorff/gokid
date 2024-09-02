patch:
	git tag $(shell git-semver -target patch)
	git push origin --tags

minor:
	git tag $(shell git-semver -target minor)
	git push origin --tags

fix:
	make minor

major:
	git tag $(shell git-semver -target major)
	git push origin --tags

feat:
	make major