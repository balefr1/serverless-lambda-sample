sed -ie "s/gqmdi3ux67.execute-api.eu-south-1.amazonaws.com/$1/g" frontend/react/build/static/js/*
cd frontend/react/build
aws s3 sync . s3://$2
