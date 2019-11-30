.PHONY: standard all report frontpage cleanfrontpage cleanreport cleanall

standard: report

all: frontpage report

report:
	pdflatex -shell-escape master.tex
	xdg-open master.pdf

frontpage:
	cd frontpage && pdflatex report.tex

cleanfrontpage:
	cd frontpage && make clean && rm frontpage.pdf

clean:
	rm master.aux master.log master.toc

cleanall: cleanfrontpage cleanreport
	rm master.pdf
