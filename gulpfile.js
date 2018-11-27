const fsx = require('fs-extra')
const gulp = require('gulp')
const del = require('del')
const exec = require('child_process').exec
const sequence = require('gulp-sequence')

const dest = 'build'
const env = process.env.GIN_MODE || 'local'

console.log(`============= ${env} =============\n`)

// 清空目录
gulp.task('clean', function () {
  fsx.emptyDirSync(dest)
  fsx.ensureDirSync(dest)
})

// 接口文档
gulp.task('build:apidoc', function (cb) {
  return exec('apidoc -i routes/ -o ./assets/public/apidoc', function (err, stdout, stderr) {
    console.log(stdout)
    console.log(stderr)
    cb(err)
  })
})

// 编译服务端
gulp.task('build:server', function (cb) {
  return exec('go build -tags=jsoniter -o web .', function (err, stdout, stderr) {
    console.log(stdout)
    console.log(stderr)
    cb(err)
  })
})

// 复制依赖文件
gulp.task('copy', function () {
  return gulp.src([
    'Dockerfile',
    '*assets/**/*',
    '*conf/**/*',
    'web'
  ]).pipe(gulp.dest(dest))
})

// 安装依赖
gulp.task('install', function (cb) {
  [
    'npm install',                              // install gulp   depend
    'npm install apidoc -g',                    // install apidoc depend
    'go get github.com/kardianos/govendor',     // install go pkg depend
    'go get github.com/json-iterator/go',       // install go build depend
    'govendor sync',                            // install project  depend
  ]
  .reduce(async (acc, curr) => {
    let some;
    if(!acc) {
      some = new Promise((res, rej) => {
        exec(curr, function (err, stdout, stderr){
          if(err) {
            rej(err);
          } else {
            console.log(stdout)
            console.log(stderr)
            res(null);
          }
        })
      })
    } else {
      some = acc;
    }
    await some;
    return new Promise((res, rej) => {
      exec(curr, function (err, stdout, stderr){
        if(err) {
          rej(err);
        } else {
          console.log(stdout)
          console.log(stderr)
          res();
        }
      })
    })
  }, null)
  .then(res => {
    cb(res)
  })
  .catch(err => {
    cb(err)
  })
})

// 清除任务
gulp.task('clean', function (cb) {
  del([
    'web',
  ], cb);
});


gulp.task('default', sequence('install', 'clean', 'build:apidoc', 'build:server', 'copy', 'clean'))
