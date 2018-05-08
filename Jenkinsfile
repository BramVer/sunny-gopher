pipeline {

    agent any

    environment {
        REPO_SERVER = 'repo.youkebox.be'
        REPO_PATH   = "/var/vhosts/repo/${env.GIT_BRANCH}-sunny-gopher"
        APPL_SERVER = '159.89.14.97'
        NAME        = 'sunny-gopher'
        VERSION     = '0.1'
        DESCRIPTION = 'Returns random quotes from IASIP.'
        ARCH        = 'x86_64'
    }

    stages {
        stage('Build') {
            steps {
                sh 'make build'
            }
        }

        stage('Package') {
            steps {
                sh "make package --environment-overrides BUILD_NO=${env.BUILD_NUMBER}"
            }
        }

        stage('Upload') {
            when {
                allOf {
                    expression { env.CHANGE_ID == null  }
                }
            }
            steps {
                sh "ssh root@${REPO_SERVER} 'mkdir -p ${REPO_PATH}/packages/'"
                sh "scp sunny-gopher-*.rpm root@${REPO_SERVER}:${REPO_PATH}/packages/"
                sh "ssh root@${REPO_SERVER} 'cd ${REPO_PATH}/packages/ && rm -rf \$(ls ${REPO_PATH}/packages/ -1t | grep ${NAME}-${VERSION} | tail -n +4)'"
                sh "ssh root@${REPO_SERVER} 'createrepo --update ${REPO_PATH}'"
            }
        }

        stage('Deploy') {
            agent any
            when {
                allOf {
                    expression { env.CHANGE_ID == null  }
                }
            }
            steps {
                 sh "ssh root@${APPL_SERVER} 'yum makecache; yum update sunny-gopher -y'"
                 sh "ssh root@${APPL_SERVER} 'systemctl restart sunny-gopher'"
            }
        }
    }
}
