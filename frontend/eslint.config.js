import js from '@eslint/js';
import globals from 'globals';
import vue from 'eslint-plugin-vue';
import { defineConfig } from 'eslint/config';

const conf = [
    {
        files: ['**/*.{js,mjs,cjs,vue}'],
        plugins: { js, vue },
        extends: ['js/recommended'],
        languageOptions: {
            globals: globals.browser,
            ecmaVersion: 2020,
            sourceType: 'module',
        },
        rules: {
            // 关闭组件名校验 Component name xxx should always be multi-word
            // 'vue/multi-word-component-names': 'off',

            // 空格
            indent: [
                'error',
                4,
                {
                    SwitchCase: 1, // switch case 缩进 1 级
                    VariableDeclarator: 1, // 变量声明缩进
                    MemberExpression: 1, // 成员表达式缩进
                    FunctionDeclaration: {
                        parameters: 1, // 函数参数缩进
                    },
                },
            ],
            quotes: [
                'error',
                'single',
                {
                    avoidEscape: true, // 允许在包含单引号的字符串中使用双引号
                    allowTemplateLiterals: true, // 允许使用模板字符串
                },
            ],
            semi: ['error', 'always'], // 必须分号
            'comma-dangle': ['error', 'never'], // 禁止尾随逗号
            'object-curly-spacing': ['error', 'always'], // 对象花括号内空格
            'array-bracket-spacing': ['error', 'never'], // 数组括号内无空格
            // 去除代码行尾的空白
            'no-trailing-spaces': 'error',

            // 花括号 { 前添加空格
            'space-before-blocks': 'error',

            // if( 括号后添加1格空格
            'keyword-spacing': ['error', { before: true, after: true }],
            'arrow-spacing': ['error', { before: true, after: true }],
            'space-infix-ops': 'error',
            'comma-spacing': ['error', { before: false, after: true }],
            'semi-spacing': ['error', { before: false, after: true }],
            'block-spacing': 'error',
            'padding-line-between-statements': [
                'error',
                // 在变量声明（const, let, var）之后，除了 import 和 export，都需要空行
                { blankLine: 'always', prev: ['const', 'let', 'var'], next: '*' },
                { blankLine: 'any', prev: ['const', 'let', 'var'], next: ['const', 'let', 'var'] },
                // 在函数声明、函数表达式、类声明、export 之后，需要空行
                { blankLine: 'always', prev: ['function', 'class', 'export'], next: '*' },
                { blankLine: 'always', prev: '*', next: ['function', 'class', 'export'] },
                // 在 import 语句之后，需要空行
                { blankLine: 'always', prev: 'import', next: '*' },
                { blankLine: 'any', prev: 'import', next: 'import' },
                // 在 return 语句之前，需要空行
                { blankLine: 'always', prev: '*', next: 'return' },
                // 在块语句（block）之后，需要空行？这通常不强制，但可以根据需要配置
                // 注意：块语句（block）是指由 {} 括起来的代码块，但这里我们可能不希望在块内的语句之间强制空行，所以需要谨慎。
                // 以下配置可选的，根据你的喜好添加
                // 在 if 语句、for 语句、while 语句等之后，通常不强制空行，但有时在它们之后和下一个语句之间需要空行。
                // 例如：在 if 语句块之后和接下来的语句之间空一行
                // { blankLine: "always", prev: "block", next: "*" },
                // 但是，注意这样可能会在代码中插入很多空行，所以请根据团队风格决定。
            ],
            curly: ['error', 'all'], // 所有控制语句都必须使用花括号
        },
    },
    // 应用 Vue 官方推荐的规则集
    vue.configs['flat/essential'],
    vue.configs['flat/recommended'],
];
export default defineConfig(conf);
