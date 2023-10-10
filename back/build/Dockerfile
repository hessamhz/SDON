# Base image
FROM python:3.10.6-alpine

# Create the appropriate directories
ENV APP_HOME=/usr/src/app
RUN mkdir -p $APP_HOME

# Set the working directory
WORKDIR $APP_HOME

# Copy the requirements file
COPY requirements/common.txt $APP_HOME/requirements/

# Copy the entrypoint file
COPY build/entrypoint.sh $APP_HOME/entrypoint.sh

# Copy the .env file
COPY wordcounter/.env $APP_HOME/

# Install dependencies
RUN pip install --no-cache-dir -r $APP_HOME/requirements/common.txt

# Copy the Django project
COPY wordcounter $APP_HOME/wordcounter

# Switch to the non-root user
RUN addgroup -S app && adduser -S -G app app
RUN chown -R app:app $APP_HOME
USER app

# Set the entrypoint script
RUN chmod +x $APP_HOME/entrypoint.sh
ENTRYPOINT ["/usr/src/app/entrypoint.sh"]

# Expose the port that Gunicorn runs on
EXPOSE 8000

# Command to run Gunicorn
# TODO I usually use gunicorn here, but it wasn't needed in this context
# CMD ["gunicorn", "-w", "4", "-b", "0.0.0.0:8000", "wordcounter.core.asgi:application", "-k", "wordcounter.core.worker.Worker"]

# Command to run Django
CMD ["python", "wordcounter/manage.py", "runserver", "0.0.0.0:8000"]