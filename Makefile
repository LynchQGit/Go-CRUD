
add:
	git add -A

push:
	git config --local user.name "LynchQGit"
	git config --local user.email "hlq_git@163.com"
	git remote set-url origin https://github.com/LynchQGit/Go-CRUD.git
	git pull --rebase
	git push --set-upstream origin HEAD

fmt:
	gofmt -w -l .