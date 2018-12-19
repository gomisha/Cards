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
        withSonarQubeEnv ('SonarQube') {
          sh 'mvn clean package sonar:sonar'
        }
        echo 'really finished testing2'
      }
    }
    stage('Deployment') {
      steps {
        echo 'deploying...'
      }
    }
  }
}
