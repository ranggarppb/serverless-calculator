#!/bin/sh

npm install -g @commitlint/cli @commitlint/config-conventional
echo "module.exports = {extends: ['@commitlint/config-conventional']};" > commitlint.config.js
pre-commit install
