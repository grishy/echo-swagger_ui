default: update_swagger

SWAGGER_UI_DIR=./swagger_ui
SWAGGER_UI_VERSION=3.50.0
SWAGGER_UI_URL=https://github.com/swagger-api/swagger-ui/archive/refs/tags/v${SWAGGER_UI_VERSION}.tar.gz

update_swagger: download_swagger remove_mapfile_swagger

download_swagger:
	@rm -rf ${SWAGGER_UI_DIR} ./download
	@mkdir -p ${SWAGGER_UI_DIR} ./download
	@wget -c ${SWAGGER_UI_URL} -O - | tar -xz -C ./download
	@cp -ap  "./download/swagger-ui-${SWAGGER_UI_VERSION}/dist/" ${SWAGGER_UI_DIR}
	@rm -rf ./download

remove_mapfile_swagger:
	@rm ${SWAGGER_UI_DIR}/*.map
