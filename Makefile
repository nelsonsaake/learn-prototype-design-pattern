.PHONY: run commit
.SILENT:
.ONESHELL:

run:
	cls
	cd 2
	go run .
	cd ..

commit:
	git add .
	git commit -m "chore: commit everything"
	git push origin main