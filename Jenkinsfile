pipeline {
    agent any
    stages {
        stage('deploy') {
            steps {
                sh "docker-compose up -d"
            }
        }
    }
}