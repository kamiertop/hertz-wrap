pipeline {
    // 任何一个代理可用就可以执行
    agent any
    // 定义流水线的加工流程, 流水线的所有阶段
    stages {
        stage('build code') {
            steps {
                echo "start build project"
            }
        }
        stage('test code') {
            steps {
                sh '/usr/local/go/bin/go version'
                sh '/usr/local/go/bin/go env'
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