
add:
	git add -A

push:
	git config --local user.name "LynchQGit"
	git config --local user.email "hlq_git@163.com"
	git pull --rebase
	git push --set-upstream origin HEAD