#!/usr/bin/env ts-node

import shell from 'shelljs';
import path from 'path';

const ProjectName = 'hookexec';

if (!shell.which('go')) {
  shell.echo('Sorry, this script requires go');
  shell.exit(1);
}

if (!shell.which('npm')) {
  shell.echo('Sorry, this script requires npm');
  shell.exit(1);
}

const currentDir = process.cwd();
const frontendPath = path.join(currentDir, 'frontend');

// 前端先打包一次
shell.cd(frontendPath);
console.info('进入前端目录:', shell.pwd().toString());
shell.exec('npm run build');

// 再处理后端
shell.cd(currentDir);
console.info('当前执行命令的目录:', shell.pwd().toString());
shell.exec('go mod tidy');

shell.rm('-rf', 'bin');
shell.mkdir('bin');

// 判断系统类型，设置不同的输出文件名
if (process.platform === 'win32') {
  shell.exec(`go build -o bin\\${ProjectName}.exe`);
} else {
  shell.exec(`go build -o bin/${ProjectName}`);
}

process.exit(0);
