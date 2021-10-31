pipeline {
    agent any
    stages {
        stage('build') {
            steps {
                sh "docker build -t ilkerispir/xds ."
                sh "docker build -t ilkerispir/resource ./resource"
            }
        }
        stage('push') {
            steps {
                sh "docker push ilkerispir/xds"
                sh "docker push ilkerispir/resource"
            }
        }
        stage('deploy') {
            steps {
                sh "docker-compose up -d"
            }
        }
    }
}