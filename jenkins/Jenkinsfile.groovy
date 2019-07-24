@Library('jenkins-integration-module@master') _
def artiMavenMap = [:] //Required for mavenArti* helpers
pipeline {
	
    agent { label 'dubinx2-sba-slave' }

    options { 
        buildDiscarder(logRotator(numToKeepStr: '2')) //Number of run results to keep
        disableConcurrentBuilds() //Ensure only one
        timeout(time: 1, unit: 'HOURS') //Timeout for the whole job
        retry(0) //Retry on failure
        timestamps() //Prepend all console output with timestamps
    }
   
    environment {
    //Use Pipeline Utility Steps plugin to read information from pom.xml into env variables
        VERSION = readMavenPom().getVersion()
    } 

    
    triggers { pollSCM('*/5 * * * *') }

    parameters {
        booleanParam(name: 'PERFORM_RELEASE', defaultValue: false, description: 'Performs release if not SNAPSHOT version')
        string(name: 'RELEASE_VERSION', defaultValue: '', description: 'Version to be applied on the release artifact. Overrides version in POM')
        booleanParam(name: 'WRITE_POM', defaultValue: true, description: 'Commits and pushes to git, RELEASE_VERSION into version tag in POM file')
        string(name: 'RELEASE_TAG', defaultValue: '', description: 'If not empty, create git tag named as specified ')
        string(name: 'RELEASE_BRANCH', defaultValue: '', description: 'If not empty, create git branch named as specified ')
        string(name: 'NEXT_VERSION', defaultValue: '', description: 'Next version to be set inside the pom files.')
    }

    stages {
        stage('Init') {
            steps {
                script {
                    currentBuild.displayName = "${GIT_BRANCH}_${VERSION}"
                    mavenArtiDefaults(artiMavenMap)
                    artiMavenMap['maventool'] = "Maven 3.3.3"
                    artiMavenMap['deployreleaseRepo'] = "modules-release-local"
                    artiMavenMap['deploysnapshotRepo'] = "modules-snapshot-local"
                    mavenArtiInit(artiMavenMap)
                }
                sh 'printenv'
            }
        }
        stage('pre-build cleanup') {
             steps {
                sh 'if [ -d "$WORKSPACE/.repository" ]; then rm -Rf $WORKSPACE/.repository; fi'    
             }
        }
        stage('build') {
          steps {
              mavenArtiRun artiMavenMap, "clean install"
          }
        }
        stage('Publish Maven Artifact') {
            steps {
                mavenArtiDeploy artiMavenMap
                mavenArtiNextVersion artiMavenMap
            }
        }
    }
    post {
        always {
            deleteDir()
        }
        regression {
            archiveArtifacts '**/*'
        }

        failure {
           slackSend color: 'danger', message: "Branch <$GIT_BRANCH> Build <${RUN_DISPLAY_URL}|${BUILD_DISPLAY_NAME}> failed for job `<${JOB_DISPLAY_URL}|${JOB_NAME}>`"
        }
        fixed {
           slackSend color: 'good', message: "Branch <$GIT_BRANCH> Build is now passing for job `<${JOB_DISPLAY_URL}|${JOB_NAME}>` (since <${RUN_DISPLAY_URL}|${BUILD_DISPLAY_NAME}>)"
        }
    }
}
