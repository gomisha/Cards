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
          sh '/var/jenkins_home/sonar-scanner/sonar-scanner-3.2.0.1227-linux/bin/sonar-scanner'
        }
        echo 'really finished testing2'
      }
    }
    stage("Quality Gate") {
        steps {
            timeout(time: 1, unit: 'HOURS') {
              waitForQualityGate abortPipeline: true
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
