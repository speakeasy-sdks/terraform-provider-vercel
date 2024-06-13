.PHONY: speakeasy
speakeasy:
	speakeasy run

openapi-trimmed.yaml: vercel-api.yaml Makefile
	speakeasy openapi transform filter-operations --operations createDeployment --out openapi-trimmed.yaml -s vercel-api.yaml
