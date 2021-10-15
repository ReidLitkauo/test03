docker build -t gcr.io/test00-328920/default .

# To run the docker server:
# docker run -p 8080:8080 gcr.io/test00-328920/default:latest

# To open a terminal inside the docker container:
# docker ps (to find the CONTAINER ID on the far left)
# docker exec -ti XXXXXXXXXXXX /bin/sh

# To push to Container Registry (is this necessary???):
# docker push gcr.io/test00-328920/default:latest

# To deploy image:
# gcloud app deploy --image-url=gcr.io/test00-328920/default:latest