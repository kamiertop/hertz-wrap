pipeline {
    // 任何一个代理可用就可以执行
    agent any
    // 定义流水线的加工流程, 流水线的所有阶段
    stages {
        stage('build code') {
            agent {
                docker {image 'golang:latest'}
            }
            steps {
                echo "start build project"
                sh 'go version'
                sh 'go env -w GOPROXY=https://goproxy.cn,direct'
                sh 'go mod tidy'
                sh 'go build -o hertz-api main.go'
            }
        }
        stage('test code') {
            agent {
                docker {image 'golang:latest'}
            }
            steps {
                echo "start test project"

            }
        }
        stage('run code') {
            steps {
                echo "start run code"
                sh 'nohup ./hertz-api &'
            }
        }
    }
}