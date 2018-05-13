pipeline {

    agent any

    environment {
        SERVER = '159.89.14.97'
        REPO_PATH   = "/var/repo/sunny-gopher"
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
                sh "ssh root@${SERVER} 'mkdir -p ${REPO_PATH}/packages/'"
                sh "scp sunny-gopher-*.rpm root@${SERVER}:${REPO_PATH}/packages/"
                sh "ssh root@${SERVER} 'cd ${REPO_PATH}/packages/ && rm -rf \$(ls ${REPO_PATH}/packages/ -1t | grep ${NAME}-${VERSION} | tail -n +4)'"
                sh "ssh root@${SERVER} 'createrepo --update ${REPO_PATH}'"
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
                 sh "ssh root@${SERVER} 'yum makecache; yum update sunny-gopher -y'"
                 sh "ssh root@${SERVER} 'systemctl restart sunny-gopher'"
            }
        }
    }
}
