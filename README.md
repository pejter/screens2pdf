screens2tex
===========

`screens2tex` is a utility tool that takes the screenshots documenting your work and 
makes a `.tex` files neatly arranging them.  
All that's left is rendering the final document.

### If you're lazy like me and want 1 command to make it a pdf
*Disclaimer: `blang/latex` image is about 3.6GB so be prepared*  
If you have **docker** installed, just `cd` into the exercise directory and run:
`docker run --rm -i --user="$(id -u):$(id -g)" -v $PWD:/data blang/latex pdflatex ./final.tex`

### MAKE IT EVEN SHORTER DAMN IT! I HAVE DOCKER TO MAKE THINGS SIMPLE!
FINE! Here it is, just make sure you're in the directory with the screenshots:
`docker run --rm -i --user="$(id -u):$(id -g)" -v $PWD:/data pejter/screens2tex -d /data -s /data -t "<your title here>" &&
docker run --rm -i --user="$(id -u):$(id -g)" -v $PWD:/data blang/latex pdflatex 
./final.tex`  
> `final.pdf` should appear in your directory.