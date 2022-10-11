sed -ie "s/api.example.com/$1/g" frontend/react/build/static/js/*
cd frontend/react/build
aws s3 sync . s3://$2
