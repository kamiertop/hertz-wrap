pipeline {
    // 任何一个代理可用就可以执行
    agent any
    // 定义流水线的加工流程, 流水线的所有阶段
    stages {
        stage('build code') {
            agent {
                docker {image 'golang:1.22'}
            }
            steps {
                echo "start build project"
                sh 'go version'
            }
        }
        stage('test code') {
            steps {
                echo "start test project"
            }
        }
        stage('run code') {
            steps {
                echo "start run code"
            }
        }
    }
}