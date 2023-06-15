GITLEAKS_DOWNLOAD_DIR="${HOME}/bin/gitleaks-files"
TARGETOS ?=$(shell uname -s | tr '[:upper:]' '[:lower:]')
ifeq ($(shell uname -m), x86_64)
	TARGETARCH="x64"
else
	TARGETARCH=$(shell uname -m)
endif

URL=$(shell curl -sSfL https://api.github.com/repos/zricethezav/gitleaks/releases/latest |\
	grep browser_download_url |\
	cut -d '"' -f 4 |\
	grep ${TARGETOS}"_"${TARGETARCH})

GIT_SUPPORT_PATH=${HOME}/.git-support
HOOKS=${GIT_SUPPORT_PATH}/hooks
PRECOMMIT=${GIT_SUPPORT_PATH}/hooks/pre-commit
PATTERNS=${GIT_SUPPORT_PATH}/gitleaks.toml
GITLEAKS= ${HOME}/bin/gitleaks


INSTALL_TARGETS= ${PATTERNS} ${PRECOMMIT} ${GITLEAKS}

install: $(INSTALL_TARGETS) global_hooks version

clean:
	/bin/rm -rf ${GIT_SUPPORT_PATH}
	git config --global --unset hooks.gitleaks
	git config --global --unset core.hooksPath
	/bin/rm -rf ${GITLEAKS}

global_hooks:
	git config --global hooks.gitleaks true
	git config --global core.hooksPath ${GIT_SUPPORT_PATH}/hooks

version:
	@if [ -x "$$(command -v ${GITLEAKS})" ]; then \
		echo "Gitleaks version $(shell ${GITLEAKS} version) OS=${TARGETOS} ARCH=${TARGETARCH}"; \
	else \
		echo "Gitleaks not installed"; \
	fi

detect:
	${GITLEAKS} detect -v

${PATTERNS}: ${GIT_SUPPORT_PATH}
	wget -q -O $@ https://raw.githubusercontent.com/zricethezav/gitleaks/master/config/gitleaks.toml

${PRECOMMIT}: pre-commit.sh ${HOOKS}
	install -m 0755 -cv $< $@

${GIT_SUPPORT_PATH} ${HOOKS}:
	mkdir -p $@

${GITLEAKS}:
	mkdir -p ${GITLEAKS_DOWNLOAD_DIR}
	wget -q ${URL} -O - | tar -xzf - -C ${GITLEAKS_DOWNLOAD_DIR}
	cp ${GITLEAKS_DOWNLOAD_DIR}/gitleaks ${GITLEAKS}
	rm -rf ${GITLEAKS_DOWNLOAD_DIR}
	chmod 755 $@
	@echo "Gitleaks installed to ${GITLEAKS}"
