VERSION=210303

echo "Building Docker image..."
sudo docker build -t dataspects/mwmapi:${VERSION} .

# sudo docker login
# sudo docker push dataspects/mwmapi:${VERSION}