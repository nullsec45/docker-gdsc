FROM nginx:latest

COPY index.html /usr/share/nginx/html

ADD https://wordpress.org/latest.zip /usr/share/nginx/html

# RUN mkdir /usr/share/nginx/wordpress/
# RUN mv /usr/share/nginx/html/latest.zip /usr/share/nginx/html/wordpress/
# RUN unzip -q latest.zip 

RUN mkdir /usr/share/nginx/html/simplerun

CMD ["cat","/usr/share/nginx/html/index.html"]