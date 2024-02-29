# pull official base image
FROM python:3.10-alpine

ENV APP_HOME=/usr/src/app

RUN mkdir $APP_HOME
RUN mkdir $APP_HOME/static
RUN mkdir $APP_HOME/media
WORKDIR $APP_HOME

ENV PYTHONDONTWRITEBYTECODE 1

ENV PYTHONUNBUFFERED 1

RUN apk update \
    && apk add libcurl curl-dev python3-dev libc-dev postgresql-dev build-base gcc python3-dev musl-dev libffi-dev \
    py3-pillow freetype-dev libpng-dev openblas-dev g++ \
    jpeg-dev zlib-dev lcms2-dev openjpeg-dev tiff-dev tk-dev tcl-dev libxslt-dev

RUN pip install --upgrade pip

COPY ./requirements/dev.txt $APP_HOME/dev.txt
RUN pip install -r dev.txt
RUN addgroup -S app && adduser -S app -G app

COPY napback $APP_HOME

COPY ./build/entrypoint.sh /entrypoint.sh
RUN sed -i 's/\r$//g' /entrypoint.sh
RUN chmod +x /entrypoint.sh

RUN chown -R app:app $APP_HOME
RUN chmod -R 755 $APP_HOME/media
RUN chmod -R 755 $APP_HOME/static

# change to the app user
USER app

ENTRYPOINT [ "/entrypoint.sh" ]


# TODO:
CMD [ "python", "manage.py", "runserver", "0.0.0.0:8000" ]
# CMD ["gunicorn", "--bind", "0.0.0.0:8000", "--workers", "3", "src.core.wsgi:application"]