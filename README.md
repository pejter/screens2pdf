screens2pdf
===========

`screens2pdf` is a utility tool that takes the screenshots documenting your work and
makes a `.tex` files neatly arranging them.  
If `pdflatex` is installed and in PATH it will also render the document to `.pdf`

### If you have **docker** installed, just `cd` into the exercise directory and run:
`docker run --rm -i --user="$(id -u):$(id -g)" -v $PWD:/data pejter/screens2pdf -d /data -s /data -t "<your title here>"`
