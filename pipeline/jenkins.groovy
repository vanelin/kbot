pipeline {
    agent any
    parameters {
        choice(name: 'OS', choices: ['linux', 'darwin', 'windows', 'all'], description: 'Pick OS')
        choice(name: 'TARGETOSARCH', choices: ['amd64', 'arm64'], description: 'Pick architecture')
    }
    environment{
        REPO = 'https://github.com/vanelin/kbot'
        BRANCH = 'develop'
        REGISTRY = 'vanelin'
        DOCKERHUB_CREDENTIALS=credentials('dockerhub')
        MACOSHOST_CREDENTIALS=credentials('machost')
    }
    stages {
        stage('clone') {
            steps {
                echo 'Clone repo'
                git branch: "${BRANCH}", url: "${REPO}"
            }
        }
        stage('Login to DockerHUB') {

            steps {
                sh 'security unlock-keychain -p $MACOSHOST_CREDENTIALS_PSW'
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
            }
        }
        stage('test') {
            steps {
                echo 'run test'
                sh 'make test'
            }
        }
        stage('build') {
            parallel {
                stage('Build Linux') {
                    when { expression { params.OS == 'linux' || params.OS == 'all' } }
                    steps {
                        echo 'Building for Linux'
                        sh 'make image TARGETOS=linux'
                    }
                }
                stage('Build Darwin') {
                    when { expression { params.OS == 'darwin' || params.OS == 'all' } }
                    steps {
                        echo 'Building for macos'
                        sh 'make image TARGETOS=macos'
                    }
                }
                stage('Build Windows') {
                    when { expression { params.OS == 'windows'  || params.OS == 'all' } }
                    steps {
                        echo 'Building for Windows'
                        sh 'make image TARGETOS=windows'
                    }
                }
            }
        }
        stage('push') {
            parallel {
                stage('Push Linux to dockerhub') {
                    when { expression { params.OS == 'linux' || params.OS == 'all' } }
                    steps {
                        sh 'make push TARGETOS=linux'
                    }
                }
                stage('Push Darwin to dockerhub') {
                    when { expression { params.OS == 'darwin' || params.OS == 'all' } }
                    steps {
                        sh 'make push TARGETOS=macos'
                    }
                }
                stage('Push Windows to dockerhub') {
                    when { expression { params.OS == 'windows' || params.OS == 'all' } }
                    steps {
                        sh 'make push TARGETOS=windows'
                    }
                }
            }
        }
        stage('clean') {
            parallel {
                stage('Clean Linux image on MacOS host ') {
                    when { expression { params.OS == 'linux' || params.OS == 'all' } }
                    steps {
                        sh 'make clean TARGETOS=linux'
                    }
                }
                stage('Clean Darwin image on MacOS host') {
                    when { expression { params.OS == 'darwin' || params.OS == 'all' } }
                    steps {
                        sh 'make clean TARGETOS=macos'
                    }
                }
                stage('Clean Windows image on MacOS host') {
                    when { expression { params.OS == 'windows' || params.OS == 'all' } }
                    steps {
                        sh 'make clean TARGETOS=windows'
                    }
                }
            }
        }
    }
}
