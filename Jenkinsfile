pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        echo 'building...'
      }
    }
    stage('Test') {
      steps {
        echo 'testing...'
        withSonarQubeEnv 'SonarQube'
        echo 'reallyl finished testing'
      }
    }
    stage('Deployment') {
      steps {
        echo 'deploying...'
      }
    }
  }
}