.PHONY: standard all report frontpage cleanfrontpage cleanreport cleanall open

standard: report

all: frontpage report

report:
	pdflatex -shell-escape master.tex 

frontpage:
	cd frontpage && pdflatex report.tex

cleanfrontpage:
	cd frontpage && make clean && rm frontpage.pdf

clean:
	rm master.aux master.log master.toc

cleanall: cleanfrontpage cleanreport
	rm master.pdf
open: report
	xdg-open master.pdf
