(TeX-add-style-hook
 "bias-variance-dart"
 (lambda ()
   (TeX-add-to-alist 'LaTeX-provided-class-options
                     '(("minimal" "dvipsnames")))
   (TeX-add-to-alist 'LaTeX-provided-package-options
                     '(("preview" "pdftex" "active" "tightpage")))
   (TeX-run-style-hooks
    "latex2e"
    "minimal"
    "minimal10"
    "tikz"
    "preview"
    "pgfplots"
    "filecontents")))

