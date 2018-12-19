pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        echo 'building...'
      }
    }
    stage('Test') {
      parallel {
        stage('Test') {
          steps {
            echo 'testing...'
          }
        }
        stage('SonarQube') {
          steps {
            withSonarQubeEnv 'SonarQube'
            waitForQualityGate true
          }
        }
        stage('') {
          steps {
            echo 'Finishing testing...'
          }
        }
      }
    }
    stage('Deployment') {
      steps {
        echo 'deploying...'
      }
    }
  }
}