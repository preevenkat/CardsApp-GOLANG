pipeline {
    agent {
        docker {
            image 'golang:1.13.4'
        }
    }

    triggers {
        gitlab(triggerOnPush: true, triggerOnMergeRequest: true, branchFilterType: 'All')
    }

    post {
        failure {
            updateGitlabCommitStatus name: 'jenkins-build', state: 'failed'
        }
        success {
            updateGitlabCommitStatus name: 'jenkins-build', state: 'success'
        }
    }

    environment {
        GOPATH = "${pwd}"
    }

    stages {
        stage('Build') {
            steps {
                updateGitlabCommitStatus name: 'jenkins_file', state: 'pending'
                echo 'Linting...'
                sh 'go version'
                sh 'go get -u -v https://github.com/smartybrains/CardsApp_Go'
                sh 'golint -set_exit_status ./...'
                echo 'Building...'
                sh '''
                cd my_app
                go build
                '''
            }
        }
    }
}
