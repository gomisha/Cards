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
          sh "/opt/sonar-scanner/bin/sonar-scanner -D sonar-project.properties"
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
